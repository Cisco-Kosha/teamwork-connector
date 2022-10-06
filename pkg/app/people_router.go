package app

import (
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
	"github.com/kosha/teamwork-connector/pkg/models"

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
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} models.People
// @Router /api/v1/people [get]
func (a *App) getPeople(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var people []*models.People

	respHeaders, _ := httpclient.GetPeople(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

	//get page range data from headers
	pageStart, pageEnd, err := getPageRange(r.URL.Query(), respHeaders, 0)
	if err != nil {
		a.Log.Errorf("Error getting page range", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//get page data
	params := r.URL.Query()
	for i := pageStart; i <= pageEnd; i++ {
		params["page"] = append(r.URL.Query()["page"], strconv.Itoa(i))
		_, p := httpclient.GetPeople(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		people = append(people, p)
	}

	respondWithJSON(w, http.StatusOK, people)

}

// getPeopleMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags people
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/people/metadata [get]
func (a *App) getPeopleMetadata(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	respHeaders, _ := httpclient.GetPeople(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
	pageCount, err := strconv.Atoi(respHeaders.Get("X-Pages"))
	if err != nil {
		a.Log.Errorf("Error getting x-pages header", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	endpointMetadata := models.EndpointMetadata{
		PageCount: pageCount,
	}
	respondWithJSON(w, http.StatusOK, endpointMetadata)
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
