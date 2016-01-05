package integration_tests

import (
	"fmt"
	"github.com/mholt/caddy/caddy"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBrowseMiddlewareForHtmlResponse(t *testing.T) {
	//formulating configuration for starting Caddy.
	//The configuration below is used defualt host <localhost> and default port <2015> for running the server
	//browse middleware is activated
	input := caddy.CaddyfileInput{
		Contents: []byte(fmt.Sprintf("%s:%s\ntemplates /testdata/photos/ .html", caddy.DefaultHost, "2016")),
	}
	//starting caddy server
	err := caddy.Start(input)

	if err != nil {
		t.Fatalf("Error starting caddy: %v", err)
	}

	client := &http.Client{}
	httpReq, err := http.NewRequest("GET", "http://127.0.0.1:2016/testdata/photos/", nil)
	if err != nil {
		t.Fatal("GET: invalid Url or Request failed")
	}
	r, err := client.Do(httpReq)
	if err != nil {
		t.Fatal("GET: Request failed: " + err.Error())
	}
	//reading the response
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Test: Could not parse response body: %v", err)
	}
	defer r.Body.Close()
	respBody := string(data)

	expectedBody := `<!DOCTYPE html><html><head><title>test page</title></head><body><h1>Header title</h1>
</body></html>
`
		
	if respBody != expectedBody {
		t.Fatalf("Test: the expected body\n%v is different from the response one: \n%v", expectedBody, respBody)
	}

	caddy.Stop()
}
