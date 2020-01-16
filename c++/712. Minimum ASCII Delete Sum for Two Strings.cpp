class Solution {
public:
    int minimumDeleteSum(string s1, string s2) {
        int n1 = s1.size(), n2 = s2.size();
        vector<vector<int>> dp(n1+1, vector<int>(n2+1, 0));
        for (int i = 1; i <= n1; i++)
            dp[i][0] = dp[i-1][0] + int(s1[i-1]);
        for (int i = 1; i <= n2; i++)
            dp[0][i] = dp[0][i-1] + int(s2[i-1]);

        for (int i = 1; i <= n1; i++) {
            for (int j = 1; j <= n2; j++) {
                if (s1[i-1] == s2[j-1]) {
                    dp[i][j] = dp[i-1][j-1];
                } else {
                    dp[i][j] = min(dp[i-1][j] + int(s1[i-1]),
                                dp[i][j-1] + int(s2[j-1])); 
                }
            }
        }

        return dp[n1][n2];
    }
};

///二维辅助数组 dp[i][j]，其含义代表字符串 s1 的子串 s1.sub(0,i-1) 和 s2 的子串 s2.sub(0,j-1) 需删除字符的ASCII值的最小和。

// 1.s1[i-1] == s2[j-1]，新增的两个字符相等的情况下，没有必要删除之前的结果，因此dp[i][j] = dp[i-1][j-1]
// 2.s1[i-1] != s2[j-1]，取三者的最小值
// （1）保留s2串，删除s1串的字符，dp[i][j] = dp[i-1][j] + s1.charAt(i-1)
// （2）保留s1串，删除s2串的字符，dp[i][j] = dp[i][j-1] + s1.charAt(j-1)
// （3）删除s1、s2串的字符，dp[i][j] = dp[i-1][j-1] + s1.charAt(i-1) + s2.charAt(j-1)