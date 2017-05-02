class Solution {
public:
    vector<vector<int>> combine(int n, int k) {
        vector<vector<int>> res;
        vector<int> nums;
        dfs(1, n, k, nums, res);
        return res;
    }
    
    void dfs(int idx, int n, int k, vector<int> &nums, vector<vector<int>> &res) {
        if (nums.size() == k) {
            res.push_back(nums);
            return;
        }
        
        for (int i = idx; i <= n + 1 - k + nums.size(); i++) {
            nums.push_back(i);
            dfs(i+1, n, k, nums, res);
            nums.pop_back();
        }
    }
};