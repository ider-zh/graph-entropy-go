package graph

import (
	"math"
	"testing"
)

const MIN = 0.000001

// MIN 为用户自定义的比较精度
func IsEqual(a, b float64) bool {
	return math.Abs(a-b) <= MIN
}

func TestGetDegreeEntropy(t *testing.T) {
	var edges = []Edge[int]{
		{0, 1}, {1, 2}, {2, 1}, {3, 4},
	}

	graph := NewGraphFromList(edges)
	ret := graph.DegreeEntropy()
	if !IsEqual(ret.EntropyIN, 0.9182958340544893) {
		t.Error("inE error")
	}
	if !IsEqual(ret.EntropyOut, 0) {
		t.Error("inE error")
	}
	if !IsEqual(ret.Entropy, 0.7219280948873623) {
		t.Error("inE error")
	}
}

func TestGetDegreeEntropyChan(t *testing.T) {
	var edges = []Edge[int]{
		{0, 1}, {1, 2}, {2, 1}, {3, 4},
	}

	edgeChan := make(chan Edge[int], 20)
	for _, v := range edges {
		edgeChan <- v
	}
	close(edgeChan)
	graph := NewGraphFromChan(edgeChan)
	ret := graph.DegreeEntropy()
	if !IsEqual(ret.EntropyIN, 0.9182958340544893) {
		t.Error("inE error")
	}
	if !IsEqual(ret.EntropyOut, 0) {
		t.Error("inE error")
	}
	if !IsEqual(ret.Entropy, 0.7219280948873623) {
		t.Error("inE error")
	}
}

func TestGetStructEntropy(t *testing.T) {
	var edges = []Edge[int]{
		{0, 1}, {1, 2}, {2, 1}, {3, 4},
	}

	graph := NewGraphFromList(edges)
	obj := graph.StructEntropy()
	if !IsEqual(obj.EntropyIN, 1.5) {
		t.Error("inE error")
	}
	if !IsEqual(obj.EntropyOut, 2) {
		t.Error("inE error")
	}
	if !IsEqual(obj.Entropy, 2.25162916738782) {
		t.Error("inE error")
	}
	if !IsEqual(obj.NormalizeEntropyIn, 0.6460148371100) {
		t.Error("inE error")
	}
	if !IsEqual(obj.NormalizeEntropyOut, 0.86135311614678) {
		t.Error("inE error")
	}
	if !IsEqual(obj.NormalizeEntropy, 0.7816315860094917) {
		t.Error("inE error")
	}
}

func TestGetStructEntropyChan(t *testing.T) {
	var edges = []Edge[int]{
		{0, 1}, {1, 2}, {2, 1}, {3, 4},
	}

	edgeChan := make(chan Edge[int], 20)
	for _, v := range edges {
		edgeChan <- v
	}
	close(edgeChan)
	graph := NewGraphFromChan(edgeChan)

	obj := graph.StructEntropy()
	if !IsEqual(obj.EntropyIN, 1.5) {
		t.Error("inE error")
	}
	if !IsEqual(obj.EntropyOut, 2) {
		t.Error("inE error")
	}
	if !IsEqual(obj.Entropy, 2.25162916738782) {
		t.Error("inE error")
	}
	if !IsEqual(obj.NormalizeEntropyIn, 0.6460148371100) {
		t.Error("inE error")
	}
	if !IsEqual(obj.NormalizeEntropyOut, 0.86135311614678) {
		t.Error("inE error")
	}
	if !IsEqual(obj.NormalizeEntropy, 0.7816315860094917) {
		t.Error("inE error")
	}
}
