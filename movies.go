package main

type Movie struct {
	name   string
	rating float64
}

type Node struct {
	value  Movie
	left   *Node
	right  *Node
	parent *Node
	depth  int
}
type Mbts struct {
	root     *Node
	numNodes int
}

func (root *Node) insert(newnode *Node) {
	if newnode.value.name > root.value.name {
		if root.right == nil {
			root.right = newnode
		} else {
			root.right.insert(newnode)
		}
	} else if newnode.value.name < root.value.name {
		if root.left == nil {
			root.left = newnode
		} else {
			root.left.insert(newnode)
		}
	}
}

func (tree *Mbst) Insert(value Movie) {
	if tree.root == nil {
		tree.root = &Node{nil, nil, value}
	}
	tree.size++
	tree.root.insert(&Node{nil, nil, value})
}
