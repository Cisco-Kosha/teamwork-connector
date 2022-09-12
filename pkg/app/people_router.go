package app

import (
	// "encoding/json"
	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
	// "github.com/kosha/teamwork-connector/pkg/models"
	// "fmt"
	"net/http"
)

// getCurrentPerson godoc
// @Summary Get current person's details
// @Description Get current person's details
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/de01076fde3a5-get-current-user-details for more parameter options.
// @Tags people
// @Accept json
// @Produce json
// @Success 200 {object} models.SinglePerson
// @Router /api/v1/me [get]
func (a *App) getCurrentPerson(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	p := httpclient.GetCurrentPerson(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, p)
}

// getPeople godoc
// @Summary Get all people from current user's company.
// @Description Get a list of all people in the current user's company.
// @Tags people
// @Accept json
// @Produce json
// @Param page query string false "Page number"
// @Success 200 {object} models.People
// @Router /api/v1/people [get]
func (a *App) getPeople(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	p := httpclient.GetPeople(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, p)
}

// getPerson godoc
// @Summary Get specific person
// @Description Get specific persons' details
// @Tags people
// @Accept json
// @Produce json
// @Param id path string false "Enter person id"
// @Success 200 {object} models.SinglePerson
// @Router /api/v1/people/{id} [get]
func (a *App) getPerson(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	p := httpclient.GetSinglePerson(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, p)
}

// getPersonsProjects godoc
// @Summary List projects for specified person
// @Description List projects for specified person
// @Tags people
// @Accept json
// @Produce json
// @Param id path string false "Enter person id"
// @Success 200 {object} models.MultiProject
// @Router /api/v1/people/{id}/projects [get]
func (a *App) getPersonsProjects(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	p := httpclient.GetPersonsProjects(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, p)
}
