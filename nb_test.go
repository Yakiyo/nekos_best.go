package nekos_best

import (
	"testing"
)

func TestFetchImage(t *testing.T) {
	res, err := Fetch("neko")
	if err != nil {
		t.Fatalf("Received error %v", err)
	}
	if res.Url == "" || res.Artist_name == "" {
		t.Fatalf("Missing required properties in response. Received %v", res)
	}
}

func TestFetchGif(t *testing.T) {
	res, err := Fetch("baka")
	if err != nil {
		t.Fatalf("Received error %v", err)
	}
	if res.Url == "" || res.Artist_name != "" {
		t.Fatalf("Wrong type of properties in response. Received %v", res)
	}
}
