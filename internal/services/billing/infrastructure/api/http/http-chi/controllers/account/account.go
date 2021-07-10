package account

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/batazor/shortlink/internal/services/api/application/http-chi/helpers"
	account_application "github.com/batazor/shortlink/internal/services/billing/application/account"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/v1"
)

type AccoutAPI struct {
	jsonpb protojson.MarshalOptions

	accountService account_application.AccountService
}

func New(accountService *account_application.AccountService) (*AccoutAPI, error) {
	return &AccoutAPI{
		accountService: *accountService,
	}, nil
}

// Routes creates a REST router
func (api *AccoutAPI) Routes(r chi.Router) {
	r.Get("/accounts", api.list)
	r.Get("/account/{hash}", api.get)
	r.Post("/account", api.add)
	r.Delete("/account/{hash}", api.delete)
}

// add ...
func (a *AccoutAPI) add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request billing.Account
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	// inject spanId in response header
	w.Header().Add("span-id", helpers.RegisterSpan(r.Context()))

	newAccount, err := a.accountService.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	res, err := a.jsonpb.Marshal(newAccount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint errcheck
}

// get ...
func (a *AccoutAPI) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// list ...
func (a *AccoutAPI) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}

// delete ...
func (a *AccoutAPI) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{}`)) // nolint errcheck
}
