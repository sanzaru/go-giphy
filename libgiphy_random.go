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

func (g *Giphy) GetRandom(tag string) (*giphyDataSingle, error) {
	body, err := g._fetch(g._buildUrl("random") + "&tag=" + url.QueryEscape(tag))
	if err != nil {
		return nil, err
	}

	var data giphyDataSingle
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
