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

//CreateType insert a type on database.
func CreateType(w http.ResponseWriter, r *http.Request) {

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var typ model.Type
	if erro = json.Unmarshal(bodyRequest, &typ); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = typ.Prepare("registration"); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoType(db)
	typ.ID, erro = repository.Create(typ)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, typ)

}

//FindTypes search all types on database.
func FindTypes(w http.ResponseWriter, r *http.Request) {
	typ := strings.ToLower(r.URL.Query().Get("typ"))
	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoType(db)
	types, erro := repository.Find(typ)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, types)
}

//FindTypeById search for as specific type by id.
func FindTypeById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	typeID, erro := strconv.ParseUint(parameters["typeId"], 10, 64)
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

	repository := repositories.NewRepoType(db)
	typ, erro := repository.FindById(typeID)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, typ)
}

//UpdateType update a type by id.
func UpdateType(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	typeID, erro := strconv.ParseUint(parameters["typeId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	bodyRequisition, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var typ model.Type
	if erro = json.Unmarshal(bodyRequisition, &typ); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = typ.Prepare("update"); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoType(db)
	erro = repository.Update(typeID, typ)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//DeleteType delete a type by id.
func DeleteType(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)
	typeID, erro := strconv.ParseUint(parameters["typeId"], 10, 64)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
	}

	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoType(db)
	if erro = repository.Delete(typeID); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
