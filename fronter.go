package federate

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Member struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Pronouns    string `json:"pronouns"`
	AvatarURL   string `json:"avatar_url"`
}

type Fronter struct {
	Members []Member `json:"members"`
}

func GetFronter() (Fronter, error) {
	resp, err := http.Get("https://api.lmhd.me/v1/front.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var fronter Fronter
	err = json.Unmarshal(b, &fronter)
	if err != nil {
		log.Fatal(err)
	}

	return fronter, nil
}
