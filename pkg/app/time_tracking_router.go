package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
	"github.com/kosha/teamwork-connector/pkg/models"
)

// createProjectTimeEntry godoc
// @Summary Create a time entry for a project
// @Description Create a time entry for a project in the system
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/0a02f6155b979-create-a-time-entry for more parameter options.
// @Tags timeentry
// @Accept json
// @Produce json
// @Param project body models.CreateTimeEntry false "Enter time entry properties"
// @Param id path string false "Enter project id"
// @Success 200
// @Router /api/v1/projects/{project_id}/timeentry [post]
func (a *App) createProjectTimeEntry(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var t models.CreateTimeEntry
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	project_id := vars["project_id"]

	res, err := httpclient.CreateProjectTimeEntry(a.Cfg.GetTeamworkURL(), project_id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), &t, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error creating a project time entry", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}

// getAllTimeEntries godoc
// @Summary Get all time entries
// @Description List all time entries across projects and tasks
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/4ea39d569915b-retrieve-all-time-entries-across-all-projects for more parameter options.
// @Tags timeentry
// @Accept  json
// @Produce  json
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} models.ReturnedTimeEntries
// @Router /api/v1/timeentries [get]
func (a *App) getAllTimeEntries(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var timeEntries []*models.ReturnedTimeEntries

	respHeaders, _ := httpclient.GetAllTimeEntries(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

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
		_, t := httpclient.GetAllTimeEntries(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		timeEntries = append(timeEntries, t)
	}

	respondWithJSON(w, http.StatusOK, timeEntries)
}

// getAllTimeEntriesMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags timeentry
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/timeentries/metadata [get]
func (a *App) getAllTimeEntriesMetadata(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	respHeaders, _ := httpclient.GetAllTimeEntries(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
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

// getProjectTimeEntries godoc
// @Summary Get all time entries associated with specified project
// @Description List all time entries in the project
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/aa65992887407-retrieve-all-time-entries-for-a-project for more parameter options.
// @Tags timeentry
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} models.ReturnedTimeEntries
// @Router /api/v1/projects/{project_id}/timeentry [get]
func (a *App) getProjectTimeEntries(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["project_id"]

	var timeEntries []*models.ReturnedTimeEntries

	respHeaders, _ := httpclient.GetProjectTimeEntries(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

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
		_, t := httpclient.GetProjectTimeEntries(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		timeEntries = append(timeEntries, t)
	}

	respondWithJSON(w, http.StatusOK, timeEntries)
}

// getProjectTimeEntriesMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags timeentry
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/projects/{project_id}/timeentry/metadata [get]
func (a *App) getProjectTimeEntriesMetadata(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	pageCount := 1
	var err error

	respHeaders, _ := httpclient.GetProjectTimeEntries(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
	if respHeaders.Get("X-Pages") != "" {
		pageCount, err = strconv.Atoi(respHeaders.Get("X-Pages"))
		if err != nil {
			a.Log.Errorf("Error getting x-pages header", err)
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	endpointMetadata := models.EndpointMetadata{
		PageCount: pageCount,
	}
	respondWithJSON(w, http.StatusOK, endpointMetadata)
}

// updateTimeEntry godoc
// @Summary Update a time entry
// @Description Add an update to a time entry
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/e1320843e680e-update-a-time-entry for more parameter options.
// @Tags timeentry
// @Accept  json
// @Produce  json
// @Param project body models.CreateTimeEntry false "Enter time entry properties"
// @Param id path string false "Enter time entry id"
// @Success 200 {object} models.CreatedTimeEntry
// @Router /api/v1/timeentries/{id} [put]
func (a *App) updateTimeEntry(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var t models.CreateTimeEntry
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	project_id := vars["id"]

	res, err := httpclient.UpdateTimeEntry(a.Cfg.GetTeamworkURL(), project_id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), &t, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error creating a project time entry", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}

// deleteTimeEntry godoc
// @Summary Delete a time entry
// @Description Remove a time entry
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/c16cf6a573f70-delete-a-time-entry for more parameter options.
// @Tags timeentry
// @Accept  json
// @Produce  json
// @Param id path string false "Enter time entry id"
// @Success 200
// @Router /api/v1/timeentries/{id} [delete]
func (a *App) deleteTimeEntry(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	t, err := httpclient.DeleteTimeEntry(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error deleting a time entry", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, t)
}
