package graph

import "math"

func sumNumbers[T int | int32 | int64 | int16 | float64](array []T) T {
	var result T = 0
	for _, v := range array {
		result += v
	}
	return result
}

func calStructEntropy(degrees []int) float64 {
	sum_degree := sumNumbers(degrees)

	cache_array := make([]float64, len(degrees))
	for i := range degrees {
		cache_array[i] = float64(degrees[i]) / float64(sum_degree)
	}

	for i := range cache_array {
		cache_array[i] = -cache_array[i] * math.Log2(cache_array[i])
	}
	var E = sumNumbers(cache_array)
	return E
}

func (c *Graph[T]) StructEntropy() *StructEntropyStats {
	var in_degree []int
	var out_degree []int
	var undirected_degree []int

	for _, node := range c.Nodes {
		InIDsCount := node.LinksIn.Size()
		if InIDsCount > 0 {
			in_degree = append(in_degree, InIDsCount)
		}

		OutIDsCount := node.LinksOut.Size()
		if OutIDsCount > 0 {
			out_degree = append(out_degree, OutIDsCount)
		}

		undirected_degree_count := node.LinksIn.Union(node.LinksOut).Size()
		if undirected_degree_count > 0 {
			undirected_degree = append(undirected_degree, undirected_degree_count)
		}
	}

	in_E := calStructEntropy(in_degree)
	out_E := calStructEntropy(out_degree)
	undirected_E := calStructEntropy(undirected_degree)

	undirected_E_min := math.Log2(4*float64(len(undirected_degree)-1)) / 2
	undirected_E_max := math.Log2(float64(len(undirected_degree)))

	return &StructEntropyStats{
		EntropyIN:           in_E,
		EntropyOut:          out_E,
		Entropy:             undirected_E,
		NormalizeEntropyIn:  in_E / math.Log2(float64(len(undirected_degree))),
		NormalizeEntropyOut: out_E / math.Log2(float64(len(undirected_degree))),
		NormalizeEntropy:    (undirected_E - undirected_E_min) / (undirected_E_max - undirected_E_min),
	}

}