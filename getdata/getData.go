package getdata

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type FullData struct {
	Id             int                 `json:"id"`
	Name           string              `json:"name"`
	Image          string              `json:"image"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Relation       string              `json:"relations"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetData(link string) []byte {
	r, err := http.Get(link)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return body
}

func BindData(link string) []FullData {
	data := GetData(link)
	artists := []FullData{}

	err := json.Unmarshal(data, &artists)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	for i := 0; i < len(artists); i++ {
		r := Relations{}
		json.Unmarshal(GetData(artists[i].Relation), &r)
		artists[i].DatesLocations = r.DatesLocations
	}
	return artists
}
