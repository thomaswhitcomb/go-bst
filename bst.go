package bst

type Node struct{
  data interface{}
  parent *Node
  left *Node
  right *Node
}

type comparer func(interface{},interface{}) int

type BST struct {
  compare comparer
}

func New(compare comparer) BST{
  return BST{compare}
}

func (bst BST) search(tree *Node,data interface{}) *Node {
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
func (bst BST) traverse(tree *Node) []*Node{
  if tree == nil {
    return []*Node{}
  }
  
  trees1 := bst.traverse(tree.left)
  trees2 := append(trees1,tree)
  trees3 := bst.traverse(tree.right)
  return append(trees2,trees3...)
}
func (bst BST) insert(tree *Node,data interface{}) *Node {
  if tree == nil {
    return &Node{data,nil,nil,nil}
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

func (bst BST) delete(tree *Node,data interface{}) *Node{
 
  target := bst.search(tree,data)
  if target.left == nil && target.right == nil {
    if target.parent.left == target {
      target.parent.left = nil
    } else{
      target.parent.right = nil
    }
  }
  return tree
}
