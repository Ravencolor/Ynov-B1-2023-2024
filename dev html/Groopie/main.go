package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	initialise()

	fmt.Println("Starting Server at Port 1237")
	fmt.Println("now open a browser and enter: localhost:1236 into the URL")

	http.HandleFunc("/", homePage)
	http.HandleFunc("/artistInfo", artistPage)
	http.HandleFunc("/locations", returnAllLocations)
	http.HandleFunc("/dates", returnAllDates)
	http.HandleFunc("/relation", returnAllRelation)
	http.HandleFunc("/search", handleSearch)
	http.HandleFunc("/egg", eggPage) // Ajout de la nouvelle route pour /egg
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.ListenAndServe(":1237", nil)
}

// eggPage gère les requêtes vers /egg
func eggPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Egg Page")
	// Chargez le modèle ou effectuez toute autre logique nécessaire pour cette page
	t, err := template.ParseFiles("egg.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Println("Endpoint Hit: returnAllArtists")
	data := ArtistData()
	t, err := template.ParseFiles("index.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func artistPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artistInfo" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Println("Endpoint Hit: Artist's Page")
	value := r.FormValue("ArtistName")
	if value == "" {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}
	a := collectData()
	var b Data
	for i, ele := range collectData() {
		if value == ele.A.Name {
			b = a[i]
		}
	}

	// Recherche des meilleurs titres de l'artiste sur Spotify
	previewID, err := searchTopTracks(value)
	if err != nil {
		log.Println("Erreur lors de la recherche des meilleurs titres sur Spotify:", err)
		// Handle the error as needed
	}

	// Ajout de l'URL de prévisualisation au modèle de données
	b.A.PreviewURL = previewID

	t, err := template.ParseFiles("artist.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	t.Execute(w, b)
}

func returnAllLocations(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllLocations")
	json.NewEncoder(w).Encode(LocationData())
}

func returnAllDates(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllDates")
	json.NewEncoder(w).Encode(DatesData())
}

func returnAllRelation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllRelation")
	json.NewEncoder(w).Encode(RelationData())
}

// rassemble les données extraites de toutes les structures API.
type Data struct {
	A Artist
	R Relation
	L Location
	D Date
}

// stocke les données de la structure API de artist.
type Artist struct {
	Id           uint     `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate uint     `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	PreviewURL   string   `json:"previewURL"`
}

// stocke les données de la structure API de location.
type Location struct {
	Locations []string `json:"locations"`
}

// stocke les données de la structure API de date.
type Date struct {
	Dates []string `json:"dates"`
}

// stocke les données de la structure API de relation.
type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Text struct {
	ErrorNum int
	ErrorMes string
}

// les slices de structs servent à indexer les données de chaque artiste depuis les API.
// les variables map[string]json.RawMessage sont utilisées pour désorganiser une autre couche
// lorsque plusieurs couches sont présentes.
var (
	artistInfo   []Artist
	locationMap  map[string]json.RawMessage
	locationInfo []Location
	datesMap     map[string]json.RawMessage
	datesInfo    []Date
	relationMap  map[string]json.RawMessage
	relationInfo []Relation
)

// ArtistData récupère et stocke les données de l'API Artist
func ArtistData() []Artist {
	artist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal()
	}
	artistData, err := ioutil.ReadAll(artist.Body)
	if err != nil {
		log.Fatal()
	}
	json.Unmarshal(artistData, &artistInfo)
	return artistInfo
}

// LocationData récupère et stocke les données de l'API Location
func LocationData() []Location {
	var bytes []byte
	location, err2 := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err2 != nil {
		log.Fatal()
	}
	locationData, err3 := ioutil.ReadAll(location.Body)
	if err3 != nil {
		log.Fatal()
	}
	err := json.Unmarshal(locationData, &locationMap)
	if err != nil {
		fmt.Println("error :", err)
	}
	for _, m := range locationMap {
		for _, v := range m {
			bytes = append(bytes, v)
		}
	}
	err = json.Unmarshal(bytes, &locationInfo)
	if err != nil {
		fmt.Println("error :", err)
	}
	return locationInfo
}

// DatesData récupère et stocke les données de l'API Dates
func DatesData() []Date {
	var bytes []byte
	dates, err2 := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err2 != nil {
		log.Fatal()
	}
	datesData, err3 := ioutil.ReadAll(dates.Body)
	if err3 != nil {
		log.Fatal()
	}
	err := json.Unmarshal(datesData, &datesMap)
	if err != nil {
		fmt.Println("error :", err)
	}
	for _, m := range datesMap {
		for _, v := range m {
			bytes = append(bytes, v)
		}
	}
	err = json.Unmarshal(bytes, &datesInfo)
	if err != nil {
		fmt.Println("error :", err)
	}
	return datesInfo
}

