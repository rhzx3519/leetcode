/**
@author ZhengHao Lou
Date    2022/04/20
*/
package main

import (
    "fmt"
    "strings"
)

/**
https://leetcode-cn.com/problems/longest-absolute-file-path/
思路：栈
*/
func lengthLongestPath(s string) int {
    var mx int
    var st []string
    n := len(s)
    var bytes []byte
    var i int
    for i = 0; i <= n; {
        if i == n || s[i] == '\n' {
            name := string(bytes)
            st = append(st, name)
            fmt.Println(st)
            if strings.ContainsRune(name, '.') {
                var l int
                for i := range st {
                    l += len(st[i])
                }
                l += len(st) - 1
                if l > mx {
                    mx = l
                }
            }

            bytes = []byte{}

            var j = i + 1
            for j < n && s[j] == '\t' {
                j++
            }
            lv := j - i - 1
            i = j
            st = st[:lv]
        } else {
            bytes = append(bytes, s[i])
            i++
        }
    }

    fmt.Println(len(s), i, string(bytes), mx)
    return mx
}

func main() {
    //lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext")
    //lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext")
    //lengthLongestPath("a")
    //lengthLongestPath("file1.txt\nfile2.txt\nlongfile.txt")
    //lengthLongestPath("a\n\tb1\n\t\tf1.txt\n\taaaaa\n\t\tf2.txt") // 14
}
