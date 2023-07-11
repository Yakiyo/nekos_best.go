package nekos_best

import (
	"fmt"
	"testing"
)

func TestFetchMulti(t *testing.T) {
	length := 3
	res, err := FetchMany("neko", length)
	if err != nil {
		t.Fatalf("Received error %v", err)
	}
	if len(res) != length {
		t.Fatalf("Response result mismatched with expected length. Receive %v responses", fmt.Sprint(len(res)))
	}
}

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
