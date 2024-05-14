package Q2244;

import java.util.HashMap;
import java.util.Map;

public class Q2244 {
}


/**
 * https://leetcode.cn/problems/minimum-rounds-to-complete-all-tasks/?envType=daily-question&envId=2024-05-14
 */
class Solution {
    public int minimumRounds(int[] tasks) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int task : tasks) {
            map.put(task, map.getOrDefault(task, 0) + 1);
        }
        int tot = 0;
        for (int task : map.keySet()) {
            int c = map.get(task);
            if (c == 1) {
                return -1;
            }
            tot += (c + 2)/3;
        }
        return tot;
    }
}