package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
	"github.com/kosha/teamwork-connector/pkg/models"
	"net/http"
)

// getAllProjects godoc
// @Summary Get all projects
// @Description List all projects
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/df7d06840ecdd-returns-a-list-of-projects for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Success 200 {object} models.MultiProject
// @Router /api/v1/projects [get]
func (a *App) getAllProjects(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	p := httpclient.GetAllProjects(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, p)
}

// getSingleProject godoc
// @Summary Get single project
// @Description List single project based on project ID
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/9daa306fff1d2-returns-a-project for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Success 200 {object} models.SingleProject
// @Router /api/v1/projects/{id} [get]
func (a *App) getSingleProject(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	p := httpclient.GetSingleProject(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, p)
}

// createProject godoc
// @Summary Create new project
// @Description Create single project in the system
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/a0c29f3b330bd-create-project for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Param project body models.SingleProject false "Enter project properties"
// @Success 200
// @Router /api/v1/projects [post]
func (a *App) createProject(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var p models.SingleProject
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	res, err := httpclient.CreateProject(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), &p, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error creating a project", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}

// deleteProject godoc
// @Summary Delete single project
// @Description Delete single project based on the project id logged in the system
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/df61dac913b37-delete-project for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string true "Enter project id"
// @Success 200
// @Router /api/v1/projects/{id} [delete]
func (a *App) deleteProject(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	t, err := httpclient.RemoveProject(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error deleting a project", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success", "resp": t})
}

// getLatestActivityAllProjects godoc
// @Summary List latest activity across all projects
// @Description Lists the latest activity across all projects ordered chronologically.
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/f207f625bd76e-latest-activity-all-projects for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Success 200 {object} models.MultiActivity
// @Router /api/v1/projects/activity [get]
func (a *App) getLatestActivityAllProjects(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	p := httpclient.GetLatestActivityAllProjects(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, p)
}

// getProjectTasklists godoc
// @Summary Retrieve a project's task list
// @Description Lists task lists based on project ID
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/7ee4827082194-get-all-task-lists-for-a-project for more parameter options.
// @Tags tasklists
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Success 200 {object} models.MultiTaskList
// @Router /api/v1/projects/{id}/tasklists [get]
func (a *App) getProjectTasklists(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	p := httpclient.GetProjectTaskLists(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, p)
}

// createTaskList godoc
// @Summary Create new tasklist
// @Description Create single tasklist in the system
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/36c110772a850-create-task-list for more parameter options.
// @Tags tasklists
// @Accept  json
// @Produce  json
// @Param tasklist body models.NewTaskList false "Enter tasklist properties"
// @Success 200
// @Router /api/v1/projects/{id}/tasklists [post]
func (a *App) createTaskList(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var t models.NewTaskList
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	res, err := httpclient.CreateTaskList(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), &t, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error creating a tasklist", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}
