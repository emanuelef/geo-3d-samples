package main

import (
	"encoding/csv"
	"io"
	"strconv"

	//"encoding/json"
	"fmt"
	"log"
	"os"

	geo "github.com/emanuelef/go-geo-3d"
)

func main() {
	csvFile, err := os.Open(fmt.Sprintf("point-segment.csv"))
	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	firstRow := true
	headerMap := make(map[string]int)

	csvReader := csv.NewReader(csvFile)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if firstRow {
			for i, v := range rec {
				headerMap[v] = i
			}
			firstRow = false
			continue
		}

		posLat, _ := strconv.ParseFloat(rec[headerMap["posLat"]], 64)
		posLon, _ := strconv.ParseFloat(rec[headerMap["posLon"]], 64)
		posAlt, _ := strconv.ParseFloat(rec[headerMap["posAlt"]], 64)

		posA := geo.NewCoord3d(posLat, posLon, posAlt)

		startLat, _ := strconv.ParseFloat(rec[headerMap["startLat"]], 64)
		startLon, _ := strconv.ParseFloat(rec[headerMap["startLon"]], 64)
		startAlt, _ := strconv.ParseFloat(rec[headerMap["startAlt"]], 64)

		start := geo.NewCoord3d(startLat, startLon, startAlt)

		endLat, _ := strconv.ParseFloat(rec[headerMap["endLat"]], 64)
		endLon, _ := strconv.ParseFloat(rec[headerMap["endLon"]], 64)
		endAlt, _ := strconv.ParseFloat(rec[headerMap["endAlt"]], 64)

		end := geo.NewCoord3d(endLat, endLon, endAlt)

		minPoint, _ := posA.ClosestPointOnLine(start, end)
		distance := geo.Distance3D(posA, minPoint)

		fmt.Println(distance)
	}
}
