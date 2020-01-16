class Solution {
public:
    vector<vector<int>> combinationSum3(int k, int n) {
        vector<vector<int>> res;
        vector<int> v;
        dfs(1, k, n, v, res);
        return res;
    }
    
    void dfs(int idx, int k, int n, vector<int> &v, vector<vector<int>> &res) {
        if (n < 0 || k < v.size()) return;
        
        if (n == 0 && k == v.size()) {
            res.push_back(v);
            return;
        }
        
        for (int i = idx; i <= 9; i++) {
            v.push_back(i);
            dfs(i+1, k, n - i, v, res);
            v.pop_back();
        }
    }
};