package server

import (
	"fmt"
	"gp/getdata"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var ArtistsData []getdata.FullData

var link = "https://groupietrackers.herokuapp.com/api/artists"

var Data = getdata.BindData(link)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	// checking if the path is not correct and returning 400
	if r.URL.Path != "/" && strings.Index(r.URL.Path, "/artists/") != 0 {
		printError(w, r, "404 - Page not found")
		return
	}
	// checking if the data is loaded correctly
	if Data == nil {
		fmt.Println("Error getting data")
		return
	}
	// parsing the html file
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprint(w, "500 - Interal Server Error")
		return
	}
	t.Execute(w, Data)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/artists/") {
		// getting the artist number from the url
		artistId := strings.TrimPrefix(r.URL.Path, "/artists/")
		if artistId == "" {
			printError(w, r, "404 - Page not found")
			return
		}
		id, _ := strconv.Atoi(artistId)
		if id < 0  {
			printError(w, r, "404 - Page not found")
			return
		}
		// getting the artist data
		artist := GetArtistById(id)
		if artist.Id == 0{
			printError(w, r, "404 - Page not found")
			return
		}
		

		// parsing the html file
		t, err := template.ParseFiles("templates/artist.html")
		if err != nil {
			printError(w, r, "500 - Interal Server Error")
			return
		}
		t.Execute(w, artist)
	}
}

func GetArtistById(id int) getdata.FullData {
	for _, artist := range Data {
		if artist.Id == id {
			return artist
		}
	}
	return getdata.FullData{}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// parsing the html file
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		printError(w, r, "500 - Interal Server Error")
		return
	}
	t.Execute(w, Data)
}

func printError(w http.ResponseWriter, r *http.Request, errMsg string) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		fmt.Fprint(w, "500 - Interal Server Error")
		return
	}
	pageData := make(map[string]string)
	pageData["error_msg"] = errMsg

	t.Execute(w, pageData)
}

