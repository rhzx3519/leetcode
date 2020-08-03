package code;

import java.util.Queue;

class Solution {
    public int kthSmallest(int[][] matrix, int k) {
        Queue<int[]> que = new PriorityQueue<>(
            (a, b) -> matrix[a[0]][a[1]] - matrix[b[0]][b[1]]
        );
        for (int i = 0; i < matrix.length; i++) {
            que.offer(new int[]{i, 0});
        }
        for (int i = k-1; i > 0; i--) {
            int[] tmp = que.poll();
            if (tmp[1] + 1 < matrix[0].length) {
                que.offer(new int[]{tmp[0], tmp[1] + 1});
            }
        }
        int[] tmp = que.peek();
        return matrix[tmp[0]][tmp[1]];
    }
}
// 使用归并的思想，相当于将N行看作N个有序数组进行归并, 每行使用一个指针
// 为了避免每次比较N个数的大小，可以使用小根堆来保存这N个指针