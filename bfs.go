package main

import (
	"fmt"
	"strconv"
)

type QueueItem struct {
	next *QueueItem
	item int
}

type QueueList struct {
	head *QueueItem
	tail *QueueItem
}

func (q *QueueList) IsEmpty() bool {
	return q.head == nil
}

func (q *QueueList) Enqueue(item int) {
	if q.IsEmpty() {
		q.head = new(QueueItem)
		q.head.item = item
		q.tail = q.head
	} else {
		q.tail.next = new(QueueItem)
		q.tail.next.item = item
		q.tail = q.tail.next
	}
}

func (q *QueueList) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Cannot dequeue. QueueList is empty.")
	}

	item := q.head.item
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return item, nil
}

func (q *QueueList) PrintQueue() {
	fmt.Println("---- Inqueue ----")
	for i := q.head; i != nil; i = i.next {
		fmt.Println(i.item)
	}
}

type Router struct {
	name string
}

type Vertice struct {
	router        Router
	adjacencyList QueueList
}

type Graph struct {
	size     int
	vertices []*Vertice
}

func (g *Graph) Init(size int) {
	g.size = size
	g.vertices = make([]*Vertice, size)

	for i := 0; i < size; i++ {
		v := new(Vertice)
		v.router.name = "Router: " + strconv.Itoa(i)
		g.vertices[i] = v
	}
}

func (g *Graph) AddEdge(v int, w int) {
	g.vertices[v].adjacencyList.Enqueue(w)
}

func (g *Graph) Bfs(entry int, target string) (*Router, error) {
	visited := make([]bool, g.size)
	for i, _ := range visited {
		visited[i] = false
	}

	queue := new(QueueList)
	queue.Enqueue(entry)

	for queue.IsEmpty() == false {
		i, err := queue.Dequeue()
		if err != nil {
			break
		}

		if visited[i] == true {
			continue
		}

		fmt.Println("Visiting:", g.vertices[i].router.name)
		visited[i] = true

		if g.vertices[i].router.name == target {
			return &g.vertices[i].router, nil
		}

		for j := g.vertices[i].adjacencyList.head; j != nil; j = j.next {
			queue.Enqueue(j.item)
		}
	}

	return new(Router), fmt.Errorf("Cannot find: ", target)
}

func main() {
	graph := new(Graph)
	graph.Init(10)
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 9)
	graph.AddEdge(4, 7)
	graph.AddEdge(4, 8)
	graph.AddEdge(5, 4)
	graph.AddEdge(6, 3)
	graph.AddEdge(6, 9)
	graph.AddEdge(8, 6)
	graph.AddEdge(9, 9)

	result, err := graph.Bfs(0, "Router: 8")
	if err == nil {
		fmt.Println(result.name)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
