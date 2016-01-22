package integration_tests

import (
	"encoding/base64"
	"fmt"
	"github.com/mholt/caddy/caddy"
	"net/http"
	"testing"
)

func TestBrowseMiddlewareForHtmlResponse(t *testing.T) {
	//formulating configuration for starting Caddy.
	//The configuration below is used defualt host <localhost> and default port <2015> for running the server
	//basicauth middleware is activated
	input := caddy.CaddyfileInput{
		Contents: []byte(fmt.Sprintf("%s:%s\nbasicauth /protected/ me you", "127.0.0.1", "2015")),
	}
	//starting caddy server
	err := caddy.Start(input)

	if err != nil {
		t.Fatalf("Error starting caddy: %v", err)
	}

	tests := []struct {
		from           string
		responseStatus int    //expected response status for the http request
		cred           string //credentials for http basicauth
	}{
		{"/protected", http.StatusUnauthorized, "my-password:my-username"},
		{"/protected", http.StatusOK, "me:you"},
		{"/protected", http.StatusUnauthorized, ""},
	}

	for i, test := range tests {
		//formulate the http request
		client := &http.Client{}
		httpReq, err := http.NewRequest("GET", "http://127.0.0.1:2015/protected/", nil)
		if err != nil {
			t.Fatal("GET: invalid Url or Request failed")
		}
		//set the auth headers
		auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(test.cred))
		httpReq.Header.Set("Authorization", auth)

		r, err := client.Do(httpReq)
		if err != nil {
			t.Fatal("GET: Request failed: " + err.Error())
		}
		//assert the response status code
		if r.StatusCode != test.responseStatus {
			t.Errorf("Test %d: Expected Header '%d' but was '%d'",
				i, test.responseStatus, r.StatusCode)
		}

	}
	caddy.Stop()
}
