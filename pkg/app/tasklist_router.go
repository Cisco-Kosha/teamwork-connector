package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
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
// @Success 200 {object} models.Tasks
// @Router /api/v1/tasklists/{id}/tasks [get]
func (a *App) getTasks(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	p := httpclient.GetTasks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, p)
}
