package main

import (
	"github.com/gazebo-web/gazebosim-web-backend/controllers"
	"gitlab.com/ignitionrobotics/web/ign-go"
)

// Routes declares the routes
var Routes = ign.Routes{

	// Route for all models
	ign.Route{
		Name:        "Libs",
		Description: "Information about all Libs",
		URI:         "/libs",
		Headers:     ign.AuthHeadersOptional,
		Methods: ign.Methods{
			// swagger:route GET /libs libs listLib
			//
			// Get a list of libraries
			//
			//   Produces:
			//   - application/json
			//
			//   Schemes: https
			//
			//   Responses:
			//     200: Libraries
			ign.Method{
				Type:        "GET",
				Description: "Get all libraries",
				Handlers: ign.FormatHandlers{
					ign.FormatHandler{
						Extension: ".json",
						Handler:   ign.JSONListResult("", controllers.LibsList),
					},

					ign.FormatHandler{
						Extension: "",
						Handler:   ign.JSONListResult("", controllers.LibsList),
					},
				},
			},
			ign.Method{
				Type:        "POST",
				Description: "Add a new libary",
				Handlers: ign.FormatHandlers{
					ign.FormatHandler{
						Extension: ".json",
						Handler:   ign.JSONListResult("", controllers.LibsCreate),
					},

					ign.FormatHandler{
						Extension: "",
						Handler:   ign.JSONListResult("", controllers.LibsCreate),
					},
				},
			},
		},
		SecureMethods: ign.SecureMethods{},
	},

	// Route for versions
	ign.Route{
		Name:        "Versions",
		Description: "Routes for lib versions",
		URI:         "/versions",
		Headers:     ign.AuthHeadersOptional,
		Methods: ign.Methods{
			// swagger:route GET /libs libs listLib
			//
			// Get a list of libraries
			//
			//   Produces:
			//   - application/json
			//
			//   Schemes: https
			//
			//   Responses:
			//     200: Libraries
			ign.Method{
				Type:        "POST",
				Description: "Add a new version",
				Handlers: ign.FormatHandlers{
					ign.FormatHandler{
						Extension: ".json",
						Handler:   ign.JSONListResult("", controllers.VersionCreate),
					},

					ign.FormatHandler{
						Extension: "",
						Handler:   ign.JSONListResult("", controllers.VersionCreate),
					},
				},
			},
		},
		SecureMethods: ign.SecureMethods{},
	},

	// Route for tutorials
	ign.Route{
		Name:        "Tutorials",
		Description: "Routes for tutorials",
		URI:         "/tutorials",
		Headers:     ign.AuthHeadersOptional,
		Methods: ign.Methods{
			// swagger:route GET /tutorials tutorials listTutorials
			//
			// Get a list of tutorial pages
			//
			//   Produces:
			//   - application/json
			//
			//   Schemes: https
			//
			//   Responses:
			//     200: Libraries
			ign.Method{
				Type:        "GET",
				Description: "Get list of tutorials pages",
				Handlers: ign.FormatHandlers{
					ign.FormatHandler{
						Extension: ".json",
						Handler:   ign.JSONListResult("", controllers.Tutorials),
					},

					ign.FormatHandler{
						Extension: "",
						Handler:   ign.JSONListResult("", controllers.Tutorials),
					},
				},
			},
		},
		SecureMethods: ign.SecureMethods{},
	},
	// Route for documentation
	ign.Route{
		Name:        "Docs",
		Description: "Routes for documentation",
		URI:         "/docs",
		Headers:     ign.AuthHeadersOptional,
		Methods: ign.Methods{
			// swagger:route GET /docs docs listDocs
			//
			// Get a list of documentation pages
			//
			//   Produces:
			//   - application/json
			//
			//   Schemes: https
			//
			//   Responses:
			//     200: Libraries
			ign.Method{
				Type:        "GET",
				Description: "Get list of documentation pages",
				Handlers: ign.FormatHandlers{
					ign.FormatHandler{
						Extension: ".json",
						Handler:   ign.JSONListResult("", controllers.Docs),
					},

					ign.FormatHandler{
						Extension: "",
						Handler:   ign.JSONListResult("", controllers.Docs),
					},
				},
			},
		},
		SecureMethods: ign.SecureMethods{},
	},
	// Route for documentation
	ign.Route{
		Name:        "Docs",
		Description: "Routes for documentation",
		URI:         "/docs/{version}/{page}/{subpage}",
		Headers:     ign.AuthHeadersOptional,
		Methods: ign.Methods{
			// swagger:route GET /libs libs listLib
			//
			// Get a list of libraries
			//
			//   Produces:
			//   - application/json
			//
			//   Schemes: https
			//
			//   Responses:
			//     200: Libraries
			ign.Method{
				Type:        "GET",
				Description: "Get documentation",
				Handlers: ign.FormatHandlers{
					ign.FormatHandler{
						Extension: ".json",
						Handler:   ign.JSONListResult("", controllers.DocsPage),
					},

					ign.FormatHandler{
						Extension: "",
						Handler:   ign.JSONListResult("", controllers.DocsPage),
					},
				},
			},
		},
		SecureMethods: ign.SecureMethods{},
	},
	// Route for library benchmarks
	ign.Route{
		Name:        "Benchmark",
		Description: "Routes for benchmarks",
		URI:         "/benchmarks",
		Headers:     ign.AuthHeadersOptional,
		Methods: ign.Methods{
			// swagger:route GET /benchmarks benchmark benchmarkLib
			//
			// Get the set of benchmark series for a library
			//
			//   Produces:
			//   - application/json
			//
			//   Schemes: https
			//
			//   Responses:
			//     200: Benchmarks
			ign.Method{
				Type:        "GET",
				Description: "Get benchmark data",
				Handlers: ign.FormatHandlers{
					ign.FormatHandler{
						Extension: ".json",
						Handler:   ign.JSONListResult("", controllers.BenchmarkSummary),
					},

					ign.FormatHandler{
						Extension: "",
						Handler:   ign.JSONListResult("", controllers.BenchmarkSummary),
					},
				},
			},
		},
		SecureMethods: ign.SecureMethods{},
	},
	// Route for library benchmarks
	ign.Route{
		Name:        "Benchmark",
		Description: "Routes for benchmarks",
		URI:         "/benchmarks/{library}",
		Headers:     ign.AuthHeadersOptional,
		Methods: ign.Methods{
			// swagger:route GET /benchmarks benchmark benchmarkLib
			//
			// Get the set of benchmark series for a library
			//
			//   Produces:
			//   - application/json
			//
			//   Schemes: https
			//
			//   Responses:
			//     200: Benchmarks
			ign.Method{
				Type:        "GET",
				Description: "Get benchmark data",
				Handlers: ign.FormatHandlers{
					ign.FormatHandler{
						Extension: ".json",
						Handler:   ign.JSONListResult("", controllers.Benchmarks),
					},

					ign.FormatHandler{
						Extension: "",
						Handler:   ign.JSONListResult("", controllers.Benchmarks),
					},
				},
			},
		},
		SecureMethods: ign.SecureMethods{
			// swagger:route POST /benchmark benchmark benchmarkLib
			//
			// Post new benchmark data
			//
			//   Produces:
			//   - application/json
			//
			//   Schemes: https
			//
			//   Responses:
			//     200: Benchmarks
			ign.Method{
				Type:        "POST",
				Description: "Post benchmark data",
				Handlers: ign.FormatHandlers{
					ign.FormatHandler{
						Extension: ".json",
						Handler:   ign.JSONListResult("", controllers.BenchmarkCreate),
					},

					ign.FormatHandler{
						Extension: "",
						Handler:   ign.JSONListResult("", controllers.BenchmarkCreate),
					},
				},
			},
		},
	},
	// Route for library benchmarks
	ign.Route{
		Name:        "Benchmark dates",
		Description: "Routes for benchmarks",
		URI:         "/benchmarks/{library}/dates",
		Headers:     ign.AuthHeadersOptional,
		Methods: ign.Methods{
			// swagger:route GET /benchmarks benchmark benchmarkLib
			//
			// Get the set of benchmark series for a library
			//
			//   Produces:
			//   - application/json
			//
			//   Schemes: https
			//
			//   Responses:
			//     200: Benchmarks
			ign.Method{
				Type:        "GET",
				Description: "Get benchmark data",
				Handlers: ign.FormatHandlers{
					ign.FormatHandler{
						Extension: ".json",
						Handler:   ign.JSONListResult("", controllers.BenchmarkDates),
					},

					ign.FormatHandler{
						Extension: "",
						Handler:   ign.JSONListResult("", controllers.BenchmarkDates),
					},
				},
			},
		},
		SecureMethods: ign.SecureMethods{},
	},
} // routes
