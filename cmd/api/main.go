package main

import (
	"net/http"

	"github.com/albuquerquehugo/golang-intensivo/internal/entity"
	"github.com/labstack/echo/v4"
)

func main() {
	// http.HandleFunc("/order", OrderHandler)
	// http.ListenAndServe(":8888", nil)

	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", OrderHandler)
	// http.ListenAndServe(":8888", r)

	e := echo.New()
	e.GET("/order", OrderHandler)
	e.Logger.Fatal(e.Start(":8888"))
}

// func OrderHandler(w http.ResponseWriter, r *http.Request) {
// 	order, _ := entity.NewOrder("1", 10, 1)
// 	err := order.CalculateFinalPrice()
// 	if err != nil { 
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	json.NewEncoder(w).Encode(order)
// }

func OrderHandler(c echo.Context) error {
	order, _ := entity.NewOrder("1", 10, 1)
	err := order.CalculateFinalPrice()
	if err != nil {
		// return c.JSON(500, err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}