package integration_tests

import (
	"fmt"
	"github.com/mholt/caddy/caddy"
	"net/http"
	"testing"
)

func TestHeaderMiddleware(t *testing.T) {
	//formulating configuration for starting Caddy.
	//The configuration below is used defualt host <localhost> and default port <2015> for running the server
	//browse middleware is activated
	var customHeader, customHeaderValue string
	customHeader = "X-Custom-Header"
	customHeaderValue = "My-Value"
	input := caddy.CaddyfileInput{
		Contents: []byte(fmt.Sprintf("%s:%s\nheader /header %s %s  ", caddy.DefaultHost, "2016", customHeader, customHeaderValue)),
	}
	//starting caddy server
	err := caddy.Start(input)

	if err != nil {
		t.Fatalf("Error starting caddy: %v", err)
	}

	client := &http.Client{}
	httpReq, err := http.NewRequest("GET", "http://127.0.0.1:2016/header", nil)
	if err != nil {
		t.Fatal("GET: invalid Url or Request failed")
	}
	r, err := client.Do(httpReq)
	if err != nil {
		t.Fatal("GET: Request failed: " + err.Error())
	}
	defer r.Body.Close()
	headerFound := r.Header.Get(customHeader)
	if headerFound != customHeaderValue {
		t.Fatalf("Expected Header %s to be set to %s, found %s", customHeader, customHeaderValue, headerFound)
	}
	caddy.Stop()
}
