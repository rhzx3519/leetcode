package Q377;

import java.util.Arrays;

public class Q377 {
}

/**
 * https://leetcode.cn/problems/combination-sum-iv/description/?envType=daily-question&envId=2024-04-22
 * 思路：dp
 */
class Solution {
    public int combinationSum4(int[] nums, int target) {
        int[] f = new int[target + 1];
        f[0] = 1;
        for (int i = 1; i <= target; i++) {
            for (int num : nums) {
                if (i >= num) {
                    f[i] += f[i - num];
                }
            }
        }
        return f[target];
    }
}