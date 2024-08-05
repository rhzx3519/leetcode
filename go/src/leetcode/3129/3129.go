package main

/**
https://leetcode.cn/problems/find-all-possible-stable-binary-arrays-i/?envType=daily-question&envId=2024-08-06
*/
func numberOfStableArrays(zero, one, limit int) (ans int) {
    const mod = 1_000_000_007
    f := make([][][2]int, zero+1)
    for i := range f {
        f[i] = make([][2]int, one+1)
    }
    for i := 1; i <= min(limit, zero); i++ {
        f[i][0][0] = 1
    }
    for j := 1; j <= min(limit, one); j++ {
        f[0][j][1] = 1
    }
    for i := 1; i <= zero; i++ {
        for j := 1; j <= one; j++ {
            f[i][j][0] = (f[i-1][j][0] + f[i-1][j][1]) % mod
            if i > limit {
                // + mod 保证答案非负
                f[i][j][0] = (f[i][j][0] - f[i-limit-1][j][1] + mod) % mod
            }
            f[i][j][1] = (f[i][j-1][0] + f[i][j-1][1]) % mod
            if j > limit {
                f[i][j][1] = (f[i][j][1] - f[i][j-limit-1][0] + mod) % mod
            }
        }
    }
    return (f[zero][one][0] + f[zero][one][1]) % mod
}
