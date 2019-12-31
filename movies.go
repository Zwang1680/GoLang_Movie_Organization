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
}
type Mbst struct {
	root     *Node
	numNodes int
	size     int
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
		tree.root = &Node{value, nil, nil, nil}
	}
	tree.size++
	tree.root.insert(&Node{value, nil, nil, nil})
}

func (tree *Mbst) searchSpecific(s string) Movie {
	n := searchSpecificHelper(s, tree.root)
	empty := Movie{"", 0.0}
	if n != nil {
		return n.value
	} else {
		return empty
	}
}

func searchSpecificHelper(s string, r *Node) *Node {
	if r == nil {
		return nil
	}
	if r.value.name == s {
		return r
	}
	if s < r.value.name {
		return searchSpecificHelper(s, r.left)
	}
	if s > r.value.name {
		return searchSpecificHelper(s, r.right)
	}
	return nil
}
