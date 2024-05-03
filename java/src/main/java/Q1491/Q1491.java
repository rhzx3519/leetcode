package Q1491;

public class Q1491 {
}

/**
 * https://leetcode.cn/problems/average-salary-excluding-the-minimum-and-maximum-salary/?envType=daily-question&envId=2024-05-03
 */
class Solution {
    public double average(int[] salary) {
        int max = Integer.MIN_VALUE;
        int min = Integer.MAX_VALUE;
        int sum = 0;
        for (int x : salary) {
            if (x > max) {
                max = x;
            }
            if (x < min) {
                min = x;
            }
            sum += x;
        }
        sum -= max + min;
        return (double) sum / (double) (salary.length - 2);
    }
}