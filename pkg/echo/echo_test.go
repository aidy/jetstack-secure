package echo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jetstack/preflight/api"
)

type testInput struct {
	description string
	data        *api.DataReadingsPost
	exp         int
	method      string
}

func TestEchoServerRequestResponse(t *testing.T) {
	// create sample data in same format that would be generated by the agent
	sampleUploadCases := []testInput{
		{
			description: "correct request input should return status code 200",
			data: &api.DataReadingsPost{
				AgentMetadata: &api.AgentMetadata{
					Version:   "test suite",
					ClusterID: "test_suite_cluster",
				},
				DataGatherTime: time.Now(),
				DataReadings: []*api.DataReading{
					&api.DataReading{
						ClusterID:    "test_suite_cluster",
						DataGatherer: "dummy",
						Timestamp:    api.Time{Time: time.Now()},
						Data: map[string]string{
							"test": "test",
						},
						SchemaVersion: "2.0.0",
					},
				},
			},
			exp:    http.StatusOK,
			method: "POST",
		},
		{
			description: "sending GET request should return status code 400",
			method:      "GET",
			data:        nil,
			exp:         http.StatusBadRequest,
		},
	}

	for _, sampleUpload := range sampleUploadCases {
		// generate the JSON representation of the data to be sent to the echo server
		requestBodyJSON, err := json.Marshal(sampleUpload.data)
		if err != nil {
			t.Fatalf("[%s]\nfailed to generate JSON request body to post: %s", sampleUpload.description, err)
		}

		// generate a request to test the handler containing the JSON data as a body
		req, err := http.NewRequest(sampleUpload.method, "http://example.com/api/v1/datareadings", bytes.NewBuffer(requestBodyJSON))
		if err != nil {
			t.Fatalf("[%s]\nfailed to generate request to test echo server: %s", sampleUpload.description, err)
		}

		// create recorder to save the response
		rr := httptest.NewRecorder()

		// perform the request with the handler
		echoHandler(rr, req)

		// Check the response from the echo handler is the expected one
		response := rr.Result()
		if response.StatusCode != sampleUpload.exp {
			t.Fatalf("[%s]\necho server responded with an unexpected code: %d", sampleUpload.description, response.StatusCode)
		}
	}
}