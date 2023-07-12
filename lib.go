package nekos_best

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// Fetch multiple results from a category
func FetchMany(category string, amount int) ([]NBResponse, error) {
	if !isValidCategory(category) {
		return []NBResponse{}, fmt.Errorf("categories %v is not valid. Must be one of %v", category, categories)
	}
	if amount < 1 || amount > 20 {
		return []NBResponse{}, fmt.Errorf("amount must be between 1 and 20")
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

// Search using a query
func Search(query string, category string, amount int) ([]NBResponse, error) {
	if !isValidCategory(category) {
		return []NBResponse{}, fmt.Errorf("categories %v is not valid. Must be one of %v", category, categories)
	}
	t := "2"
	if contains(image_categories, category) {
		t = "1"
	}
	params := url.Values{
		"query":    {query},
		"amount":   {fmt.Sprint(amount)},
		"category": {category},
		"type":     {t},
	}
	res, err := http.Get(fmt.Sprintf("https://nekos.best/api/v2/%v?%v", category, params.Encode()))
	if err != nil {
		return []NBResponse{}, err
	}
	if res.StatusCode == 429 {
		remaining := res.Header.Get("x-rate-limit-remaining")
		return []NBResponse{}, fmt.Errorf("api ratelimit. Remaining: %v", remaining)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []NBResponse{}, err
	}
	v := &fullNBResponse{}
	json.Unmarshal(data, v)
	return v.Results, nil
}

// Random nekos.best category
func RandomCategory() string {
	n := rand.New(rand.NewSource(time.Now().Unix())).Intn(len(categories))
	return categories[n]
}
