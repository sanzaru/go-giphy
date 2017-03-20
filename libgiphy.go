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

// Constructor function
func NewGiphy(apiKey string) * Giphy {
    return &Giphy{
        apiKey: apiKey,
    }
}

// Build a API URL based on given action
func (g * Giphy) buildUrl(action string) string {
    return BASE_URL + action + "?api_key=" + g.apiKey
}

// Fetch API data from given URL
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