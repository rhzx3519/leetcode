package Q2391;

public class Q2391 {
}


/**
 * https://leetcode.cn/problems/minimum-amount-of-time-to-collect-garbage/?envType=daily-question&envId=2024-05-11
 */
class Solution {
    public int garbageCollection(String[] garbage, int[] travel) {
        int[] dis = new int[3];
        int collect = 0, t = 0;
        for (int i = 0; i < garbage.length; i++) {
            collect += garbage[i].length();

            if (i > 0) {
                t += travel[i-1];
                if (garbage[i].contains("G")) {
                    dis[0] = t;
                }
                if (garbage[i].contains("P")) {
                    dis[1] = t;
                }
                if (garbage[i].contains("M")) {
                    dis[2] = t;
                }
            }
        }
        return dis[0] + dis[1] + dis[2] + collect;
    }
}


