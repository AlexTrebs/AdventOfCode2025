package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coordinates struct {
	x, y, z int
}

type Edge struct {
	from, to Coordinates
	distance int
}

type UnionFind struct {
	parent map[Coordinates]Coordinates
	rank   map[Coordinates]int
	coords []Coordinates
}

func NewUnionFind(coords []Coordinates) *UnionFind {
	uf := &UnionFind{
		parent: make(map[Coordinates]Coordinates),
		rank:   make(map[Coordinates]int),
		coords: coords,
	}

	for _, coord := range coords {
		uf.parent[coord] = coord
		uf.rank[coord] = 0
	}

	return uf
}

func (uf *UnionFind) Find(coord Coordinates) Coordinates {
	if uf.parent[coord] != coord {
		uf.parent[coord] = uf.Find(uf.parent[coord])
	}
	return uf.parent[coord]
}

func (uf *UnionFind) Union(a, b Coordinates) bool {
	rootA := uf.Find(a)
	rootB := uf.Find(b)

	if rootA == rootB {
		return false
	}

	if uf.rank[rootA] < uf.rank[rootB] {
		uf.parent[rootA] = rootB
	} else if uf.rank[rootA] > uf.rank[rootB] {
		uf.parent[rootB] = rootA
	} else {
		uf.parent[rootB] = rootA
		uf.rank[rootA]++
	}

	return true
}

func (uf *UnionFind) GetCircuitSizes() []int {
	circuits := make(map[Coordinates][]Coordinates)

	for _, coord := range uf.coords {
		root := uf.Find(coord)
		circuits[root] = append(circuits[root], coord)
	}

	sizes := make([]int, 0, len(circuits))
	for _, circuit := range circuits {
		sizes = append(sizes, len(circuit))
	}

	return sizes
}

func square(i int) int {
	return i * i
}

func calculateDistance(from, to Coordinates) int {
	return square(from.x-to.x) + square(from.y-to.y) + square(from.z-to.z)
}

func getAllEdges(coords []Coordinates) []Edge {
	var edges []Edge

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			distance := calculateDistance(coords[i], coords[j])
			edges = append(edges, Edge{
				from:     coords[i],
				to:       coords[j],
				distance: distance,
			})
		}
	}

	return edges
}

func main() {
	file, _ := os.ReadFile("input.txt")
	var coords []Coordinates

	for line := range strings.SplitSeq(string(file), "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		parts := strings.Split(line, ",")

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		coords = append(coords, Coordinates{x: x, y: y, z: z})
	}

	edges := getAllEdges(coords)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	uf := NewUnionFind(coords)
	i := 0
	var edge Edge

	for len(uf.GetCircuitSizes()) > 1 {
		edge = edges[i]
		uf.Union(edge.from, edge.to)

		if i == 1000 {
			sizes := uf.GetCircuitSizes()

			sort.Slice(sizes, func(i, j int) bool {
				return sizes[i] > sizes[j]
			})

			result := sizes[0] * sizes[1] * sizes[2]
			fmt.Println("Part 1:", result)
		}

		i += 1
	}

	fmt.Println("Part 2:", edge.from.x*edge.to.x)
}
