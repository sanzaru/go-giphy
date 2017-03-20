# go-giphy

## Description
giphy.com API client implementation in Go programming language.

## Installation

    go get github.com/sanzaru/go-giphy
    
## Supported endpoints

1. Search
2. Trending
3. Translate
4. Random

## Examples

### Search
Search all Giphy GIFs for a word or phrase. Punctuation will be stripped and ignored. Use a plus or url encode for phrases.

Function parameters:
* query (string): The search query
* limit (int): Limit of entries to fetch. Set to -1 for no limit.
* offset (int): Offset of entries to fetch. Set to -1 for no offset.
* rating (string): Limit results to those rated (y,g, pg, pg-13 or r).
* lang - (string) Specify default country for regional content; format is 2-letter ISO 639-1 country code.
* fmt - (bool) Return results in html or json format (useful for viewing responses as GIFs to debug/test)

More information: https://github.com/Giphy/GiphyAPI#search-endpoint

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

More information: https://github.com/Giphy/GiphyAPI#trending-gifs-endpoint

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
The translate API draws on search, but uses the Giphy "special sauce" to handle translating from one 
vocabulary to another. In this case, words and phrases to GIFs. Example implementations of translate 
can be found in the Giphy Slack, Hipchat, Wire, or Dasher integrations. 

Function parameters:
* term (string): The term to translate
* rating (string): Limit results to those rated (y,g, pg, pg-13 or r).
* lang - (string) Specify default country for regional content; format is 2-letter ISO 639-1 country code.
* fmt - (bool) Return results in html or json format (useful for viewing responses as GIFs to debug/test)

More information: https://github.com/Giphy/GiphyAPI#translate-endpoint

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
Returns a random GIF, limited by tag. Excluding the tag parameter will return a random GIF from the Giphy catalog.

Function parameters:
* tag (string): The GIF tag to limit randomness by

More information: https://github.com/Giphy/GiphyAPI#random-endpoint

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

### Get single ID

Returns meta data about a GIF, by GIF id. In the below example, the GIF ID is "feqkVgjJpYtjy"

Function parameters:
* id (string): ID string of the GIF to search

More information: https://github.com/Giphy/GiphyAPI#get-gif-by-id-endpoint

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
    
    dataIdSingle, err := giphy.GetById("feqkVgjJpYtjy");
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("Id single data: %+v\n", dataIdSingle)
}
```

### Get multiple IDs

A multiget version of the get GIF by ID endpoint. In this case the IDs are feqkVgjJpYtjy and 3NtY188QaxDdC. 
Note the additional user metadata attached to the document that describes the second GIF in the response, 3NtY188QaxDdC.

Function parameters:
* ids ([]string): Array of strings with ids of the GIFs to search

More information: https://github.com/Giphy/GiphyAPI#get-gifs-by-id-endpoint

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
    
    ids := []string{
        "feqkVgjJpYtjy",
        "3NtY188QaxDdC",
    }

    dataIdMulti, err := giphy.GetByIds(ids);
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("Id multi data: %+v\n", dataIdMulti)
}
```
