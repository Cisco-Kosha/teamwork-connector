package httpclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/kosha/teamwork-connector/pkg/models"
)

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func makeHttpReq(username string, password string, req *http.Request, params url.Values) []byte {
	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	req.URL.RawQuery = params.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}

func GetResponseHeaders(username string, password string, path string, params url.Values) http.Header {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	req.URL.RawQuery = params.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	return resp.Header
}

func GetAccount(url string, username string, password string, params url.Values) *models.SingleAccount {
	req, err := http.NewRequest("GET", url+"account.json", nil)
	if err != nil {
		return nil
	}
	var account *models.SingleAccount

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &account)
	if err != nil {
		return nil
	}
	return account
}

func GetAllProjects(url string, username string, password string, params url.Values) *models.MultiProject {
	req, err := http.NewRequest("GET", url+"projects/api/v3/projects.json", nil)
	if err != nil {
		return nil
	}
	var project *models.MultiProject

	res := makeHttpReq(username, password, req, params)

	// Convert response body to target struct
	err = json.Unmarshal(res, &project)
	if err != nil {
		return nil
	}
	return project
}

func GetSingleProject(url string, id string, username string, password string, params url.Values) *models.SingleProject {

	req, err := http.NewRequest("GET", url+"projects/api/v3/projects/"+id+".json", nil)

	if err != nil {
		return nil
	}
	var project *models.SingleProject

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &project)
	if err != nil {
		return nil
	}
	return project
}

func GetProjectTaskLists(url string, id string, username string, password string, params url.Values) *models.MultiTaskList {

	req, err := http.NewRequest("GET", url+"projects/"+id+"/tasklists.json", nil)

	if err != nil {
		return nil
	}
	var tasklists *models.MultiTaskList

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &tasklists)
	if err != nil {
		return nil
	}
	return tasklists
}

func GetProjectTasks(url string, id string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.Tasks) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/"+id+"/tasks.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/"+id+"/tasks.json", nil)

	if err != nil {
		return nil, nil
	}
	var tasks *models.Tasks

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &tasks)
	if err != nil {
		return nil, nil
	}

	return nil, tasks
}

func GetTasks(url string, id string, username string, password string, params url.Values) *models.Tasks {
	req, err := http.NewRequest("GET", url+"tasklists/"+id+"/tasks.json", nil)

	if err != nil {
		return nil
	}
	var tasks *models.Tasks

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &tasks)
	if err != nil {
		return nil
	}
	return tasks
}

func GetLatestActivityAllProjects(url string, username string, password string, params url.Values) *models.MultiActivity {

	req, err := http.NewRequest("GET", url+"projects/api/v3/latestactivity.json", nil)
	if err != nil {
		return nil
	}
	var activity *models.MultiActivity

	res := makeHttpReq(username, password, req, params)

	// Convert response body to target struct
	err = json.Unmarshal(res, &activity)
	if err != nil {
		return nil
	}
	return activity
}

func CreateProject(url string, username string, password string, project *models.SingleProject, params url.Values) (string, error) {

	jsonReq, err := json.Marshal(project)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url+"projects.json", bytes.NewBuffer(jsonReq))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(username, password, req, params)), nil
}

func CreateTaskList(url string, id string, username string, password string, tasklist *models.NewTaskList, params url.Values) (string, error) {

	jsonReq, err := json.Marshal(tasklist)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url+"projects/"+id+"/tasklists.json", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(username, password, req, params)), nil
}

func RemoveProject(url string, id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest(http.MethodDelete, url+"/projects/"+id+".json", nil)
	if err != nil {
		return "", err
	}

	res := makeHttpReq(username, password, req, params)
	return string(res), nil
}

func RemoveTaskList(url string, id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest(http.MethodDelete, url+"/tasklists/"+id+".json", nil)
	if err != nil {
		return "", err
	}

	res := makeHttpReq(username, password, req, params)
	return string(res), nil
}

func GetSinglePerson(url string, id string, username string, password string, params url.Values) *models.SinglePerson {

	req, err := http.NewRequest("GET", url+"people/"+id+".json", nil)

	if err != nil {
		return nil
	}
	var person *models.SinglePerson

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &person)
	if err != nil {
		return nil
	}
	return person
}

func GetCurrentPerson(url string, username string, password string, params url.Values) *models.SinglePerson {

	req, err := http.NewRequest("GET", url+"me.json", nil)

	if err != nil {
		return nil
	}
	var person *models.SinglePerson

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &person)
	if err != nil {
		return nil
	}
	return person
}

func GetPeople(url string, username string, password string, params url.Values) *models.People {

	currentPerson := GetCurrentPerson(url, username, password, nil)
	companyId := currentPerson.Person.CompanyId

	req, err := http.NewRequest("GET", url+"companies/"+companyId+"/people.json", nil)

	if err != nil {
		return nil
	}
	var people *models.People

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &people)
	if err != nil {
		return nil
	}
	return people
}

