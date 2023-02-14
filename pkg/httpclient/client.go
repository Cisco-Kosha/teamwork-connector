package httpclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
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
		fmt.Println("There was an error: ", err)
		return nil
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}

func GetResponseHeaders(username string, password string, path string, params url.Values) http.Header {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	req.URL.RawQuery = params.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing http client: ", err)
		return nil
	}

	defer resp.Body.Close()
	return resp.Header
}

func GetAccount(url string, username string, password string, params url.Values) *models.SingleAccount {
	req, err := http.NewRequest("GET", url+"account.json", nil)
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil
	}
	var account *models.SingleAccount

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &account)
	if err != nil {
		fmt.Println("There is an issue with unmarshaling: ", err)
		return nil
	}
	return account
}

func GetAllProjectsV3(url string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ProjectResponseV3) {
	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/api/v3/projects.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/api/v3/projects.json", nil)
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}

	var projects *models.ProjectResponseV3

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &projects)
	if err != nil {
		fmt.Println("There is an issue with unmarshaling: ", err)
		return nil, nil
	}

	// var error *models.Error
	// d := json.NewDecoder(strings.NewReader(string(res)))
	// d.DisallowUnknownFields()

	// if err := d.Decode(&error); err == nil {
	// 	return err
	// }

	return nil, projects
}

func GetAllProjectsV1(url string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ProjectResponseV1) {
	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects.json", nil)
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}

	var projects *models.ProjectResponseV1

	res := makeHttpReq(username, password, req, params)

	// Convert response body to target struct
	err = json.Unmarshal(res, &projects)
	if err != nil {
		fmt.Println("There is an issue with unmarshaling")
		return nil, nil
	}

	return nil, projects
}

func GetSingleProjectV3(url string, id string, username string, password string, params url.Values) *models.SingleProject {

	req, err := http.NewRequest("GET", url+"projects/api/v3/projects/"+id+".json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
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

func GetSingleProjectV1(url string, id string, username string, password string, params url.Values) *models.SingleProjectResponseV1 {

	req, err := http.NewRequest("GET", url+"projects/"+id+".json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil
	}
	var project *models.SingleProjectResponseV1

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &project)
	if err != nil {
		return nil
	}
	return project
}

func GetProjectTaskLists(url string, id string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.MultiTaskList) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/"+id+"/tasklists.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/"+id+"/tasklists.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var tasklists *models.MultiTaskList

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &tasklists)
	if err != nil {
		return nil, nil
	}
	return nil, tasklists
}

func GetProjectTasks(url string, id string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.Tasks) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/"+id+"/tasks.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/"+id+"/tasks.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
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

func GetTasks(url string, id string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.Tasks) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"tasklists/"+id+"/tasks.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"tasklists/"+id+"/tasks.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
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

func GetLatestActivityAllProjectsV3(url string, username string, password string, params url.Values) *models.MultiActivity {

	req, err := http.NewRequest("GET", url+"projects/api/v3/latestactivity.json", nil)
	if err != nil {
		fmt.Println("Error creating new request: ", err)
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

func GetLatestActivityAllProjectsV1(url string, username string, password string, params url.Values) *models.ActivityResponseV1 {

	req, err := http.NewRequest("GET", url+"latestactivity.json", nil)
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil
	}
	var activity *models.ActivityResponseV1

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
		fmt.Println("Error creating new request: ", err)
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
		fmt.Println("Error marshaling json")
		return "", err
	}
	req, err := http.NewRequest("POST", url+"projects/"+id+"/tasklists.json", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("Error creating new request: ", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(username, password, req, params)), nil
}

func RemoveProject(url string, id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest(http.MethodDelete, url+"/projects/"+id+".json", nil)
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return "", err
	}

	res := makeHttpReq(username, password, req, params)
	return string(res), nil
}

func RemoveTaskList(url string, id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest(http.MethodDelete, url+"/tasklists/"+id+".json", nil)
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return "", err
	}

	res := makeHttpReq(username, password, req, params)
	return string(res), nil
}

func GetSinglePerson(url string, id string, username string, password string, params url.Values) *models.SinglePerson {

	req, err := http.NewRequest("GET", url+"people/"+id+".json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
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
		fmt.Println("Error creating new request: ", err)
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

func GetPeople(url string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.People) {

	currentPerson := GetCurrentPerson(url, username, password, nil)
	companyId := currentPerson.Person.CompanyId

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"companies/"+companyId+"/people.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"companies/"+companyId+"/people.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var people *models.People

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &people)
	if err != nil {
		return nil, nil
	}
	return nil, people
}

func GetPersonsProjects(url string, id string, username string, password string, params url.Values) *models.ProjectResponseV3 {

	customParams := make(map[string][]string)
	customParams["fields[users]"] = append(customParams["fields[users]"], id)
	_, projects := GetAllProjectsV3(url, username, password, customParams, false)
	return projects
}

func CreateProjectTimeEntry(url string, project_id string, username string, password string, timeentry *models.CreateTimeEntry, params url.Values) (*models.CreatedTimeEntry, error) {

	jsonReq, err := json.Marshal(timeentry)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url+"projects/"+project_id+"/time_entries.json", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, err
	}
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

func GetAllTimeEntries(url string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ReturnedTimeEntries) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"time_entries.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"time_entries.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var timeEntries *models.ReturnedTimeEntries

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &timeEntries)
	if err != nil {
		return nil, nil
	}
	return nil, timeEntries
}

