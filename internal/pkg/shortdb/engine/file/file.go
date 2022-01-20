package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/shortdb/engine/options"
	"github.com/batazor/shortlink/internal/pkg/shortdb/io_uring"
	v1 "github.com/batazor/shortlink/internal/pkg/shortdb/query/v1"
	table "github.com/batazor/shortlink/internal/pkg/shortdb/table/v1"
)

type file struct {
	mc sync.Mutex

	name string
	path string

	database *table.DataBase
}

func New(opts ...options.Option) (*file, error) {
	viper.AutomaticEnv()
	viper.SetDefault("SHORTDB_DEFAULT_DATABASE", "public.db") // ShortDB default database
	viper.SetDefault("SHORTDB_PAGE_SIZE", 2)                  // ShortDB default page of size

	var err error
	f := &file{
		name: viper.GetString("SHORTDB_DEFAULT_DATABASE"),
		database: &table.DataBase{
			Tables: make(map[string]*table.Table),
		},
	}

	for _, opt := range opts {
		if errApplyOptions := opt(f); errApplyOptions != nil {
			panic(errApplyOptions)
		}
	}

	// if not set path, set temp directory
	if f.path == "" {
		f.path, err = ioutil.TempDir(os.TempDir(), "shortdb_")
		if err != nil {
			return nil, err
		}
	}

	// init db
	err = f.init()
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *file) Exec(query *v1.Query) (interface{}, error) {
	switch query.Type {
	case v1.Type_TYPE_UNSPECIFIED:
		return nil, fmt.Errorf("exec: incorret type")
	case v1.Type_TYPE_SELECT:
		return f.Select(query)
	case v1.Type_TYPE_UPDATE:
		return nil, f.Update(query)
	case v1.Type_TYPE_INSERT:
		return nil, f.Insert(query)
	case v1.Type_TYPE_DELETE:
		return nil, f.Delete(query)
	case v1.Type_TYPE_CREATE_TABLE:
		return nil, f.CreateTable(query)
	case v1.Type_TYPE_DROP_TABLE:
		return nil, f.DropTable(query.TableName)
	}

	return nil, nil
}

func (f *file) init() error {
	f.mc.Lock()
	defer f.mc.Unlock()

	path := fmt.Sprintf("%s/%s", f.path, f.name)

	// create directory if not exist
	err := os.MkdirAll(f.path, os.ModePerm)
	if err != nil {
		return err
	}

	// create file if not exist
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	// init io_uring
	err = io_uring.Init()
	if err != nil {
		return err
	}
	defer io_uring.Cleanup()

	go func() {
		for errIOUring := range io_uring.Err() {
			fmt.Println(errIOUring)
		}
	}()

	payload := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	// Read a file.
	err = io_uring.ReadFile(path, func(buf []byte) {
		payload = buf
		wg.Done()
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	io_uring.Poll()
	wg.Wait()

	if len(payload) != 0 {
		err = proto.Unmarshal(payload, f.database)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *file) Close() error {
	f.mc.Lock()
	defer f.mc.Unlock()

	path := fmt.Sprintf("%s/%s", f.path, f.name)

	// create file if not exist
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	payload, err := proto.Marshal(f.database)
	if err != nil {
		return err
	}

	// init io_uring
	err = io_uring.Init()
	if err != nil {
		return err
	}
	defer io_uring.Cleanup()

	go func() {
		for err := range io_uring.Err() {
			fmt.Println(err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	// Write something
	err = io_uring.WriteFile(path, payload, 0644, func(n int) {
		defer wg.Done()
		// handle n
	})

	// Call Poll to let the kernel know to read the entries.
	io_uring.Poll()
	// Wait till all callbacks are done.
	wg.Wait()

	return nil
}
