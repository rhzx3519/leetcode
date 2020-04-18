class Solution {
public:
    int findPaths(int m, int n, int N, int i, int j) {
        long mod = long(1e9 + 7);
        vector<vector<vector<int>>> dp(N+1, vector<vector<int>>(m, vector<int>(n, 0)));
        for (int k = 1; k <= N; k++) {
            for (int i = 0; i < m; i++) {
                for (int j = 0; j < n; j++) {
                    long k1 = i==0?1:dp[k-1][i-1][j];
                    long k2 = j==0?1:dp[k-1][i][j-1];
                    long k3 = i==m-1?1:dp[k-1][i+1][j];
                    long k4 = j==n-1?1:dp[k-1][i][j+1];
                    dp[k][i][j] = (k1 + k2 + k3 + k4) % mod;
                }
            }
        }

        return dp[N][i][j];
    }
};

// 对于一个起始点为i，j，N步可以走出的点的路径个数，等于该点周围的4个点，N-1步可以走出的路径个数之和，知道了这个之后，我们就可以以这个公式作为状态转移方程。