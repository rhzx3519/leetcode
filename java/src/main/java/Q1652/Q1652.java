package Q1652;

public class Q1652 {
}

/**
 * https://leetcode.cn/problems/defuse-the-bomb/?envType=daily-question&envId=2024-05-05
 */
class Solution {
    public int[] decrypt(int[] code, int k) {
        int n = code.length;
        int[] sums = new int[n+1];
        for (int i = 0; i < n; i++) {
            sums[i+1] = sums[i] + code[i];
        }
        if (k == 0) {
            return new int[n];
        }
        int[] res = new int[n];
        if (k > 0) {
            for (int i = 0; i < n; i++) {
                if (i + k + 1 <= n) {
                    res[i] = sums[i+k+1] - sums[i+1];
                } else {
                    res[i] = sums[n] - sums[i+1] + sums[i+k+1 - n];
                }
            }
        } else {
            k = -k;
            for (int i = 0; i < n; i++) {
                if (i-k >= 0) {
                    res[i] = sums[i] - sums[i-k];
                } else {
                    // k - i
                    res[i] = sums[i] + sums[n] - sums[n - (k - i)];
                }
            }
        }

        return res;
    }

    public static void main(String[] args) {

    }
}

