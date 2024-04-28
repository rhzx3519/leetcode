package Q1017;

public class Q1017 {
}

/**
 * https://leetcode.cn/problems/convert-to-base-2/?envType=daily-question&envId=2024-04-28
 */
class Solution {
    public String baseNeg2(int n) {
        if (n == 0) {
            return "0";
        }
        if (n == 1) {
            return "1";
        }
        StringBuffer sb = new StringBuffer();
        for (int i = n; i != 0; ) {
            int d = i&1;
            if (d == 0) {
                sb.append("0");
            } else {
                sb.append("1");
            }
            i -= d;
            i /= -2;
        }
        return sb.reverse().toString();
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        s.baseNeg2(2);
    }
}
