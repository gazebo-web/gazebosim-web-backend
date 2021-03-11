package main

// Import this file's dependencies
import (
	"gitlab.com/ignitionrobotics/web/web-server/globals"
	"gitlab.com/ignitionrobotics/web/web-server/models"
	"log"
	"time"
)

// DBMigrate auto migrates database tables
func DBMigrate() {
	// Note about Migration from GORM doc:
	// http://jinzhu.me/gorm/database.html#migration
	//
	// WARNING: AutoMigrate will ONLY create tables,
	// missing columns and missing indexes,
	// and WON'T change existing column's type or delete
	// unused columns to protect your data.
	if globals.Server.Db != nil {
		globals.Server.Db.AutoMigrate(&models.Lib{}, &models.Version{},
			&models.BenchmarkLibs{}, &models.Benchmark{}, &models.BenchmarkSeriesInstance{})
		//&models.Benchmark{}, &models.BenchmarkInstance{},
		//&models.BenchmarkContext{}, &models.BenchmarkCache{})
	}
}

func getTime(t string) time.Time {
	result, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// DBAddDefaults adds default values to the database
func DBAddDefaults() {

	libs := models.Libs{
		{
			Name: "Common",
			Versions: models.Versions{
				{
					Major:       0,
					Minor:       1,
					Patch:       0,
					ReleaseDate: getTime("2016-07-29T03:28:00+02:00"),
				},
				{
					Major:       0,
					Minor:       2,
					Patch:       0,
					ReleaseDate: getTime("2017-04-10T19:31:00+02:00"),
				},
				{
					Major:       0,
					Minor:       3,
					Patch:       0,
					ReleaseDate: getTime("2017-05-19T20:46:00+02:00"),
				},
				{
					Major:       0,
					Minor:       4,
					Patch:       0,
					ReleaseDate: getTime("2017-05-30T17:22:00-07:00"),
				},
			},

			Repo:        "https://github.com/ignitionrobotics/ign-common",
			Description: "A collection of useful classes and functions for handling many command tasks. This includes parsing 3D mesh files, managing console output, and using PID controllers.",
		},
		{
			Name: "Math",
			Versions: models.Versions{
				{
					Major:       0,
					Minor:       1,
					Patch:       0,
					ReleaseDate: getTime("2014-06-17T02:36:00+02:00"),
				},
				{
					Major:       0,
					Minor:       2,
					Patch:       0,
					ReleaseDate: getTime("2014-10-29T17:13:00-07:00"),
				},
				{
					Major:       1,
					Minor:       0,
					Patch:       0,
					ReleaseDate: getTime("2015-03-06T10:51:00-08:00"),
				},
				{
					Major:       2,
					Minor:       0,
					Patch:       0,
					ReleaseDate: getTime("2015-04-17T08:51:00-07:00"),
				},
				{
					Major:       2,
					Minor:       0,
					Patch:       1,
					ReleaseDate: getTime("2015-04-17T13:29:00-07:00"),
				},
				{
					Major:       2,
					Minor:       0,
					Patch:       2,
					ReleaseDate: getTime("2015-05-04T14:29:00+02:00"),
				},
				{
					Major:       2,
					Minor:       0,
					Patch:       3,
					ReleaseDate: getTime("2015-05-21T10:47:00-07:00"),
				},
				{
					Major:       2,
					Minor:       1,
					Patch:       0,
					ReleaseDate: getTime("2015-05-21T18:08:00-07:00"),
				},
				{
					Major:       2,
					Minor:       1,
					Patch:       1,
					ReleaseDate: getTime("2015-05-22T07:42:00-07:00"),
				},
				{
					Major:       2,
					Minor:       2,
					Patch:       0,
					ReleaseDate: getTime("2015-07-27T16:49:00-07:00"),
				},
				{
					Major:       2,
					Minor:       2,
					Patch:       1,
					ReleaseDate: getTime("2015-07-28T07:44:00-07:00"),
				},
				{
					Major:       2,
					Minor:       2,
					Patch:       2,
					ReleaseDate: getTime("2015-08-03T17:22:00-07:00"),
				},
				{
					Major:       2,
					Minor:       2,
					Patch:       3,
					ReleaseDate: getTime("2015-09-18T00:07:00+02:00"),
				},
				{
					Major:       2,
					Minor:       3,
					Patch:       0,
					ReleaseDate: getTime("2016-01-15T20:34:00+01:00"),
				},
				{
					Major:       2,
					Minor:       4,
					Patch:       0,
					ReleaseDate: getTime("2016-06-17T15:01:00-07:00"),
				},
				{
					Major:       2,
					Minor:       4,
					Patch:       1,
					ReleaseDate: getTime("2016-07-08T11:48:00-07:00"),
				},
				{
					Major:       2,
					Minor:       5,
					Patch:       0,
					ReleaseDate: getTime("2016-07-21T08:48:00-07:00"),
				},
				{
					Major:       2,
					Minor:       6,
					Patch:       0,
					ReleaseDate: getTime("2016-11-02T08:29:00-07:00"),
				},
				{
					Major:       2,
					Minor:       7,
					Patch:       0,
					ReleaseDate: getTime("2017-01-04T10:18:00-08:00"),
				},
				{
					Major:       2,
					Minor:       8,
					Patch:       0,
					ReleaseDate: getTime("2017-02-14T14:59:00-08:00"),
				},
				{
					Major:       3,
					Minor:       0,
					Patch:       0,
					ReleaseDate: getTime("2017-01-05T11:46:00-08:00"),
				},
				{
					Major:       3,
					Minor:       1,
					Patch:       0,
					ReleaseDate: getTime("2017-04-12T02:15:00+02:00"),
				},
				{
					Major:       3,
					Minor:       2,
					Patch:       0,
					ReleaseDate: getTime("2017-05-15T17:09:00-07:00"),
				},
			},

			Repo:        "https://github.com/ignitionrobotics/ign-math",
			Description: "A small, fast, and high performance math library. This library is a self-contained set of classes and functions suitable for robot applications.",
		},
		{
			Name: "Msgs",
			Versions: models.Versions{
				{
					Major:       0,
					Minor:       1,
					Patch:       0,
					ReleaseDate: getTime("2014-07-20T21:32:00+02:00"),
				},
				{
					Major:       0,
					Minor:       2,
					Patch:       0,
					ReleaseDate: getTime("2016-05-03T14:12:00-07:00"),
				},
				{
					Major:       0,
					Minor:       3,
					Patch:       0,
					ReleaseDate: getTime("2016-05-04T12:00:00-07:00"),
				},
				{
					Major:       0,
					Minor:       3,
					Patch:       1,
					ReleaseDate: getTime("2016-06-08T08:13:00-07:00"),
				},
				{
					Major:       0,
					Minor:       4,
					Patch:       0,
					ReleaseDate: getTime("2016-07-26T02:28:00+02:00"),
				},
				{
					Major:       0,
					Minor:       5,
					Patch:       0,
					ReleaseDate: getTime("2016-08-12T10:26:00-07:00"),
				},
				{
					Major:       0,
					Minor:       6,
					Patch:       0,
					ReleaseDate: getTime("2016-11-01T15:35:00-07:00"),
				},
				{
					Major:       0,
					Minor:       6,
					Patch:       1,
					ReleaseDate: getTime("2016-12-06T15:37:00-08:00"),
				},
				{
					Major:       0,
					Minor:       7,
					Patch:       0,
					ReleaseDate: getTime("2017-01-19T14:56:00+01:00"),
				},
			},
			Repo:        "https://github.com/ignitionrobotics/ign-msgs",
			Description: "Standard set of message definitions, used by Ignition Transport, and other applications.",
		},
		{
			Name: "Transport",
			Versions: models.Versions{
				{
					Major:       0,
					Minor:       3,
					Patch:       0,
					ReleaseDate: getTime("2014-08-13T22:43:00+02:00"),
				},
				{
					Major:       0,
					Minor:       3,
					Patch:       1,
					ReleaseDate: getTime("2014-08-14T01:28:00+02:00"),
				},
				{
					Major:       0,
					Minor:       4,
					Patch:       0,
					ReleaseDate: getTime("2014-10-27T18:15:00-07:00"),
				},
				{
					Major:       0,
					Minor:       4,
					Patch:       1,
					ReleaseDate: getTime("2014-10-29T09:52:00-07:00"),
				},
				{
					Major:       0,
					Minor:       4,
					Patch:       2,
					ReleaseDate: getTime("2014-11-03T18:33:00-08:00"),
				},
				{
					Major:       0,
					Minor:       4,
					Patch:       3,
					ReleaseDate: getTime("2014-11-04T13:52:00-08:00"),
				},
				{
					Major:       0,
					Minor:       4,
					Patch:       4,
					ReleaseDate: getTime("2014-11-06T14:34:00-08:00"),
				},
				{
					Major:       0,
					Minor:       5,
					Patch:       0,
					ReleaseDate: getTime("2014-12-23T20:31:00+01:00"),
				},
				{
					Major:       0,
					Minor:       6,
					Patch:       0,
					ReleaseDate: getTime("2015-01-29T00:03:00+01:00"),
				},
				{
					Major:       0,
					Minor:       7,
					Patch:       0,
					ReleaseDate: getTime("2015-02-19T23:23:00+01:00"),
				},
				{
					Major:       0,
					Minor:       8,
					Patch:       0,
					ReleaseDate: getTime("2015-03-31T14:30:00-07:00"),
				},
				{
					Major:       0,
					Minor:       8,
					Patch:       1,
					ReleaseDate: getTime("2015-04-02T01:45:00+02:00"),
				},
				{
					Major:       0,
					Minor:       9,
					Patch:       0,
					ReleaseDate: getTime("2015-10-07T19:27:00+02:00"),
				},
				{
					Major:       1,
					Minor:       0,
					Patch:       0,
					ReleaseDate: getTime("2016-02-05T18:20:00+01:00"),
				},
				{
					Major:       1,
					Minor:       0,
					Patch:       1,
					ReleaseDate: getTime("2016-02-07T16:29:00-08:00"),
				},
				{
					Major:       1,
					Minor:       1,
					Patch:       0,
					ReleaseDate: getTime("2016-03-08T16:29:00-08:00"),
				},
				{
					Major:       1,
					Minor:       2,
					Patch:       0,
					ReleaseDate: getTime("2016-05-03T21:01:00-07:00"),
				},
				{
					Major:       1,
					Minor:       3,
					Patch:       0,
					ReleaseDate: getTime("2016-07-20T11:18:00-07:00"),
				},
				{
					Major:       1,
					Minor:       4,
					Patch:       0,
					ReleaseDate: getTime("2016-10-05T12:36:00-07:00"),
				},
				{
					Major:       2,
					Minor:       0,
					Patch:       0,
					ReleaseDate: getTime("2016-08-08T17:59:00+02:00"),
				},
				{
					Major:       2,
					Minor:       1,
					Patch:       0,
					ReleaseDate: getTime("2016-10-05T13:22:00-07:00"),
				},
				{
					Major:       3,
					Minor:       0,
					Patch:       0,
					ReleaseDate: getTime("2016-12-16T06:51:00-08:00"),
				},
				{
					Major:       3,
					Minor:       0,
					Patch:       1,
					ReleaseDate: getTime("2017-01-09T23:56:00+01:00"),
				},
			},

			Repo:        "https://github.com/ignitionrobotics/ign-transport",
			Description: "The transport library combines ZeroMQ with Protobufs to create a fast and efficient message passing system. Asynchronous message publication and subscription is provided along with service calls and discovery.",
		},
		{
			Name: "Physics",
			Versions: models.Versions{
				{
					Major:       0,
					Minor:       1,
					Patch:       0,
					ReleaseDate: time.Now(),
				},
			},

			Repo:        "https://github.com/ignitionrobotics/ign-physics",
			Description: "A plugin based interface to physics engines, such as ODE, Bullet, and DART.",
		},
		{
			Name: "Rendering",
			Versions: models.Versions{
				{
					Major:       0,
					Minor:       1,
					Patch:       0,
					ReleaseDate: time.Now(),
				},
			},

			Repo:        "https://github.com/ignitionrobotics/ign-rendering",
			Description: "A plugin based interface to rendering engines, such as OGRE and Optix.",
		},
		{
			Name: "GUI",
			Versions: models.Versions{
				{
					Major:       0,
					Minor:       1,
					Patch:       0,
					ReleaseDate: time.Now(),
				},
			},
			Repo:        "https://github.com/ignitionrobotics/ign-gui",
			Description: "A framework for graphical user interfaces centered around QT. Each component in Ignition GUI is an independent plugin",
		},
		{
			Name: "Sensors",
			Versions: models.Versions{
				{
					Major:       0,
					Minor:       1,
					Patch:       0,
					ReleaseDate: time.Now(),
				},
			},
			Repo:        "https://github.com/ignitionrobotics/ign-sensors",
			Description: "A large set of sensor and noise models suitable for generating realistic data in simulation.",
		},
		{
			Name: "SDFormat",
			Versions: models.Versions{
				{
					Major:       0,
					Minor:       1,
					Patch:       0,
					ReleaseDate: time.Now(),
				},
			},
			Repo:        "https://github.com/osrf/sdformat",
			Description: "Simulation Description Format parser and description files.",
		},
		{
			Name: "Utils",
			Versions: models.Versions{
				{
					Major:       0,
					Minor:       1,
					Patch:       0,
					ReleaseDate: time.Now(),
				},
			},

			Repo:        "https://github.com/ignitionrobotics/ign-utils",
			Description: "General purpose classes and functions with minimal dependencies. It includes command line parsing, a helper class to implement the PIMPL pattern, macros to suppress warnings, etc.",
		},
	}

	// Insert the libraries if they do not exist
	for _, lib := range libs {
		var tLib models.Lib
		if err := globals.Server.Db.Model(&models.Lib{}).Where("name = ?", lib.Name).Preload("Versions").First(&tLib).Error; err != nil && err.Error() == "record not found" {
			globals.Server.Db.Create(&lib)
		} else {

			// Update versions
			for _, v1 := range lib.Versions {
				found := false
				for j, v2 := range tLib.Versions {
					if v2.Major == v1.Major &&
						v2.Minor == v1.Minor &&
						v2.Patch == v1.Patch {

						found = true
						tLib.Versions[j].ReleaseDate = v1.ReleaseDate
					}
				}

				if !found {
					tLib.Versions = append(tLib.Versions, v1)
				}
			}

			tLib.Repo = lib.Repo
			tLib.Description = lib.Description
			globals.Server.Db.Save(&tLib)
		}
	}
}
