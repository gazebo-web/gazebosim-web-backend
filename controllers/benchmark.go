package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"time"

	"encoding/json"
	"github.com/jinzhu/gorm"
	"gitlab.com/ignitionrobotics/web/ign-go"
	"gitlab.com/ignitionrobotics/web/web-server/models"
	"net/http"
)

type singleDateResult struct {
	Date time.Time `json:"-"`
}

func (b *singleDateResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Date.Format("2006-01-02"))
}

// getBenchmarkDate is a helper function that returns data for a particular library and date
func getBenchmarkDate(tx *gorm.DB, w http.ResponseWriter, r *http.Request, benchmarkLibID uint, date string) (interface{}, *ign.ErrMsg) {
	type Results struct {
		Name     string  `json:"name"`
		CPUTime  float64 `json:"cpu_time"`
		RealTime float64 `json:"real_time"`
	}

	var results []Results
	tx.Raw("select benchmarks.name, benchmark_series_instances.cpu_time, benchmark_series_instances.real_time FROM benchmark_series_instances INNER JOIN benchmarks on benchmark_series_instances.benchmark_id = benchmarks.id where date(benchmark_series_instances.date) = ?", date).Scan(&results)

	return results, nil
}

// getBenchmarkSeries is a helper function that returns chart data for a particular library and benchmark name
func getBenchmarkSeries(tx *gorm.DB, w http.ResponseWriter, r *http.Request, benchmarkLibID uint, benchName string) (interface{}, *ign.ErrMsg) {
	// Convert the data to a format the suitable for ngx-charts.
	type DataPoint struct {
		Name  time.Time `json:"name"`
		Value float64   `json:"value"`
	}
	type ChartSeries struct {
		Name   string      `json:"name"`
		Series []DataPoint `json:"series"`
	}
	type ChartData []ChartSeries

	var chartData ChartData

	var benchmark models.Benchmark
	if err := tx.Model(&models.Benchmark{}).Preload("Series").Where("name = ? && benchmark_libs_id = ?", benchName, benchmarkLibID).First(&benchmark).Error; err != nil {
		return nil, ign.NewErrorMessageWithBase(ign.ErrorNameNotFound, err)
	}

	var cpuSeries ChartSeries
	var realSeries ChartSeries
	cpuSeries.Name = benchmark.Name + ": CPU Time"
	realSeries.Name = benchmark.Name + ": Real Time"

	for _, seriesInstance := range benchmark.Series {
		cpuSeries.Series = append(cpuSeries.Series, DataPoint{
			Name:  seriesInstance.Date,
			Value: seriesInstance.CPUTime,
		})

		realSeries.Series = append(realSeries.Series, DataPoint{
			Name:  seriesInstance.Date,
			Value: seriesInstance.RealTime,
		})
	}
	chartData = append(chartData, cpuSeries)
	chartData = append(chartData, realSeries)

	return chartData, nil
}

// Benchmarks returns either the set of benchmark test names, or data series for a particular benchmark.
// You can request all the benchmark test names for a library using:
// curl -k -X GET http://localhost:8000/1.0/benchmarks/{library}
//
// You can request the data series for a benchmark using:
// curl -k -X GET http://localhost:8000/1.0/benchmarks/{library}?benchmark={benchmark_name}
func Benchmarks(tx *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, *ign.ErrMsg) {
	// Get the library name from the URL
	params := mux.Vars(r)
	libName, valid := params["library"]
	if !valid {
		return nil, ign.NewErrorMessage(ign.ErrorNameNotFound)
	}

	// Get the benchmark lib data
	var benchmarkLib models.BenchmarkLibs
	if err := tx.Model(&models.BenchmarkLibs{}).Preload("Benchmarks").Where("name = ?", libName).First(&benchmarkLib).Error; err != nil {
		return nil, ign.NewErrorMessageWithBase(ign.ErrorNameNotFound, err)
	}

	// Check if a benchmark query parameter was set
	queryP := r.URL.Query()
	benchNames, hasBenchName := queryP["benchmark"]
	if hasBenchName {
		return getBenchmarkSeries(tx, w, r, benchmarkLib.ID, benchNames[0])
	}

	// Check if a date query parameter was set
	dates, hasDate := queryP["date"]
	if hasDate {
		return getBenchmarkDate(tx, w, r, benchmarkLib.ID, dates[0])
	}

	return benchmarkLib.Benchmarks, nil
}

