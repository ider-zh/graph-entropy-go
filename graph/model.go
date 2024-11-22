package graph

import "github.com/emirpasic/gods/v2/sets/hashset"

type PointType interface {
	int | int16 | int32 | int64 | string
}

type Node[T PointType] struct {
	ID       T
	LinksIn  *hashset.Set[T]
	LinksOut *hashset.Set[T]
}

type Edge[T PointType] struct {
	From T
	To   T
}

type Graph[T PointType] struct {
	Nodes map[T]*Node[T]
}

type StructEntropyStats struct {
	EntropyIN           float64
	EntropyOut          float64
	Entropy             float64
	NormalizeEntropyIn  float64
	NormalizeEntropyOut float64
	NormalizeEntropy    float64
}

type DegreeEntropyStats struct {
	EntropyIN  float64
	EntropyOut float64
	Entropy    float64
}
