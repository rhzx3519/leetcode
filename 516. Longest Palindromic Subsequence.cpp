class Solution {
public:
    int longestPalindromeSubseq(string s) {
        int n = s.size();
        vector<vector<int>> dp(n, vector<int>(n, 0));
        for (int i = 0; i < n; i++)
            dp[i][i] = 1;
        
        for (int i = n-1; i >= 0; i--) {
            for (int j = i+1; j < n; j++) {
                if (s[i]==s[j])
                    dp[i][j] = dp[i+1][j-1] + 2;
                else
                    dp[i][j] = max(dp[i+1][j], dp[i][j-1]);
            }
        }

        return dp[0][n-1];
    }
};

// **思路：** dp[i][j]表示s[i..j]代表的字符串的最长回文子序列

// d[i][i]=1
// dp[i][j] = dp[i+1][j-1]+2 当s[i]=s[j]
// dp[i][j]=max(dp[i+1][j],dp[i][j-1]) 当s[i]!=s[j] 取s[i+1..j] 和s[i..j-1]中最长的 由于dp[i][j]需要dp[i+1][j]所以需要逆序枚举s的长度，而又因为j是递增的，所以在求解dp[i][j]时,dp[i][j-1]肯定已经求解过了