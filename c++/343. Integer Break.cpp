class Solution {
public:
    int integerBreak(int n) {
<<<<<<< HEAD
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
=======
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
>>>>>>> 837bde34d06bb470ffcea088785f591f742870d7
