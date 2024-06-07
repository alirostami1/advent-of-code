package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

func main() {
	distances := map[string]map[string]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		var (
			start, end string
			dist       int
		)
		fmt.Sscanf(input, "%s to %s = %d", &start, &end, &dist)
		if distances[start] == nil {
			distances[start] = map[string]int{}
		}
		distances[start][end] = dist
		if distances[end] == nil {
			distances[end] = map[string]int{}
		}
		distances[end][start] = dist
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	closest := ClosestRoute(distances)

	fmt.Printf("%+v\n", distances)
	fmt.Println(closest)
}

func ClosestRoute(distances map[string]map[string]int) int {
	closestRoute := math.MaxInt
	for start, connections := range distances {
		fmt.Printf("starting from %s with connections %+v\n", start, connections)
		visited := make([]string, 0, len(distances))
		visited = append(visited, start)
		ClosestRouteFromCity(start, visited, connections, distances)
		visited = visited[0 : len(visited)-1]
	}
	return closestRoute
}

func ClosestRouteFromCity(cityName string, visited []string, neighbors map[string]int, distances map[string]map[string]int) int {
	closestRoute := math.MaxInt
	hasUnvisitedNeighbors := false
	for dest, distance := range neighbors {
		if !slices.Contains(visited, dest) {
			hasUnvisitedNeighbors = true
			for i := 0; i < len(visited); i++ {
				fmt.Printf("+")
			}
			fmt.Printf("visited = %+v\n", visited)
			visited = append(visited, dest)
			closestRouteFromThisCity := ClosestRouteFromCity(dest, visited, distances[dest], distances)
			fmt.Printf("reaching city %s with distance = %d, closestRoute =  %d, closestRouteFromThisCity = %d\n", dest, distance, closestRoute, closestRouteFromThisCity)
			if closestRouteFromThisCity+distance < closestRoute {
				closestRoute = closestRouteFromThisCity
			}
			visited = visited[0 : len(visited)-1]
		}
	}

	if !hasUnvisitedNeighbors {
		return 0
	}

	return closestRoute
}
