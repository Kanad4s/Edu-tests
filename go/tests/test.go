package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Node represents a computation step: either Mul or Sum.
type Node struct {
	ID    string   `json:"id"`
	Op    string   `json:"op"`   // Mul or Sum
	Args  []int    `json:"args"` // [i, k, j] or [i, j, k]
	Edges []string `json:"edges"`
}

func main() {
	var nodes []Node

	AHeight := 2
	AWidth := 3
	BWidth := 6

	// Step 1: Create Mul nodes
	for i := 0; i < AHeight; i++ {
		for j := 0; j < BWidth; j++ {
			for k := 0; k < AWidth; k++ {
				id := fmt.Sprintf("Mul_%d_%d_%d", i, k, j)
				nodes = append(nodes, Node{
					ID:   id,
					Op:   "Mul",
					Args: []int{i, k, j},
				})
			}
		}
	}

	// Step 2: Create Sum nodes with dependencies
	for i := 0; i < AHeight; i++ {
		for j := 0; j < BWidth; j++ {
			for k := 0; k < AWidth; k++ {
				id := fmt.Sprintf("Sum_%d_%d_%d", i, j, k)
				edges := []string{}
				if k == 0 {
					edges = append(edges, fmt.Sprintf("Mul_%d_%d_%d", i, k, j))
				} else {
					edges = append(edges,
						fmt.Sprintf("Sum_%d_%d_%d", i, j, k-1),
						fmt.Sprintf("Mul_%d_%d_%d", i, k, j),
					)
				}
				nodes = append(nodes, Node{
					ID:    id,
					Op:    "Sum",
					Args:  []int{i, j, k},
					Edges: edges,
				})
			}
		}
	}

	// Step 3: Export to JSON
	file, err := os.Create("execution_plan.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(nodes)
	if err != nil {
		panic(err)
	}

	fmt.Println("Execution plan exported to execution_plan.json")
}
