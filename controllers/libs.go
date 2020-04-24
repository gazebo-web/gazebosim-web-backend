package controllers

import (
	"gitlab.com/ignitionrobotics/web/ign-go"
	"gitlab.com/ignitionrobotics/web/web-server/globals"
	"gitlab.com/ignitionrobotics/web/web-server/models"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"net/http"
	"sort"
	"strings"
)

type byName []models.Lib

func (p byName) Len() int {
	return len(p)
}

func (p byName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byName) Less(i, j int) bool {
	return strings.Compare(p[i].Name, p[j].Name) < 0
}

type byVersion []models.Version

func (p byVersion) Len() int {
	return len(p)
}

func (p byVersion) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byVersion) Less(i, j int) bool {
	return p[i].Major > p[j].Major ||
		(p[i].Major == p[j].Major && p[i].Minor > p[j].Minor) ||
		(p[i].Major == p[j].Major && p[i].Minor == p[j].Minor && p[i].Patch > p[j].Patch)
}

// LibsList returns the list of models from a team/user. The returned value
// will be of type "fuel.Models"
// You can request this method with the following curl request:
//     curl -k -X GET --url https://localhost:4430/1.0/models
// or  curl -k -X GET --url https://localhost:4430/1.0/models.proto
// or  curl -k -X GET --url https://localhost:4430/1.0/models.json
// or  curl -k -X GET --url https://localhost:4430/1.0/{owner}/models with all the
// above format variants.
func LibsList(tx *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, *ign.ErrMsg) {

	var libs models.Libs
	if err := globals.Server.Db.Model(&models.Lib{}).Preload("Versions").Find(&libs).Error; err != nil {
		return nil, ign.NewErrorMessageWithBase(ign.ErrorIDNotFound, err)
	}

	sort.Sort(byName(libs))

	for _, lib := range libs {
		sort.Sort(byVersion(lib.Versions))
	}
	return libs, nil
}

// LibsCreate creates a new library
// You can request this method with the following curl request:
// curl -k -X POST -d '{"name":"math", "repo":"https://github.com/ignitionrobotics/ign-math", "description":"Math description","password":"secret"}' http://localhost:8000/1.0/libs
func LibsCreate(tx *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, *ign.ErrMsg) {

	type Data struct {
		models.Lib
		Password string `json:"password"`
	}

	var lib Data

	// Decode the data
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&lib)
	if err != nil {
		return nil, ign.NewErrorMessageWithArgs(ign.ErrorUnmarshalJSON, err,
			[]string{"Incorrect POST data"})
	}

	// Check the password
	if lib.Password != globals.VersionPassword {
		return nil, ign.NewErrorMessageWithBase(ign.ErrorForm, errors.New("Invalid password"))
	}

	var tLib models.Lib
	if err := globals.Server.Db.Model(&models.Lib{}).Where("name = ?", lib.Lib.Name).First(&tLib).Error; err == nil {
		return nil, ign.NewErrorMessageWithArgs(ign.ErrorUnmarshalJSON, err,
			[]string{"Library already exists"})
	}

	globals.Server.Db.Create(&lib.Lib)
	return lib.Lib, nil
}
