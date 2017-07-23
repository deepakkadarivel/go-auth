package product

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

var products = []Product{
	{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Shoot your way to the top on 14 different hoverboards"},
	{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

var StatusHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Api is up and running."))
})

var ProductsHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	payload, _ := json.Marshal(products)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(payload))
})

var AddFeedbackHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	var requestedProduct Product
	vars := mux.Vars(req)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			requestedProduct = p
		}
	}

	res.Header().Set("Content-Type", "application/json")
	if requestedProduct.Slug != "" {
		payload, _ := json.Marshal(requestedProduct)
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(payload))
	} else {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Product not found."))
	}
})
