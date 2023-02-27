package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	var apiV1 = "/api/v1"

	// specification routes
	a.Router.HandleFunc(apiV1+"/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/specification/test", a.testConnectorSpecification).Methods("POST", "OPTIONS")

	// tasklists
	a.Router.HandleFunc(apiV1+"/projects/{id}/tasklists/metadata", a.getProjectTasklistsMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}/tasklists", a.getProjectTasklists).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}/tasklists", a.createTaskList).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/tasklists/{id}/tasks/metadata", a.getTasksMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/tasklists/{id}/tasks", a.getTasks).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/tasklists/{id}", a.deleteTaskList).Methods("DELETE", "OPTIONS")

	//people
	a.Router.HandleFunc(apiV1+"/me", a.getCurrentPerson).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/people/metadata", a.getPeopleMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/people/{id}/projects", a.getPersonsProjects).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/people/{id}", a.getPerson).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/people", a.getPeople).Methods("GET", "OPTIONS")

	// time tracking
	a.Router.HandleFunc(apiV1+"/timeentries/{id}", a.updateTimeEntry).Methods("PUT", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/timeentries/{id}", a.deleteTimeEntry).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/timeentries/metadata", a.getAllTimeEntriesMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/timeentries", a.getAllTimeEntries).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{project_id}/timeentry/metadata", a.getProjectTimeEntriesMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{project_id}/timeentry", a.getProjectTimeEntries).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{project_id}/timeentry", a.createProjectTimeEntry).Methods("POST", "OPTIONS")

	// risks
	a.Router.HandleFunc(apiV1+"/projects/{id}/risks/metadata", a.getProjectRisksMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}/risks", a.createProjectRisk).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}/risks", a.getProjectRisksV1).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/risks/metadata", a.getAllRisksMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/risks/{id}", a.updateRisks).Methods("PUT", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/risks/{id}", a.deleteRisks).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/risks", a.getAllRisksV1).Methods("GET", "OPTIONS")

	//milestones
	a.Router.HandleFunc(apiV1+"/milestones/metadata", a.getAllMilestonesMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/milestones/{id}", a.getSingleMilestone).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/milestones", a.getAllMilestones).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}/milestones/metadata", a.getProjectMilestonesMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}/milestones", a.getProjectMilestones).Methods("GET", "OPTIONS")

	// projects
	a.Router.HandleFunc(apiV1+"/projects/metadata", a.getAllProjectsMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/activity", a.getLatestActivityAllProjectsV1).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/tasks/metadata", a.getTasksAllProjectsMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/tasks", a.getTasksAllProjects).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/updates", a.getAllProjectUpdatesV3).Methods("GET", "OPTIONS")

	//TODO: Project Updates is a paid feature on Teamwork so we can't see pagination headers for this endpoint
	//a.Router.HandleFunc(apiV1+"/projects/updates/metadata", a.getAllProjectUpdatesMetadata).Methods("GET", "OPTIONS")
	//a.Router.HandleFunc(apiV1+"/projects/{id}/updates/metadata", a.getSingleProjectUpdatesMetadata).Methods("GET", "OPTIONS")

	a.Router.HandleFunc(apiV1+"/projects/{id}/tasks/metadata", a.getProjectTasksMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}/tasks", a.getProjectTasks).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}/updates", a.getSingleProjectUpdatesV1).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}/update", a.createProjectUpdate).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/update/{id}", a.modifyProjectUpdate).Methods("PUT", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/update/{id}", a.deleteProjectUpdate).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}", a.getSingleProjectV1).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects", a.getAllProjectsV1).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}", a.deleteProject).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects", a.createProject).Methods("POST", "OPTIONS")

	// tags
	// a.Router.HandleFunc(apiV1+"/tags/{tag_id}}", a.getSingleTag).Methods("GET", "OPTIONS")
	// a.Router.HandleFunc(apiV1+"/tags", a.getAllTags).Methods("GET", "OPTIONS")
	// a.Router.HandleFunc(apiV1+"/tags", a.createSingleTag).Methods("CREATE", "OPTIONS")
	// a.Router.HandleFunc(apiV1+"/tags/{tag_id}}", a.updateSingleTag).Methods("PATCH", "OPTIONS")
	// a.Router.HandleFunc(apiV1+"/tags/{tag_id}}", a.deleteSingleTag).Methods("DELETE", "OPTIONS")

	// stats
	a.Router.HandleFunc(apiV1+"/stats/people/{id}", a.getStatsForPerson).Methods("GET", "OPTIONS")

	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
