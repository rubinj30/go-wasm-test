package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetPokemon() {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Print("HELLO TEST")
	// json.NewDecoder(body)
	// log.Println(fmt.Sprintf("%q\n", body))
	var data interface{} // TopTracks
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(data)
}
