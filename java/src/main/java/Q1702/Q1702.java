package Q1702;

public class Q1702 {
}


/**
 * https://leetcode.cn/problems/maximum-binary-string-after-change/?envType=daily-question&envId=2024-04-10
 */
class Solution {
    public String maximumBinaryString(String binary) {
        int n = binary.length();
        int i = binary.indexOf("0");
        if (i == -1) {
            return binary;
        }
        int c1 = 0;
        for (; i < n; i++) {
            if (binary.charAt(i) == '1') {
                c1++;
            }
        }

        StringBuffer sb = new StringBuffer();
        for (i = 0; i < n - c1 - 1; i++) {
            sb.append('1');
        }
        sb.append('0');
        for (i = 0; i < c1; i++) {
            sb.append('1');
        }
        return sb.toString();
    }
}

