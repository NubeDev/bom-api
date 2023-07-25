package bom

import "time"

type Search struct {
	Metadata struct {
		ResponseTimestamp time.Time `json:"response_timestamp"`
	} `json:"metadata"`
	Data []struct {
		Geohash  string `json:"geohash"`
		Id       string `json:"id"`
		Name     string `json:"name"`
		Postcode string `json:"postcode"`
		State    string `json:"state"`
	} `json:"data"`
}

type Observations struct {
	Metadata struct {
		ResponseTimestamp time.Time `json:"response_timestamp"`
		IssueTime         time.Time `json:"issue_time"`
		ObservationTime   time.Time `json:"observation_time"`
		Copyright         string    `json:"copyright"`
	} `json:"metadata"`
	Data struct {
		Temp          float64 `json:"temp"`
		TempFeelsLike float64 `json:"temp_feels_like"`
		Wind          struct {
			SpeedKilometre int    `json:"speed_kilometre"`
			SpeedKnot      int    `json:"speed_knot"`
			Direction      string `json:"direction"`
		} `json:"wind"`
		Gust struct {
			SpeedKilometre int `json:"speed_kilometre"`
			SpeedKnot      int `json:"speed_knot"`
		} `json:"gust"`
		MaxGust struct {
			SpeedKilometre int       `json:"speed_kilometre"`
			SpeedKnot      int       `json:"speed_knot"`
			Time           time.Time `json:"time"`
		} `json:"max_gust"`
		MaxTemp struct {
			Time  time.Time `json:"time"`
			Value float64   `json:"value"`
		} `json:"max_temp"`
		MinTemp struct {
			Time  time.Time `json:"time"`
			Value float64   `json:"value"`
		} `json:"min_temp"`
		RainSince9Am int `json:"rain_since_9am"`
		Humidity     int `json:"humidity"`
		Station      struct {
			BomId    string `json:"bom_id"`
			Name     string `json:"name"`
			Distance int    `json:"distance"`
		} `json:"station"`
	} `json:"data"`
}
