package Q1463;

public class Q1463 {
}

/**
 * https://leetcode.cn/problems/cherry-pickup-ii/description/?envType=daily-question&envId=2024-05-07
 */
class Solution {
    public int cherryPickup(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[][][] f = new int[m + 1][n + 2][n + 2];
        for (int i = m - 1; i >= 0; i--) {
            for (int j = 0; j < Math.min(n, i + 1); j++) {
                for (int k = Math.max(j + 1, n - 1 - i); k < n; k++) {
                    f[i][j + 1][k + 1] = max(
                            f[i + 1][j][k], f[i + 1][j][k + 1], f[i + 1][j][k + 2],
                            f[i + 1][j + 1][k], f[i + 1][j + 1][k + 1], f[i + 1][j + 1][k + 2],
                            f[i + 1][j + 2][k], f[i + 1][j + 2][k + 1], f[i + 1][j + 2][k + 2]
                    ) + grid[i][j] + grid[i][k];
                }
            }
        }
        return f[0][1][n];
    }

    private int max(int x, int... y) {
        for (int v : y) {
            x = Math.max(x, v);
        }
        return x;
    }
}


