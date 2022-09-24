package main

import "fmt"

var m = make(map[string]int)

var graph = make(map[string]map[string]bool)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdges(from, to string) bool {
	return graph[from][to]
}

func main() {
	Add([]string{"a", "b", "c", "f"})
	Add([]string{"aa", "bb", "cc", "ff"})
	Add([]string{"aa", "bb", "cc", "ff"})
	fmt.Println(Count([]string{"a", "b", "c", "f"}))     //	1
	fmt.Println(Count([]string{"aa", "bb", "cc", "ff"})) //	2
	fmt.Println(m)                                       //	map[["a" "b" "c" "f"]:1 ["aa" "bb" "cc" "ff"]:2]

	addEdge("a", "b")
	addEdge("b", "a")
	addEdge("b", "c")
	addEdge("c", "c")

	fmt.Println(hasEdges("a", "b")) //	true
	fmt.Println(hasEdges("b", "a")) //	true
	fmt.Println(hasEdges("b", "c")) //	true
	fmt.Println(hasEdges("c", "b")) //	false
	fmt.Println(hasEdges("c", "c")) //	true

	fmt.Println(graph) //	map[a:map[b:true] b:map[a:true c:true] c:map[c:true]]
}
