package app

import (
	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
	"github.com/kosha/teamwork-connector/pkg/models"
	"net/http"
)

// getStatsForPerson godoc
// @Summary Get statistics for a particular person
// @Description Get statistics for a particular person id
// @Tags people
// @Accept json
// @Produce json
// @Param id path string false "Enter person id"
// @Success 200 {object} models.OverallPerson
// @Router /api/v1/stats/people/{id} [get]
func (a *App) getStatsForPerson(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	person := httpclient.GetSinglePerson(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())
	permissions := person.Person.Permissions
	projects := httpclient.GetPersonsProjects(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

	overallPerson := models.OverallPerson{
		Person:      person.Person,
		Projects:    projects.Projects,
		Permissions: permissions,
	}

	respondWithJSON(w, http.StatusOK, overallPerson)
	return
}
