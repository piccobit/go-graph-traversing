package main

import (
	"fmt"

	graphalgos "github.com/piccobit/go-graph-traversing/graphalgos"
)

func main() {
	var r []string

	fmt.Println("GoLang, Graph DFS and BFS implementation")
	fmt.Println("DFS: Depth First Search")
	fmt.Println("BFS: Breadth First Search")

	g := graphalgos.NewGraph()

	fmt.Println("\nAdding vertexes...")
	g.AddVertex("alpha")
	g.AddVertex("bravo")
	g.AddVertex("charlie")
	g.AddVertex("delta")
	g.AddVertex("echo")
	g.AddVertex("foxtrot")
	g.AddVertex("golf")
	g.AddVertex("hotel")

	fmt.Println("\nAdding edges...")
	g.AddEdge("alpha", "bravo")
	g.AddEdge("alpha", "charlie")
	g.AddEdge("charlie", "delta")
	g.AddEdge("bravo", "echo")
	g.AddEdge("echo", "bravo")
	g.AddEdge("bravo", "foxtrot")

	fmt.Println("\nDFS Recursive: alpha")
	r = g.DFSRecursive("alpha")
	for i, n := range r {
		fmt.Printf("%d: %s\n", i, n)
	}

	fmt.Println("\nDFS Iterative: alpha")
	r = g.DFSIterative("alpha")
	for i, n := range r {
		fmt.Printf("%d: %s\n", i, n)
	}

	fmt.Println("\nBFS: alpha")
	r = g.BFS("alpha")
	for i, n := range r {
		fmt.Printf("%d: %s\n", i, n)
	}

	fmt.Println("\nDFS Iterative: bravo")
	r = g.DFSIterative("bravo")
	for i, n := range r {
		fmt.Printf("%d: %s\n", i, n)
	}

	fmt.Println("\nCreate path:")
	g.CreatePath("bravo", "echo")
}
