package main
import (
	"errors"
)
type Data struct{
	id string
	answer string
	question string
}
type Node struct {
	Value string
	Data  Data
	Left  *Node
	Right *Node
}

func NavigateNode(n *Node, b string) *Node{
	if(b == "yes"){
		return n.Left
	}else{
		return n.Right
	}
}

func (n *Node) Insert(value string, data Data) error {

	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	//If the data is already in the tree, return.
	case value == n.Value:
		return nil
		//If the data value is less than the current node’s value, and if the left child node is nil, insert a new left child node. Else call Insert on the left subtree.
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)
		//If the data value is greater than the current node’s value, do the same but for the right subtree.
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

// todo change to get
func (n *Node) Find(s string) (Data, bool) {

	if n == nil {
		d1 := Data{"","",""}
		return d1, false
	}

	switch {
	//If the current node contains the value, return the node.
	case s == n.Value:
		return n.Data, true
		//If the data value is less than the current node’s value, call Find for the left child node,
	case s < n.Value:
		return n.Left.Find(s)
		//else call Find for the right child node.
	default:
		return n.Right.Find(s)
	}
}





//TREE STRUCT
func GetRoot(t *Tree)*Node{
	return t.Root
}
type Tree struct {
	Root *Node
}
//Insert calls Node.Insert unless the root node is nil
func (t *Tree) Insert(value string, data Data) error {
	//If the tree is empty, create a new node,…
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}
	//…else call Node.Insert.
	return t.Root.Insert(value, data)
}
//Find calls Node.Find unless the root node is nil
func (t *Tree) Find(s string) (Data, bool) {
	if t.Root == nil {
		//d := Data("","")
		d1 := Data{"","",""}
		return d1, false
	}
	return t.Root.Find(s)
}
//Delete has one special case: the empty tree. (And deleting from an empty tree is an error.) In all other cases, it calls Node.Delete.
func (t *Tree) Delete(s string) error {

	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}
//	CallNode.Delete. Passing a “fake” parent node here almost avoids having to treat the root node as a special case, with one exception.
fakeParent := &Node{Right: t.Root}
err := t.Root.Delete(s, fakeParent)
if err != nil {
return err
}
//If the root node is the only node in the tree, and if it is deleted, then it only got removed from fakeParent. t.Root still points to the old node. We rectify this by setting t.Root to nil.
if fakeParent.Right == nil {
t.Root = nil
}
return nil
}
//Traverse is a simple method that traverses the tree in left-to-right order (which, by pure incidence ;-), is the same as traversing from smallest to largest value) and calls a custom function on each node.
func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}



//TREE STRUCT





























// TREE
func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}
//replaceNode replaces the parent’s child pointer to n with a pointer to the replacement node. parent must not be nil.
func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}
//Delete removes an element from the tree. It is an error to try deleting an element that does not exist. In order to remove an element properly, Delete needs to know the node’s parent node. parent must not be nil.
func (n *Node) Delete(s string, parent *Node) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}
	//Search the node to be deleted.
	switch {
	case s < n.Value:
		return n.Left.Delete(s, n)
	case s > n.Value:
		return n.Right.Delete(s, n)
	default:
		//We found the node to be deleted. If the node has no children, simply remove it from its parent.
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}
		//If the node has one child: Replace the node with its child.
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}
		//If the node has two children: Find the maximum element in the left subtree…
		replacement, replParent := n.Left.findMax(n)
		//…and replace the node’s value and data with the replacement’s value and data.
		n.Value = replacement.Value
		n.Data = replacement.Data
		//Then remove the replacement node.
		return replacement.Delete(replacement.Value, replParent)
	}
}








