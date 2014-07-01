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
  
func compareChar(first interface{},second interface{}) int{
  if first.(int32) < second.(int32){
    return -1
  } else if first.(int32) > second.(int32) {
    return 1
  } else {
    return 0
  }
}
func verifyTree(bst *BST,tree *Node) bool {
  trees := bst.traverse(tree)
  var min interface{} = nil
  for _,value := range(trees){
    if min == nil {
      min = value.data
    } else{
      if bst.compare(min,value.data) != -1 {
        return false
      }
      min = value.data
    }
  }
  return true
}  
func buildTree() (*Node,*BST) {
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
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("Test1 - VerifyTree failed"))
  }
  tree = bst.search(tree,66)
  b := tree.data == 66 && tree.parent.data == 70 && tree.parent.parent.data == 50 && tree.parent.parent.parent.data == 25
  if !b {
   t.Error(fmt.Sprintf("Test1 - missing or bad order"))
  }

}
func Test2(t *testing.T){
  tree,bst := buildTree()
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("Test2 - VerifyTree failed"))
  }
  tree = bst.search(tree,100)

  if tree != nil {
   t.Error(fmt.Sprintf("Test2 - expected nil on first trees element. found %d",tree.data))
  }

}
func Test3(t *testing.T){
  tree,bst := buildTree()
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("Test3 - VerifyTree failed"))
  }
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
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("Test4 - VerifyTree failed"))
  }
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
func TestDelete1(t *testing.T){
  bst := New(compareChar)
  tree := bst.insert(nil,'a')
  bst.insert(tree,'b')
  bst.insert(tree,'c')
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete1 - VerifyTree 1 failed"))
  }
  bst.delete(tree,'b')
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete1 - VerifyTree 1 failed"))
  }
}
func TestDelete2(t *testing.T){
  bst := New(compareChar)
  tree := bst.insert(nil,'c')
  bst.insert(tree,'b')
  bst.insert(tree,'a')
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete2 - VerifyTree 1 failed"))
  }
  bst.delete(tree,'b')
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete2 - VerifyTree 2 failed"))
  }
}
func TestDelete3(t *testing.T){
  bst := New(compareInt)
  tree := bst.insert(nil,5)
  bst.insert(tree,8)
  bst.insert(tree,7)
  bst.insert(tree,9)
  bst.insert(tree,10)
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete3 - VerifyTree 1 failed"))
  }
  bst.delete(tree,8)
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete3 - VerifyTree 2 failed"))
  }
}

func TestDelete4(t *testing.T){
  tree,bst := buildTree()
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete4 - VerifyTree 1 failed"))
  }
  bst.delete(tree,70)
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete4 - VerifyTree 2 failed"))
  }
}
func TestDelete5(t *testing.T){
  tree,bst := buildTree()
  bst.print(tree,0)
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete5 - VerifyTree 1 failed"))
  }
  bst.delete(tree,15)
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete5 - VerifyTree 2 failed"))
  }
  fmt.Println("========================")
  bst.print(tree,0)
}

func TestDelete6(t *testing.T){
  tree,bst := buildTree()
  bst.print(tree,0)
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete6 - VerifyTree 1 failed"))
  }
  fmt.Println("========================")
  bst.delete(tree,15)
  bst.print(tree,0)
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete5 - VerifyTree 2 failed"))
  }
  fmt.Println("========================")
  bst.delete(tree,18)
  bst.print(tree,0)
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete5 - VerifyTree 3 failed"))
  }
  fmt.Println("========================")
  bst.delete(tree,10)
  bst.print(tree,0)
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete5 - VerifyTree 4 failed"))
  }
  fmt.Println("========================")
  bst.delete(tree,25)
  bst.print(tree,0)
  if ! verifyTree(bst,tree) {
     t.Error(fmt.Sprintf("TestDelete5 - VerifyTree 5 failed"))
  }
}



