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
	a.Router.HandleFunc(apiV1+"/projects/{id}/tasklists", a.getProjectTasklists).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}/tasklists", a.createTaskList).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/tasklists/{id}/tasks", a.getTasks).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/tasklists/{id}", a.deleteTaskList).Methods("DELETE", "OPTIONS")

	//people
	a.Router.HandleFunc(apiV1+"/me", a.getCurrentPerson).Methods("GET", "OPTIONS")
	// a.Router.HandleFunc(apiV1+"/people/{id}/tasks", a.getPersonsTasks).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/people/{id}/projects", a.getPersonsProjects).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/people/{id}", a.getPerson).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/people", a.getPeople).Methods("GET", "OPTIONS")

	// time tracking
	a.Router.HandleFunc(apiV1+"/timeentries/{id}", a.updateTimeEntry).Methods("PUT", "OPTIONS")	
	a.Router.HandleFunc(apiV1+"/timeentries/{id}", a.deleteTimeEntry).Methods("DELETE", "OPTIONS")		
	a.Router.HandleFunc(apiV1+"/timeentries", a.getAllTimeEntries).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{project_id}/timeentry", a.getProjectTimeEntries).Methods("GET", "OPTIONS")	
	a.Router.HandleFunc(apiV1+"/projects/{project_id}/timeentry", a.createProjectTimeEntry).Methods("POST", "OPTIONS")	


	// risks
	a.Router.HandleFunc(apiV1+"/projects/{id}/risks", a.createProjectRisk).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}/risks", a.getProjectRisks).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/risks/{id}", a.updateRisks).Methods("PUT", "OPTIONS")	
	a.Router.HandleFunc(apiV1+"/risks/{id}", a.deleteRisks).Methods("DELETE", "OPTIONS")	
	a.Router.HandleFunc(apiV1+"/risks", a.getAllRisks).Methods("GET", "OPTIONS")		


	// projects
	a.Router.HandleFunc(apiV1+"/projects/activity", a.getLatestActivityAllProjects).Methods("GET", "OPTIONS")
	// a.Router.HandleFunc(apiV1+"/projects/updates", a.getAllProjectUpdates).Methods("GET", "OPTIONS")		
	a.Router.HandleFunc(apiV1+"/projects/{id}/update", a.createProjectUpdate).Methods("POST", "OPTIONS")	
	// a.Router.HandleFunc(apiV1+"/projects/update", a.modifyProjectUpdate).Methods("PUT", "OPTIONS")	
	// a.Router.HandleFunc(apiV1+"/projects/update", a.deleteProjectUpdate).Methods("DELETE", "OPTIONS")	
	a.Router.HandleFunc(apiV1+"/projects/{id}", a.getSingleProject).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects", a.getAllProjects).Methods("GET", "OPTIONS")
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
