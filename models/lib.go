package models

// Lib defines a library
type Lib struct {
	Model
	Name        string   `json:"name"`
	Versions    Versions `json:"versions"`
	Repo        string   `json:"repo"`
	Description string   `json:"description"`
}

// Libs defines an array of libraries
type Libs []Lib
