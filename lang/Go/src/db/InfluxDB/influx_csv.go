package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	influxdb "github.com/influxdata/influxdb1-client/v2"
)

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// Record represents the structure of a CSV record
// type Record struct {
// 	Timestamp string  `csv:"timestamp"`
// 	AgeT   string  `csv:"aget"`
// 	AgeV    float64 `csv:"agev"`
// 	// Field2    float64 `csv:"field2"`
// }

const (
	DateTimeMs = "2006-01-02 15:04:05.000"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	influxAddr = "http://localhost:8086"
	influxUser = "yourusername"
	influxPass = "yourpassword"
	influxDBNm = "influx"
	influxMeas = "example_measurement"
	influxSCsv = "data.csv" // source CSV
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	// Open the CSV file
	file, err := os.Open(influxSCsv)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	// Read the CSV header
	header, err := reader.Read()
	if err != nil {
		log.Fatalf("Error reading CSV header: %v", err)
	}

	// Validate header
	expectedHeader := []string{"timestamp", "aget", "agev"}
	for i, v := range expectedHeader {
		if header[i] != v {
			log.Fatalf("Invalid CSV header, expected %s but got %s", expectedHeader, header)
		}
	}

	// Read the CSV records
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV records: %v", err)
	}

	// Connect to InfluxDB
	client, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr:     influxAddr,
		Username: influxUser,
		Password: influxPass,
	})
	if err != nil {
		log.Fatalf("Error creating InfluxDB client: %v", err)
	}
	defer client.Close()

	// Create a new batch of points
	bp, err := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
		Database:  influxDBNm,
		Precision: "ms",
	})
	if err != nil {
		log.Fatalf("Error creating batch points: %v", err)
	}

	// Iterate over CSV records and add them to the batch
	for _, record := range records {
		// Parse the timestamp
		timestamp, err := time.Parse(DateTimeMs, record[0])
		if err != nil {
			log.Fatalf("Error parsing timestamp: %v", err)
		}

		// Parse the fields
		agev, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatalf("could not convert age to int: %v", err)
		}

		// field2, err := strconv.ParseFloat(record[4], 64)
		// if err != nil {
		// 	log.Fatalf("Error parsing field2: %v", err)
		// }

		// Create a point
		tags := map[string]string{
			"aget": record[1],
			// "tagKey2": record[2],
		}
		fields := map[string]interface{}{
			"agev": agev,
			// "field2": field2,
		}

		pt, err := influxdb.NewPoint(influxMeas, tags, fields, timestamp)
		if err != nil {
			log.Fatalf("Error creating point: %v", err)
		}

		// Add the point to the batch
		bp.AddPoint(pt)
	}

	// Write the batch to InfluxDB
	if err := client.Write(bp); err != nil {
		log.Fatalf("Error writing batch points to InfluxDB: %v", err)
	}

	fmt.Println("Data successfully written to InfluxDB!")
}
