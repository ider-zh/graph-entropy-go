package graph

import "github.com/emirpasic/gods/v2/sets/hashset"

func NewGraphFromChan[T PointType](edgeIn chan Edge[T]) *Graph[T] {
	graph := make(map[T]*Node[T])
	for edge := range edgeIn {
		if _, ok := graph[edge.From]; !ok {
			graph[edge.From] = &Node[T]{
				ID:       edge.From,
				LinksIn:  hashset.New[T](),
				LinksOut: hashset.New[T](edge.To),
			}
		} else {
			graph[edge.From].LinksOut.Add(edge.To)
		}

		if _, ok := graph[edge.To]; !ok {
			graph[edge.To] = &Node[T]{
				ID:       edge.To,
				LinksIn:  hashset.New[T](edge.From),
				LinksOut: hashset.New[T](),
			}
		} else {
			graph[edge.To].LinksIn.Add(edge.From)
		}

	}
	return &Graph[T]{Nodes: graph}
}

func NewGraphFromList[T PointType](edgeIn []Edge[T]) *Graph[T] {
	graph := make(map[T]*Node[T])
	for _, edge := range edgeIn {
		if _, ok := graph[edge.From]; !ok {
			graph[edge.From] = &Node[T]{
				ID:       edge.From,
				LinksIn:  hashset.New[T](),
				LinksOut: hashset.New[T](edge.To),
			}
		} else {
			graph[edge.From].LinksOut.Add(edge.To)
		}

		if _, ok := graph[edge.To]; !ok {
			graph[edge.To] = &Node[T]{
				ID:       edge.To,
				LinksIn:  hashset.New[T](edge.From),
				LinksOut: hashset.New[T](),
			}
		} else {
			graph[edge.To].LinksIn.Add(edge.From)
		}

	}
	return &Graph[T]{Nodes: graph}
}
