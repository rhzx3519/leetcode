package Q2952;

import java.util.Arrays;

/**
 * https://leetcode.cn/problems/minimum-number-of-coins-to-be-added/description/?envType=daily-question&envId=2024-03-30
 */
public class Q2952 {
    public int minimumAddedCoins(int[] coins, int target) {
        Arrays.sort(coins);
        int ans = 0;
        int s = 1;
        int i = 0;
        while (s <= target) {
            if (i < coins.length && coins[i] <= s) {
                s += coins[i];
                i++;
            } else {
                s *= 2;
                ans ++;
            }
        }
        return ans;
    }
}
