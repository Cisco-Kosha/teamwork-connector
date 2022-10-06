package app

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
	"github.com/kosha/teamwork-connector/pkg/models"
)

// deleteTaskList godoc
// @Summary Delete single tasklist
// @Description Delete single tasklist based on the tasklist id logged in the system.
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/498ee4e9c4fa5-delete-a-task-list for more parameter options.
// @Tags tasklists
// @Accept  json
// @Produce  json
// @Param id path string true "Enter tasklist id"
// @Success 200
// @Router /api/v1/tasklists/{id} [delete]
func (a *App) deleteTaskList(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	t, err := httpclient.RemoveTaskList(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error deleting a tasklist", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success", "resp": t})
}

// getTasks godoc
// @Summary Retrieve a tasklist's tasks
// @Description Lists tasks based on tasklist Id
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/a9b6130385c3e-get-all-tasks-on-a-given-task-list for more parameter options.
// @Tags tasklists
// @Accept  json
// @Produce  json
// @Param id path string false "Enter tasklist id"
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} models.Tasks
// @Router /api/v1/tasklists/{id}/tasks [get]
func (a *App) getTasks(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	var tasks []*models.Tasks

	respHeaders, _ := httpclient.GetTasks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

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
		_, t := httpclient.GetTasks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		tasks = append(tasks, t)
	}

	respondWithJSON(w, http.StatusOK, tasks)
}

// getTasksMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags tasklists
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/tasklists/{id}/tasks/metadata [get]
func (a *App) getTasksMetadata(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	respHeaders, _ := httpclient.GetTasks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
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
