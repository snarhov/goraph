package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

// Use Pointer when we need to update the struct with receiver
// https://golang.org/doc/faq#methods_on_values_or_pointers

// Work Flow
// 1. Create a graph `Data`.
// 2. Create a `Vertex`.
// 3. Add a `Vertex` to a graph Data.
// 4. Connect with an Edge with `AddEdge`

// Data contains graph data, represented in adjacency list and slice.
type Data struct {
	Vertices []*Vertex

	sync.Mutex

	// OutEdges maps each Vertex to its outgoing edges
	OutEdges map[*Vertex][]Edge

	// InEdges maps each Vertex to its incoming edges
	InEdges map[*Vertex][]Edge
}

// Vertex is a vertex(node) in Graph.
type Vertex struct {
	// ID of Vertex is assumed to be unique between vertices.
	ID string

	// Color is used for graph traversal.
	Color string

	sync.Mutex

	// StamMap stores stamp records for several graph algorithms.
	StampMap map[string]float64
}

// Edge is an edge(arc) in a graph that has direction from one to another vertex.
type Edge struct {
	// Vtx can be either source or destination
	Vtx *Vertex

	// Weight contains the weight value in float64.
	Weight float64
}

// NewData returns a new Data.
func NewData() *Data {
	return &Data{
		Vertices: []*Vertex{},
		OutEdges: make(map[*Vertex][]Edge),
		InEdges:  make(map[*Vertex][]Edge),
	}
}

// NewVertex returns a new Vertex.
func NewVertex(id string) *Vertex {
	return &Vertex{
		ID:      id,
		Color:   "",
		StamMap: make(map[string]float64),
	}
}

// AddVertex adds a vertex to a graph Data.
func (d *Data) AddVertex(vtx *Vertex) {
	d.Vertices = append(d.Vertices, vtx)
}

// Connect adds an edge from src to dst Vertex, to a graph Data.
func (d *Data) Connect(src, dst *Vertex, weight float64) {
	edgeSrc := Edge{
		Vtx:    src,
		Weight: weight,
	}
	edgeDst := Edge{
		Vtx:    dst,
		Weight: weight,
	}
	d.Mutex.Lock()
	if _, ok := d.OutEdges[src]; !ok {
		d.OutEdges[src] = []Edge{edgeDst}
	} else {
		// if OutEdges already exists
		duplicate := false
		for _, elem := range d.OutEdges[src] {
			// if there is a duplicate(parallel) edge
			if elem.Vtx == src {
				log.Println("Duplicate(Parallel) Edge Found. Overwriting the Weight value.")
				log.Printf("%v --> %v + %v\n", elem.Weight, elem.Weight, weight)
				elem.Weight += weight
				duplicate = true
				break
			}
		}
		// if this is just another edge from `src` Vertex
		if !duplicate {
			d.OutEdges[src] = append(d.OutEdges[src], edgeDst)
		}
	}
	if _, ok := d.InEdges[dst]; !ok {
		d.InEdges[dst] = []Edge{edgeSrc}
	} else {
		// if InEdges already exists
		duplicate := false
		for _, elem := range d.InEdges[dst] {
			// if there is a duplicate(parallel) edge
			if elem.Vtx == dst {
				log.Println("Duplicate(Parallel) Edge Found. Overwriting the Weight value.")
				log.Printf("%v --> %v + %v\n", elem.Weight, elem.Weight, weight)
				elem.Weight += weight
				duplicate = true
				break
			}
		}
		// if this is just another edge to `dst` Vertex
		if !duplicate {
			d.InEdges[dst] = append(d.InEdges[dst], edgeSrc)
		}
	}
	d.Mutex.Unlock()
}

// Init initializes the graph Data.
func (d *Data) Init() {
	// (X) d = NewData()
	// this only updates the pointer
	//
	// Do this.
	*d = *NewData()
}

// GetVertexSize returns the size of Vertex of the graph Data.
func (d Data) GetVertexSize() int64 {
	return int64(len(d.Vertices))
}

// String describes the graph Data.
func (d Data) String() string {
	if d.GetVertexSize() == 0 {
		return "Graph is empty."
	}
	slice := []string{}
	for _, vtx := range d.Vertices {
		slice = append(slice, fmt.Sprintf("Vertex: %s", vtx.ID))
		d.Mutex.Lock()
		if _, ok := d.OutEdges[vtx]; !ok {
			slice = append(slice, fmt.Sprintf("No Outgoing Edge from %s", vtx.ID))
		} else {
			for _, edge := range d.OutEdges[vtx] {
				slice = append(slice, fmt.Sprintf("Outgoing Edges: [%s] -- %f --> [%s]\n", edge.Vtx.ID, edge.Weight, vtx.ID))
			}
		}
		if _, ok := d.InEdges[vtx]; !ok {
			slice = append(slice, fmt.Sprintf("No Incoming Edge from %s", vtx.ID))
		} else {
			for _, edge := range d.InEdges[vtx] {
				slice = append(slice, fmt.Sprintf("Incoming Edges: [%s] -- %f --> [%s]\n", edge.Vtx.ID, edge.Weight, vtx.ID))
			}
		}
		slice = append(slice, "\n")
		d.Mutex.Unlock()
	}
	return strings.Join(slice, "\n")
}

// FindVertexByID finds a Vertex by ID.
func (d Data) FindVertexByID(id string) *Vertex {
}

// DeleteVertex deletes a Vertex from the graph Data.
func (d *Data) DeleteVertex(vtx *Vertex) {

}

// DeleteEdge deletes an Edge from src to dst from the graph Data.
func (d *Data) DeleteEdge(src, dst *Vertex) {

}

// Clone clones the graph Data.
// It does `Deep Copy`.
// That is, changing the cloned Data would not affect the original Data.
func (d *Data) Clone() *Data {

}