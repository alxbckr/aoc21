package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Vertex struct {
	name string
}

type Set map[Vertex]struct{}

type Graph struct {
	adjacency map[Vertex]*Set
}

type Path struct {
	nodes        []Vertex
	visitedNodes map[Vertex]int
}

func (s *Set) Add(elem Vertex) {
	(*s)[elem] = struct{}{}
}

func NewSet() *Set {
	return &Set{}
}

func NewGraph() *Graph {
	return &Graph{adjacency: map[Vertex]*Set{}}
}

func (g *Graph) AddVertex(v Vertex) {
	if _, exists := g.adjacency[v]; !exists {
		g.adjacency[v] = NewSet()
	}
}

func (g *Graph) AddEdge(v1, v2 Vertex) {
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.adjacency[v1].Add(v2)
	g.adjacency[v2].Add(v1)
}

func (g *Graph) GetStart() *Vertex {
	for v := range g.adjacency {
		if v.name == "start" {
			return &v
		}
	}
	return &Vertex{}
}

func (g *Graph) Dump() {
	for vertex, edges := range g.adjacency {
		fmt.Printf("[%v]: ", vertex.name)
		for edge := range *edges {
			fmt.Printf("[%v] ", edge.name)
		}
		fmt.Println()
	}
}

func NewPath() *Path {
	return &Path{nodes: []Vertex{}, visitedNodes: map[Vertex]int{}}
}

func (p *Path) Add(v Vertex) {
	p.nodes = append(p.nodes, v)
	p.visitedNodes[v] += 1
}

func (p *Path) canRevisit(v Vertex) bool {
	return p.visitedNodes[v] == 0 || checkUppercase(v.name)
}

func checkUppercase(s string) bool {
	return s >= "A" && s <= "Z"
}

func buildPaths(g *Graph, start *Vertex, paths *[]Path, currPath Path) {
	for e := range *(g.adjacency[*start]) {
		if !currPath.canRevisit(e) {
			continue
		}
		pathNew := NewPath()
		for _, v := range currPath.nodes {
			pathNew.Add(v)
		}
		pathNew.Add(e)
		if e.name == "end" {
			*paths = append(*paths, *pathNew)
			continue
		}
		buildPaths(g, &e, paths, *pathNew)
	}
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer f.Close()

	g := NewGraph()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		vertexes := strings.Split(input, "-")
		g.AddEdge(Vertex{name: vertexes[0]},
			Vertex{name: vertexes[1]})
	}

	g.Dump()

	paths := []Path{}
	startPath := NewPath()
	startPath.Add(*g.GetStart())
	buildPaths(g, g.GetStart(), &paths, *startPath)
	for _, path := range paths {
		for _, v := range path.nodes {
			fmt.Printf("%v,", v.name)
		}
		fmt.Println()
	}

	fmt.Printf("Num paths: %v\n", len(paths))

}
