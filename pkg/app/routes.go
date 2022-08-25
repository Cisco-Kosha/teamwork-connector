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

	// projects routes
	a.Router.HandleFunc(apiV1+"/projects/activity", a.getLatestActivityAllProjects).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects", a.getAllProjects).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}", a.getSingleProject).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects/{id}", a.deleteProject).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/projects", a.createProject).Methods("POST", "OPTIONS")	

	// stats
	a.Router.HandleFunc(apiV1+"/stats/people/{id}", a.getStatsForPerson).Methods("GET", "OPTIONS")

	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
