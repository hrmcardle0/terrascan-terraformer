package helpers

import (
	"encoding/json"
	//"fmt"
	"log"
	//"time"
	//"errors"
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
	FilterValue string        `json:"bucketName"`
}

type NewEvent struct {
	Source  string `json:"source"`
	Account string `json:"account"`
}

func ToJson(s string) (NewEvent, error) {
	/*
		var e Event
		var event interface{}
		err := json.Unmarshal([]byte(s), &event)
		if err != nil {
			log.Printf("Error unmarshaling data: %v\n", err)
			return e, err
		}
	*/
	var event interface{}
	err := json.Unmarshal([]byte(s), &event)
	if err != nil {
		log.Printf("Error unmarshaling data: %v\n", err)
		return NewEvent{}, err
	}

	eventParsed := event.(map[string]interface{})

	return NewEvent{
		Source:  eventParsed["source"].(string),
		Account: eventParsed["account"].(string),
	}, nil

	/*
		//eventParsedDetail := eventParsed["detail"].(map[string]interface{})
		//eventParsedRP := eventParsedDetail["requestParameters"].(map[string]interface{})
		//BucketName := eventParsedDetail["requestParameters"].(map[string]interface{})["bucketName"].(string)
		//log.Println(BucketName)

		log.Println("Parsing event")

		if _, exists := eventParsed["detail"].(map[string]interface{})["errorCode"].(string); exists {
			log.Println("Error Code Found. Skipping")
			return e, errors.New("Error code found in event")
		}

		switch eventParsed["source"] {

		// s3
		case "aws.s3":
			e = Event{
				Id:          eventParsed["id"].(string),
				DetailType:  eventParsed["detail-type"].(string),
				Source:      "s3",
				Account:     eventParsed["account"].(string),
				Time:        eventParsed["time"].(string),
				EventName:   eventParsed["detail"].(map[string]interface{})["eventName"].(string),
				EventSource: eventParsed["detail"].(map[string]interface{})["eventSource"].(string),
				SourceIP:    eventParsed["detail"].(map[string]interface{})["sourceIPAddress"].(string),
				FilterValue: "Name=id;Value=" + eventParsed["detail"].(map[string]interface{})["requestParameters"].(map[string]interface{})["bucketName"].(string),
				//SourceIP:    eventParsed["detail"].(map[string]interface{})["sourceIPAddress"].(string),
			}

		// api gateway
		case "aws.apigateway":
			eventName := eventParsed["detail"].(map[string]interface{})["eventName"].(string)

			// api gateway rest
			if eventName == "CreateRestApi" {
				e = Event{
					Id:          eventParsed["id"].(string),
					DetailType:  eventParsed["detail-type"].(string),
					Source:      "api_gateway",
					Account:     eventParsed["account"].(string),
					Time:        eventParsed["time"].(string),
					EventName:   eventParsed["detail"].(map[string]interface{})["eventName"].(string),
					EventSource: eventParsed["detail"].(map[string]interface{})["eventSource"].(string),
					SourceIP:    eventParsed["detail"].(map[string]interface{})["sourceIPAddress"].(string),
					FilterValue: "Name=id;Value=" + eventParsed["detail"].(map[string]interface{})["responseElements"].(map[string]interface{})["restapiUpdate"].(map[string]interface{})["restApiId"].(string),
					//SourceIP:    eventParsed["detail"].(map[string]interface{})["sourceIPAddress"].(string),
				}
			}
			// api gateway rest
			if eventName == "CreateStage" {
				apiId := eventParsed["detail"].(map[string]interface{})["requestParameters"].(map[string]interface{})["restApiId"].(string)
				stageName := eventParsed["detail"].(map[string]interface{})["requestParameters"].(map[string]interface{})["stageName"].(string)
				filterString := fmt.Sprintf("Name=id;Value=%s/%s", apiId, stageName)
				e = Event{
					Id:          eventParsed["id"].(string),
					DetailType:  eventParsed["detail-type"].(string),
					Source:      "api_gateway",
					Account:     eventParsed["account"].(string),
					Time:        eventParsed["time"].(string),
					EventName:   eventParsed["detail"].(map[string]interface{})["eventName"].(string),
					EventSource: eventParsed["detail"].(map[string]interface{})["eventSource"].(string),
					SourceIP:    eventParsed["detail"].(map[string]interface{})["sourceIPAddress"].(string),
					FilterValue: filterString,
					//SourceIP:    eventParsed["detail"].(map[string]interface{})["sourceIPAddress"].(string),
				}
			}

		// iam
		case "aws.iam":
			if eventName == "CreateUser" {
				apiId := eventParsed["detail"].(map[string]interface{})["requestParameters"].(map[string]interface{})["restApiId"].(string)
				stageName := eventParsed["detail"].(map[string]interface{})["requestParameters"].(map[string]interface{})["stageName"].(string)
				filterString := fmt.Sprintf("Name=id;Value=%s/%s", apiId, stageName)
				e = Event{
					Id:          eventParsed["id"].(string),
					DetailType:  eventParsed["detail-type"].(string),
					Source:      "",
					Account:     eventParsed["account"].(string),
					Time:        eventParsed["time"].(string),
					EventName:   eventParsed["detail"].(map[string]interface{})["eventName"].(string),
					EventSource: eventParsed["detail"].(map[string]interface{})["eventSource"].(string),
					SourceIP:    eventParsed["detail"].(map[string]interface{})["sourceIPAddress"].(string),
					FilterValue: filterString,
					//SourceIP:    eventParsed["detail"].(map[string]interface{})["sourceIPAddress"].(string),
				}

		}

		return e, nil
	*/

}
