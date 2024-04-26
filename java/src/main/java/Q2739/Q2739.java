package Q2739;

public class Q2739 {
}

/**
 * https://leetcode.cn/problems/total-distance-traveled/?envType=daily-question&envId=2024-04-25
 */
class Solution {
    public int distanceTraveled(int mainTank, int additionalTank) {
        return ( Math.min(Math.min(1, mainTank/5) + Math.max(0, mainTank-5)/4, additionalTank) + mainTank) * 10;
    }
}


