// Copyright 2017 Martin Albrecht <martin.albrecht@javacoffee.de>. All rights reserved.
// Use of this source code is governed by the MIT-license
// that can be found in the LICENSE file.

//
// Translate endpoint implementation of the giphy.com API
// More information: https://github.com/Giphy/GiphyAPI#translate-endpoint
//

package libgiphy

import (
    "net/url"
)

func (g * Giphy) GetTranslate(term, rating, lang string, doFmt bool) (*giphyDataSingle, error) {
    apiUrl := g._buildUrl("translate") + "&s=" + url.QueryEscape(term);

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

    return g._parseDataSingle(body)
}
