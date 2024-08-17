package main

/**
https://leetcode.cn/problems/minimum-number-of-operations-to-make-word-k-periodic/?envType=daily-question&envId=2024-08-17
*/
func minimumOperationsToMakeKPeriodic(word string, k int) int {
    n := len(word)
    cnt := map[string]int{}
    for i := k; i <= n; i += k {
        cnt[word[i-k:i]]++
    }
    mx := 0
    for _, c := range cnt {
        mx = max(mx, c)
    }
    return n/k - mx
}
