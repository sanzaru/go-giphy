// Copyright 2017 Martin Albrecht <martin.albrecht@javacoffee.de>. All rights reserved.
// Use of this source code is governed by the MIT-license
// that can be found in the LICENSE file.

package libgiphy

import (
    "net/http"
    "io/ioutil"
)

const (
    BASE_URL = "http://api.giphy.com/v1/gifs/"
)


//
// Image data types
//
type gipyImageData struct {
    Url, Width, Height string
}

type gipyImageDataSized struct {
    Url, Width, Height, Size string
}

type gipyImageDataExtended struct {
    Url, Width, Height, Size, Mp4, Mp4_size, Webp, Webp_size string
}

type gipyImageDataExtendedFrames struct {
    Url, Width, Height, Size, Frames, Mp4, Mp4_size, Webp, Webp_size string
}


//
// General Giphy class
//
type Giphy struct {
    apiKey string
}


func (g * Giphy) buildUrl(action string) string {
    return BASE_URL + action + "?api_key=" + g.apiKey
}


func (g * Giphy) fetch(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}


func NewGiphy(apiKey string) * Giphy {
    return &Giphy{
        apiKey: apiKey,
    }
}


/*func (g * Giphy) GetTrending() (*giphyData, error) {
    body, err := g.getData(g.buildUrl("trending") + "&limit=1")
    if err != nil {
        return nil, err
    }

    return g.parseData(body);
}*/
