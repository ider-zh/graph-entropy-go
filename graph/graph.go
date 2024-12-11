package graph

func NewGraphFromChan[T PointType](edgeIn chan *Edge[T]) *Graph[T] {
	graph := make(map[T]*Node[T])
	for edge := range edgeIn {
		if _, ok := graph[edge.From]; !ok {
			graph[edge.From] = &Node[T]{
				ID:       edge.From,
				LinksIn:  []T{},
				LinksOut: []T{edge.To},
			}
		} else {
			graph[edge.From].LinksOut = append(graph[edge.From].LinksOut, edge.To)
		}

		if _, ok := graph[edge.To]; !ok {
			graph[edge.To] = &Node[T]{
				ID:       edge.To,
				LinksIn:  []T{edge.From},
				LinksOut: []T{},
			}
		} else {
			graph[edge.To].LinksIn = append(graph[edge.To].LinksIn, edge.From)
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
				LinksIn:  []T{},
				LinksOut: []T{edge.To},
			}
		} else {
			graph[edge.From].LinksOut = append(graph[edge.From].LinksOut, edge.To)
		}

		if _, ok := graph[edge.To]; !ok {
			graph[edge.To] = &Node[T]{
				ID:       edge.To,
				LinksIn:  []T{edge.From},
				LinksOut: []T{},
			}
		} else {
			graph[edge.To].LinksIn = append(graph[edge.To].LinksIn, edge.From)
		}
	}
	return &Graph[T]{Nodes: graph}
}
