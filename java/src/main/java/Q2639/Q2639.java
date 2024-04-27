package Q2639;

public class Q2639 {
}

class Solution {
    private int width(int x) {
        int ans = 0;
        if (x < 0) {
            ans++;
            x = -x;
        }
        if (x == 0) {
            ans++;
            return ans;
        }
        for (; x != 0; x /= 10) {
           ans++;
        }
        return ans;
    }

    public int[] findColumnWidth(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[] ans = new int[n];
        for (int j = 0; j < n; j++) {
            for (int i = 0; i < m; i++) {
                ans[j] = Math.max(ans[j], width(grid[i][j]));
            }
        }
        return ans;
    }
}