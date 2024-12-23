package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func part1(network map[string][]string) {
	now := time.Now()
	connections := make(map[[3]string]struct{})
	for node, children := range network {
		for i := 0; i < len(children)-2; i++ {
			for j := i + 1; j < len(children); j++ {
				child1 := children[i]
				child2 := children[j]
				if node[0] != 't' && child1[0] != 't' && child2[0] != 't' {
					continue
				}
				if slices.Contains(network[child1], child2) {
					connection := [3]string{node, child1, child2}
					slices.Sort(connection[:])
					connections[connection] = struct{}{}
				}
			}
		}
	}

	fmt.Println("Part 1:", len(connections))
	fmt.Println("Time:", time.Since(now))
}

func bronKerbosch(g map[string][]string, r, p, x []string, cliques *[][]string) {
	if len(p) == 0 && len(x) == 0 {
		*cliques = append(*cliques, append([]string(nil), r...))
		return
	}
	u := choosePivot(append(p, x...), g)
	for _, v := range difference(p, g[u]) {
		n := g[v]
		bronKerbosch(g, append(r, v), intersection(p, n), intersection(x, n), cliques)
		p = difference(p, []string{v})
		x = append(x, v)
	}
}

func part2(network map[string][]string) {
	now := time.Now()
	var cliques [][]string
	nodes := []string{}
	for node := range network {
		nodes = append(nodes, node)
	}
	bronKerbosch(network, []string{}, nodes, []string{}, &cliques)
	largestClique := []string{}
	for _, clique := range cliques {
		if len(clique) > len(largestClique) {
			largestClique = clique
		}
	}
	slices.Sort(largestClique)
	password := strings.Join(largestClique, ",")
	fmt.Println("Part 2:", password)
	fmt.Println("Time:", time.Since(now))
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	network := make(map[string][]string)

	for _, line := range lines {
		split := strings.Split(line, "-")
		network[split[0]] = append(network[split[0]], split[1])
		network[split[1]] = append(network[split[1]], split[0])
	}

	part1(network)
	part2(network)
}

func intersection(a, b []string) []string {
	set := make(map[string]bool)
	for _, v := range b {
		set[v] = true
	}
	var result []string
	for _, v := range a {
		if set[v] {
			result = append(result, v)
		}
	}
	return result
}

func difference(a, b []string) []string {
	set := make(map[string]bool)
	for _, v := range b {
		set[v] = true
	}
	var result []string
	for _, v := range a {
		if !set[v] {
			result = append(result, v)
		}
	}
	return result
}

func choosePivot(nodes []string, g map[string][]string) string {
	maxDegree := -1
	var pivot string
	for _, node := range nodes {
		if len(g[node]) > maxDegree {
			maxDegree = len(g[node])
			pivot = node
		}
	}
	return pivot
}
