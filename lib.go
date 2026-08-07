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

var client = &http.Client{}
var userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/151.0.0.0 Safari/537.36"

// Fetch multiple results from a category
func FetchMany(category string, amount int) ([]NBResponse, error) {
	if !isValidCategory(category) {
		return []NBResponse{}, fmt.Errorf("categories %v is not valid. Must be one of %v", category, categories)
	}
	if amount < 1 || amount > 20 {
		return []NBResponse{}, fmt.Errorf("amount must be between 1 and 20")
	}

	req, err := http.NewRequest("GET", "https://nekos.best/api/v2/" + category + "?amount=" + fmt.Sprint(amount), nil)
	req.Header.Set("User-Agent", userAgent)

    res, err := client.Do(req)
	if err != nil {
		fmt.Printf("error while fetching neko: %v\n", err)
		return []NBResponse{}, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return []NBResponse{}, err
	}
	r := &fullNBResponse{}
	json.Unmarshal(bytes, r)

	fmt.Printf("Response: %v\n", res)
	fmt.Printf("Results: %v\n", r.Results)

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

	req, err := http.NewRequest("GET", fmt.Sprintf("https://nekos.best/api/v2/%v?%v", category, params.Encode()), nil)
	req.Header.Set("User-Agent", userAgent)

	res, err := client.Do(req)
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
