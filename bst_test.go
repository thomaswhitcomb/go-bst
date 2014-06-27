package bst

import "testing"
import "fmt"

func compareInt(first interface{},second interface{}) int{
  if first.(int) < second.(int){
    return -1
  } else if first.(int) > second.(int) {
    return 1
  } else {
    return 0
  }
}
  
func buildTree() (*Node,BST) {
  bst := New(compareInt)
  tree := bst.insert(nil,25)
  bst.insert(tree,15)
  bst.insert(tree,50)
  bst.insert(tree,10)
  bst.insert(tree,22)
  bst.insert(tree,35)
  bst.insert(tree,70)
  bst.insert(tree,4)
  bst.insert(tree,12)
  bst.insert(tree,18)
  bst.insert(tree,24)
  bst.insert(tree,31)
  bst.insert(tree,44)
  bst.insert(tree,66)
  bst.insert(tree,90)
  return tree,bst
}

func Test1(t *testing.T){
  tree,bst := buildTree()
  tree = bst.search(tree,66)
  b := tree.data == 66 && tree.parent.data == 70 && tree.parent.parent.data == 50 && tree.parent.parent.parent.data == 25
  if !b {
   t.Error(fmt.Sprintf("Test1 - missing or bad order"))
  }

}
func Test2(t *testing.T){
  tree,bst := buildTree()
  tree = bst.search(tree,100)

  if tree != nil {
   t.Error(fmt.Sprintf("Test2 - expected nil on first trees element. found %d",tree.data))
  }

}
func Test3(t *testing.T){
  tree,bst := buildTree()
  trees := bst.traverse(tree)
  var s string = ""
  for _,tree := range trees{
    s = s + fmt.Sprintf("%d ",tree.data)
  }
  if s != "4 10 12 15 18 22 24 25 31 35 44 50 66 70 90 " {
    t.Error(fmt.Sprintf("Test2 - bad traverse. got %s",s))
  }
}

func Test4(t *testing.T){
  tree,bst := buildTree()
  bst.delete(tree,90)
  trees := bst.traverse(tree)
  var s string = ""
  for _,tree := range trees{
    s = s + fmt.Sprintf("%d ",tree.data)
  }
  if s != "4 10 12 15 18 22 24 25 31 35 44 50 66 70 " {
    t.Error(fmt.Sprintf("Test4 - deleted failed. got %s",s))
  }

  
}
