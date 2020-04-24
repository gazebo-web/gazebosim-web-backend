package controllers

import (
	"gitlab.com/ignitionrobotics/web/ign-go"
	"gitlab.com/ignitionrobotics/web/web-server/globals"
	"gitlab.com/ignitionrobotics/web/web-server/models"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// VersionCreate creates a new version
// curl -k -X POST -d '{"libName":"common", "version":"1.2.4", "releaseDate":"2017-11-14T12:02:54-07:00","password":"secret"}' http://localhost:8000/1.0/versions
func VersionCreate(tx *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, *ign.ErrMsg) {
	var err error
	var major uint64
	var minor uint64
	var patch uint64

	// var libName string
	// var version string

	type Data struct {
		LibName     string    `json:"libName"`
		Version     string    `json:"version"`
		ReleaseDate time.Time `json:"releaseDate"`
		Password    string    `json:"password"`
	}
	var data Data

	// Decode the data
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, ign.NewErrorMessageWithArgs(ign.ErrorUnmarshalJSON, err,
			[]string{
				"Failed to parse POST data. Maybe the release date is invalid. The ReleaseDate should have the following format: 2017-11-14T12:02:54-07:00"})
	}

	// Check the password
	if data.Password != globals.VersionPassword {
		return nil, ign.NewErrorMessageWithBase(ign.ErrorForm, errors.New("Invalid password"))
	}

	// Split the version data
	versionParts := strings.Split(data.Version, ".")

	// Check that the full version is given
	if len(versionParts) != 3 {
		return nil, ign.NewErrorMessageWithBase(ign.ErrorForm,
			errors.New("Version must have major.minor.patch"))
	}

	// Get the major number
	major, err = strconv.ParseUint(versionParts[0], 10, 64)
	if err != nil {
		return nil, ign.NewErrorMessageWithBase(ign.ErrorForm, err)
	}

	// Get the minor number
	minor, err = strconv.ParseUint(versionParts[1], 10, 64)
	if err != nil {
		return nil, ign.NewErrorMessageWithBase(ign.ErrorForm, err)
	}

	// Get the patch number
	patch, err = strconv.ParseUint(versionParts[2], 10, 64)
	if err != nil {
		return nil, ign.NewErrorMessageWithBase(ign.ErrorForm, err)
	}

	// Find the library
	var lib models.Lib
	if err := globals.Server.Db.Preload("Versions").Where("name = ?", data.LibName).First(&lib).Error; err != nil {
		return nil, ign.NewErrorMessageWithArgs(ign.ErrorNameNotFound, err,
			[]string{"Library does not exist with given name"})
	}

	// Check that the version doesn't exist
	for _, v := range lib.Versions {
		if v.Major == major && v.Minor == minor && v.Patch == patch {
			return nil, ign.NewErrorMessageWithArgs(ign.ErrorNameNotFound,
				errors.New("Version already exists"),
				[]string{"Version already exists"})
		}
	}

	// Create the version
	lib.Versions = append(lib.Versions,
		models.Version{
			Major:       major,
			Minor:       minor,
			Patch:       patch,
			ReleaseDate: data.ReleaseDate,
		},
	)

	// Save the new version
	globals.Server.Db.Save(&lib)

	return "{status: success, info: Upload documentation to S3 using ...}", nil
}
