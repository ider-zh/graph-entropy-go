# Graph Entropy Calculation
This Go project is designed to calculate the entropy of a graph, which is a measure of the uncertainty or randomness in the graph structure. The project provides functions to calculate both the degree entropy and the structural entropy of a graph.


# Usage

```
package main

import (
	"fmt"
	"github.com/ider-zh/graph-entropy-go/graph"
)

func main() {
	// Create a new graph
	graph := graph.NewGraphFromList([]graph.Edge[int]{
		{0, 1}, {1, 2}, {2, 1}, {3, 4},
	})

	// Calculate the structural entropy
	entropy := graph.StructEntropy()

	// Print the results
	fmt.Printf("Entropy IN: %.4f\n", entropy.EntropyIN)
	fmt.Printf("Entropy OUT: %.4f\n", entropy.EntropyOut)
	fmt.Printf("Entropy: %.4f\n", entropy.Entropy)
	fmt.Printf("Normalized Entropy IN: %.4f\n", entropy.NormalizeEntropyIn)
	fmt.Printf("Normalized Entropy OUT: %.4f\n", entropy.NormalizeEntropyOut)
	fmt.Printf("Normalized Entropy: %.4f\n", entropy.NormalizeEntropy)
}
```


# Functions
`NewGraphFromList`
Creates a new graph from a list of edges.

`NewGraphFromChan`
Creates a new graph from a channel of edges.

`DegreeEntropy`
Calculates the degree entropy of a graph.

`StructEntropy`
Calculates the structural entropy of a graph.

# Examples
The entropy_test.go file contains several examples of how to use the functions in this project. You can run these tests using the following command:


sh
```
go test
```