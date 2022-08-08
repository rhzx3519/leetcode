package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
https://leetcode-cn.com/problems/delete-duplicate-folders-in-system/
思路：字典树，子树序列化

 */
type folder struct {
	child map[string]*folder
	val string
	del bool
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	root := &folder{
		child: map[string]*folder{},
		val: "/",
	}
	for _, path := range paths {
		f := root
		for _, s := range path {
			if f.child == nil {
				f.child = map[string]*folder{}
			}
			if f.child[s] == nil {
				f.child[s] = &folder{
					child: map[string]*folder{},
					val: s,
				}
			}
			f = f.child[s]
		}
	}

	foldersMap := make(map[string][]*folder)
	var dfs func(f *folder) string
	dfs = func(f *folder) string {
		if f.child == nil || len(f.child) == 0 {
			return "(" + f.val + ")"
		}
		exper := []string{}
		for _, c := range f.child {
			exper = append(exper, dfs(c))
		}
		sort.Strings(exper)
		key := strings.Join(exper, " ")
		foldersMap[key] = append(foldersMap[key], f)
		return "(" + f.val + key + ")"
	}
	dfs(root)
	for _, folders := range foldersMap {
		if len(folders) > 1 {
			for _, f := range folders {
				f.del = true
			}
		}
	}

	var ans = [][]string{}
	var d2 func(f *folder, path []string)
	d2 = func(f *folder, path []string) {
		if f.del {
			return
		}

		path = append(path, f.val)
		ans = append(ans, append([]string{}, path[1:]...))
		for _, c := range f.child {
			d2(c, path)
		}
		path = path[:len(path) - 1]
	}
	d2(root, []string{})

	fmt.Println(ans[1:])
	return ans[1:]
}

func main() {
	deleteDuplicateFolder([][]string{{"a"},{"c"},{"d"},{"a","b"},{"c","b"},{"d","a"}})
	deleteDuplicateFolder([][]string{{"a"},{"c"},{"a","b"},{"c","b"},{"a","b","x"},{"a","b","x","y"},{"w"},{"w","y"}})
}