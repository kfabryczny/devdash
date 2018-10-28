package main

import (
	"bytes"
	"fmt"

	"github.com/Phantas0s/devdash/internal"
	"github.com/spf13/viper"
)

const (
	// services
	ga = "ga"
	// keys
	kquit = "C-c"
)

type config struct {
	General  General   `mapstructure "general"`
	Projects []Project `mapstructure "projects"`
}

type General struct {
	Keys    map[string]string
	Refresh int64 `mapstructure:"refresh"`
}

type Project struct {
	Name     string   `mapstructure:"name"`
	Services Services `mapstructure "services"`
	Widgets  []Row    `mapstructure "widgets"`
}

type Row struct {
	Row []Column `mapstructure "row"`
}

type Column struct {
	Col []Widgets `mapstructure "Col"`
}

type Widgets struct {
	Size     string            `json:"size"`
	Elements []internal.Widget `json:"elements"`
}

type Services struct {
	GoogleAnalytics GoogleAnalytics `mapstructure:"google_analytics"`
}

type GoogleAnalytics struct {
	Keyfile string `mapstructure:"keyfile"`
	ViewID  string `mapstructure:"view_id"`
}

func mapConfig(data []byte) config {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		panic(fmt.Errorf("could not read config data %s: %s", string(data), err))
	}

	var cfg config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return cfg
}

func (c config) KQuit() string {
	if ok := c.General.Keys["quit"]; ok != "" {
		return c.General.Keys["quit"]
	}

	return kquit
}
