```go
const (
    N int = 1e5 + 10
    M int = 10
    P int = 131313
)

var s string
n := len(s)
// 构造与字符串s等长的哈希数组h []int 和 次方数组p []int
var (
    h, p = make([]int, n+1), make([]int, n+1)
)
p[0] = 1
for i := 1; i <= n; i++ {	// 将s转换为hash数组
    h[i] = h[i - 1]*P + int(s[i-1])
    p[i] = p[i - 1]*P
}

// 求字串s[i-1:j]的hash值
hash := h[j] - h[i-1]*p[j-i+1]

```