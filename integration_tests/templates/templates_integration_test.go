package integration_tests

import (
	"fmt"

	"github.com/mholt/caddy/caddy"
	//"github.com/mholt/caddy/middleware"
	//"encoding/json"
	//"github.com/mholt/caddy/middleware/browse"
	"io/ioutil"
	"net/http"
	//"net/url"
	//"os"
	//"path/filepath"
	"testing"
	//"time"
	//"time"
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
	httpReq, err := http.NewRequest("GET", "http://127.0.0.1:2016/testdata/photos/test.html", nil)
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
	//asserting the response status code
	// if r.StatusCode != http.StatusOK {
	// 	t.Fatalf("Wrong status, expected %d, got %d", http.StatusOK, r.StatusCode)
	// }
	// 	expectedBody := `<!DOCTYPE html>
	// <html>
	// <head>
	// <title>Template</title>
	// </head>
	// <body>
	// <h1>Header</h1>
	// <h1>/testdata/photos/</h1>
	//
	// <a href="test.html">test.html</a><br>
	//
	// <a href="test2.html">test2.html</a><br>
	//
	// <a href="test3.html">test3.html</a><br>
	//
	// </body>
	// </html>
	// `
	t.Logf("\nresponse %s\n ", respBody)
	// if respBody != expectedBody {
	// 	t.Fatalf("Expected body: %v got: %v", expectedBody, respBody)
	// }
	caddy.Stop()
}
