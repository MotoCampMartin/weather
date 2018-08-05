package cmd

import (
	"fmt"

	"github.com/Useless-Double-E/weather/constants"
	"github.com/Useless-Double-E/weather/models"
)

func Parseflags(flags map[string]string) (models.Config, error) {
	var cfg models.Config

	citylong := flags[constants.CityLong]
	cityshort := flags[constants.CityShort]
	ziplong := flags[constants.ZipLong]
	zipshort := flags[constants.ZipShort]
	city := shorthandcheck(citylong, cityshort)
	zip := shorthandcheck(ziplong, zipshort)

	if city == "" && zip == "" {
		return cfg, fmt.Errorf("City and Zip Flags Missing")
	}

	cfg.City = city
	cfg.Zip = zip

	return cfg, nil
}

func shorthandcheck(long, short string) string {
	if long != "" && short == "" {
		return long
	}
	if short != "" && long == "" {
		return short
	}
	return ""
}
