package main

import "fmt"

// Graph represents a directed graph using an adjacency list.
type Graph struct {
	vertices []*Vertex
}

// Vertex represents a node in the graph.
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a new vertex with the given key to the graph.
func (g *Graph) AddVertex(k int) {
	if g.contains(g.vertices, k) {
		fmt.Printf("vertex %d not added: key already exists\n", k)
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

// AddEdge adds a directed edge from one vertex to another.
func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	switch {
	case fromVertex == nil || toVertex == nil:
		fmt.Printf("Invalid edge (%d --> %d): vertex not found\n", from, to)
	case g.contains(fromVertex.adjacent, to):
		fmt.Printf("Duplicate edge (%d --> %d): already exists\n", from, to)
	default:
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// Print displays the adjacency list of the graph.
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("Vertex %d:", v.key)
		for _, neighbor := range v.adjacent {
			fmt.Printf(" %d", neighbor.key)
		}
		fmt.Println()
	}
}

func (g *Graph) GetVertex(k int) *Vertex {
	for _, v := range g.vertices {
		if v.key == k {
			return v
		}
	}
	return nil
}

func (g *Graph) contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}
	}
	return false
}

func main() {
	g := &Graph{}

	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}

	g.AddEdge(1, 2)
	g.AddEdge(6, 2)
	g.AddEdge(3, 2)
	g.AddEdge(3, 2)
	g.Print()
}
