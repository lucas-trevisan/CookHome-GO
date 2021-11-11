package controllers

import (
	"encoding/json"
	"go/src/database"
	"go/src/model"
	"go/src/repositories"
	"go/src/responses"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//CreatePortion insert a portion on database.
func CreatePortion(w http.ResponseWriter, r *http.Request) {

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var portion model.Portion
	if erro = json.Unmarshal(bodyRequest, &portion); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = portion.Prepare("registration"); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoPortion(db)
	portion.ID, erro = repository.Create(portion)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, portion)

}

//FindPortions search all portions on database.
func FindPortions(w http.ResponseWriter, r *http.Request) {
	quantity := strings.ToLower(r.URL.Query().Get("portion"))
	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoPortion(db)
	portions, erro := repository.Find(quantity)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, portions)
}

//FindPortionById search for as specific portion by id.
func FindPortionById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	portionID, erro := strconv.ParseUint(parameters["portionId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoPortion(db)
	portion, erro := repository.FindById(portionID)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, portion)
}

//UpdatePortion update a portion by id.
func UpdatePortion(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	portionID, erro := strconv.ParseUint(parameters["portionId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	bodyRequisition, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var portion model.Portion
	if erro = json.Unmarshal(bodyRequisition, &portion); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = portion.Prepare("update"); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoPortion(db)
	erro = repository.Update(portionID, portion)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//DeletePortion delete a portion by id.
func DeletePortion(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)
	portionID, erro := strconv.ParseUint(parameters["portionId"], 10, 64)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
	}

	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoPortion(db)
	if erro = repository.Delete(portionID); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
