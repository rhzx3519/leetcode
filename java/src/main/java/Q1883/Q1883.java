package Q1883;

public class Q1883 {
}


/**
 * https://leetcode.cn/problems/minimum-skips-to-arrive-at-meeting-on-time/?envType=daily-question&envId=2024-04-19
 */
class Solution {
    public int minSkips(int[] dist, int speed, int hoursBefore) {
        int sumDist = 0;
        for (int d : dist) {
            sumDist += d;
        }
        if (sumDist > speed * hoursBefore) {
            return -1;
        }
        int n = dist.length;
        int[][] f = new int[n][];   // f[i][j] 表示最多跳过i次，到达j时的最小时间, 再乘以speed
        for (int i = 0; i < n; i++) {
            f[i] = new int[n];
            for (int j = 0; j < n - 1; j++) {
                int d = dist[j];
                f[i][j+1] = (f[i][j] + d + speed - 1) / speed * speed; // (a + b - 1)/b*b = [a/b]向上取整
                if (i > 0) {
                    f[i][j+1] = Math.min(f[i][j+1], f[i-1][j] + d);
                }
            }
            if (f[i][n-1] + dist[n-1] <= speed * hoursBefore) {
                return i;
            }
        }
        return -1;
    }
}