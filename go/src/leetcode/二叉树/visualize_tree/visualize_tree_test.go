package visualize_tree

import (
	"fmt"
	"github.com/rhzx3519/leetcode/go/src/leetcode/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTree(t *testing.T) {
	arr := []types.T{3,9,20,nil,nil,15,7,nil,nil,6}
	root := buildTree(arr)
	fmt.Println(root)
}

func TestDepth(t *testing.T)  {
	arr := []types.T{3,9,20,nil,nil,15,7,nil,nil,6}
	root := buildTree(arr)
	assert.Equal(t, 4, depth(root))
}

func TestDrawTree_simple(t *testing.T) {
	arr := []types.T{1,2,3}
	DrawTree(arr)
}

func TestDrawTree(t *testing.T) {
	arr := []types.T{3,9,20,nil,nil,15,7,nil,nil,6}
	DrawTree(arr)
}
