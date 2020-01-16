class Solution {
public:
    int numIslands(vector<vector<char>>& grid) {
        if (grid.empty()) return 0;
        int m = grid.size(), n = grid[0].size(), res = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '0')
                    continue;
                res++;
                dfs(grid, i, j);
            }
        }
        
        return res;
    }
    
    void dfs(vector<vector<char>>& grid, int i, int j) {
        grid[i][j] = '0';
        int m = grid.size(), n = grid[0].size();
        
        for (int k = 0; k < 4; k++) {
            int x = i + dx[k], y = j + dy[k];
            if (x < 0 || x > m-1 || y < 0 || y > n-1)
                continue;
            if (grid[x][y] == '1')
                dfs(grid, x, y);
        }
    }
    
private:
    int dx[4] = {1,-1,0,0};
    int dy[4] = {0,0,1,-1};
};