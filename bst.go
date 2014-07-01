package bst
import "fmt"
import "strings"

type Node struct{
  data interface{}
  parent *Node
  left *Node
  right *Node
}

type comparer func(interface{},interface{}) int

type BST struct {
  compare comparer
  root *Node
}

func New(compare comparer) *BST{
  return &BST{compare,nil}
}

func (bst *BST) search(tree *Node,data interface{}) *Node {
  if tree == nil {
    return nil
  }
  if tree.data == data {
    return tree
  }
  if bst.compare(data,tree.data) > 0 {
    return bst.search(tree.right,data)
  } else {
    return bst.search(tree.left,data)
  }
  
}

func (bst *BST) print(tree *Node,level int){
  if tree == nil {
    return
  }
  bst.print(tree.right,level+1)
  fmt.Printf("%s",strings.Repeat(".",level * 6))
  fmt.Println(tree.data)
  bst.print(tree.left,level+1)
}

func (bst *BST) traverse(tree *Node) []*Node{
  if tree == nil {
    return []*Node{}
  }
  
  trees1 := bst.traverse(tree.left)
  trees2 := append(trees1,tree)
  trees3 := bst.traverse(tree.right)
  return append(trees2,trees3...)
}
func (bst *BST) insert(tree *Node,data interface{}) *Node {
  if tree == nil {
    var node *Node = &Node{data,nil,nil,nil}
    if bst.root == nil {
      bst.root = node
    }
    return node
  }
  if bst.compare(data,tree.data) < 0 {
    tree.left = bst.insert(tree.left,data)
    tree.left.parent = tree
     
  } else if bst.compare(data,tree.data) > 0 {
    tree.right = bst.insert(tree.right,data)
    tree.right.parent = tree
  }
  return tree 
}
func (bst *BST) transplant(to *Node,from *Node){
  if to.parent == nil {
    bst.root = from
  } else if to == to.parent.left {
    to.parent.left = from
  } else{
    to.parent.right = from
  }
  if from != nil {
    from.parent = to.parent
  }
}

func (bst *BST) minimum(tree *Node) *Node{
  for tree.left != nil {
    tree = tree.left
  }
  return tree
}

func (bst *BST) maximum(tree *Node) *Node{
  for tree.right != nil {
    tree = tree.right
  }
  return tree
}

func (bst *BST) delete(tree *Node,data interface{}) *Node{
 
  target := bst.search(tree,data)
  if target.left == nil {
    bst.transplant(target,target.right)
  } else if target.right == nil {
    bst.transplant(target,target.left)
  } else {
    y := bst.minimum(target.right)
    if y.parent != target {
      bst.transplant(y,y.right)
      y.right = target.right
      y.right.parent = y
    }
    bst.transplant(target,y)
    y.left = target.left
    y.left.parent = y 
  }
  return tree
}
