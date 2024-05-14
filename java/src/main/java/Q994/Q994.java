package Q994;

import java.util.Deque;
import java.util.LinkedList;
import java.util.Queue;

public class Q994 {
}

/**
 * https://leetcode.cn/problems/rotting-oranges/?envType=daily-question&envId=2024-05-13
 */
class Solution {
    private int[][] offset = new int[][]{{0, 1}, {1, 0}, {0, -1}, {-1, 0}};

    public int orangesRotting(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        Deque<Integer> q = new LinkedList<>();
        int fresh = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 2) {
                    q.offerLast(i*n + j);
                    grid[i][j] = 0;
                } else if (grid[i][j] == 1) {
                    fresh++;
                }
            }
        }
        if (fresh == 0) {
            return 0;
        }
        int tot = 0;
        while (!q.isEmpty()) {
            for (int sz = q.size(); sz != 0; sz--) {
                int r = q.pollFirst();
                int a = r/n;
                int b = r%n;
                for (int[] d : offset)  {
                    int i = a + d[0];
                    int j = b + d[1];
                    if (i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1) {
                        grid[i][j] = 0;
                        fresh--;
                        q.offerLast(i*n + j);
                    }
                }
            }
            tot++;
        }
        if (fresh != 0) {
            return -1;
        }
        return tot - 1;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        s.orangesRotting(new int[][]{{2,1,1},{1,1,0},{0,1,1}});
    }
}

