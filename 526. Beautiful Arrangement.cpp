class Solution {
public:
    int countArrangement(int N) {
        vector<int> vis(N+1, 0);
        int res = 0;
        dfs(1, res, N, vis);
        return res;
    }

    void dfs(int idx, int &res, int N, vector<int> &vis) {
        if (idx == N+1) {
            res++;
            return;
        }

        for (int i = 1; i <= N; i++) {
            if (vis[i]==1) continue;
            vis[i] = 1;
            if (i%idx==0 || idx%i==0)
                dfs(idx+1, res, N, vis);
            vis[i] = 0;
        }
    }
};