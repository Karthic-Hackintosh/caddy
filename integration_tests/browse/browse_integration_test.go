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
		Contents: []byte(fmt.Sprintf("%s:%s\nbrowse /testdata/photos/ ./testdata/photos.tpl", caddy.DefaultHost, "2016")),
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
	//asserting the response status code
	// if r.StatusCode != http.StatusOK {
	// 	t.Fatalf("Wrong status, expected %d, got %d", http.StatusOK, r.StatusCode)
	// }
	expectedBody := `<!DOCTYPE html>
<html>
<head>
<title>Template</title>
</head>
<body>
<h1>Header</h1>
<h1>/testdata/photos/</h1>

<a href="test.html">test.html</a><br>

<a href="test2.html">test2.html</a><br>

<a href="test3.html">test3.html</a><br>

</body>
</html>
`
	if respBody != expectedBody {
		t.Fatalf("Expected body: %v got: %v", expectedBody, respBody)
	}
	caddy.Stop()
}

// func TestBrowseMiddlewareForJsonResponse(t *testing.T) {
// 	//formulating configuration for starting Caddy.
// 	//The configuration below is used defualt host <localhost> and default port <2015> for running the server
// 	//browse middleware is activated
// 	input := caddy.CaddyfileInput{
// 		Contents: []byte(fmt.Sprintf("%s:%s\nbrowse /testdata/photos/", caddy.DefaultHost, caddy.DefaultPort)),
// 	}

// 	//Getting the listing from the ./testdata/photos, the listing returned will be used to validate test results
// 	testDataPath := "./testdata/photos/"
// 	file, err := os.Open(testDataPath)
// 	if err != nil {
// 		if os.IsPermission(err) {
// 			t.Fatalf("Os Permission Error")
// 		}
// 	}
// 	defer file.Close()

// 	files, err := file.Readdir(-1)
// 	if err != nil {
// 		t.Fatalf("Unable to Read Contents of the directory")
// 	}
// 	var fileinfos []browse.FileInfo

// 	for i, f := range files {
// 		name := f.Name()

// 		// Tests fail in CI environment because all file mod times are the same for
// 		// some reason, making the sorting unpredictable. To hack around this,
// 		// we ensure here that each file has a different mod time.
// 		chTime := f.ModTime().Add(-(time.Duration(i) * time.Second))
// 		if err := os.Chtimes(filepath.Join(testDataPath, name), chTime, chTime); err != nil {
// 			t.Fatal(err)
// 		}

// 		if f.IsDir() {
// 			name += "/"
// 		}

// 		url := url.URL{Path: name}

// 		fileinfos = append(fileinfos, browse.FileInfo{
// 			IsDir:   f.IsDir(),
// 			Name:    f.Name(),
// 			Size:    f.Size(),
// 			URL:     url.String(),
// 			ModTime: chTime,
// 			Mode:    f.Mode(),
// 		})
// 	}
// 	listing := browse.Listing{Items: fileinfos}
// 	//starting the caddy server for Testing the browse Endpoint
// 	err = caddy.Start(input)

// 	if err != nil {
// 		t.Fatalf("Failed to start Caddy: %v", err)
// 	}

