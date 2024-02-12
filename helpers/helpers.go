package helpers

import (
	"encoding/json"
	"log"
	//"time"
)

var Response string

type Event struct {
	Version     string        `json:"version"`
	Id          string        `json:"id"`
	DetailType  string        `json:"detail-type"`
	Source      string        `json:"source"`
	Account     string        `json:"account"`
	Time        string        `json:"time"`
	Region      string        `json:"region"`
	Resources   []interface{} `json:"resources"`
	EventName   string        `json:"eventName"`
	EventSource string        `json:"eventSource"`
	SourceIP    string        `json:"sourceIPAddress"`
	BucketName  string        `json:"bucketName"`
}

func ToJson(s string) (Event, error) {
	var e Event
	var event interface{}
	err := json.Unmarshal([]byte(s), &event)
	if err != nil {
		log.Printf("Error unmarshaling data: %v\n", err)
		return e, err
	}

	eventParsed := event.(map[string]interface{})
	//eventParsedDetail := eventParsed["detail"].(map[string]interface{})
	//eventParsedRP := eventParsedDetail["requestParameters"].(map[string]interface{})
	//BucketName := eventParsedDetail["requestParameters"].(map[string]interface{})["bucketName"].(string)
	//log.Println(BucketName)
	e = Event{
		Id:          eventParsed["id"].(string),
		DetailType:  eventParsed["detail-type"].(string),
		Source:      eventParsed["source"].(string),
		Account:     eventParsed["account"].(string),
		Time:        eventParsed["time"].(string),
		EventName:   eventParsed["detail"].(map[string]interface{})["eventName"].(string),
		EventSource: eventParsed["detail"].(map[string]interface{})["eventSource"].(string),
		SourceIP:    eventParsed["detail"].(map[string]interface{})["sourceIPAddress"].(string),
		BucketName:  eventParsed["detail"].(map[string]interface{})["requestParameters"].(map[string]interface{})["bucketName"].(string),
		//SourceIP:    eventParsed["detail"].(map[string]interface{})["sourceIPAddress"].(string),
	}

	//eventName := eventParsed["detail"].(map[string]interface{})["eventName"].(string)
	//eventName := requestParameters["eventName"].(string)
	//e.EventName = eventName

	return e, nil

}
