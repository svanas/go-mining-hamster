package mininghamster

import (
	"encoding/json"
	"time"
)

type Signal struct {
	Market     string    `json:"market"`
	LastPrice  float64   `json:"lastprice,string"`
	SignalMode string    `json:"signalmode"`
	Exchange   string    `json:"exchange"`
	BaseVolume float64   `json:"basevolume"`
	Time       time.Time `json:"time"`
}

func (s *Signal) UnmarshalJSON(data []byte) error {
	var err error
	type Alias Signal
	aux := &struct {
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if s.Time, err = time.Parse("2006-01-02 15:04:05", aux.Time); err != nil {
		return err
	}
	return nil
}
