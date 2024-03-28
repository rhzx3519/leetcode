package Q1997;

import java.util.HashMap;
import java.util.Map;

public class Q1997 {
    private final int mod = 1_000_000_007;

    public int firstDayBeenInAllRooms(int[] nextVisit) {
        int n = nextVisit.length;
        long[] s= new long[n];
        for (int i = 0; i < n - 1; i ++) {
            int j = nextVisit[i];
            s[i+1] = (s[i] * 2 - s[j] + 2 + mod)  % mod;
        }
        return (int) s[n-1];
    }

    public static void main(String[] args) {
        Q1997 q = new Q1997();
        System.out.println(q.firstDayBeenInAllRooms(new int[]{0,0}));
        System.out.println(q.firstDayBeenInAllRooms(new int[]{0,0,2}));
    }
}
