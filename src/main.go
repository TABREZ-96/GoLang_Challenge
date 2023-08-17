package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "https://graph.facebook.com/v17.0/1300740743881123/events"                                                                                                                                                                                       //Replace with your pixel id
	accessToken := "EAAJT0IZCxLsABOZBrzZAtSMOBxnJOlAEvIZCBZBzVhxQbZAkmGkhFE2cuIoiA36AQ2YAC0goEW2lm2DH91E4b6FEfFAZAls178hDadRv0cHb2hvlBbLqVJquEtgFCGQpdN7zsNv5CeGI1LoirjcUbUSR23wnpTUEZALWi0h12LxL3Dj8ZCdkLwU1vKqYDvZAmW87z2qeP8lz7D8iOsOB5kKDgiVrlFZCwZDZD" //replace with your token

	// Prepare the event data
	eventData := map[string]interface{}{
		"data": []map[string]interface{}{
			{
				"event_name": "Purchase",
				"event_time": 1690317359,
				"user_data": map[string]string{
					"fbc": "fb.1.1558571054389.1098115397",
					"fbp": "fb.1.1558571054389.1098115397",
					"em":  "f660ab912ec121d1b1e928a0bb4bc61b15f5ad44d5efdc4e1c92a25e99b8e44a",
				},
				"custom_data": map[string]string{
					"currency": "USD",
					"value":    "19.99",
				},
				"event_source_url": "https://example.com/product123",
				"action_source":    "website",
			},
		},
	}

	// Convert the event data to JSON
	jsonData, err := json.Marshal(eventData)
	if err != nil {
		fmt.Println("Error marshaling JSON data:", err)
		return
	}

	// Send the HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Print the response
	fmt.Println(result)
}
