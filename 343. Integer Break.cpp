class Solution {
public:
    int integerBreak(int n) {
        vector<int> dp = {0, 0, 1};
        while (dp.size() <= n) {
            next(dp);
        }
        
        return dp[n];
    }
    
    void next(vector<int> &dp) {
        int n = dp.size();
        dp.push_back(dp.back());
        for (int i = 2; i <= n; i++) 
            dp[n] = max(dp[n], i*max(n-i, dp[n-i]));
    }
};