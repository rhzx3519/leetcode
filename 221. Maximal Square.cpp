class Solution {
public:
    int maximalSquare(vector<vector<char>>& matrix) {
        if (matrix.empty()) return 0;
        int m = matrix.size(), n = matrix[0].size();
        vector<vector<int>> dp(m, vector<int>(n, 0));
        int x = 0;
        
        for (int i = 0; i < m; i++) {
            dp[i][0] = matrix[i][0] - '0';
            x = max(dp[i][0], x);
        }
            
        for (int i = 0; i < n; i++) {
            dp[0][i] = matrix[0][i] - '0';
            x = max(dp[0][i], x);
        }
        
        
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                dp[i][j] = matrix[i][j] == '1'? min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1 : 0;
                x = max(dp[i][j], x);
            }
        }
        
        return x*x;
    }
};