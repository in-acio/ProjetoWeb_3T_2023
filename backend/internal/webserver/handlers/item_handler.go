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
	"strconv"
	"strings"

	"github.com/dchest/uniuri"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type ItemHandler struct {
	ItemDB database.ItemInterface
    ImagesFolder string
}

func NewItemHandler(itemDB database.ItemInterface, imgFolder string) *ItemHandler{
	return &ItemHandler{
		ItemDB: itemDB,
        ImagesFolder: imgFolder,
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

    fileExt := strings.ToLower(filepath.Ext(handler.Filename))

    if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
        handleBadRequest(w, "A imagem escolhida não é válida")
        return
    }

    imgName := uniuri.NewLen(30) + fileExt
    f, err := os.OpenFile(h.ImagesFolder + imgName, os.O_WRONLY|os.O_CREATE, 0666)

    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }
    defer f.Close()


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

    i, err := h.ItemDB.FindAll(0, 0, "desc")
    
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(&i)
}

func (h *ItemHandler) Show(w http.ResponseWriter, r *http.Request){
    id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32);
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    item, err := h.ItemDB.FindById(uint(id))
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(&item)
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

func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request){
    id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32);
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    item, err := h.ItemDB.FindById(uint(id))
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    err = h.ItemDB.Delete(uint(id))
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    err = os.Remove(h.ImagesFolder + item.Img) 
    if err != nil { 
        handleBadRequest(w, err.Error())
        return
    } 

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
}

func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request){
    id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32);
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    existingItem, err := h.ItemDB.FindById(uint(id))
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }
    
    r.ParseMultipartForm(32 << 20)
    file, handler, err := r.FormFile("file")

    if err != nil && err != http.ErrMissingFile {
        handleBadRequest(w, err.Error())
        return
    }

    var imgName string
    if file != nil {
        defer file.Close()

        err = os.Remove(h.ImagesFolder + existingItem.Img) 
        if err != nil { 
            handleBadRequest(w, err.Error())
            return
        }

        imgName = uniuri.NewLen(30) + filepath.Ext(handler.Filename)
        f, err := os.OpenFile(h.ImagesFolder + imgName, os.O_WRONLY|os.O_CREATE, 0666)

        if err != nil {
            handleBadRequest(w, err.Error())
            return
        }
        defer f.Close()

        _, err = io.Copy(f, file)
        if err != nil {
            handleBadRequest(w, err.Error())
            return
        }
    } else {
        imgName = existingItem.Img
    }

    name := r.FormValue("name")

    item := entity.NewItem(name, imgName)
    item.ID = existingItem.ID

    err = h.ItemDB.Update(item)
    if err != nil {
        handleBadRequest(w, err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
}