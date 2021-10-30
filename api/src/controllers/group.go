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

//CreateGroup insert a group on database.
func CreateGroup(w http.ResponseWriter, r *http.Request) {

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var group model.Group
	if erro = json.Unmarshal(bodyRequest, &group); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = group.Prepare("registration"); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoGroup(db)
	group.ID, erro = repository.Create(group)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, group)

}

//FindGroups search all groups on database.
func FindGroups(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Query().Get("group"))
	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoGroup(db)
	groups, erro := repository.Find(name)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, groups)
}

//FindGroupById search for as specific group by id.
func FindGroupById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	groupID, erro := strconv.ParseUint(parameters["groupId"], 10, 64)
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

	repository := repositories.NewRepoGroup(db)
	group, erro := repository.FindById(groupID)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, group)
}

//UpdateGroup update a group by id.
func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	groupID, erro := strconv.ParseUint(parameters["groupId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	bodyRequisition, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var group model.Group
	if erro = json.Unmarshal(bodyRequisition, &group); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = group.Prepare("update"); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoGroup(db)
	erro = repository.Update(groupID, group)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//DeleteGroup delete a group by id.
func DeleteGroup(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)
	groupID, erro := strconv.ParseUint(parameters["groupId"], 10, 64)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
	}

	db, erro := database.Conn()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepoGroup(db)
	if erro = repository.Delete(groupID); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
