package headers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mholt/caddy/middleware"
)

func TestHeaders(t *testing.T) {
	for i, test := range []struct {
		from  string
		name  string
		value string
	}{
		{"/a", "Foo", "Bar"},
		{"/a", "Bar", ""},
		{"/a", "Baz", ""},
		{"/b", "Foo", ""},
		{"/b", "Bar", "Removed in /a"},
	} {
		he := Headers{
			Next: middleware.HandlerFunc(func(w http.ResponseWriter, r *http.Request) (int, error) {
				return 0, nil
			}),
			Rules: []Rule{
				{Path: "/a", Headers: []Header{
					{Name: "Foo", Value: "Bar"},
					{Name: "-Bar"},
				}},
			},
		}

		req, err := http.NewRequest("GET", test.from, nil)
		if err != nil {
			t.Fatalf("Test %d: Could not create HTTP request: %v", i, err)
		}

		rec := httptest.NewRecorder()
		rec.Header().Set("Bar", "Removed in /a")

		he.ServeHTTP(rec, req)

		if got := rec.Header().Get(test.name); got != test.value {
			t.Errorf("Test %d: Expected %s header to be %q but was %q",
				i, test.name, test.value, got)
		}
	}
}

func BenchmarkHeaders(b *testing.B) {
	he := Headers{
		Next: middleware.HandlerFunc(func(w http.ResponseWriter, r *http.Request) (int, error) {
			return 0, nil
		}),
		Rules: []Rule{
			{Path: "/a", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "-Bar"},
			}},
		},
	}

	req, err := http.NewRequest("GET", "/a", nil)
	if err != nil {
		b.Fatalf(" Could not create HTTP request")
	}

	rec := httptest.NewRecorder()
	rec.Header().Set("Bar", "Removed in /a") //this is to make sure that the header remove ops is also considered, see the rule -Bar
	for n := 0; n < b.N; n++ {
		_, err := he.ServeHTTP(rec, req)
		if err != nil {
			b.Fatal(err.Error())
		}
	}

}

func BenchmarkHeadersWithMoreRules(b *testing.B) { //To see the performance difference with more rules
	he := Headers{
		Next: middleware.HandlerFunc(func(w http.ResponseWriter, r *http.Request) (int, error) {
			return 0, nil
		}),
		Rules: []Rule{
			{Path: "/a", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/b", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/c", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/d", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/e", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/f", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/g", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/h", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/i", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/j", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/k", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/l", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
			{Path: "/m", Headers: []Header{
				{Name: "Foo", Value: "Bar"},
				{Name: "Foo1", Value: "Bar"},
				{Name: "Foo2", Value: "Bar"},
				{Name: "Foo3", Value: "Bar"},
				{Name: "Foo4", Value: "Bar"},
				{Name: "-Bar"},
			}},
		},
	}

	req, err := http.NewRequest("GET", "/a", nil)
	if err != nil {
		b.Fatalf(" Could not create HTTP request")
	}

	rec := httptest.NewRecorder()
	rec.Header().Set("Bar", "Removed in /a") //this is to make sure that the header remove ops is also considered, see the rule -Bar
	for n := 0; n < b.N; n++ {
		_, err := he.ServeHTTP(rec, req)
		if err != nil {
			b.Fatal(err.Error())
		}
	}

}
