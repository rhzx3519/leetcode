package Q741;

import java.util.Arrays;

public class Q741 {
}

/**
 * https://leetcode.cn/problems/cherry-pickup/?envType=daily-question&envId=2024-05-06
 */
class Solution {
    public int cherryPickup(int[][] grid) {
        int n = grid.length;
        int[][][] f = new int[n+1][n+1][n+1];
        for (int i = 0; i <= n; i++) {
            for (int j = 0; j <= n; j++) {
                for (int k = 0; k <= n; k++) {
                    f[i][j][k] = -1;
                }
            }
        }

        f[1][1][1] = grid[0][0];
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                for (int k = 1; k <= Math.min(i+j, n); k++) {
                    int l = i + j - k;
                    if (l > n || l < 1) continue;
                    if (grid[i-1][j-1] == -1 || grid[k-1][l-1] == -1) continue;

                    int res = Math.max(f[i-1][j][k-1], f[i-1][j][k]);
                    res = Math.max(res, f[i][j-1][k]);
                    res = Math.max(res, f[i][j-1][k-1]);
                    if (res < 0) continue;
                    res += grid[i-1][j-1];
                    if (i != k || j != l) {
                        res += grid[k-1][l-1];
                    }
                    f[i][j][k] = res;
                }
            }
        }
        return Math.max(f[n][n][n], 0);
    }
}

