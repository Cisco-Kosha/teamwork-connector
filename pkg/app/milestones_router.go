package app

import (
	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
	"github.com/kosha/teamwork-connector/pkg/models"
	"net/http"
	"strconv"
)

// getAllMilestones godoc
// @Summary Get all milestones.
// @Description Get a list of all milestones for which the current user has access.
// @Tags milestones
// @Accept json
// @Produce json
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} models.ReturnedMilestones
// @Router /api/v1/milestones [get]
func (a *App) getAllMilestones(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var milestones []*models.ReturnedMilestones

	respHeaders, _ := httpclient.GetAllMilestones(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

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
		_, m := httpclient.GetAllMilestones(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		milestones = append(milestones, m)
	}

	respondWithJSON(w, http.StatusOK, milestones)

}

// getAllMilestonesMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags milestones
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/milestones/metadata [get]
func (a *App) getAllMilestonesMetadata(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	respHeaders, _ := httpclient.GetAllMilestones(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
	var pageCount int
	var err error

	if respHeaders.Get("X-Pages") != "" {
		pageCount, err = strconv.Atoi(respHeaders.Get("X-Pages"))
	} else {
		pageCount = 1
	}
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

// getProjectMilestones godoc
// @Summary Get all milestones for a specified project.
// @Description Get a list of all milestones for a specified project for which the current user has access.
// @Tags milestones
// @Accept json
// @Produce json
// @Param id path string false "Enter project id"
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} models.ReturnedMilestones
// @Router /api/v1/projects/{id}/milestones [get]
func (a *App) getProjectMilestones(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var milestones []*models.ReturnedMilestones

	vars := mux.Vars(r)
	id := vars["id"]

	respHeaders, _ := httpclient.GetProjectMilestones(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

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
		_, m := httpclient.GetProjectMilestones(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		milestones = append(milestones, m)
	}

	respondWithJSON(w, http.StatusOK, milestones)

}

// getProjectMilestonesMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags milestones
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Success 200
// @Router /api/v1/projects/{id}/milestones [get]
func (a *App) getProjectMilestonesMetadata(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	respHeaders, _ := httpclient.GetProjectMilestones(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
	var pageCount int
	var err error

	if respHeaders.Get("X-Pages") != "" {
		pageCount, err = strconv.Atoi(respHeaders.Get("X-Pages"))
	} else {
		pageCount = 1
	}
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

// getSingleMilestone godoc
// @Summary Get the specified milestone.
// @Description Get milestone information for the id specified.
// @Tags milestones
// @Accept json
// @Produce json
// @Param id path string false "Enter milestone id"
// @Success 200 {object} models.ReturnedMilestone
// @Router /api/v1/milestones/{id} [get]
func (a *App) getSingleMilestone(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var milestones []*models.ReturnedMilestone

	vars := mux.Vars(r)
	id := vars["id"]

	respHeaders, _ := httpclient.GetSingleMilestone(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

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
		_, m := httpclient.GetSingleMilestone(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		milestones = append(milestones, m)
	}

	respondWithJSON(w, http.StatusOK, milestones)

}