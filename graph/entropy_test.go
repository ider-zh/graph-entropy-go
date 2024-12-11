package graph

import (
	"math"
	"testing"
)

var edges = []Edge[int]{{0, 1}, {1, 2}, {2, 1}, {3, 4}, {4, 5}, {6, 4}, {7, 4}}

// 0 {'in': set(), 'out': {1}}
// 1 {'in': {0, 2}, 'out': {2}}
// 2 {'in': {1}, 'out': {1}}
// 3 {'in': set(), 'out': {4}}
// 4 {'in': {3, 6, 7}, 'out': {5}}
// 5 {'in': {4}, 'out': set()}
// 6 {'in': set(), 'out': {4}}
// 7 {'in': set(), 'out': {4}}

// 去重边, 不计算 0 度节点的度分布熵
// {'in': 1.5, 'out': 0.0, 'all': 1.061278124459133}

const MIN = 0.000001

// MIN 为用户自定义的比较精度
func IsEqual(a, b float64) bool {
	return math.Abs(a-b) <= MIN
}

func TestGetDegreeEntropy(t *testing.T) {

	graph := NewGraphFromList(edges)
	ret := graph.DegreeEntropy()
	if !IsEqual(ret.EntropyIN, 1.750000000000000) {
		t.Error("inE error")
	}
	if !IsEqual(ret.EntropyOut, 0.543564443199596) {
		t.Error("inE error")
	}
	if !IsEqual(ret.Entropy, 1.548794940695398) {
		t.Error("inE error", ret.Entropy)
	}
}

func TestGetDegreeEntropyChan(t *testing.T) {

	edgeChan := make(chan *Edge[int], 20)
	for _, v := range edges {
		edgeChan <- &v
	}
	close(edgeChan)
	graph := NewGraphFromChan(edgeChan)
	ret := graph.DegreeEntropy()
	if !IsEqual(ret.EntropyIN, 1.750000000000000) {
		t.Error("inE error")
	}
	if !IsEqual(ret.EntropyOut, 0.543564443199596) {
		t.Error("inE error")
	}
	if !IsEqual(ret.Entropy, 1.548794940695398) {
		t.Error("inE error")
	}
	if ret.NodeCount != 8 {
		t.Error("inE error", ret.NodeCount)
	}
}

func TestGetStructEntropy(t *testing.T) {

	graph := NewGraphFromList(edges)
	obj := graph.StructEntropy()
	if !IsEqual(obj.EntropyIN, 1.8423709931771084) {
		t.Error("EntropyIN error")
	}
	if !IsEqual(obj.EntropyOut, 2.807354922057604) {
		t.Error("EntropyOut error")
	}
	if !IsEqual(obj.Entropy, 2.753434386188785) {
		t.Error("Entropy error")
	}
	if !IsEqual(obj.NormalizeEntropyIn, 0.24041077205417438) {
		t.Error("NormalizeEntropyIn error")
	}
	if !IsEqual(obj.NormalizeEntropyOut, 1) {
		t.Error("NormalizeEntropyOut error")
	}
	if !IsEqual(obj.NormalizeEntropy, 0.586523068142618) {
		t.Error("NormalizeEntropy error")
	}
}

func TestGetStructEntropyChan(t *testing.T) {

	edgeChan := make(chan *Edge[int], 20)
	for _, v := range edges {
		edgeChan <- &v
	}
	close(edgeChan)
	graph := NewGraphFromChan(edgeChan)

	obj := graph.StructEntropy()
	if !IsEqual(obj.EntropyIN, 1.8423709931771084) {
		t.Error("EntropyIN error", obj.EntropyIN)
	}
	if !IsEqual(obj.EntropyOut, 2.807354922057604) {
		t.Error("EntropyOut error", obj.EntropyOut)
	}
	if !IsEqual(obj.Entropy, 2.753434386188785) {
		t.Error("Entropy error", obj.Entropy)
	}
	if !IsEqual(obj.NormalizeEntropyIn, 0.24041077205417438) {
		t.Error("NormalizeEntropyIn error", obj.NormalizeEntropyIn)
	}
	if !IsEqual(obj.NormalizeEntropyOut, 1) {
		t.Error("NormalizeEntropyOut error", obj.NormalizeEntropyOut)
	}
	if !IsEqual(obj.NormalizeEntropy, 0.586523068142618) {
		t.Error("NormalizeEntropy error", obj.NormalizeEntropy)
	}

	if obj.NodeCount != 8 {
		t.Error("NodeCount error", obj.NodeCount)
	}
	if obj.NodeInCount != 4 {
		t.Error("NodeInCount error", obj.NodeInCount)
	}
	if obj.NodeOutCount != 7 {
		t.Error("NodeOutCount error", obj.NodeOutCount)
	}
}
