package api

import (
	"flag"
	"log"
	"strconv"

	"where-to-eat/configuration"
	"where-to-eat/model"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

// PlaceAPI represents a place API
type PlaceAPI struct{}

var (
	apiKey    = flag.String("key", "", "API Key for using Google Maps API.")
	clientID  = flag.String("client_id", "", "ClientID for Maps for Work API access.")
	signature = flag.String("signature", "", "Signature for Maps for Work API access.")
	location  = flag.String("location", "", "The latitude/longitude around which to retrieve place information. This must be specified as latitude,longitude.")
	radius    = flag.Uint("radius", 0, "Defines the distance (in meters) within which to bias place results. The maximum allowed radius is 50,000 meters.")
	keyword   = flag.String("keyword", "", "A term to be matched against all content that Google has indexed for this place, including but not limited to name, type, and address, as well as customer reviews and other third-party content.")
	language  = flag.String("language", "", "The language in which to return results.")
	minPrice  = flag.String("minprice", "", "Restricts results to only those places within the specified price level.")
	maxPrice  = flag.String("maxprice", "", "Restricts results to only those places within the specified price level.")
	name      = flag.String("name", "", "One or more terms to be matched against the names of places, separated with a space character.")
	openNow   = flag.Bool("open_now", false, "Restricts results to only those places that are open for business at the time the query is sent.")
	rankBy    = flag.String("rankby", "", "Specifies the order in which results are listed. Valid values are prominence or distance.")
	placeType = flag.String("type", "", "Restricts the results to places matching the specified type.")
	pageToken = flag.String("pagetoken", "", "Set to retrieve the next page of results.")
)

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

// LocationString returns a string in a format that the api needs
func LocationString(l model.Location) string {
	locationString := strconv.FormatFloat(l.Latitude, 'f', 7, 64)
	locationString += ","
	locationString += strconv.FormatFloat(l.Longitude, 'f', 7, 64)

	return locationString
}

// FindFood returns places based on a location
func FindFood(locations model.Location) maps.PlacesSearchResponse {
	var err error
	config, err := configuration.Config("conf.json")
	check(err)

	*apiKey = config.APIKey
	*location = LocationString(locations)
	*radius = 2000
	*placeType = "meal_takeaway"

	var client *maps.Client

	if *apiKey != "" {
		client, err = maps.NewClient(maps.WithAPIKey(*apiKey))
	} else if *clientID != "" || *signature != "" {
		client, err = maps.NewClient(maps.WithClientIDAndSignature(*clientID, *signature))
	}

	check(err)

	r := &maps.NearbySearchRequest{
		Radius:   *radius,
		Keyword:  *keyword,
		Language: *language,
		Name:     *name,
		OpenNow:  *openNow,
	}

	parseLocation(*location, r)
	parsePriceLevels(*minPrice, *maxPrice, r)
	parseRankBy(*rankBy, r)
	parsePlaceType(*placeType, r)

	resp, err := client.NearbySearch(context.Background(), r)
	check(err)

	return resp
}

func parseLocation(location string, r *maps.NearbySearchRequest) {
	if location != "" {
		l, err := maps.ParseLatLng(location)
		check(err)
		r.Location = &l
	}
}

func parsePriceLevel(priceLevel string) maps.PriceLevel {
	switch priceLevel {
	case "0":
		return maps.PriceLevelFree
	case "1":
		return maps.PriceLevelInexpensive
	case "2":
		return maps.PriceLevelModerate
	case "3":
		return maps.PriceLevelExpensive
	case "4":
		return maps.PriceLevelVeryExpensive
	}
	return maps.PriceLevelFree
}

func parsePriceLevels(minPrice string, maxPrice string, r *maps.NearbySearchRequest) {
	if minPrice != "" {
		r.MinPrice = parsePriceLevel(minPrice)
	}

	if maxPrice != "" {
		r.MaxPrice = parsePriceLevel(minPrice)
	}
}

func parseRankBy(rankBy string, r *maps.NearbySearchRequest) {
	switch rankBy {
	case "prominence":
		r.RankBy = maps.RankByProminence
		return
	case "distance":
		r.RankBy = maps.RankByDistance
		return
	case "":
		return
	}
}

func parsePlaceType(placeType string, r *maps.NearbySearchRequest) {
	if placeType != "" {
		t, err := maps.ParsePlaceType(placeType)
		check(err)

		r.Type = t
	}
}
