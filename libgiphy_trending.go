// Copyright 2017 Martin Albrecht <martin.albrecht@javacoffee.de>. All rights reserved.
// Use of this source code is governed by the MIT-license
// that can be found in the LICENSE file.

//
// Trending endpoint implementation of the giphy.com API
// More information: https://github.com/Giphy/GiphyAPI#trending-gifs-endpoint
//

package libgiphy

import (
    "encoding/json"
)

type giphyDataTrending struct {
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



func (g * Giphy) GetTrending() (*giphyDataTrending, error) {
    body, err := g.fetch(g.buildUrl("trending"))
    if err != nil {
        return nil, err
    }

    var data giphyDataTrending
    err = json.Unmarshal(body, &data)
    if err != nil {
        return nil, err
    }

    return &data, nil;
}
