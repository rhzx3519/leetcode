package Q2580;

import Q518.Q518;

import java.util.Arrays;
import java.util.Comparator;

public class Q2580 {
    private final int MOD = 1_000_000_000 + 7;
    public int countWays(int[][] ranges) {
        Arrays.sort(ranges, Comparator.comparingInt(a -> a[0]));
        int ans = 2;
        int maxR = ranges[0][1];
        for (int[] range : ranges) {
            if (range[0] > maxR) {
                ans = (ans<<1) % MOD;
            }
            maxR = Math.max(maxR, range[1]);
        }
        return ans;
    }

    public static void main(String[] args) {
        Q2580 q = new Q2580();
        q.countWays(new int[][]{{1,3},{10,20},{2,5},{4,8}});
    }
}
