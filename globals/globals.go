package globals

import (
	"bitbucket.org/ignitionrobotics/ign-go"
	"github.com/go-playground/form"
	"gopkg.in/go-playground/validator.v9"
)

/////////////////////////////////////////////////
/// Define global variables here

// Server encapsulates database, router, and auth0
var Server *ign.Server

// APIVersion is route api version.
// See also routes and routers
// \todo: Add support for multiple versions.
var APIVersion = "1.0"

// Validate references the global structs validator.
// See https://github.com/go-playground/validator.
// We use a single instance of validator, as it caches struct info
var Validate *validator.Validate

// FormDecoder holds a reference to the global Form Decoder.
// See https://github.com/go-playground/form.
// We use a single instance of Decoder, as it caches struct info
var FormDecoder *form.Decoder

// VersionPassword is the secreate API password
var VersionPassword string
