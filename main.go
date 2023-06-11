package main

import (
	"fmt"
	"sync"
	"time"

	geo "github.com/emanuelef/go-geo-3d"
)

func main() {
	// Coordinates are in degrees and altitude in metres
	start := geo.NewCoord3d(51.39674, -0.36148, 1104.9)
	end := geo.NewCoord3d(51.38463, -0.36819, 1219.2)

	// Distance in metres between two 3D coordinates
	distance := geo.Distance3D(start, end)

	fmt.Printf("Distance 3D line from start to end: %.3fm\n", distance)

	posA := geo.NewCoord3d(51.3909, -0.364, 15)
	// Minimum distance in metres from one 3D point to a project line in 3D coordinates
	minPoint, _ := posA.ClosestPointOnLine(start, end)
	distance = geo.Distance3D(posA, minPoint)

	fmt.Printf("Distance from one point to a line: %.3fm\n", distance)

	startTime := time.Now().UTC()

	var wg sync.WaitGroup

	for i := 0; i <= 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			minPoint, _ := posA.ClosestPointOnLine(start, end)
			distance = geo.Distance3D(posA, minPoint)
		}()
	}
	wg.Wait()

	totalTime := time.Now().UTC().Sub(startTime)

	fmt.Printf("Time taken running 1M requests %d ms\n", int(totalTime.Milliseconds()))
}
