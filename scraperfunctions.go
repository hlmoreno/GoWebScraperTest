package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetLatestProductPriceJumbo gets the latest blog title headings from the url
// given and returns them as a list.
func GetLatestProductPriceJumbo(url string) (DataWS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return DataWS{}, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return DataWS{}, err
	}

	products := DataWS{Almacen: "Jumbo"}
	doc.Find(".nm-product-item").Each(func(i int, s *goquery.Selection) {

		image, _ := s.Find(".nm-product-img-container img").Attr("src")
		producto := ProductoWS{
			Nombre: s.Find(".nm-product-name a").Text(),
			Precio: strings.ReplaceAll(strings.Trim(s.Find(".nm-price-value").Text(), "$"), ".", ""),
			ImgURL: image,
		}
		products.Productos = append(products.Productos, producto)
	})
	return products, nil
}

// GetLatestProductPriceExito gets the latest blog title headings from the url
// given and returns them as a list.
func GetLatestProductPriceExito(url string) (DataWS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return DataWS{}, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	data1 := dataExito{}
	err = json.Unmarshal(body, &data1)
	if err != nil {
		log.Fatal(err)
	}

	products := DataWS{Almacen: "Exito"}
	for _, prod := range data1.Data.SearchResults.Products {
		producto := ProductoWS{Nombre: prod.ProductName, Precio: strconv.Itoa(prod.PriceRange.SellingPrice.HighPrice), ImgURL: prod.Items[0].Images[0].ImageURL}

		products.Productos = append(products.Productos, producto)
	}

	return products, nil
}

// GetLatestProductPriceAlkosto gets the latest blog title headings from the url
// given and returns them as a list.
func GetLatestProductPriceAlkosto(url string) (DataWS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return DataWS{}, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	data1 := []dataAlkosto{}
	err = json.Unmarshal(body, &data1)
	if err != nil {
		log.Fatal(err)
	}

	products := DataWS{Almacen: "Alkosto"}
	for _, prod := range data1[0].Products {
		precio := fmt.Sprintf("%0.0f", prod.Price)
		producto := ProductoWS{Nombre: prod.ProductBrand.Name + "-" + prod.ProductName + " " + prod.Package, Precio: precio, ImgURL: prod.ImageURL}

		products.Productos = append(products.Productos, producto)
	}

	return products, nil
}

// GenerateSearchURLStringJumbo converts search string to base64
func GenerateSearchURLStringJumbo(searchStr string) string {

	strURL := "https://busqueda.tiendasjumbo.co/busca?q=" + url.QueryEscape(searchStr) + "\""

	return strURL
}

// GenerateSearchURLStringExito converts search string to base64
func GenerateSearchURLStringExito(searchStr string) string {

	strEnc := GenerateBase64StringExito(searchStr)
	strURL := "https://www.exito.com/_v/segment/graphql/v1?workspace=master&maxAge=short&appsEtag=remove&domain=store&locale=es-CO&__bindingId=2f829b4f-604f-499c-9ffb-2c5590f076db&operationName=searchResult&variables=%7B%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%22dcf550c27cd0bbf0e6899e3fa1f4b8c0b977330e321b9b8304cc23e2d2bad674%22%2C%22sender%22%3A%22vtex.search%400.x%22%2C%22provider%22%3A%22vtex.search%400.x%22%7D%2C%22variables%22%3A%22" + strEnc + "%3D%22%7D"

	return strURL
}

// GenerateSearchURLStringAlkosto converts search string to base64
func GenerateSearchURLStringAlkosto(searchStr string) string {

	strEnc := strings.ReplaceAll(searchStr, " ", "+")
	strURL := "https://cornershopapp.com/api/v1/branches/4901/search?query=" + strEnc

	return strURL
}

// GenerateBase64StringExito converts search string to base64
func GenerateBase64StringExito(searchStr string) string {

	strResult := "{\"productOrigin\":\"VTEX\",\"indexingType\":\"API\",\"query\":\"" + searchStr + "\",\"page\":1,\"attributePath\":\"\",\"sort\":\"\",\"count\":20,\"leap\":false}"

	strEnc := base64.StdEncoding.EncodeToString([]byte(strResult))

	return strEnc
}
