# go-giphy

## Description
giphy.com API implementation in Go

## Installation

    go get github.com/sanzaru/go-giphy
    
## Supported endpoints

1. Search (https://github.com/Giphy/GiphyAPI#search-endpoint)
2. Trending (https://github.com/Giphy/GiphyAPI#trending-gifs-endpoint)
3. Translate (https://github.com/Giphy/GiphyAPI#translate-endpoint)
4. Random (https://github.com/Giphy/GiphyAPI#random-endpoint)

## Examples

### Search
```go
package main

import (
    "fmt"
    "github.com/sanzaru/go-giphy"
)

const (
    API_KEY = "YOUR_API_KEY"
)

func main() {
    giphy := libgiphy.NewGiphy(API_KEY)
    
    dataSearch, err := giphy.GetSearch("fun", -1, -1, "", "", false);
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("Search data: %+v\n", dataSearch)
}
```

### Trending
```go
package main

import (
    "fmt"
    "github.com/sanzaru/go-giphy"
)

const (
    API_KEY = "YOUR_API_KEY"
)

func main() {
    giphy := libgiphy.NewGiphy(API_KEY)
    
    dataTrending, err := giphy.GetTrending();
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("Trending data: %+v\n", dataTrending)
}
```

### Translate
```go
package main

import (
    "fmt"
    "github.com/sanzaru/go-giphy"
)

const (
    API_KEY = "YOUR_API_KEY"
)

func main() {
    giphy := libgiphy.NewGiphy(API_KEY)
    
    dataTranslate, err := giphy.GetTranslate("hello", "", "", false);
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("Translate data: %+v\n", dataTranslate)
}
```

### Random
```go
package main

import (
    "fmt"
    "github.com/sanzaru/go-giphy"
)

const (
    API_KEY = "YOUR_API_KEY"
)

func main() {
    giphy := libgiphy.NewGiphy(API_KEY)
    
    dataRandom, err := giphy.GetRandom("very funny")
    if err != nil {
        fmt.Println("error:", err)
    }
    
    fmt.Printf("Random data: %+v\n", dataRandom)
}
```
