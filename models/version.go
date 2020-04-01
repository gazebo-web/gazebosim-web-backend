package models

import "time"

// Version is a single version of a library
type Version struct {
	Model

	LibID uint64 `json:"libId"`

	Major uint64 `json:"major"`
	Minor uint64 `json:"minor"`
	Patch uint64 `json:"patch"`

	ReleaseDate time.Time `json:"releaseDate"`
}

// Versions is an array of versions
type Versions []Version
