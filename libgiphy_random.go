// Copyright 2017 Martin Albrecht <martin.albrecht@javacoffee.de>. All rights reserved.
// Use of this source code is governed by the MIT-license
// that can be found in the LICENSE file.

//
// Random endpoint implementation of the giphy.com API
// More information: https://github.com/Giphy/GiphyAPI#random-endpoint
//

package libgiphy

import (
    "encoding/json"
    "net/url"
)

//
// Data type for API data
//
type giphyDataRandom struct {
    Data struct {
         Type, Id, Url, Image_original_url, Image_url,
         Image_mp4_url, Image_frames, Image_width,
         Image_height, Fixed_height_downsampled_url,
         Fixed_height_downsampled_width, Fixed_height_downsampled_height,
         Fixed_width_downsampled_url, Fixed_width_downsampled_width,
         Fixed_width_downsampled_height, Fixed_height_small_url,
         Fixed_height_small_still_url, Fixed_height_small_width,
         Fixed_height_small_height, Fixed_width_small_url,
         Fixed_width_small_still_url, Fixed_width_small_width,
         Fixed_width_small_height string
    }

    Meta struct {
        Status int
        Msg string
    }
}


func (g * Giphy) GetRandom(tag string) (*giphyDataRandom, error) {
    body, err := g._fetch(g._buildUrl("random") + "&tag=" + url.QueryEscape(tag))
    if err != nil {
        return nil, err
    }

    var data giphyDataRandom
    err = json.Unmarshal(body, &data)
    if err != nil {
        return nil, err
    }

    return &data, nil
}