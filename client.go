package nekos_best

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// Fetch multiple results from a category
func FetchMany(category string, amount int) ([]NBResponse, error) {
	if !isValidCategory(category) {
		return []NBResponse{}, fmt.Errorf("categories %v is not valid. Must be one of %v", category, categories)
	}
	res, err := http.Get("https://nekos.best/api/v2/" + category + "?amount=" + fmt.Sprint(amount))
	if err != nil {
		return []NBResponse{}, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return []NBResponse{}, err
	}
	r := &fullNBResponse{}
	json.Unmarshal(bytes, r)

	return r.Results, nil
}

// Fetch a single result from a category
func Fetch(category string) (NBResponse, error) {
	res, err := FetchMany(category, 1)
	if err != nil {
		return NBResponse{}, err
	}
	return res[0], nil
}

// Fetch a random image file from a category
func FetchFile(category string) (NBBufferResponse, error) {
	res, err := Fetch(category)
	if err != nil {
		return NBBufferResponse{}, err
	}
	resp, err := http.Get(res.Url)
	if err != nil {
		return NBBufferResponse{}, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return NBBufferResponse{}, err
	}
	return NBBufferResponse{
		Data: bytes,
		Body: res,
	}, nil
}

// Random nekos.best category
func RandomCategory() string {
	n := rand.New(rand.NewSource(time.Now().Unix())).Intn(len(categories))
	return categories[n]
}
