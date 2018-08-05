package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Useless-Double-E/weather/cmd"
	"github.com/Useless-Double-E/weather/constants"
)

func main() {
	var citylong = flag.String("city", "", "City for weather request")
	var cityshort = flag.String("c", "", "City for weather request")
	var ziplong = flag.String("zip", "", "Zip code for weather request")
	var zipshort = flag.String("z", "", "Zip code for weather request")

	flag.Parse()
	cfg, err := cmd.Parseflags(map[string]string{
		constants.CityShort: *cityshort,
		constants.CityLong:  *citylong,
		constants.ZipShort:  *zipshort,
		constants.ZipLong:   *ziplong,
	})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(cfg)
}
