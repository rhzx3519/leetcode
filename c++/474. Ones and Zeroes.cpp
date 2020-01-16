class Solution {
public:
    int findMaxForm(vector<string>& strs, int m, int n) {
        vector<vector<int>> dp(m+1, vector<int>(n+1, 0)); //多维费用问题的0-1背包问题。dp[i][j]表示使用i个0和j个1能表示的字符串的最大数量
        for (string &str : strs) {
            int n0 = 0, n1 = 0;
            count(str, n0, n1);
            for (int i = m; i >= n0; i--)
                for (int j = n; j >= n1; j--)
                    dp[i][j] = max(dp[i][j], 1+dp[i-n0][j-n1]); //状态转移方程：dp[i][j] = Math.max(dp[i][j],1+dp[i-numZero][j-numOne])
        }

        return dp[m][n];
    }

    void count(string &s, int &n0, int &n1){
        for (char c : s) {
            if (c=='0')
                n0++;
            else if (c=='1')
                n1++;
        }
    }
};