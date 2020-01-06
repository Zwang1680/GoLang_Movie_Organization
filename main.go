package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Movie struct {
	name   string
	rating float64
}

/*type Node struct {
	value  Movie
	left   *Node
	right  *Node
	parent *Node
}
/*type Mbst struct {
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
	}
	return empty
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
}*/

func main() {
	//var movies Mbst
	csvfile, err := os.Open("input_20_ordered.csv")
	if err != nil {
		log.Fatalln("Could not open file", err)
	}
	var movies0 []Movie
	var temp Movie
	read := csv.NewReader(csvfile)
	for {
		// Read each record from csv
		record, error := read.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		/*movies.Insert(Movie{
			name: line[0]
			rating: strconv.AppendFloat(line[1])
		})*/
		temp.name = record[0]
		rating, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		temp.rating = rating
		movies0 = append(movies0, temp)
	}
	moviesJSON, __ := json.Marshal(movies0)
	if __ != nil {
		log.Fatal(__)
	}
	fmt.Println(string(moviesJSON))
	fmt.Println(movies0[1])
}
