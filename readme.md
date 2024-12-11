# Graph Entropy Calculation
This Go project is designed to calculate the entropy of a graph, which is a measure of the uncertainty or randomness in the graph structure. The project provides functions to calculate both the degree entropy and the structural entropy of a graph.

# theory

度分布熵
1. 包括0度的度数
2. 不包括度数为0的度值
3. 重复边不去重,包括双向记2条边

标准结构熵
1. 有向网络中, 忽略数量为0的度数
2. 计算标准结构熵, 熵的最小值状态为星型网络, 在有向网络中, 这个网络大小就是0, 还是符合标准结构熵的理论
3. 计算标准结构熵, 有向网络还是使用无向网络理论中的熵最大最小值进行归一化

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