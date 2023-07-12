# nekos_best.go
[![Go Reference](https://pkg.go.dev/badge/github.com/Yakiyo/nekos_best.go.svg)](https://pkg.go.dev/github.com/Yakiyo/nekos_best.go) [![ci](https://github.com/Yakiyo/nekos_best.go/actions/workflows/ci.yml/badge.svg)](https://github.com/Yakiyo/nekos_best.go/actions/workflows/ci.yml)

Go module for nekos.best

## Getting started
Add the module to your project
```shell
$ go get -u github.com/Yakiyo/nekos_best.go
```

## Usage
Import it in your code
```go
import (
    nb github.com/Yakiyo/nekos_best.go
)
```
Full reference: https://pkg.go.dev/github.com/Yakiyo/nekos_best.go

### Fetch a single entry
```go
res, err := nb.Fetch("neko")
if err != nil {
    // handle err
}
fmt.Println(res.Url, res.Artist_name, res.Artist_href)
```

### Fetch multiple entries
```go
res, _ := nb.FetchMany("baka", 3)

fmt.Println(res[0].Url)
```

### Fetch a file
```go
import "os"

res, _ := nb.FetchFile("pat")

os.WriteFile("image.png", res.Data, 0644)

// print associated entry for the image
fmt.Println(res.Body)
```

### Search for an entry
```go
res, _ := nb.Search("Gochuumon wa Usagi Desuka??", "baka", 3)

fmt.Println(res)
```

## Author

**nekos_best.go** © [Yakiyo](https://github.com/Yakiyo). Authored and maintained by Yakiyo.

Released under [MIT](https://opensource.org/licenses/MIT) License