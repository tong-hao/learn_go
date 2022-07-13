package main

import "fmt"

func sort(edges *map[int][]int) []int {
	degrees := indegree(edges)
	fmt.Println("inDegree", degrees)
	startNodes := make([]int, 0)
	for node, degree := range degrees {
		if degree == 0 {
			startNodes = append(startNodes, node)
		}
	}
	orders := make([]int, 0)
	queue := startNodes
	for len(queue) > 0 {
		fmt.Println(queue)
		cur := queue[0]
		queue = queue[1:]
		orders = append(orders, cur)

		for _, dst := range (*edges)[cur] {
			degrees[dst]--
			if degrees[dst] == 0 {
				queue = append(queue, dst)
			}
		}
	}

	return orders
}

func indegree(edges *map[int][]int) map[int]int {
	degree := make(map[int]int)
	for src, dsts := range *edges {
		degree[src] = 0
		for _, dst := range dsts {
			degree[dst] = 0
		}
	}
	for _, dsts := range *edges {
		for _, dst := range dsts {
			degree[dst]++
		}
	}
	return degree
}

func main() {
	edges := make(map[int][]int)
	edges[0] = []int{2, 3}
	edges[2] = []int{4}
	edges[3] = []int{4}
	edges[4] = []int{5}
	fmt.Println("orders", sort(&edges))

	graph2 := make(map[int][]int)
	graph2[5] = []int{11}
	graph2[11] = []int{2, 9, 10}
	graph2[7] = []int{11}
	graph2[8] = []int{9}
	graph2[3] = []int{8,10}

	fmt.Println("orders", sort(&graph2))

}
