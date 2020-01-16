class Solution {
public:
    vector<vector<int>> permute(vector<int>& nums) {
        vector<vector<int>> res;
        dfs(0, nums, res);
        
        return res;
    }
    
    void dfs(int idx, vector<int>& nums, vector<vector<int>>& res)
    {
        if (idx == nums.size())
        {
            res.push_back(nums);
            return;
        }
        
        for (int i = idx; i < nums.size(); i++)
        {
            swap(nums[i], nums[idx]);
            dfs(idx+1, nums, res);
            swap(nums[i], nums[idx]);
        }
    }
};