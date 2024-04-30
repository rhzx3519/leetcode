package Q2798;

public class Q2798 {
}

/**
 * https://leetcode.cn/problems/number-of-employees-who-met-the-target/?envType=daily-question&envId=2024-04-30
 */
class Solution {
    public int numberOfEmployeesWhoMetTarget(int[] hours, int target) {
        int tot = 0;
        for (int h : hours) {
            if (h >= target) {
               tot++;
            }
        }
        return tot;
    }
}

