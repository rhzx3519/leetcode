class Solution {
public:
    vector<vector<int>> combine(int n, int k) {
        vector<vector<int>> res;
        vector<int> nums;
        dfs(0, n, k, nums, res);
        return res;
    }

    void dfs(int idx, int n, int k, vector<int>& nums, vector<vector<int>>& res)
    {
        if (idx == n)
        {
            if (nums.size() == k)
                res.push_back(nums);
            return;
        }

        nums.push_back(idx+1);
        dfs(idx+1, n, k, nums, res);
        nums.pop_back();
        dfs(idx+1, n, k, nums, res);
        
    }
};