class Solution {
public:
    int numSquares(int n) {
        int res = 0;
        vector<int> dp(n+1, 0);
        
        for (int i = 1; i <= n; i++) {
            dp[i] = INT_MAX;
            for (int j = sqrt(i); j > 0; j--) {
                dp[i] = min(dp[i], dp[i - j*j] + 1);
            }
        }
        
        return dp[n];
    }
};