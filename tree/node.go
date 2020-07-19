package tree

import "fmt"

// Node ...
type Node struct {
	Value       int
	Left, Right *Node
}

// Print ...Print
func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

// SetValue ...SetValue
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

// CreateNode ...CreateNode
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
