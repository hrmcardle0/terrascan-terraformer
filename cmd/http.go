package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"gitlab.com/secops/development/aws/terrascan/helpers"
)

func InitHttp(terraformString string) error {
	// call tfsec endpoint
	url := "http://localhost:8081/scan"
	data := map[string]string{"content": terraformString}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return err
	}

	// Create a new POST request with the JSON data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Println("Error: unexpected response status code:", resp.StatusCode)
		return err
	}

	// Read the response body
	var result string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Error decoding JSON response:", err)
		return err
	}

	// Print the response
	log.Println("Response:", result)
	helpers.Response = result

	return nil

}
