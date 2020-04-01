// Package main Ignition Fuel Server RESET API
//
// This package provides a REST API to the Ignition Fuel server.
//
// Schemes: https
// Host: staging-api.ignitionfuel.org
// BasePath: /1.0
// Version: 0.1.0
// License: Apache 2.0
// Contact: info@openrobotics.org
//
// swagger:meta
// go:generate swagger generate spec
package main

// \todo Add in the following to the comments at the top of this file to enable
// security
//
// SecurityDefinitions:
//   token:
//     type: apiKey
//     name: authorization
//     in: header
//     description: Ignition Fuel token
//   auth0:
//     type: apiKey
//     name: authorization
//     in: header
//     description: Auth0 token. Note, It must start with 'Bearer '
//

// Import this file's dependencies
import (
	"net/http"
	"path/filepath"

	"bitbucket.org/ignitionrobotics/ign-go"
	"bitbucket.org/ignitionrobotics/ign-webserver/globals"

	// "context"
	"flag"
	"log"
	"strconv"

	"github.com/caarlos0/env"
	"github.com/go-playground/form"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gopkg.in/go-playground/validator.v9"
)

// Impl note: we move this as a constant as it is used by tests.
const sysAdminForTest = "rootfortests"

type appConfig struct {
	SysAdmin          string `env:"IGN_WEBSERVER_SYSTEM_ADMIN"`
	isGoTest          bool
	Auth0RsaPublickey string `env:"AUTH0_RSA256_PUBLIC_KEY"`
	VersionPassword   string `env:"IGN_VERSION_PASSWORD"`
	SSLport           string `env:"IGN_WEBSERVER_SSL_PORT" envDefault:":4430"`
	HTTPport          string `env:"IGN_WEBSERVER_HTTP_PORT" envDefault:":8000"`
	logger            ign.Logger
}

/////////////////////////////////////////////////
/// Initialize this package
///
/// See readme for environment variables.
func init() {

	// Using ENV approach to allow multiple layers of configuration.
	// See https://github.com/joho/godotenv
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file. %+v\n", err)
	}

	cfg := appConfig{}
	// Also using env-to-struct approach to read configuration
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing environment into appConfig struct. %+v\n", err)
	}

	logger := initLogging(cfg)
	// logCtx := ign.NewContextWithLogger(context.Background(), logger)

	cfg.isGoTest = flag.Lookup("test.v") != nil

	// Get the auth0 credentials.
	if cfg.Auth0RsaPublickey == "" {
		logger.Info("Missing AUTH0_RSA256_PUBLIC_KEY env variable. Authentication will not work.")
	}

	// Get the version password
	if cfg.VersionPassword == "" {
		log.Fatalf("Missing IGN_VERSION_PASSWORD env variable.")
		return
	}

	var err error
	globals.Server, err = ign.Init(cfg.Auth0RsaPublickey)

	if err != nil {
		logger.Critical(err)
		log.Fatalf("Error while initializing app. %+v\n", err)
		return
	}

	// Override ports
	globals.Server.HTTPPort = cfg.HTTPport
	globals.Server.SSLport = cfg.SSLport
	log.Printf("Using HTTP port [%s] and SSL port [%s]", globals.Server.HTTPPort, globals.Server.SSLport)

	// Create the main Router and set it to the server.
	// Note: here it is the place to define multiple APIs
	s := globals.Server
	mainRouter := ign.NewRouter()
	apiPrefix := "/" + globals.APIVersion
	r := mainRouter.PathPrefix(apiPrefix).Subrouter()
	s.ConfigureRouterWithRoutes(apiPrefix, r, Routes)

	// Special file server route
	imgRoute := "/" + globals.APIVersion + "/images/{version}/{file:.+}"
	mainRouter.HandleFunc(imgRoute, func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		var file string
		if params["version"] != "all" {
			file = filepath.Join("./docs/", params["version"], filepath.Clean(params["file"]))
		} else {
			file = filepath.Join("./docs/", filepath.Clean(params["file"]))
		}
		http.ServeFile(w, req, file)
	})
	globals.Server.SetRouter(mainRouter)

	globals.Validate = initValidator(cfg)
	globals.FormDecoder = form.NewDecoder()
	globals.VersionPassword = cfg.VersionPassword
	// globals.Permissions = initPermissions(cfg)

	logger.Info("[application.go] Started using database: ", globals.Server.DbConfig.Name)

	// Migrate database tables
	DBMigrate()
	DBAddDefaults()
}

/////////////////////////////////////////////////
// Run the router and server
func main() {
	globals.Server.Run()
}

func initValidator(cfg appConfig) *validator.Validate {
	validate := validator.New()
	return validate
}

func initLogging(cfg appConfig) ign.Logger {
	verbosity := ign.VerbosityWarning
	if verbStr, verr := ign.ReadEnvVar("IGN_WEBSERVER_VERBOSITY"); verr == nil {
		verbosity, _ = strconv.Atoi(verbStr)
	}

	logStd := ign.ReadStdLogEnvVar()
	logger := ign.NewLogger("init", logStd, verbosity)
	return logger
}
