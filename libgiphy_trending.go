// Copyright 2017 Martin Albrecht <martin.albrecht@javacoffee.de>. All rights reserved.
// Use of this source code is governed by the MIT-license
// that can be found in the LICENSE file.

//
// Trending endpoint implementation of the giphy.com API
// More information: https://github.com/Giphy/GiphyAPI#trending-gifs-endpoint
//

package libgiphy

func (g * Giphy) GetTrending() (*giphyDataArray, error) {
    body, err := g._fetch(g._buildUrl("trending"))
    if err != nil {
        return nil, err
    }

    return g._parseDataArray(body)
}
