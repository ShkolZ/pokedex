package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func GetMaps(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func GetPokemons(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting pokemons")
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading")
	}

	return data, nil
}

func GetPokemon(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error with request(GetPokemon)")
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error with reading(GetPokemon)")
	}
	return data, nil
}

//explore canalave-city-area
//{"encounter_method_rates":
// [
// {"encounter_method":{
// "name":"walk",
// "url":"https://pokeapi.co/api/v2/encounter-method/1/"},
// "version_details":[
// 		{"rate":10,
// 		"version":
//		{	"name":
// 			"diamond","url":"https://pokeapi.co/api/v2/version/12/"}},{"rate":10,"version":{"name":"pearl","url":"https://pokeapi.co/api/v2/version/13/"}},{"rate":10,"version":{"name":"platinum","url":"https://pokeapi.co/api/v2/version/14/"}}]}
