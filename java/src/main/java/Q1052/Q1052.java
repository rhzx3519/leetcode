package Q1052;

public class Q1052 {
}

/**
 * https://leetcode.cn/problems/grumpy-bookstore-owner/?envType=daily-question&envId=2024-04-23
 */
class Solution {
    public int maxSatisfied(int[] customers, int[] grumpy, int minutes) {
        int n = customers.length;
        int unsatisfied = 0;
        int mxUnsatisfied = 0;
        int total = 0;
        for (int i = 0; i < n; i++) {
            if (grumpy[i] == 1) {
               unsatisfied += customers[i];
            }
            if (grumpy[i] == 0) {
                total += customers[i];
            }
            if (i >= minutes) {
                if (grumpy[i - minutes] == 1) {
                    unsatisfied -= customers[i - minutes];
                }
            }
            mxUnsatisfied = Math.max(mxUnsatisfied, unsatisfied);
        }
        return total + mxUnsatisfied;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.maxSatisfied(new int[] { 1,0,1,2,1,1,7,5 }, new int[] {0,1,0,1,0,1,0,1 }, 3));
        System.out.println(s.maxSatisfied(new int[] { 1 }, new int[] {0 }, 1));
    }
}