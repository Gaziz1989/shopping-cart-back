package handler

import (
	"net/http"
	// "fmt"
	"encoding/json"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"landing-back/useCase/order"
	"landing-back/entities"
)

func listOrders(service order.UseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		orders, err := service.ListOrders()
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		data := map[string]interface{}{}
		data["message"] = map[string]interface{}{
			"status": true,
			"text":   "OK",
			"orders": orders,
		}
		json.NewEncoder(w).Encode(data)
	}
}

func addOrder(service order.UseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		input := entities.Order{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = service.CreateOrder(input)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		data := map[string]interface{}{}
		data["message"] = map[string]interface{}{
			"status": true,
			"text":   "OK",
			"newOrder": input,
		}
		json.NewEncoder(w).Encode(data)
	}
}

func deleteOrder(service order.UseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		err = service.DeleteOrder(id)
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

//MakeOrderHandlers make url handlers
func MakeOrderHandlers(router *httprouter.Router, service order.UseCase) {
	router.GET("/api/v1/order", listOrders(service))
	router.POST("/api/v1/order", addOrder(service))
	router.DELETE("/api/v1/order/:id", deleteOrder(service))
}