package nekos_best

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Fetch from a category
func Fetch(cat string) (NBResponse, error) {
	if !isValidCategory(cat) {
		return NBResponse{}, fmt.Errorf("categories %v is not valid. Must be one of %v", cat, categories)
	}
	res, err := http.Get("https://nekos.best/api/v2/" + cat)
	if err != nil {
		return NBResponse{}, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	
	if err != nil {
		return NBResponse{}, err
	}
	r := &fullNBResponse{}
	json.Unmarshal(bytes, r)

	return r.Results[0], nil
}
