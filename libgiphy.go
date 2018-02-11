// Copyright 2017 Martin Albrecht <martin.albrecht@javacoffee.de>. All rights reserved.
// Use of this source code is governed by the MIT-license
// that can be found in the LICENSE file.

package libgiphy

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
    "fmt"
    "errors"
)

const (
    BASE_URL = "http://api.giphy.com/v1/gifs"
)

// gipyImageData is the basic image data struct
type gipyImageData struct {
    Url, Width, Height string
}

// gipyImageSute is the basic image data struct
// with additional size field
type gipyImageDataSized struct {
    Url, Width, Height, Size string
}

// gipyImageDataExtended is the extended image data struct
// with additional video fields
type gipyImageDataExtended struct {
    Url, Width, Height, Size, Mp4, Mp4_size, Webp, Webp_size string
}

// gipyImageDataExtendedFrames is the extended image data struct
// with additional frames field
type gipyImageDataExtendedFrames struct {
    Url, Width, Height, Size, Frames, Mp4, Mp4_size, Webp, Webp_size string
}

// giphyDataArray is a struct holding multiple API result entries
type giphyDataArray struct {
    Data []struct {
        Type, Id, Slug, Url, Bitly_gif_url, Bitly_url, Embed_url,
        Username, Source, Rating, Caption, Content_url, Source_tld,
        Source_post_url, Import_datetime, Trending_datetime string

        Images struct {
            Fixed_height gipyImageDataExtended
            Fixed_height_still gipyImageData
            Fixed_height_downsampled gipyImageDataExtended
            Fixed_width gipyImageDataExtended
            Fixed_width_still gipyImageData
            Fixed_width_downsampled gipyImageDataExtended
            Fixed_height_small gipyImageDataExtended
            Fixed_height_small_still gipyImageData
            Fixed_width_small gipyImageDataExtended
            Fixed_width_small_still gipyImageData
            Downsized gipyImageDataSized
            Downsized_still gipyImageData
            Downsized_large gipyImageDataSized
            Original gipyImageDataExtendedFrames
            Original_still gipyImageData
        }

        Meta struct {
            Status int
            Msg string
        }

        Pagination struct {
            Total_count, Count, Offset int
        }
    }
}

// giphyDataArray is a struct holding a single API result entry
type giphyDataSingle struct {
    Data struct {
        Type, Id, Slug, Url, Bitly_gif_url, Bitly_url, Embed_url,
        Username, Source, Rating, Caption, Content_url, Source_tld,
        Source_post_url, Import_datetime, Trending_datetime string

        Images struct {
            Fixed_height gipyImageDataExtended
            Fixed_height_still gipyImageData
            Fixed_height_downsampled gipyImageDataExtended
            Fixed_width gipyImageDataExtended
            Fixed_width_still gipyImageData
            Fixed_width_downsampled gipyImageDataExtended
            Fixed_height_small gipyImageDataExtended
            Fixed_height_small_still gipyImageData
            Fixed_width_small gipyImageDataExtended
            Fixed_width_small_still gipyImageData
            Downsized gipyImageDataSized
            Downsized_still gipyImageData
            Downsized_large gipyImageDataSized
            Original gipyImageDataExtendedFrames
            Original_still gipyImageData
        }

        Meta struct {
            Status int
            Msg string
        }

        Pagination struct {
            Total_count, Count, Offset int
        }
    }
}


// General Giphy class
type Giphy struct {
    apiKey string
}


// Build a API URL based on given action
func (g * Giphy) _buildUrl(action string) string {
    if action != "" && action[0] != '/' {
        action = "/" + action
    }
    return BASE_URL + action + "?api_key=" + g.apiKey
}

// Fetch API data from given URL
func (g * Giphy) _fetch(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, errors.New(fmt.Sprintf("Error HTTP status %d: %s", resp.StatusCode, resp.Status))
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}

// _parseDataSingle is a function to parse JSON data for single data
// entries from byte array
func (g * Giphy) _parseDataSingle(body []byte) (*giphyDataSingle, error) {
    var data giphyDataSingle

    if len(body) <= 0 {
        return nil, errors.New("_parseDataArray: No data in response body to parse.")
    }

    err := json.Unmarshal(body, &data)
    if err != nil {
        return nil, err
    }

    return &data, nil
}

// _parseDataSingle is a function to parse JSON data for multiple data
// entries from byte array
func (g * Giphy) _parseDataArray(body []byte) (*giphyDataArray, error) {
    var data giphyDataArray

    if len(body) <= 0 {
        return nil, errors.New("_parseDataArray: No data in response body to parse.")
    }

    err := json.Unmarshal(body, &data)
    if err != nil {
        return nil, err
    }

    return &data, nil
}

// Constructor function
func NewGiphy(apiKey string) * Giphy {
    return &Giphy{
        apiKey: apiKey,
    }
}

// GetById fetches GIF image information based on given ID string
// More information: https://github.com/Giphy/GiphyAPI#get-gif-by-id-endpoint
func (g * Giphy) GetById(id string) (*giphyDataSingle, error) {
    body, err := g._fetch(g._buildUrl(id))
    if err != nil {
        return nil, err
    }

    return g._parseDataSingle(body)
}

// GetById fetches GIF image information for multiple ids based
// on given ID string array
// More information: https://github.com/Giphy/GiphyAPI#get-gifs-by-id-endpoint
func (g * Giphy) GetByIds(ids []string) (*giphyDataArray, error) {
    body, err := g._fetch(g._buildUrl("") + "&ids=" + strings.Join(ids, ","))
    if err != nil {
        return nil, err
    }

    return g._parseDataArray(body)
}