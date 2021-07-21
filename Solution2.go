
package main
import (
	"fmt"
	"strings"
)
type Node struct {
	data string
	left *Node
	right *Node
}

func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.data+ " ")
	preOrder(node.left)
	preOrder(node.right)
}
func postOrder(node *Node) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Print(node.data+ " ")
}

func createTree(arr []string) *Node {
	if len(arr)==0 {
		return nil
	}
	if !checkOperator(arr[0]) {
		node := Node{arr[0], nil, nil}
		head := createTree(arr[1:])
		if head == nil {
			return &node
		}
		head.left = &node
		return head
	} else {
		node := Node{arr[0], nil, nil}
		right := createTree(arr[1:])
		node.right = right
		return &node
	}
}
func checkOperator(s string) bool {
	if s=="+" || s=="-" {
		return true
	}
	return false
}


func main() {
	s := "a + b - c"
	arr :=  strings.Split(s, " ")
	root := createTree(arr)
	preOrder(root)
	fmt.Println()
	postOrder(root)
	fmt.Println()
}