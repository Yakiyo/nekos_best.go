package nekos_best

import (
	"fmt"
	"io"
	"net/http"
)

// Fetch from a category
func Fetch(cat string) (string, error) {
	res, err := http.Get("https://nekos.best/api/v2/" + cat)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	fmt.Println(body)
	return "", nil
}
