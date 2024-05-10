package Q2960;

public class Q2960 {
}

/**
 * https://leetcode.cn/problems/count-tested-devices-after-test-operations/?envType=daily-question&envId=2024-05-10
 */
class Solution {
    public int countTestedDevices(int[] batteryPercentages) {
        int d = 0;
        for (int i = 0; i < batteryPercentages.length; i++) {
            if (batteryPercentages[i] - d > 0 ) {
                d++;
            }
        }
        return d;
    }
}
