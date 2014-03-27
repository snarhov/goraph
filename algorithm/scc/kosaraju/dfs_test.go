package kosaraju

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

// Order is a bit different than ParseToGraph.
func Test_JSON_DFS(test *testing.T) {
	g4 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.004")
	g4s := DFS(g4)
	g4c := "S → B → A → T → F → E → C → D"
	if g4s != g4c {
		test.Errorf("Should be same but\n%v\n%v", g4s, g4c)
	}

	allvisited4 := true
	g4vts := g4.GetVertices()
	for _, vtx := range *g4vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gsd.Vertex).Color) {
			allvisited4 = false
		}
	}
	if !allvisited4 {
		test.Errorf("All vertices should be marked black")
	}

	g5 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.005")
	g5s := DFS(g5)
	g5c := "A → B → C → D → E → F"
	if g5s != g5c {
		test.Errorf("Should be same but\n%v\n%v", g5s, g5c)
	}

	allvisited5 := true
	g5vts := g5.GetVertices()
	for _, vtx := range *g5vts {
		if "black" != fmt.Sprintf("%v", vtx.(*gsd.Vertex).Color) {
			allvisited5 = false
		}
	}
	if !allvisited5 {
		test.Errorf("All vertices should be marked black")
	}
}