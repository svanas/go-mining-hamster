package mininghamster

import (
	"encoding/json"
	"strings"
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

func (signal *Signal) BaseCurrency() string {
	return strings.Split(signal.Market, "-")[1]
}

func (signal *Signal) QuoteCurrency() string {
	return strings.Split(signal.Market, "-")[0]
}

func (signal *Signal) UnmarshalJSON(data []byte) error {
	var err error
	type Alias Signal
	aux := &struct {
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(signal),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var loc *time.Location
	if loc, err = time.LoadLocation("CET"); err != nil {
		return err
	}
	if signal.Time, err = time.ParseInLocation("2006-01-02 15:04:05", aux.Time, loc); err != nil {
		return err
	}
	return nil
}

type (
	Signals []Signal
)

func (signals Signals) IndexOf(signal *Signal) int {
	for i, s := range signals {
		if (s.Market == signal.Market) && (s.SignalMode == signal.SignalMode) && (s.Exchange == signal.Exchange) {
			return i
		}
	}
	return -1
}
