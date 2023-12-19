package util

import (
	"errors"
)

// BFSNotFound is returned by BreadthFirstSearch.Run when the queue empties
// without ever triggering the Graph.IsFinal function.
var BFSNotFound = errors.New("bfs: unable to find final state")

// BreadthFirstSearch is a struct that holds the state of a breadth first
// search. It holds a Queue, the visited nodes, the previous node for each
// visited node, and the distance (number of steps) to get from the start
// to each node.
type BreadthFirstSearch[T comparable] struct {
	Queue    []T
	Visited  map[T]bool
	Previous map[T]T
	Distance map[T]uint64
}

// A graph should implement the Graph interface in order to be exporable
// in a breadth first search.
type Graph[T comparable] interface {
	// GetInitial returns the initial node
	GetInitial() T

	// GetNeighbors returns the neighbors of the given node
	GetNeighbors(T) []T

	// IsFinal returns true if the given node is the final node.
	IsFinal(T) bool
}

// NewBFS creates a new BreadthFirstSearch struct initialized to being a search.
func NewBFS[T comparable]() *BreadthFirstSearch[T] {
	bfs := BreadthFirstSearch[T]{[]T{}, make(map[T]bool), make(map[T]T), make(map[T]uint64)}
	return &bfs
}

// Run a breadth first search on Graph g returning the final state, or an error
// if the final node is not found.
func (b *BreadthFirstSearch[T]) Run(g Graph[T]) (T, error) {
	initState := g.GetInitial()
	b.Queue = append(b.Queue, initState)
	b.Distance[initState] = 0

	// Do bfs
	for len(b.Queue) > 0 {
		s := b.Queue[0]
		b.Queue = b.Queue[1:]
		d := b.Distance[s]

		if b.Visited[s] {
			continue
		}
		b.Visited[s] = true

		for _, ns := range g.GetNeighbors(s) {
			if b.Visited[ns] {
				continue
			}

			b.Distance[ns] = d + 1
			b.Previous[ns] = s
			if g.IsFinal(ns) {
				return ns, nil
			}

			b.Queue = append(b.Queue, ns)
		}
	}

	return initState, BFSNotFound
}

// GetPath return the shortest path from the initial node to the final given node.
func (b *BreadthFirstSearch[T]) GetPath(s T) []T {
	ret := []T{s}

	for {
		ns, ok := b.Previous[s]
		if !ok {
			break
		}
		ret = append(ret, ns)
		s = ns
	}

	// reverse
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return ret
}
