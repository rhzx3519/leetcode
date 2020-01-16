class Solution {
public:
    int findCircleNum(vector<vector<int>>& M) {
        int m = M.size();
        vector<int> vis(m, 0);
        int res = 0;
        for (int i = 0; i < m; i++) {
            if (vis[i] == 0) {
                dfs(i, M, vis);
                res++;
            }
                
        }
        return res;
    }

    void dfs(int i, vector<vector<int>>& M, vector<int> &vis) {
        for (int j = 0; j < M.size(); j++) {
            if (vis[j] == 0 && M[i][j]==1) {
                vis[j] = 1;
                dfs(j, M, vis);
            }
        }
    }
};