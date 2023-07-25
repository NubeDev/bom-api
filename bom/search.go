package bom

import (
	"fmt"
)

/*
Search

Returns search result list or [] if no matches or if the search string was ''.

Example https://api.weather.bom.gov.au/v1/locations?search=3130
https://api.weather.bom.gov.au/v1/locations?search=Parkville+VIC

Multiple matching records are returned. If suburb and state are specified then
selecting the first element should be sufficient.

If no match is found an empty list is returned and self.gohash is None.

data:[{geohash, id, name, postcode, state},]

geohash     e.g. 'r1r143n'
id          e.g. 'Parkville (Vic.)-r1r143n'
name        e.g. 'Parkville'
state       e.g. 'VIC'

Location is returned as a 6 character precision geohash such as '1r143n'.
(https://en.wikipedia.org/wiki/Geohash)
*/

func (inst *Client) SearchByTown(town, state string) (*Search, string, error) {
	url := fmt.Sprintf("/locations?search=%s+%s", town, state)
	resp, err := FormatRestyResponse(inst.Rest.R().
		SetResult(&Search{}).
		Get(url))
	if err != nil {
		return nil, "", err
	}
	data := resp.Result().(*Search)
	return data, getGeo(data), nil
}

func (inst *Client) SearchByZip(postCode string) (*Search, string, error) {
	url := fmt.Sprintf("/locations?search=%s", postCode)
	resp, err := FormatRestyResponse(inst.Rest.R().
		SetResult(&Search{}).
		Get(url))
	if err != nil {
		return nil, "", err
	}
	data := resp.Result().(*Search)
	return data, getGeo(data), nil
}

func getGeo(data *Search) string {
	var geoHash string
	if data != nil {
		for i, d := range data.Data {
			if i == 0 {
				geoHash = d.Geohash // get first geoHash
				if len(geoHash) > 6 {
					geoHash = geoHash[0:6]
				}
			}
		}
	}
	return geoHash
}
