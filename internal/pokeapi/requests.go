package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetMaps(url string) (error, Locations) {

	res, err := http.Get(url)
	if err != nil {
		return err, Locations{}
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err, Locations{}
	}
	locations := Locations{}

	if err := json.Unmarshal(data, &locations); err != nil {
		return err, Locations{}
	}

	return nil, locations

}
