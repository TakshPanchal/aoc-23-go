package day8

import "fmt"

type Node struct {
	Name        string
	Right, Left *Node
}

// type Instructions string
type Network map[string]*Node

func parseNetwork(nodes []string) Network {
	net := Network{}

	for _, node := range nodes {
		if len(node) > 0 {
			name := node[:3]
			right, left := node[12:15], node[7:10]

			for _, node := range [3]string{name, left, right} {
				if _, ok := net[node]; !ok {
					net[node] = &Node{Name: node}
				}
			}
			net[name].Left = net[left]
			net[name].Right = net[right]
		}

	}
	return net
}

func CalculateSteps(lines []string) int {
	instruction := lines[0]
	net := parseNetwork(lines[2:])

	currNodes := getNodesEndingWith(net, 'A')
	currDir := instruction[0]
	steps := 0

	fmt.Printf("%v \n", currDir)
	for i := range currNodes {
		fmt.Printf("%v ", currNodes[i].Name)
	}
	fmt.Println()
	for !isReached(currNodes) {
		for i := range currNodes {
			if currDir == 'L' {
				currNodes[i] = currNodes[i].Left
			} else {
				currNodes[i] = currNodes[i].Right
			}
		}
		steps++
		currDir = instruction[steps%len(instruction)]
	}

	return steps
}

// isReached returns true if all current nodes are reached to nodes ending with Z
func isReached(nodes []*Node) bool {
	for _, n := range nodes {
		fmt.Printf("%v ", n.Name)
		if n.Name[len(n.Name)-1] != 'Z' {
			fmt.Printf("\n")
			return false
		}
	}
	fmt.Printf("\n")
	return true
}

func getNodesEndingWith(net Network, endChar rune) (nodes []*Node) {
	for name, node := range net {
		if name[len(name)-1] == byte(endChar) {
			nodes = append(nodes, node)
		}
	}
	return
}
