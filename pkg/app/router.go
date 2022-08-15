package app

import (
	"encoding/json"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
	"github.com/kosha/teamwork-connector/pkg/models"
	"net/http"
)

// listConnectorSpecification godoc
// @Summary Get connector specification details
// @Description Retrieve necessary environment variables
// @Tags specification
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/v1/specification/list [get]
func (a *App) listConnectorSpecification(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	respondWithJSON(w, http.StatusOK, map[string]string{
		"USERNAME":    "Teamwork Username",
		"PASSWORD":    "Teamwork Password",
		"DOMAIN_NAME": "Teamwork Domain Name",
	})
}

// testConnectorSpecification godoc
// @Summary Test auth against the specification
// @Description Check if domain account can be verified
// @Tags specification
// @Accept  json
// @Produce  json
// @Param text body models.Specification false "Enter auth and domain name properties"
// @Success 200 {object} models.SingleAccount
// @Router /api/v1/specification/test [post]
func (a *App) testConnectorSpecification(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	if (*r).Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	var s models.Specification
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	url := "https://" + s.DomainName + ".teamwork.com/"
	account := httpclient.GetAccount(url, s.Username, s.Password, r.URL.Query())
	if account != nil {
		respondWithJSON(w, http.StatusOK, account)
	} else {
		respondWithError(w, http.StatusBadRequest, "Account not verified")
	}
}