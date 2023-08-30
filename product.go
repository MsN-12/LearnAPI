package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Get All Products
func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, products)
}

// Get One Product
func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid product ID")
		return
	}
	p := Product{ID: key}
	err = p.getProduct(app.DB)
	if err != nil {
		if err == sql.ErrNoRows {
			sendError(w, http.StatusNotFound, "Product not found")
		}
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, p)
}

// Create One Product
func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	err = p.createProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusCreated, p)
}

// Update One Product
func (app *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid product ID")
		return
	}
	var p Product
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	p.ID = key
	err = p.updateProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	}
	sendResponse(w, http.StatusOK, p)
}

// Delete One Product
func (app *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid product ID")
		return
	}
	p := Product{ID: key}
	err = p.deleteProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, map[string]string{"result": "delete was successful"})
}