func GetProjectTimeEntries(url string, project_id string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ReturnedTimeEntries) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/"+project_id+"/time_entries.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/"+project_id+"/time_entries.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var timeEntries *models.ReturnedTimeEntries

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &timeEntries)
	if err != nil {
		return nil, nil
	}
	return nil, timeEntries
}

func UpdateTimeEntry(url string, id string, username string, password string, timeentry *models.CreateTimeEntry, params url.Values) (*models.CreatedTimeEntry, error) {

	jsonReq, err := json.Marshal(timeentry)
	if err != nil {
		fmt.Println("Error creating new request: ", err)
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
		fmt.Println("Error creating new request: ", err)
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
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(username, password, req, params)), nil
}

func GetAllProjectUpdatesV3(url string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ProjectUpdateResponse) {
	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/api/v3/projects/updates.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/api/v3/projects/updates.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var updates *models.ProjectUpdateResponse

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &updates)
	if err != nil {
		return nil, nil
	}
	return nil, updates
}

func GetAllProjectUpdatesV1(url string, id string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ProjectUpdateResponseV1) {
	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/"+id+"/updates.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/"+id+"/updates.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var updates *models.ProjectUpdateResponseV1

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &updates)
	if err != nil {
		return nil, nil
	}
	return nil, updates
}

func ModifyProjectUpdate(url string, id string, username string, password string, update *models.ProjectUpdate, params url.Values) (string, error) {

	jsonReq, err := json.Marshal(update)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("PUT", url+"projects/updates/"+id+".json", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(username, password, req, params)), nil
}

func DeleteProjectUpdate(url string, id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest("DELETE", url+"projects/updates/"+id+".json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return "", err
	}
	res := makeHttpReq(username, password, req, params)
	return string(res), nil
}

func CreateProjectRisk(url string, id string, username string, password string, risk *models.CreateRisk, params url.Values) (string, error) {
	jsonReq, err := json.Marshal(risk)
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return "", err
	}
	req, err := http.NewRequest("POST", url+"projects/"+id+"/risks.json", bytes.NewBuffer(jsonReq))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(username, password, req, params)), nil
}

func GetAllRisksV3(url string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ReturnedRisks) {
	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/api/v3/risks.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/api/v3/risks.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var risks *models.ReturnedRisks

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &risks)
	if err != nil {
		return nil, nil
	}
	return nil, risks
}

func GetAllRisksV1(url string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ReturnedRisksV1) {
	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"risks.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"risks.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var risks *models.ReturnedRisksV1

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &risks)
	if err != nil {
		return nil, nil
	}
	return nil, risks
}

func GetProjectRisksV1(url string, project_id string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ReturnedRisksV1) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/"+project_id+"/risks.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/"+project_id+"/risks.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var tasks *models.ReturnedRisksV1

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &tasks)
	if err != nil {
		return nil, nil
	}
	return nil, tasks
}

func GetTasksAllProjects(url string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.Tasks) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"tasks.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"tasks.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
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


func GetProjectRisksV3(url string, project_id string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ReturnedRisks) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/api/v3/projects/"+project_id+"/risks.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/api/v3/projects/"+project_id+"/risks.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var tasks *models.ReturnedRisks

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &tasks)
	if err != nil {
		return nil, nil
	}
	return nil, tasks
}

func UpdateRisks(url string, id string, username string, password string, risk *models.CreateRisk, params url.Values) (string, error) {

	jsonReq, err := json.Marshal(risk)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("PUT", url+"risks/"+id+".json", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(username, password, req, params)), nil
}

func DeleteRisks(url string, id string, username string, password string, params url.Values) (string, error) {
	req, err := http.NewRequest("DELETE", url+"risks/"+id+".json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return "", err
	}
	res := makeHttpReq(username, password, req, params)
	return string(res), nil
}

func GetAllMilestones(url string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ReturnedMilestones) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/api/v3/milestones.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/api/v3/milestones.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var milestones *models.ReturnedMilestones

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &milestones)
	if err != nil {
		return nil, nil
	}
	return nil, milestones
}

func GetProjectMilestones(url string, id string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ReturnedMilestones) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/api/v3/projects/" + id + "/milestones.json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/api/v3/projects/" + id + "/milestones.json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var milestones *models.ReturnedMilestones

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &milestones)
	if err != nil {
		return nil, nil
	}
	return nil, milestones
}

func GetSingleMilestone(url string, id string, username string, password string, params url.Values, getRespHeaders bool) (http.Header, *models.ReturnedMilestone) {

	if getRespHeaders {
		return GetResponseHeaders(username, password, url+"projects/api/v3/milestones/" + id + ".json", params), nil
	}

	req, err := http.NewRequest("GET", url+"projects/api/v3/milestones/" + id + ".json", nil)

	if err != nil {
		fmt.Println("Error creating new request: ", err)
		return nil, nil
	}
	var milestones *models.ReturnedMilestone

	res := makeHttpReq(username, password, req, params)
	// Convert response body to target struct
	err = json.Unmarshal(res, &milestones)
	if err != nil {
		return nil, nil
	}
	return nil, milestones
}