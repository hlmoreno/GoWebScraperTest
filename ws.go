package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func searchPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		searchStr := r.FormValue("searchText")
		http.Redirect(w, r, "/search/"+searchStr, http.StatusSeeOther)
	} else {
		tmpl := template.Must(template.ParseFiles("search.html"))

		tmpl.Execute(w, nil)
	}

}

func searchJSONResultsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	searchTxt := vars["searchText"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := publishPriceResults(searchTxt)
	resJSON, _ := json.Marshal(res)
	w.Write(resJSON)

}

func searchWebResultsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	searchTxt := vars["searchText"]

	tmpl := template.Must(template.ParseFiles("results.html"))
	res, _ := publishPriceResults(searchTxt)
	data := ResultDataPage{
		Title:      "Búsqueda: ",
		SearchText: searchTxt,
		Data:       res,
	}

	tmpl.Execute(w, data)

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", searchPage)
	router.HandleFunc("/api/search/{searchText}", searchJSONResultsPage)
	router.HandleFunc("/search/{searchText}", searchWebResultsPage)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func main() {

	//strSearch1 := "salsa piña"
	//strSearch2 := "harina trigo tradicional"

	//printProductos(strSearch1)
	//printProductos(strSearch2)

	handleRequests()
}

func printProductos(searchStr string) {

	fmt.Println("Precios productos: (" + searchStr + ")")
	prodPriceJumbo, err := GetLatestProductPriceJumbo(GenerateSearchURLStringJumbo(searchStr))
	if err != nil {
		log.Println(err)
	}

	prodPriceExito, err := GetLatestProductPriceExito(GenerateSearchURLStringExito(searchStr))
	if err != nil {
		log.Println(err)
	}

	prodPriceAlkosto, err := GetLatestProductPriceAlkosto(GenerateSearchURLStringAlkosto(searchStr))
	if err != nil {
		log.Println(err)
	}

	fmt.Println(prodPriceJumbo)
	fmt.Println(prodPriceExito)
	fmt.Println(prodPriceAlkosto)
	fmt.Println("")
}

func publishPriceResults(searchStr string) ([]DataWS, error) {

	prodPrices := []DataWS{}

	prodPriceJumbo, err := GetLatestProductPriceJumbo(GenerateSearchURLStringJumbo(searchStr))
	if err != nil {
		return nil, err
	}

	prodPriceExito, err := GetLatestProductPriceExito(GenerateSearchURLStringExito(searchStr))
	if err != nil {
		return nil, err
	}
	prodPriceAlkosto, err := GetLatestProductPriceAlkosto(GenerateSearchURLStringAlkosto(searchStr))
	if err != nil {
		return nil, err
	}

	prodPrices = append(prodPrices, prodPriceJumbo)
	prodPrices = append(prodPrices, prodPriceExito)
	prodPrices = append(prodPrices, prodPriceAlkosto)

	return prodPrices, nil
}
