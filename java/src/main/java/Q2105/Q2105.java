package Q2105;

public class Q2105 {
}

/**
 * https://leetcode.cn/problems/watering-plants-ii/?envType=daily-question&envId=2024-05-09
 */
class Solution {
    public int minimumRefill(int[] plants, int capacityA, int capacityB) {
        int n = plants.length;
        int a = capacityA, b = capacityB;
        int tot = 0;
        for (int l = 0, r = n-1; l < r; l++, r--) {
            if (a < plants[l]) {
                tot++;
                a = capacityA - plants[l];
            } else {
                a -= plants[l];
            }
            if (b < plants[r]) {
                tot++;
                b = capacityB - plants[r];
            } else {
                b -= plants[r];
            }
        }

        if ((n&1) == 1) {
            int mx = Math.max(a, b);
            mx -= plants[n>>1];
            if (mx < 0) {
                tot++;
            }
        }
        return tot;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        s.minimumRefill(new int[]{1,2,4,4,5}, 6, 5);
    }
}

