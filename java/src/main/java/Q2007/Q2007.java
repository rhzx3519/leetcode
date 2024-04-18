package Q2007;

import java.util.*;
import java.util.stream.Stream;

/**
 * https://leetcode.cn/problems/find-original-array-from-doubled-array/description/?envType=daily-question&envId=2024-04-18
 */
public class Q2007 {

}


class Solution {
    public int[] findOriginalArray(int[] changed) {
        int n = changed.length;
        if ((n&1) != 0) return new int[]{};
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < n; i++) {
            map.put(changed[i], map.getOrDefault(changed[i], 0) + 1);
        }
        List<Integer> ans = new ArrayList<>();
        for (int i : map.keySet()) {
            if ((i&1) == 1) continue;
            int c = map.getOrDefault(i, 0);
            if (c == 0) {
                continue;
            }
            int cj = map.getOrDefault(i>>1, 0);
            if (cj > 0) {
                int m = Math.min(c, cj);
                if (i == i>>1) {
                    m = m>>1;
                }
                map.put(i, c - m);
                cj = map.getOrDefault(i>>1, 0);
                map.put(i>>1, cj - m);
                for (int j = 0; j < m; j++) {
                    ans.add(i>>1);
                }
            }
        }
        for (int i : map.keySet()) {
            if (map.getOrDefault(i, 0) != 0) {
                return new int[]{};
            }
        }
        return ans.stream().mapToInt(i -> i).toArray();
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[] ans = s.findOriginalArray(new int[]{0,0,0,0});
        System.out.println(ans);
    }
}


