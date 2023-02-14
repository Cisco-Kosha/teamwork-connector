package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
	"github.com/kosha/teamwork-connector/pkg/models"
)

func (a *App) getAllProjectsV3(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var projects []*models.ProjectResponseV3
	var pageStart, pageEnd int
	var err error

	_, data := httpclient.GetAllProjectsV3(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), false)
	//endpoint called
	//error code from teamwork's side
	pageStart, pageEnd, err = getPageRange(r.URL.Query(), nil, data.Meta.Page.Count)

	if err != nil {
		a.Log.Errorf("Error getting page range", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//get page data
	params := r.URL.Query()
	for i := pageStart; i <= pageEnd; i++ {
		params["page"] = append(r.URL.Query()["page"], strconv.Itoa(i))
		_, p := httpclient.GetAllProjectsV3(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		projects = append(projects, p)
	}

	respondWithJSON(w, http.StatusOK, projects)
}

// getAllProjects godoc
// @Summary Get all projects
// @Description V1 Teamwork API: list all projects
// @Description Please refer to hhttps://apidocs.teamwork.com/docs/teamwork/626f30d917e1c-retrieve-all-projects for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Param includeProjectOwner query boolean false "Adds project owner details to response"
// @Param includeTags query boolean false "Adds tag details to response"
// @Success 200 {object} models.ProjectResponseV1
// @Router /api/v1/projects [get]
func (a *App) getAllProjectsV1(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var projects []*models.ProjectResponseV1

	respHeaders, _ := httpclient.GetAllProjectsV1(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

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
		_, p := httpclient.GetAllProjectsV1(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		projects = append(projects, p)
	}

	respondWithJSON(w, http.StatusOK, projects)
}

// getAllProjectsMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags projects
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/projects/metadata [get]
func (a *App) getAllProjectsMetadata(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	respHeaders, _ := httpclient.GetAllProjectsV1(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
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

// getSingleProject godoc
// @Summary Get single project
// @Description List single project based on project ID
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/da79573311f8a-retrieve-a-single-project for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Success 200 {object} models.SingleProjectResponseV1
// @Router /api/v1/projects/{id} [get]
func (a *App) getSingleProjectV1(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	p := httpclient.GetSingleProjectV1(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

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
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/e8d592debf406-latest-activity-across-all-projects for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Param page query string false "Page number"
// @Success 200 {object} models.MultiActivity
// @Router /api/v1/projects/activity [get]
func (a *App) getLatestActivityAllProjectsV1(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	activity := httpclient.GetLatestActivityAllProjectsV1(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, activity)
}

func (a *App) getLatestActivityAllProjectsV3(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	activity := httpclient.GetLatestActivityAllProjectsV3(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	respondWithJSON(w, http.StatusOK, activity)
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
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} models.MultiTaskList
// @Router /api/v1/projects/{id}/tasklists [get]
func (a *App) getProjectTasklists(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	var tasklists []*models.MultiTaskList

	respHeaders, _ := httpclient.GetProjectTaskLists(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

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
		_, t := httpclient.GetProjectTaskLists(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		tasklists = append(tasklists, t)
	}

	respondWithJSON(w, http.StatusOK, tasklists)
}

// getProjectTasklistsMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags tasklists
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Success 200
// @Router /api/v1/projects/{id}/tasklists/metadata [get]
func (a *App) getProjectTasklistsMetadata(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	respHeaders, _ := httpclient.GetProjectTaskLists(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
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

// getTasksAllProjects godoc
// @Summary Retrieves all tasks across all projects
// @Description Lists all tasks across all projects
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/54bdf625aa2f3-get-all-tasks-across-all-projects
// @Tags projects
// @Accept  json
// @Produce  json
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} models.Tasks
// @Router /api/v1/projects/tasks [get]
func (a *App) getTasksAllProjects(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var tasks []*models.Tasks

	respHeaders, _ := httpclient.GetTasksAllProjects(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
	//fmt.Println(respHeaders)

	//get page range data from headers
	pageStart, pageEnd, err := getPageRange(r.URL.Query(), respHeaders, 0)
	if err != nil {
		a.Log.Errorf("Invalid pageStart or pageEnd header", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//get page data
	params := r.URL.Query()
	for i := pageStart; i <= pageEnd; i++ {
		params["page"] = append(r.URL.Query()["page"], strconv.Itoa(i))
		_, t := httpclient.GetTasksAllProjects(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		tasks = append(tasks, t)
	}

	respondWithJSON(w, http.StatusOK, tasks)
}

// getTasksAllProjectsMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags projects
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/projects/tasks/metadata [get]
func (a *App) getTasksAllProjectsMetadata(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	respHeaders, _ := httpclient.GetTasksAllProjects(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
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

// getProjectTasks godoc
// @Summary Retrieves all tasks for specified project
// @Description Lists all tasks for specified project
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/6e3da2c04d779-get-all-tasks-on-a-given-project
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
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
	//fmt.Println(respHeaders)

	//get page range data from headers
	pageStart, pageEnd, err := getPageRange(r.URL.Query(), respHeaders, 0)
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

// getProjectTasksMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Success 200
// @Router /api/v1/projects/{id}/tasks/metadata [get]
func (a *App) getProjectTasksMetadata(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	respHeaders, _ := httpclient.GetProjectTasks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
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

func (a *App) getAllProjectUpdatesV3(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var projectUpdates []*models.ProjectUpdateResponse
	var pageStart, pageEnd int
	var err error

	_, data := httpclient.GetAllProjectUpdatesV3(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), false)
	pageStart, pageEnd, err = getPageRange(r.URL.Query(), nil, data.Metadata.Page.Count)

	if err != nil {
		a.Log.Errorf("Error getting page range", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//get page data
	params := r.URL.Query()
	for i := pageStart; i <= pageEnd; i++ {
		params["page"] = append(r.URL.Query()["page"], strconv.Itoa(i))
		_, p := httpclient.GetAllProjectUpdatesV3(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)

		projectUpdates = append(projectUpdates, p)
	}

	respondWithJSON(w, http.StatusOK, projectUpdates)
}

// getProjectUpdates godoc
// @Summary Get all  updates for a specific project
// @Description List all updates for a specific project
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/277672affb50e-get-project-updates for more parameter options.
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} models.ReturnedRisks
// @Router /api/v1/projects/updates [get]
func (a *App) getAllProjectUpdatesV1(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var projectUpdates []*models.ProjectUpdateResponseV1
	var pageStart, pageEnd int
	// var err error

	vars := mux.Vars(r)
	id := vars["id"]

	respHeaders, _ := httpclient.GetAllProjectUpdatesV1(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

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
		_, p := httpclient.GetAllProjectUpdatesV1(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)

		projectUpdates = append(projectUpdates, p)
	}

	respondWithJSON(w, http.StatusOK, projectUpdates)
}

// getAllProjectUpdatesMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags projects
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Success 200
// @Router /api/v1/projects/updates/metadata [get]
func (a *App) getAllProjectUpdatesMetadata(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	respHeaders, _ := httpclient.GetAllProjectUpdatesV1(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)
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
