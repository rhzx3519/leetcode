package Q2924;

public class Q2924 {
}

class Solution {
    public int findChampion(int n, int[][] edges) {
        int[] ind = new int[n];
        for (int[] e : edges) {
            ind[e[1]]++;
        }
        int champian = -1;
        for (int i = 0; i < n; i++) {
            if (ind[i] == 0) {
                if (champian != -1) {
                    return -1;
                }
                champian = i;
            }
        }
        return champian;
    }
}
