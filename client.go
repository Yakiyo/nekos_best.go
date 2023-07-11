package nekos_best

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Fetch multiple results from a category
func FetchMany(cat string, amount int) ([]NBResponse, error) {
	if !isValidCategory(cat) {
		return []NBResponse{}, fmt.Errorf("categories %v is not valid. Must be one of %v", cat, categories)
	}
	res, err := http.Get("https://nekos.best/api/v2/" + cat + "?amount=" + fmt.Sprint(amount))
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
func Fetch(cat string) (NBResponse, error) {
	res, err := FetchMany(cat, 1)
	if err != nil {
		return NBResponse{}, err
	}
	return res[0], nil
}
