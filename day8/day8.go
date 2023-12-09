package day8

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

	currNode := net["AAA"]
	currDir := instruction[0]
	steps := 0
	for currNode.Name != "ZZZ" {
		if currDir == 'L' {
			currNode = currNode.Left
		} else {
			currNode = currNode.Right
		}
		steps++
		currDir = instruction[steps%len(instruction)]
	}

	return steps
}
