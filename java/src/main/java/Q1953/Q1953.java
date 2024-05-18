package Q1953;

public class Q1953 {
}

/**
 * https://leetcode.cn/problems/maximum-number-of-weeks-for-which-you-can-work/?envType=daily-question&envId=2024-05-16
 */
class Solution {
    public long numberOfWeeks(int[] milestones) {
        long sum = 0;
        long mx = 0;
        for (int m : milestones) {
           sum += m;
           mx = Math.max(mx, (long) m);
        }
        long rest = sum - mx;
        if (mx > rest + 1) {
            return rest*2 + 1;
        }
        return sum;
    }
}