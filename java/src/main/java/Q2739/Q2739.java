package Q2739;

public class Q2739 {
}

/**
 * https://leetcode.cn/problems/total-distance-traveled/?envType=daily-question&envId=2024-04-25
 */
class Solution {
    public int distanceTraveled(int mainTank, int additionalTank) {
        return (Math.min(mainTank/5, additionalTank) + mainTank) * 10;
    }
}


