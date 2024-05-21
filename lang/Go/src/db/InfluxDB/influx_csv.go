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

	influxRunID = ""
)

//////////////////////////////////////////////////////////////////////////
// interface/structure definitions

type csvI interface {
	getExpectedHeader() []string
	process(d csvI)
	parseRecord(record []string) (tags map[string]string, fields map[string]interface{})
}

//==========================================================================
// csvS

type csvS struct {
	expectedHeader []string
}

func (c *csvS) getExpectedHeader() []string {
	return c.expectedHeader
}

func (c *csvS) process(d csvI) {
	// Init vars from system environment
	if ev, ok := os.LookupEnv("c_cfg_Influxdb"); ok {
		influxAddr = ev
	}
	if ev, ok := os.LookupEnv("c_cfg_DbNmae"); ok {
		influxDBNm = ev
	}
	if ev, ok := os.LookupEnv("c_cfg_Measure"); ok {
		influxMeas = ev
	}
	if ev, ok := os.LookupEnv("c_cfg_CSV"); ok {
		influxSCsv = ev
	}
	influxRunID = os.Getenv("c_cfg_RunID")

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
	expectedHeader := d.getExpectedHeader()
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

		tags, fields := d.parseRecord(record)

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

//==========================================================================
// ages

type ages struct {
	csvS
}

func newAges() ages {
	return ages{csvS: csvS{expectedHeader: []string{"timestamp", "aget", "agev"}}}
}

func (a *ages) process(d csvI) {
	a.csvS.process(d)
}

func (a *ages) getExpectedHeader() []string {
	return a.expectedHeader
}

func (a *ages) parseRecord(record []string) (tags map[string]string, fields map[string]interface{}) {

	// Parse the fields
	agev, err := strconv.Atoi(record[2])
	if err != nil {
		log.Fatalf("could not convert age to int: %v", err)
	}

	// Create a point
	tags = map[string]string{
		"runId": influxRunID,
		"aget":  record[1],
	}
	fields = map[string]interface{}{
		"agev": agev,
	}

	return //tags, fields
}

//==========================================================================
// cpu

type cpu struct {
	csvS
}

func newCpu() cpu {
	return cpu{csvS: csvS{expectedHeader: []string{"timeStamp", "elapsed", "label"}}}
}

func (c *cpu) process(d csvI) {
	c.csvS.process(d)
}

func (c *cpu) getExpectedHeader() []string {
	return c.expectedHeader
}

func (c *cpu) parseRecord(record []string) (tags map[string]string, fields map[string]interface{}) {

	// Parse the fields
	cpuv, err := strconv.Atoi(record[1])
	if err != nil {
		log.Fatalf("could not convert cpu to int: %v", err)
	}

	// Create a point
	tags = map[string]string{
		"runId": influxRunID,
		"label": record[2],
	}
	fields = map[string]interface{}{
		"cpu": float64(cpuv) / 1000,
	}

	return //tags, fields
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func testAges() {
	a := newAges()
	a.process(&a)
}

func testCpu() {
	c := newCpu()
	c.process(&c)
}

// Function main
func main() {
	testCpu()
}
