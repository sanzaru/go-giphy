// Copyright 2017 Martin Albrecht <martin.albrecht@javacoffee.de>. All rights reserved.
// Use of this source code is governed by the MIT-license
// that can be found in the LICENSE file.

//
// Search endpoint implementation of the giphy.com API
// More information: https://github.com/Giphy/GiphyAPI#search-endpoint
//

package libgiphy

import (
    "net/url"
    "strconv"
)

func (g * Giphy) GetSearch(query string, limit, offset int, rating, lang string, doFmt bool) (*giphyDataArray, error) {
    apiUrl := g._buildUrl("search") + "&q=" + url.QueryEscape(query);

    if limit > 0 {
        apiUrl += "&limit=" + strconv.Itoa(limit)
    }

    if offset > 0 {
        apiUrl += "&offset=" + strconv.Itoa(offset)
    }

    if rating != "" {
        apiUrl += "&rating=" + rating
    }

    if lang != "" {
        apiUrl += "&lang=" + lang
    }

    if doFmt == true {
        apiUrl += "&fmt"
    }

    body, err := g._fetch(apiUrl)
    if err != nil {
        return nil, err
    }

    return g._parseDataArray(body)
}
