package controllers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gitlab.com/ignitionrobotics/web/ign-go"
	"gopkg.in/yaml.v2"
)

// Docs returns a map of version names to array of tutorial names.
func Docs(tx *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, *ign.ErrMsg) {

	// Get all subdirectories in the "docs" folder
	// versionDirs, err := ioutil.ReadDir("docs")

	type Page struct {
		Name        string `json:"name" yaml:"name"`
		Title       string `json:"title" yaml:"title"`
		File        string `json:"file" yaml:"file"`
		Description string `json:"description" yaml:"description"`
		Children    []Page `json:"children" yaml:"children"`
	}
	type Pages struct {
		Pages []Page `json:"pages" yaml:"pages"`
	}

	type RootIndex struct {
		Pages    []Page   `json:"pages" yaml:"pages"`
		Releases []string `json:"releases" yaml:"releases"`
	}

	type DocsInfo struct {
		Versions []string          `json:"versions"`
		Pages    map[string][]Page `json:"pages" yaml:"pages"`
	}
	var result DocsInfo

	result.Pages = make(map[string][]Page)

	var rootIndex RootIndex
	file := filepath.Join("docs", "index.yaml")
	rootData, _ := ioutil.ReadFile(file)
	err := yaml.Unmarshal(rootData, &rootIndex)

	result.Pages["all"] = make([]Page, len(rootIndex.Pages))
	copy(result.Pages["all"], rootIndex.Pages)

	if err == nil {
		// Iterate over the version directories
		for _, v := range rootIndex.Releases {
			// Get all files in the versioned directory
			result.Versions = append(result.Versions, v)
			file := filepath.Join("docs", v, "index.yaml")
			data, _ := ioutil.ReadFile(file)

			var pageData Pages
			_ = yaml.Unmarshal(data, &pageData)
			result.Pages[v] = make([]Page, len(pageData.Pages))
			copy(result.Pages[v], pageData.Pages)
		}
	} else {
		return result, ign.NewErrorMessageWithBase(ign.ErrorFileNotFound, err)
	}

	return result, nil
}

// DocsPage returns the specified markdown file
func DocsPage(tx *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, *ign.ErrMsg) {
	params := mux.Vars(r)

	page := params["page"]
	version, valid := params["version"]
	if version == "all" {
		version = ""
	}

	var doc = "# 404"
	if valid {
		file := filepath.Join("docs", version, filepath.Clean(page))
		data, err := ioutil.ReadFile(file)
		if err == nil {
			doc = string(data)
		}
	}
	return doc, nil
}
