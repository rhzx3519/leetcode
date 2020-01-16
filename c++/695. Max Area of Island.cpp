class Solution {
public:
    int current = 0;
    int max_val = 0;

    int maxAreaOfIsland(vector<vector<int>>& grid) {
        for (int i = 0; i < grid.size(); i++) {
            for (int j = 0; j < grid[i].size(); j++) {
                if (grid[i][j] == 1) {
                    current = 0;
                    dfs(i, j, grid);
                    max_val = max(current, max_val);
                }
            }
        }
        return max_val;
    }

    void dfs(int i, int j, vector<vector<int>>& grid) {
        if (i < 0 || i > grid.size()-1 || j < 0 || j > grid[i].size()-1)
            return;
        if (grid[i][j] == 1) {
            current++;
            grid[i][j] = -1;
            dfs(i+1, j, grid);
            dfs(i-1, j, grid);
            dfs(i, j+1, grid);
            dfs(i, j-1, grid);
        }
    }
};