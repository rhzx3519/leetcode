package Q1553;

import java.util.HashMap;
import java.util.Map;

public class Q1553 {
}

/**
 * https://leetcode.cn/problems/minimum-number-of-days-to-eat-n-oranges/?envType=daily-question&envId=2024-05-12
 */
class Solution {

    private Map<Integer, Integer> memo = new HashMap<>();
    public int minDays(int n) {
        if (n <= 1) {
            return 1;
        }
        if (memo.containsKey(n)) {
            return memo.getOrDefault(n, 0);
        }
        int res = Math.min(minDays(n/2) + n%2 + 1, minDays(n/3) + n%3 + 1);
        memo.put(n, res);
        return res;
    }
}