package nekos_best

import (
	"fmt"
	"os"
	"testing"
)

func TestFetchMulti(t *testing.T) {
	length := 3
	res, err := FetchMany("neko", length)
	handleErr(err, t)
	if len(res) != length {
		t.Fatalf("Response result mismatched with expected length. Receive %v responses", fmt.Sprint(len(res)))
	}
}

func TestFetchImage(t *testing.T) {
	res, err := Fetch("neko")
	handleErr(err, t)
	if res.Url == "" || res.Artist_name == "" {
		t.Fatalf("Missing required properties in response. Received %v", res)
	}
}

func TestFetchGif(t *testing.T) {
	res, err := Fetch("baka")
	handleErr(err, t)
	if res.Url == "" || res.Artist_name != "" {
		t.Fatalf("Wrong type of properties in response. Received %v", res)
	}
}

func TestFetchFile(t *testing.T) {
	res, err := FetchFile("neko")
	handleErr(err, t)
	if len(res.Data) <= 0 {
		t.Fatalf("Image body is empty")
	}
	os.WriteFile("test.png", res.Data, 0644)
}

func TestRandomCat(t *testing.T) {
	cat := RandomCategory()
	if !isValidCategory(cat) {
		t.Fatalf("%v is not a valid category", cat)
	}
}

func handleErr(e error, t *testing.T) {
	if e != nil {
		t.Fatalf("Received error %v", e)
	}
}
