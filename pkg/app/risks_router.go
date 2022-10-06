package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/httpclient"
	"github.com/kosha/teamwork-connector/pkg/models"
)

// createProjectRisk godoc
// @Summary Create a risk for a project
// @Description Create a risk for a project in the system
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/f6dfef8ed1b7f-create-a-risk-on-a-project for more parameter options.
// @Tags risks
// @Accept json
// @Produce json
// @Param project body models.CreateRisk false "Enter project risk properties"
// @Param id path string false "Enter project id"
// @Success 200
// @Router /api/v1/projects/{id}/risks [post]
func (a *App) createProjectRisk(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var t models.CreateRisk
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	project_id := vars["id"]

	res, err := httpclient.CreateProjectRisk(a.Cfg.GetTeamworkURL(), project_id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), &t, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error creating a project risk", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}

// getAllRisks godoc
// @Summary Get all risks
// @Description List all risks across projects
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/64d22ab985a58-get-all-risks for more parameter options.
// @Tags risks
// @Accept  json
// @Produce  json
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} models.ReturnedRisks
// @Router /api/v1/risks [get]
func (a *App) getAllRisks(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var risks []*models.ReturnedRisks
	var pageStart, pageEnd int
	var err error

	_, data := httpclient.GetAllRisks(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), false)
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
		_, r := httpclient.GetAllRisks(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		risks = append(risks, r)
	}

	respondWithJSON(w, http.StatusOK, risks)
}

// getAllRisksMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags risks
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/risks/metadata [get]
func (a *App) getAllRisksMetadata(w http.ResponseWriter, r *http.Request) {

	_, data := httpclient.GetAllRisks(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), false)
	endpointMetadata := models.EndpointMetadata{
		PageCount: data.Metadata.Page.Count,
	}
	respondWithJSON(w, http.StatusOK, endpointMetadata)
}

// getProjectRisks godoc
// @Summary Get all risks associated with specified project
// @Description List all risks in the project
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/3f3555d7cd5d0-get-risks-for-a-specific-project for more parameter options.
// @Tags risks
// @Accept  json
// @Produce  json
// @Param id path string false "Enter project id"
// @Param page query string false "Page number"
// @Param allPages query boolean false "Collates all pages"
// @Param pageStart query integer false "First page to collate"
// @Param pageEnd query integer false "Last page to collate"
// @Success 200 {object} models.ReturnedRisks
// @Router /api/v1/projects/{id}/risks [get]
func (a *App) getProjectRisks(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	var risks []*models.ReturnedRisks

	respHeaders, _ := httpclient.GetProjectRisks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), true)

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
		_, r := httpclient.GetProjectRisks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), params, false)
		risks = append(risks, r)
	}

	respondWithJSON(w, http.StatusOK, risks)

}

// getProjectRisksMetadata godoc
// @Summary Get number of pages and page length data
// @Description Get page metadata for endpoint
// @Tags risks
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v1/projects/{id}/risks/metadata [get]
func (a *App) getProjectRisksMetadata(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, data := httpclient.GetProjectRisks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query(), false)
	endpointMetadata := models.EndpointMetadata{
		PageCount: data.Metadata.Page.Count,
	}
	respondWithJSON(w, http.StatusOK, endpointMetadata)
}

// updateRisks godoc
// @Summary Update a project risk
// @Description Add an update to a project's risk
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/a1d50e01411b8-updating-a-risk-on-a-project for more parameter options.
// @Tags risks
// @Accept  json
// @Produce  json
// @Param project body models.CreateRisk false "Enter project risk properties"
// @Param id path string false "Enter risk id"
// @Success 200
// @Router /api/v1/risks/{id} [put]
func (a *App) updateRisks(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var t models.CreateRisk
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	project_id := vars["id"]

	res, err := httpclient.UpdateRisks(a.Cfg.GetTeamworkURL(), project_id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), &t, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error creating a project risk", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}

// deleteRisks godoc
// @Summary Delete a risk
// @Description Remove a risk
// @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/a2a40cec66701-deleting-a-risk-on-a-project for more parameter options.
// @Tags risks
// @Accept  json
// @Produce  json
// @Param id path string false "Enter risk id"
// @Success 200
// @Router /api/v1/risks/{id} [delete]
func (a *App) deleteRisks(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	id := vars["id"]

	t, err := httpclient.DeleteRisks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error deleting a risk", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, t)
}