// 	tests := []struct {
// 		QueryUrl       string
// 		SortBy         string
// 		OrderBy        string
// 		Limit          int
// 		shouldErr      bool
// 		expectedResult []browse.FileInfo
// 	}{
// 		//test case 1: testing for default sort and  order and without the limit parameter, default sort is by name and the default order is ascending
// 		{"", "", "", -1, false, listing.Items},
// 		//test case 2: limit is set to 1, orderBy and sortBy is default
// 		{"?limit=1", "", "", 1, false, listing.Items[:1]},
// 		//test case 3 : if the listing request is bigger than total size of listing then it should return everything
// 		{"?limit=100000000", "", "", 100000000, false, listing.Items},
// 		//test case 4 : testing for negative limit
// 		{"?limit=-1", "", "", -1, false, listing.Items},
// 		//test case 5 : testing with limit set to -1 and order set to descending
// 		{"?limit=-1&order=desc", "", "desc", -1, false, listing.Items},
// 		//test case 6 : testing with limit set to 2 and order set to descending
// 		{"?limit=2&order=desc", "", "desc", 2, false, listing.Items},
// 		//test case 7 : testing with limit set to 3 and order set to descending
// 		{"?limit=3&order=desc", "", "desc", 3, false, listing.Items},
// 		//test case 8 : testing with limit set to 3 and order set to ascending
// 		{"?limit=3&order=asc", "", "asc", 3, false, listing.Items},
// 		//test case 9 : testing with limit set to 1111111 and order set to ascending
// 		{"?limit=1111111&order=asc", "", "asc", 1111111, false, listing.Items},
// 		//test case 10 : testing with limit set to default and order set to ascending and sorting by size
// 		{"?order=asc&sort=size", "size", "asc", -1, false, listing.Items},
// 		//test case 11 : testing with limit set to default and order set to ascending and sorting by last modified
// 		{"?order=asc&sort=time", "time", "asc", -1, false, listing.Items},
// 		//test case 12 : testing with limit set to 1 and order set to ascending and sorting by last modified
// 		{"?order=asc&sort=time&limit=1", "time", "asc", 1, false, listing.Items},
// 		//test case 13 : testing with limit set to -100 and order set to ascending and sorting by last modified
// 		{"?order=asc&sort=time&limit=-100", "time", "asc", -100, false, listing.Items},
// 		//test case 14 : testing with limit set to -100 and order set to ascending and sorting by size
// 		{"?order=asc&sort=size&limit=-100", "size", "asc", -100, false, listing.Items},
// 	}
// 	for i, test := range tests {
// 		var marsh []byte
// 		fmt.Printf("\n%s\n", "Before Sending Get")
// 		// cfg := forest.NewConfig("/testdata/photos/", test.QueryUrl).Header("Accept", "application/json")
// 		// r := browseTest.GET(t, cfg)
// 		client := &http.Client{}
// 		httpReq, err := http.NewRequest("GET", "http://127.0.0.1:2015/testdata/photos/"+test.QueryUrl, nil)
// 		if err != nil {
// 			t.Fatal("GET: invalid Url or Request failed")
// 		}
// 		httpReq.Header.Set("Accept", "application/json")
// 		//copyHeaders(config.HeaderMap, httpReq.Header)
// 		//t.Logf("\n%v %v %v", httpReq.Method, httpReq.URL, headersString(httpReq.Header))
// 		r, err := client.Do(httpReq)
// 		if err != nil {
// 			t.Fatal("GET: Request failed")
// 		}
// 		data, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			t.Fatalf("Test: Could not parse response body: %v", err)
// 		}
// 		defer r.Body.Close()
// 		fmt.Printf("\nbody %s\n", string(data))
// 		//forest.ExpectStatus(t, r, 200)
// 		if r.StatusCode != http.StatusOK {
// 			t.Fatalf("Wrong status, expected %d, got %d", http.StatusOK, r.StatusCode)
// 		}
// 		if r.Header.Get("Content-Type") != "application/json; charset=utf-8" {
// 			t.Fatalf("Expected Content type to be application/json; charset=utf-8, but got %s ", r.Header.Get("Content-Type"))
// 		}

// 		actualJSONResponse := string(data)
// 		copyOflisting := listing
// 		if test.SortBy == "" {
// 			copyOflisting.Sort = "name"
// 		} else {
// 			copyOflisting.Sort = test.SortBy
// 		}
// 		if test.OrderBy == "" {
// 			copyOflisting.Order = "asc"
// 		} else {
// 			copyOflisting.Order = test.OrderBy
// 		}

// 		copyOflisting.ApplySort()

// 		limit := test.Limit
// 		if limit <= len(copyOflisting.Items) && limit > 0 {
// 			marsh, err = json.Marshal(copyOflisting.Items[:limit])
// 		} else { // if the 'limit' query is empty, or has the wrong value, list everything
// 			marsh, err = json.Marshal(copyOflisting.Items)
// 		}

// 		if err != nil {
// 			t.Fatalf("Unable to Marshal the listing ")
// 		}
// 		expectedJSON := string(marsh)

// 		if actualJSONResponse != expectedJSON {
// 			t.Errorf("JSON response doesn't match the expected for test number %d with sort=%s, order=%s\nExpected response %s\nActual response = %s\n",
// 				i+1, test.SortBy, test.OrderBy, expectedJSON, actualJSONResponse)
// 		}

// 	}
// 	//time.Sleep(100 * time.Second)
// 	caddy.Stop()
// }
