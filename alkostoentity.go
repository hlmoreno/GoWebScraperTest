package main

type brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type productsAlkosto struct {
	ProductID    int     `json:"id"`
	ProductBrand brand   `json:"brand"`
	ProductName  string  `json:"name"`
	Price        float32 `json:"price"`
	Package      string  `json:"package"`
	ImageURL     string  `json:"img_url"`
}

type dataAlkosto struct {
	AisleID  string            `json:"aisle_id"`
	Products []productsAlkosto `json:"products"`
}
