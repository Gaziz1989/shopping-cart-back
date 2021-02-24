package handler

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"landing-back/useCase/product"
)

func listProducts(service product.UseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		products, err := service.ListProducts()
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		data := map[string]interface{}{}
		data["message"] = map[string]interface{}{
			"status": true,
			"text":   "OK",
			"products": products,
		}
		json.NewEncoder(w).Encode(data)
	}
}

func addProduct(service product.UseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		input := struct{
			Title string `json:"title,omitempty"`
			Description string `json:"description,omitempty"`
			Image string `json:"image,omitempty"`
			Price float64 `json:"price,omitempty"`
			AvailableSizes []string `json:"availableSizes,omitempty"`
		}{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		newProduct, err := service.CreateProduct(input.Title,input.Description,input.Image,input.Price,input.AvailableSizes)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		data := map[string]interface{}{}
		data["message"] = map[string]interface{}{
			"status": true,
			"text":   "OK",
			"product": newProduct,
		}
		json.NewEncoder(w).Encode(data)
	}
}

func deleteProduct(service product.UseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Printf("%+v\n", ps.ByName("id"))
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		err = service.DeleteProduct(id)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		data := map[string]interface{}{}
		data["message"] = map[string]interface{}{
			"status": true,
			"text":   "OK",
		}
		json.NewEncoder(w).Encode(data)
	}
}

//MakeProductHandlers make url handlers
func MakeProductHandlers(router *httprouter.Router, service product.UseCase) {
	router.GET("/api/v1/product", listProducts(service))
	router.POST("/api/v1/product", addProduct(service))
	router.DELETE("/api/v1/product/:id", deleteProduct(service))
}