// BenchmarkDates returns the set of dates that have benchmark data.
// You can request the data series for a benchmark using:
// curl -k -X GET http://localhost:8000/1.0/benchmarks/{library}/dates
func BenchmarkDates(tx *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, *ign.ErrMsg) {
	// Get the library name from the URL
	params := mux.Vars(r)
	libName, valid := params["library"]
	if !valid {
		return nil, ign.NewErrorMessage(ign.ErrorNameNotFound)
	}

	// Get the benchmark lib data
	var benchmarkLib models.BenchmarkLibs
	if err := tx.Where("name = ?", libName).First(&benchmarkLib).Error; err != nil {
		return nil, ign.NewErrorMessageWithBase(ign.ErrorNameNotFound, err)
	}
	var results []singleDateResult
	tx.Raw("select distinct benchmark_series_instances.date FROM benchmark_series_instances INNER JOIN benchmarks ON benchmark_series_instances.benchmark_id = benchmarks.id WHERE benchmarks.benchmark_libs_id = ?", benchmarkLib.ID).Scan(&results)

	return results, nil
}

// BenchmarkSummary returns all the benchmark libs along with the set of available tests.
// You can request this information using:
// curl -k -X GET http://localhost:8000/1.0/benchmarks
func BenchmarkSummary(tx *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, *ign.ErrMsg) {

	var benchmarkLib []models.BenchmarkLibs
	if err := tx.Model(&models.BenchmarkLibs{}).Preload("Benchmarks").Find(&benchmarkLib).Error; err != nil {
		return nil, ign.NewErrorMessageWithBase(ign.ErrorNameNotFound, err)
	}
	return benchmarkLib, nil
}

// BenchmarkCreate creates a new benchmark metric
// You can request this method with the following curl request:
// curl -k -X POST -d @result.json http://localhost:8000/1.0/benchmarks/{library}
func BenchmarkCreate(tx *gorm.DB, w http.ResponseWriter, r *http.Request) (interface{}, *ign.ErrMsg) {

	// Get the library name from the URL
	params := mux.Vars(r)
	libName, valid := params["library"]
	if !valid {
		return nil, ign.NewErrorMessage(ign.ErrorNameNotFound)
	}

	// Decode the benchmark submission data
	var benchmarkSubmission models.BenchmarkSubmission
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&benchmarkSubmission)
	if err != nil {
		return nil, ign.NewErrorMessageWithArgs(ign.ErrorUnmarshalJSON, err,
			[]string{"Incorrect POST data"})
	}

	// Convert the date from JSON to golang time.
	dateLayout := "2006-01-02 15:04:05"
	submissionDate, timeErr := time.Parse(dateLayout, benchmarkSubmission.Context.Date)
	if timeErr != nil {
		fmt.Println(timeErr)
	}

	// Get the library entry to use, or create one if it doesn't exist.
	var benchmarkLib models.BenchmarkLibs
	if err := tx.Where("name = ?", libName).First(&benchmarkLib).Error; err != nil {
		benchmarkLib.Name = libName
		tx.Create(&benchmarkLib)
		tx.Where("name = ?", libName).First(&benchmarkLib)
	}

	// Process each submission
	for _, bSubmission := range benchmarkSubmission.Benchmarks {
		var benchmark models.Benchmark

		// Get the Benchmark based on the bSubmission.Name. Create the benchmark if the name doesn't exist.
		if err := tx.Where("name = ? && benchmark_libs_id = ?", bSubmission.Name, benchmarkLib.ID).First(&benchmark).Error; err != nil {
			benchmark.Name = bSubmission.Name
			benchmarkLib.Benchmarks = append(benchmarkLib.Benchmarks, benchmark)
			tx.Save(&benchmarkLib)
			tx.Where("name = ? && benchmark_libs_id = ?", bSubmission.Name, benchmarkLib.ID).First(&benchmark)
		}

		// Get the raw json so that we can keep the original benchmark
		var rawJSON []byte
		rawJSON, _ = json.Marshal(bSubmission)

		// Create the benchmark series instance
		benchmarkSeriesInstance := models.BenchmarkSeriesInstance{
			Date:       submissionDate,
			RunType:    bSubmission.RunType,
			Iterations: bSubmission.Iterations,
			RealTime:   bSubmission.RealTime,
			CPUTime:    bSubmission.CPUTime,
			TimeUnit:   bSubmission.TimeUnit,
			Raw:        rawJSON,
		}

		// Store and save the data
		benchmark.Series = append(benchmark.Series, benchmarkSeriesInstance)
		tx.Save(&benchmark)
	}

	return benchmarkSubmission, nil
}
