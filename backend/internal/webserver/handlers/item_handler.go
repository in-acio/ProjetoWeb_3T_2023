package handlers

import (
	"backend/internal/database"
	"backend/internal/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dchest/uniuri"
	"github.com/go-chi/jwtauth"
)

type ItemHandler struct {
	ItemDB database.ItemInterface
}

func NewItemHandler(itemDB database.ItemInterface) *ItemHandler{
	return &ItemHandler{
		ItemDB: itemDB,
	}
}

func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
    file, handler, err := r.FormFile("file")

    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    defer file.Close()

    imgName := uniuri.NewLen(30) + filepath.Ext(handler.Filename)
    f, err := os.OpenFile("./internal/webserver/images/" + imgName, os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    _, err = io.Copy(f, file)
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    name := r.FormValue("name")

    item := entity.NewItem(name, imgName)

    err = h.ItemDB.Create(item)
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (h *ItemHandler) List(w http.ResponseWriter, r *http.Request){

    _, claims, _ := jwtauth.FromContext(r.Context())
    fmt.Println(claims["userId"])

    p, err := h.ItemDB.FindAll(0, 0, "")
    
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(&p)
}

func (h *ItemHandler) Ranking(w http.ResponseWriter, r *http.Request){
    items, err := h.ItemDB.GetRanking();
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(&items);
}