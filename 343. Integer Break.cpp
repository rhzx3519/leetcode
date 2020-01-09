class Solution {
public:
    int integerBreak(int n) {
        vector<int> dp(n+1, 0);
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <i; j++) {
                dp[i] = max(dp[i], j * max(i-j, dp[i-j]));
            }
        }
        return dp[n];
    }
};

// dp[n] = max(dp[n], i*max(dp[n-i], n-i))