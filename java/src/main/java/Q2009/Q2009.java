package Q2009;

import java.util.Arrays;

public class Q2009 {

}

/**
 *https://leetcode.cn/problems/minimum-number-of-operations-to-make-array-continuous/description/?envType=daily-question&envId=2024-04-08
 */
class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length;
        int[] a = Arrays.stream(nums).sorted().distinct().toArray();
        int included = 0;
        int left = 0;
        for (int i = 0; i < a.length; i++) {
            int x = a[i];
            while (a[left] < x - n + 1) {
                left++;
            }
            included = Math.max(included, i - left + 1);
        }
        return n - included;
    }
}