// RelationData récupère et stocke les données de l'API Relation
func RelationData() []Relation {
	var bytes []byte
	relation, err2 := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err2 != nil {
		log.Fatal()
	}
	relationData, err3 := ioutil.ReadAll(relation.Body)
	if err3 != nil {
		log.Fatal()
	}
	err := json.Unmarshal(relationData, &relationMap)
	if err != nil {
		fmt.Println("error :", err)
	}

	for _, m := range relationMap {
		for _, v := range m {
			bytes = append(bytes, v)
		}
	}

	err = json.Unmarshal(bytes, &relationInfo)
	if err != nil {
		fmt.Println("error :", err)
	}
	return relationInfo
}

// collectData agrège les données de toutes les API dans une seule structure
func collectData() []Data {
	ArtistData()
	RelationData()
	LocationData()
	DatesData()
	dataData := make([]Data, len(artistInfo))
	for i := 0; i < len(artistInfo); i++ {
		dataData[i].A = artistInfo[i]
		dataData[i].R = relationInfo[i]
		dataData[i].L = locationInfo[i]
		dataData[i].D = datesInfo[i]
	}
	return dataData
}

// errorHandler gère les messages d'erreur
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, err := template.ParseFiles("error.html")
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		em := "Statut HTTP 404 : page introuvable"
		p := Text{ErrorNum: status, ErrorMes: em}
		t.Execute(w, p)
	}
	if status == http.StatusInternalServerError {
		t, err := template.ParseFiles("error.html")
		if err != nil {
			fmt.Fprint(w, "Statut HTTP 500 : erreur de serveur interne - fichier error.html manquant")
		}
		em := "Statut HTTP 500 : erreur de serveur interne"
		p := Text{ErrorNum: status, ErrorMes: em}
		t.Execute(w, p)
	}
	if status == http.StatusBadRequest {
		t, err := template.ParseFiles("error.html")
		if err != nil {
			fmt.Fprint(w, "Statut HTTP 500 : erreur de serveur interne - fichier error.html manquant")
		}
		em := "	"
		p := Text{ErrorNum: status, ErrorMes: em}
		t.Execute(w, p)
	}
}

// Pour la barre de recherche
func handleSearch(w http.ResponseWriter, r *http.Request) {
	// Extrait les paramètres de recherche de la requête
	query := r.URL.Query().Get("q")

	// Effectue la recherche
	results := performSearch(query)

	// Renvoie les résultats de la recherche
	json.NewEncoder(w).Encode(results)
}

type SearchResult struct {
	Title string `json:"title"`
	Image string `json:"image"`
	URL   string `json:"url"`
}

// permet d'afficher le(s) résultat(s) de la recherche
func performSearch(query string) []SearchResult {
	artistList := ArtistData()

	var searchResults []SearchResult

	for _, artist := range artistList {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			artistURL := "/artistInfo?ArtistName=" + url.QueryEscape(artist.Name)
			searchResults = append(searchResults, SearchResult{Title: artist.Name, Image: artist.Image, URL: artistURL})
		}
	}

	return searchResults
}

var (
	spotifyClient *spotify.Client // Déclaration du client Spotify
)

func initialise() {
	// Initialisation de la configuration OAuth2 client credentials
	config := &clientcredentials.Config{
		ClientID:     "f7fa880eeb84438ab664bda997e21af3",
		ClientSecret: "564e16a3000042f1a48bb218fb7934d2",
		TokenURL:     spotify.TokenURL,
	}

	// Obtention du token d'accès
	ctx := context.Background()
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("Impossible d'obtenir le token d'accès: %v", err)
	}

	// Création du client Spotify en utilisant le token d'accès
	client := spotify.Authenticator{}.NewClient(token)
	spotifyClient = &client // affectation de l'adresse du client Spotify
}

// recherche les meilleurs titres de l'artiste sur Spotify
func searchTopTracks(artistName string) (string, error) {
	// Use the Spotify API to search for the artist
	result, err := spotifyClient.Search(artistName, spotify.SearchTypeArtist)
	if err != nil {
		return "", fmt.Errorf("Error searching for artist on Spotify: %v", err)
	}

	// Check if any artists were found
	if len(result.Artists.Artists) == 0 {
		return "", errors.New("No artist found on Spotify")
	}

	// Retrieve the artist ID from the search results
	artistID := result.Artists.Artists[0].ID

	// Use the Spotify API to get the artist's top tracks
	tracks, err := spotifyClient.GetArtistsTopTracks(artistID, "US")
	if err != nil {
		return "", fmt.Errorf("Error getting top tracks from Spotify: %v", err)
	}

	// Check if any top tracks were found
	if len(tracks) == 0 {
		return "", errors.New("No top tracks found for the artist on Spotify")
	}
	// Retrieve the track ID
	trackID := tracks[0].ID

	// Convert the track ID to a regular string
	trackIDString := string(trackID)

	// Use the Spotify API to get full track details
	track, err := spotifyClient.GetTrack(trackID)
	if err != nil {
		return "", fmt.Errorf("Error getting track details from Spotify: %v", err)
	}

	// Retrieve the track URI
	trackURI := track.URI

	fmt.Println("Track URI:", trackURI)

	// Return the track ID as a regular string
	return trackIDString, nil
}
