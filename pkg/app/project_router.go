package app

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
	"github.com/kosha/teamwork-connector/pkg/models"
)

func (a *App) getPageRange(params url.Values, respHeaders http.Header) (int, int, error) {
	var err error
	pageStart := 1
	pageEnd := 1
	numPages := 1

	if val, ok := params["pageStart"]; ok {
		pageStart, err = strconv.Atoi(val[0])
		if err != nil {
			return 0, 0, err
		}
	}
	if val, ok := params["pageEnd"]; ok {
		pageEnd, err = strconv.Atoi(val[0])
		if err != nil {
			return 0, 0, err
		}
	}

	numPages, err = strconv.Atoi(respHeaders.Get("X-Pages"))
	if err != nil {
		return 0, 0, err
	}

	if pageStart > numPages || pageEnd > numPages || pageStart < 1 {
		return 0, 0, err
	}

	if val, ok := params["allPages"]; ok && val[0] == "True" {
		pageStart = 1
		pageEnd = numPages
	}

	return pageStart, pageEnd, nil
}

// getAllProjects godoc
// @Summary Get all projects
// @Description List all projects
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/df7d06840ecdd-returns-a-list-of-projects for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Param page query string false "Page number"
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
// @Param page query string false "Page number"
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
// @Param page query string false "Page number"
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

// getProjectTasks godoc
// @Summary Retrieves all tasks in a project
// @Description Lists all tasks based on project ID
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/6e3da2c04d779-get-all-tasks-on-a-given-project
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Param page query string false "Page number"
// @Success 200 {object} models.Tasks
// @Router /api/v1/projects/{id}/tasks [get]
func (a *App) getProjectTasks(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	var tasks []*models.Tasks

	respHeaders, _ := httpclient.GetProjectTasks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

	//get page range data from headers
	pageStart, pageEnd, err := a.getPageRange(r.URL.Query(), respHeaders)
	if err != nil {
		a.Log.Errorf("Invalid pageStart or pageEnd header", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//get page data
	params := r.URL.Query()
	for i := pageStart; i <= pageEnd; i++ {
		params["page"] = append(r.URL.Query()["page"], strconv.Itoa(i))
		_, t := httpclient.GetProjectTasks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		tasks = append(tasks, t)
	}

	respondWithJSON(w, http.StatusOK, tasks)
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

// createProjectUpdate godoc
// @Summary Add an update for a project
// @Description Update a project in the system
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/3a875e7157506-create-a-project-update for more parameter options.
// @Tags projects
// @Accept json
// @Produce json
// @Param project body models.ProjectUpdate false "Enter project update properties"
// @Param id path string false "Enter project id"
// @Success 200
// @Router /api/v1/projects/{id}/update [post]
func (a *App) createProjectUpdate(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var p models.ProjectUpdate
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	res, err := httpclient.CreateProjectUpdate(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), &p, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error creating a project update", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}

// getAllProjectUpdates godoc
// @Summary Get all project updates
// @Description List all updates across projects that the logged-in user can access.
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/2e4f8bf140cab-get-all-project-updates for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Param page query string false "Page number"
// @Success 200 {object} models.ReturnedRisks
// @Router /api/v1/projects/updates [get]
func (a *App) getAllProjectUpdates(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	p := httpclient.GetAllProjectUpdates(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, p)
}

// modifyProjectUpdate godoc
// @Summary Modify a project update
// @Description Change an update made to a project
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/aa0bc9bb0fd37-modify-a-project-update for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Param project body models.ProjectUpdate false "Enter project update properties"
// @Param id path string false "Enter project update id"
// @Success 200
// @Router /api/v1/projects/update/{id} [put]
func (a *App) modifyProjectUpdate(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var t models.ProjectUpdate
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	res, err := httpclient.ModifyProjectUpdate(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), &t, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error modifying a project update", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}

// deleteProjectUpdate godoc
// @Summary Delete a project update
// @Description Remove an update made to a project
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/cfdb775e8ade2-delete-a-project-update for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project update id"
// @Success 200
// @Router /api/v1/projects/update/{id} [delete]
func (a *App) deleteProjectUpdate(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	t, err := httpclient.DeleteProjectUpdate(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error deleting a project update", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, t)
}
