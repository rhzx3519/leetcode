class Solution {
public:
    int numSquares(int n) {
        vector<int> dp(n+1, INT_MAX);
        for (int i = 1; i*i <= n; i++)
            dp[i*i] = 1;
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j*j < i; j++) 
                dp[i] = min(dp[i], dp[i-j*j]+1);
        }
        return dp[n];
    }
};

// dp[i] 表示组成i的最少的完全平方数个数, dp[i] = min(dp[i], )
//动态方程的意思是：对于每个 i ，比 i 小一个完全平方数的那些数中最小的个数+1就是所求，也就是 dp [ i - j * j ] + 1 。