package app

// import (
// 	"encoding/json"
// 	"github.com/gorilla/mux"
// 	"github.com/kosha/teamwork-connector/pkg/httpclient"
// 	"github.com/kosha/teamwork-connector/pkg/models"
// 	"net/http"
// )

// // createSingleTag godoc
// // @Summary Creates a singular tag
// // @Description Creates a tag in the system
// // @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/7dc79ba8544b6-create-a-single-tag for more parameter options.
// // @Tags tags
// // @Accept json
// // @Produce json
// // @Param tag body models.Risk false "Enter project risk properties"
// // @Success 200
// // @Router /api/v1/project/{project_id}/risks [post]
// func (a *App) createProjectRisk(w http.ResponseWriter, r *http.Request) {
// 	//Allow CORS here By * or specific origin
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "*")
// 	var t models.CreateRisk
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&t); err != nil {
// 		a.Log.Errorf("Error parsing json payload", err)
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	defer r.Body.Close()

// 	vars := mux.Vars(r)
// 	project_id := vars["id"]

// 	res, err := httpclient.CreateProjectRisk(a.Cfg.GetTeamworkURL(), project_id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), &t, r.URL.Query())
// 	if err != nil {
// 		a.Log.Errorf("Error creating a project risk", err)
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	respondWithJSON(w, http.StatusOK, res)
// }

// // getAllRisks godoc
// // @Summary Get all risks
// // @Description List all risks across projects
// // @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/64d22ab985a58-get-all-risks for more parameter options.
// // @Tags risks
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} models.ReturnedRisks
// // @Router /api/v1/risks [get]
// func (a *App) getAllRisks(w http.ResponseWriter, r *http.Request) {

// 	//Allow CORS here By * or specific origin
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "*")

// 	p := httpclient.GetAllRisks(a.Cfg.GetTeamworkURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

// 	respondWithJSON(w, http.StatusOK, p)
// }

// // getProjectRisks godoc
// // @Summary Get all risks associated with specified project
// // @Description List all risks in the project
// // @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/3f3555d7cd5d0-get-risks-for-a-specific-project for more parameter options.
// // @Tags risks
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} models.ReturnedRisks
// // @Router /api/v1/project/{project_id}/risks [get]
// func (a *App) getProjectRisks(w http.ResponseWriter, r *http.Request) {

// 	//Allow CORS here By * or specific origin
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "*")

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	p := httpclient.GetProjectRisks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())

// 	respondWithJSON(w, http.StatusOK, p)
// }

// // updateRisks godoc
// // @Summary Update a project risk
// // @Description Add an update to a project's risk
// // @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/a1d50e01411b8-updating-a-risk-on-a-project for more parameter options.
// // @Tags risks
// // @Accept  json
// // @Produce  json
// // @Param project body models.CreateRisk false "Enter project risk properties"
// // @Success 200
// // @Router /api/v1/risks/{id} [put]
// func (a *App) updateRisks(w http.ResponseWriter, r *http.Request) {

// 	//Allow CORS here By * or specific origin
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "*")

// 	var t models.CreateRisk
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&t); err != nil {
// 		a.Log.Errorf("Error parsing json payload", err)
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	defer r.Body.Close()

// 	vars := mux.Vars(r)
// 	project_id := vars["id"]

// 	res, err := httpclient.UpdateRisks(a.Cfg.GetTeamworkURL(), project_id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), &t, r.URL.Query())
// 	if err != nil {
// 		a.Log.Errorf("Error creating a project risk", err)
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	respondWithJSON(w, http.StatusOK, res)
// }

// // deleteRisks godoc
// // @Summary Delete a risk
// // @Description Remove a risk
// // @Description Please refer to https://apidocs.teamwork.com/docs/teamwork/a2a40cec66701-deleting-a-risk-on-a-project for more parameter options.
// // @Tags risks
// // @Accept  json
// // @Produce  json
// // @Success 200
// // @Router /api/v1/risks/{id} [delete]
// func (a *App) deleteRisks(w http.ResponseWriter, r *http.Request) {
// 	//Allow CORS here By * or specific origin
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "*")

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	t, err := httpclient.DeleteRisks(a.Cfg.GetTeamworkURL(), id, a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())
// 	if err != nil {
// 		a.Log.Errorf("Error deleting a risk", err)
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 	}

// 	respondWithJSON(w, http.StatusOK, t)
// }
