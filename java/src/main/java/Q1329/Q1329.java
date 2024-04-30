package Q1329;

import java.util.*;

public class Q1329 {
}

/**
 * https://leetcode.cn/problems/sort-the-matrix-diagonally/?envType=daily-question&envId=2024-04-29
 */
class Solution {
    public int[][] diagonalSort(int[][] mat) {
        int m = mat.length;
        int n = mat[0].length;
        for (int a = 0; a < m; a++) {
            int b = 0;
            List<Integer> arr = new ArrayList<>();
            for (int i = a, j = b; i < m && j < n; i++, j++) {
                arr.add(mat[i][j]);
            }
            Collections.sort(arr);
            for (int i = a, j = b; i < m && j < n; i++, j++) {
                mat[i][j] = arr.get(j);
            }
        }
        for (int a = 1; a < n; a++) {
            int b = 0;
            List<Integer> arr = new ArrayList<>();
            for (int i = b, j = a; i < m && j < n; i++, j++) {
                arr.add(mat[i][j]);
            }
            Collections.sort(arr);
            for (int i = b, j = a; i < m && j < n; i++, j++) {
                mat[i][j] = arr.get(i);
            }
        }
        return mat;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        s.diagonalSort(new int[][]{{3,3,1,1}, {2,2,1,2}, {1,1,1,2}});
    }
}


