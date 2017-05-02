class Solution {
public:
    int longestIncreasingPath(vector<vector<int>>& matrix) {
        if (matrix.empty()) return 0;
        int m = matrix.size(), n = matrix[0].size();
        vector<vector<int>> dp(m, vector<int>(n, 0));
        int res = 0;
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                res = max(res, dfs(matrix, dp, i, j));
            }
        }
        
        return res;
    }
    
    int dfs(vector<vector<int>>& matrix, vector<vector<int>> &dp, int i, int j) {
        if (dp[i][j] != 0) 
            return dp[i][j];
        int m = matrix.size(), n = matrix[0].size();
        
        dp[i][j] = 1;
        for (int k = 0; k < 4; k++) {
            int x = i + offset[k][0];
            int y = j + offset[k][1];
            if (x < 0 || x > m-1 || y < 0 || y > n-1)
                continue;
            if (matrix[x][y] > matrix[i][j])
                dp[i][j] = max(dp[i][j], dfs(matrix, dp, x, y) + 1);
        }
        
        return dp[i][j];
    }
    
private:
    const int offset[4][2] = {{1,0}, {-1, 0}, {0, 1}, {0, -1}};
};