func GetPersonsProjects(url string, id string, username string, password string, params url.Values) *models.MultiProject {

	customParams := make(map[string][]string)
	customParams["fields[users]"] = append(customParams["fields[users]"], id)
	projects := GetAllProjects(url, username, password, customParams)
	return projects
}

func CreateProjectTimeEntry(url string, project_id string, username string, password string, timeentry *models.CreateTimeEntry, params url.Values) (*models.CreatedTimeEntry, error) {

	jsonReq, err := json.Marshal(timeentry)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url+"projects/"+project_id+"/time_entries.json", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var createReply *models.CreatedTimeEntry
	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &createReply)
	if err != nil {
		return nil, err
	}
	return createReply, nil
}

func GetAllTimeEntries(url string, username string, password string, params url.Values) *models.ReturnedTimeEntries {
	req, err := http.NewRequest("GET", url+"time_entries.json", nil)

	if err != nil {
		return nil
	}
	var tasks *models.ReturnedTimeEntries

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &tasks)
	if err != nil {
		return nil
	}
	return tasks
}

func GetProjectTimeEntries(url string, project_id string, username string, password string, params url.Values) *models.ReturnedTimeEntries {
	req, err := http.NewRequest("GET", url+"projects/"+project_id+"/time_entries.json", nil)

	if err != nil {
		return nil
	}
	var tasks *models.ReturnedTimeEntries

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &tasks)
	if err != nil {
		return nil
	}
	return tasks
}

func UpdateTimeEntry(url string, id string, username string, password string, timeentry *models.CreateTimeEntry, params url.Values) (*models.CreatedTimeEntry, error) {

	jsonReq, err := json.Marshal(timeentry)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url+"time_entries/"+id+".json", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var createReply *models.CreatedTimeEntry
	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &createReply)
	if err != nil {
		return nil, err
	}
	return createReply, nil
}

func DeleteTimeEntry(url string, id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest("DELETE", url+"time_entries/"+id+".json", nil)

	if err != nil {
		return "", err
	}
	res := makeHttpReq(username, password, req, params)
	return string(res), nil
}

func CreateProjectUpdate(url string, id string, username string, password string, update *models.ProjectUpdate, params url.Values) (string, error) {
	jsonReq, err := json.Marshal(update)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url+"projects/"+id+"/update.json", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(username, password, req, params)), nil
}

func GetAllProjectUpdates(url string, username string, password string, params url.Values) *models.ProjectUpdateResponse {
	req, err := http.NewRequest("GET", url+"projects/api/v3/projects/updates.json", nil)

	if err != nil {
		return nil
	}
	var updates *models.ProjectUpdateResponse

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &updates)
	if err != nil {
		return nil
	}
	return updates
}

func ModifyProjectUpdate(url string, id string, username string, password string, update *models.ProjectUpdate, params url.Values) (string, error) {

	jsonReq, err := json.Marshal(update)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("PUT", url+"projects/updates/"+id+".json", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(username, password, req, params)), nil
}

func DeleteProjectUpdate(url string, id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest("DELETE", url+"projects/updates/"+id+".json", nil)

	if err != nil {
		return "", err
	}
	res := makeHttpReq(username, password, req, params)
	return string(res), nil
}

func CreateProjectRisk(url string, id string, username string, password string, risk *models.CreateRisk, params url.Values) (string, error) {
	jsonReq, err := json.Marshal(risk)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url+"projects/"+id+"/risks.json", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(username, password, req, params)), nil
}

func GetAllRisks(url string, username string, password string, params url.Values) *models.ReturnedRisks {
	req, err := http.NewRequest("GET", url+"projects/api/v3/risks.json", nil)

	if err != nil {
		return nil
	}
	var risks *models.ReturnedRisks

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &risks)
	if err != nil {
		return nil
	}
	return risks
}

func GetProjectRisks(url string, project_id string, username string, password string, params url.Values) *models.ReturnedRisks {
	req, err := http.NewRequest("GET", url+"projects/api/v3/projects/"+project_id+"/risks.json", nil)

	if err != nil {
		return nil
	}
	var tasks *models.ReturnedRisks

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &tasks)
	if err != nil {
		return nil
	}
	return tasks
}

func UpdateRisks(url string, id string, username string, password string, risk *models.CreateRisk, params url.Values) (string, error) {

	jsonReq, err := json.Marshal(risk)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("PUT", url+"risks/"+id+".json", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(username, password, req, params)), nil
}

func DeleteRisks(url string, id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest("DELETE", url+"risks/"+id+".json", nil)

	if err != nil {
		return "", err
	}
	res := makeHttpReq(username, password, req, params)
	return string(res), nil
}
