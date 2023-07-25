package bom

import "fmt"

/*
Weather

http://www.bom.gov.au/places/nsw/helensburgh/

pass in geohash, The geohash string is 6 characters.

observations
Example https://api.weather.bom.gov.au/v1/locations/r1r143/observations

warnings
Example https://api.weather.bom.gov.au/v1/locations/r1r143/warnings

forecast
Example https://api.weather.bom.gov.au/v1/locations/r1r143/forecast/rain

daily
Example https://api.weather.bom.gov.au/v1/locations/r1r143/forecasts/daily
https://api.weather.bom.gov.au/v1/locations/r1r143/forecasts/hourly

*/

func (inst *Client) Observations(geo string) (*Observations, error) {
	url := fmt.Sprintf("/locations/%s/observations", geo)
	resp, err := FormatRestyResponse(inst.Rest.R().
		SetResult(&Observations{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Observations), nil
}

func (inst *Client) ObservationByTown(town, state string) (*Observations, error) {
	_, geo, err := inst.SearchByTown(town, state)
	if err != nil {
		return nil, err
	}
	return inst.Observations(geo)
}
