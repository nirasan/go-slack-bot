package app

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAppHandler_ServeHTTP(t *testing.T) {
	ts := httptest.NewServer(NewAppHandler())
	defer ts.Close()

	body := `{
		"token": "Jhj5dZrVaK7ZwHHjRyZWjbDl",
		"challenge": "3eZbrw1aBm2rZgRNFdxV2595E9CY3gmdALWMmHkvFXO7tYXAYM8P",
		"type": "url_verification"
	}`

	res, err := http.Post(ts.URL, "application/json", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != "3eZbrw1aBm2rZgRNFdxV2595E9CY3gmdALWMmHkvFXO7tYXAYM8P" {
		t.Errorf("failed to response verification: %+v", res)
	}

	t.Logf("%+v", res)
}
