package main

import (
	"gitlab.com/ignitionrobotics/web/ign-go"
	"gitlab.com/ignitionrobotics/web/web-server/controllers"
)

// Routes declares the routes
var Routes = ign.Routes{

	// Route for all models
	ign.Route{
		"Libs",
		"Information about all Libs",
		"/libs",
		ign.AuthHeadersOptional,
		ign.Methods{
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
				"GET",
				"Get all libraries",
				ign.FormatHandlers{
					ign.FormatHandler{".json",
						ign.JSONListResult("", controllers.LibsList)},

					ign.FormatHandler{"", ign.JSONListResult("", controllers.LibsList)},
				},
			},
			ign.Method{
				"POST",
				"Add a new libary",
				ign.FormatHandlers{
					ign.FormatHandler{".json",
						ign.JSONListResult("", controllers.LibsCreate)},

					ign.FormatHandler{"",
						ign.JSONListResult("", controllers.LibsCreate)},
				},
			},
		},
		ign.SecureMethods{},
	},

	// Route for versions
	ign.Route{
		"Versions",
		"Routes for lib versions",
		"/versions",
		ign.AuthHeadersOptional,
		ign.Methods{
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
				"POST",
				"Add a new version",
				ign.FormatHandlers{
					ign.FormatHandler{".json",
						ign.JSONListResult("", controllers.VersionCreate)},

					ign.FormatHandler{"",
						ign.JSONListResult("", controllers.VersionCreate)},
				},
			},
		},
		ign.SecureMethods{},
	},

	// Route for documentation
	ign.Route{
		"Docs",
		"Routes for documentation",
		"/docs",
		ign.AuthHeadersOptional,
		ign.Methods{
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
				"GET",
				"Get list of documentation pages",
				ign.FormatHandlers{
					ign.FormatHandler{".json",
						ign.JSONListResult("", controllers.Docs)},

					ign.FormatHandler{"",
						ign.JSONListResult("", controllers.Docs)},
				},
			},
		},
		ign.SecureMethods{},
	},
	// Route for documentation
	ign.Route{
		"Docs",
		"Routes for documentation",
		"/docs/{version}/{page}",
		ign.AuthHeadersOptional,
		ign.Methods{
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
				"GET",
				"Get documentation",
				ign.FormatHandlers{
					ign.FormatHandler{".json",
						ign.JSONListResult("", controllers.DocsPage)},

					ign.FormatHandler{"",
						ign.JSONListResult("", controllers.DocsPage)},
				},
			},
		},
		ign.SecureMethods{},
	},
	// Route for library benchmarks
	ign.Route{
		"Benchmark",
		"Routes for benchmarks",
		"/benchmarks",
		ign.AuthHeadersOptional,
		ign.Methods{
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
				"GET",
				"Get benchmark data",
				ign.FormatHandlers{
					ign.FormatHandler{".json",
						ign.JSONListResult("", controllers.BenchmarkSummary)},

					ign.FormatHandler{"",
						ign.JSONListResult("", controllers.BenchmarkSummary)},
				},
			},
		},
		ign.SecureMethods{},
	},
	// Route for library benchmarks
	ign.Route{
		"Benchmark",
		"Routes for benchmarks",
		"/benchmarks/{library}",
		ign.AuthHeadersOptional,
		ign.Methods{
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
				"GET",
				"Get benchmark data",
				ign.FormatHandlers{
					ign.FormatHandler{".json",
						ign.JSONListResult("", controllers.Benchmarks)},

					ign.FormatHandler{"",
						ign.JSONListResult("", controllers.Benchmarks)},
				},
			},
		},
		ign.SecureMethods{
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
				"POST",
				"Post benchmark data",
				ign.FormatHandlers{
					ign.FormatHandler{".json",
						ign.JSONListResult("", controllers.BenchmarkCreate)},

					ign.FormatHandler{"",
						ign.JSONListResult("", controllers.BenchmarkCreate)},
				},
			},
		},
	},
	// Route for library benchmarks
	ign.Route{
		"Benchmark dates",
		"Routes for benchmarks",
		"/benchmarks/{library}/dates",
		ign.AuthHeadersOptional,
		ign.Methods{
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
				"GET",
				"Get benchmark data",
				ign.FormatHandlers{
					ign.FormatHandler{".json",
						ign.JSONListResult("", controllers.BenchmarkDates)},

					ign.FormatHandler{"",
						ign.JSONListResult("", controllers.BenchmarkDates)},
				},
			},
		},
		ign.SecureMethods{},
	},
} // routes
