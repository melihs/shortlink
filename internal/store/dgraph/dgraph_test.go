package dgraph

import (
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/internal/store/mock"
)

func TestDgraph(t *testing.T) {
	store := DGraphLinkList{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	assert.Nil(t, err, "Could not connect to docker")

	// create a network with Client.CreateNetwork()
	network, err := pool.Client.CreateNetwork(docker.CreateNetworkOptions{
		Name: "shortlink-test",
	})
	if err != nil {
		assert.Errorf(t, err, "Error create docker network")
		os.Exit(1)
	}

	// pulls an image, creates a container based on it and runs it
	ZERO, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "dgraph/dgraph",
		Tag:          "v2.0.0-beta",
		Cmd:          []string{"dgraph", "zero", "--my=test-dgraph-zero:5080"},
		ExposedPorts: []string{"5080"},
		Name:         "test-dgraph-zero",
		NetworkID:    network.ID,
	})
	assert.Nil(t, err, "Could not start resource")

	ALPHA, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "dgraph/dgraph",
		Tag:        "v2.0.0-beta",
		Cmd:        []string{"dgraph", "alpha", "--my=localhost:7080", "--lru_mb=2048", fmt.Sprintf("--zero=%s:%s", "test-dgraph-zero", "5080")},
		NetworkID:  network.ID,
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(ZERO); errPurge != nil {
			assert.Errorf(t, errPurge, "Could not purge resource")
		}

		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		err = os.Setenv("STORE_DGRAPH_URI", fmt.Sprintf("localhost:%s", ALPHA.GetPort("9080/tcp")))
		if err != nil {
			assert.Errorf(t, err, "Cannot set ENV")
			return nil
		}

		err = store.Init()
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		assert.Errorf(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		// When you're done, kill and remove the container
		if err := pool.Purge(ALPHA); err != nil {
			assert.Errorf(t, err, "Could not purge resource")
		}

		// When you're done, kill and remove the container
		if err := pool.Purge(ZERO); err != nil {
			assert.Errorf(t, err, "Could not purge resource")
		}
	})

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(mock.AddLink)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(mock.GetLink.Hash)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(nil)
		assert.Nil(t, err)
		assert.Equal(t, len(links), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		assert.Nil(t, store.Delete(mock.GetLink.Hash))
	})

	t.Run("Close", func(t *testing.T) {
		assert.Nil(t, store.Close())
	})
}
