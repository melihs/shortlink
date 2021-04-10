package http_chi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/domain/link"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// Routes creates a REST router
func (api *API) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", api.Add)
	r.Get("/links", api.List)
	r.Get("/link/{hash}", api.Get)
	r.Delete("/{hash}", api.Delete)

	return r
}

// Add ...
func (api *API) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request AddLinkRequest
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	newLink := &link.Link{
		Url:      request.Url,
		Describe: request.Describe,
	}

	responseCh := make(chan interface{})

	// TODO: send []byte format
	go notify.Publish(r.Context(), api_type.METHOD_ADD, newLink, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_ADD"})

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
	case notify.Response:
		err = r.Error
		if err == nil {
			newLink = r.Payload.(*link.Link) // nolint errcheck
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := json.Marshal(newLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) // nolint errcheck
}

// Get ...
func (api *API) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var hash = chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint errcheck
		return
	}

	// Parse request
	var request = &GetLinkRequest{
		Hash: hash,
	}

	var (
		response     *link.Link
		responseLink GetLinkResponse // for custom JSON parsing
		err          error
	)

	responseCh := make(chan interface{})

	go notify.Publish(r.Context(), api_type.METHOD_GET, request.Hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_GET"})

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
	case notify.Response:
		err = r.Error
		if err == nil {
			response = r.Payload.(*link.Link) // nolint errcheck
		}
	}

	var errorLink *link.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	responseLink = GetLinkResponse{
		Link: response,
	}
	res, err := json.Marshal(responseLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

// List ...
func (api *API) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Get filter
	filter := r.URL.Query().Get("filter")

	var (
		response     []*link.Link
		responseLink GetListLinkResponse // for custom JSON parsing
		err          error
	)

	responseCh := make(chan interface{})

	go notify.Publish(r.Context(), api_type.METHOD_LIST, filter, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_LIST"})

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_LIST")
	case notify.Response:
		err = r.Error
		if err == nil {
			response = r.Payload.([]*link.Link) // nolint errcheck
		}
	}

	var errorLink *link.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	for l := range response {
		responseLink.List = append(responseLink.List, response[l])
	}

	res, err := json.Marshal(responseLink)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

// Delete ...
func (api *API) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	var hash = chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint errcheck
		return
	}

	request := DeleteLinkRequest{
		Hash: hash,
	}

	responseCh := make(chan interface{})

	go notify.Publish(r.Context(), api_type.METHOD_DELETE, request.Hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_DELETE"})

	var err error
	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err = fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
	case notify.Response:
		err = r.Error
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}
