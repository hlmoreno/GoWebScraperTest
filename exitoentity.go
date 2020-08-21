package main

type imageExito struct {
	ImageID   string `json:"imageId"`
	ImageText string `json:"imageText"`
	ImageURL  string `json:"imageUrl"`
}

type itemsExito struct {
	Name   string       `json:"name"`
	Ean    string       `json:"ean"`
	Images []imageExito `json:"images"`
}

type price struct {
	HighPrice int `json:"highPrice"`
	LowPrice  int `json:"lowPrice"`
}

type pricerange struct {
	ListPrice    price `json:"listPrice"`
	SellingPrice price `json:"sellingPrice"`
}

type productsExito struct {
	ProductID   string       `json:"productId"`
	ProductName string       `json:"productName"`
	Items       []itemsExito `json:"items"`
	PriceRange  pricerange   `json:"priceRange"`
}

type resultsExito struct {
	Products []productsExito `json:"products"`
}

type data struct {
	SearchResults resultsExito `json:"searchResult"`
}

type dataExito struct {
	Data data `json:"data"`
}
