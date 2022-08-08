package visualize_tree

import (
	"fmt"
	types2 "github.com/rhzx3519/leetcode/go/src/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTree(t *testing.T) {
	arr := []types2.T{3,9,20,nil,nil,15,7,nil,nil,6}
	root := buildTree(arr)
	fmt.Println(root)
}

func TestDepth(t *testing.T)  {
	arr := []types2.T{3,9,20,nil,nil,15,7,nil,nil,6}
	root := buildTree(arr)
	assert.Equal(t, 4, depth(root))
}

func TestDrawTree_simple(t *testing.T) {
	arr := []types2.T{1,2,3}
	DrawTree(arr)
}

func TestDrawTree(t *testing.T) {
	arr := []types2.T{3,9,20,nil,nil,15,7,nil,nil,6}
	DrawTree(arr)
}

func TestDrawTreeComplicated(t *testing.T) {
	arr := []types2.T{3,9,20,nil,2,nil,4,nil,5,6,7}
	DrawTree(arr)
}

func TestDrawTreeComplicated1(t *testing.T) {
	arr := []types2.T{3,5,1,6,2,9,8,nil,nil,7,4,6}
	DrawTree(arr)
}

func TestDrawTreeComplicated2(t *testing.T) {
	arr := []types2.T{5,4,8,11,nil,13,4,7,2,nil,nil,5,1}
	DrawTree(arr)
}
