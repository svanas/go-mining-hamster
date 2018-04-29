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

func (self *Signal) BaseCurrency() string {
	return strings.Split(self.Market, "-")[1]
}

func (self *Signal) QuoteCurrency() string {
	return strings.Split(self.Market, "-")[0]
}

func (self *Signal) UnmarshalJSON(data []byte) error {
	var err error
	type Alias Signal
	aux := &struct {
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(self),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var loc *time.Location
	if loc, err = time.LoadLocation("Europe/Vienna"); err != nil {
		return err
	}
	if self.Time, err = time.ParseInLocation("2006-01-02 15:04:05", aux.Time, loc); err != nil {
		return err
	}
	return nil
}

type (
	Signals []Signal
)

func (self Signals) IndexOf(signal *Signal) int {
	for i, s := range self {
		if s == *signal {
			return i
		}
	}
	return -1
}
