package httpclient

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	MQTT_SVC_PUBLISH_URL = "http://localhost:19000/publish" //TODO add to pkg/config/
)

//Event body to be sent in Publish API call to mqtt-svc
//When this event body is read by mqtt-svc, /trigger Api will be called on the TriggerURL, passing in TriggerData
type Event struct {
	TriggerURL  string          `json:"triggerURL"`  //connector URL to send the /trigger call to
	TriggerData json.RawMessage `json:"triggerData"` //Data to pass to /trigger API as req body (connector1 data -> connector2)
}

func ProcessCreateProjectEvent(data interface{}) error {

	/////////////////////////////////////////////////////////
	//DYNAMIC - Connector developer can decide
	publishTopicID := "freshservice_conn"
	triggerURL := "http://localhost:8012/api/v2/trigger" //freshservice connector /trigger API endpoint
	triggerData, err := json.Marshal(data)               //data passed from this connector
	if err != nil {
		return err
	}
	////////////////////////////////////////////////////////////

	url := MQTT_SVC_PUBLISH_URL + "?topicID=" + publishTopicID

	headers := http.Header(make(map[string][]string))
	headers.Add("Content-Type", "application/json")

	reqBody := Event{
		TriggerURL:  triggerURL,
		TriggerData: []byte(triggerData),
	}

	resp, err := Post(url, headers, reqBody)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return err
	}

	return nil
}

func Post(url string, headers http.Header, body interface{}) (*resty.Response, error) {

	client := resty.New()
	r := client.NewRequest()

	for k, vals := range headers {
		for _, v := range vals {
			r.Header.Add(k, v)
		}
	}

	r.SetBody(body)
	resp, err := r.Post(url)
	if err != nil {
		log.Printf("made http request [%s] [%s]. got error: %s", "POST", url, err)
		return nil, err
	}

	log.Printf("made http request [%s] [%s]. got response [%s] [%s]",
		r.Method, r.URL, resp.Status(), resp.String())

	return resp, nil
}
