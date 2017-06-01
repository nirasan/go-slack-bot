package app

import (
	"strings"
	"testing"
)

func TestDecodeJSON_urlVerification(t *testing.T) {
	payload := `
	{
    "token": "Jhj5dZrVaK7ZwHHjRyZWjbDl",
    "challenge": "3eZbrw1aBm2rZgRNFdxV2595E9CY3gmdALWMmHkvFXO7tYXAYM8P",
    "type": "url_verification"
	}`
	data, err := DecodeJSON(strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	if data.Type() != "url_verification" {
		t.Errorf("failed to decode: %+v", data)
	}
	t.Logf("%+v", data)
}
