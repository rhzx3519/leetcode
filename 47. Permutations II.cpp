class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        vector<vector<int>> res;
        dfs(nums, 0, res);
        return res;
    }
    
    void dfs(vector<int> nums, int pos, vector<vector<int>> &res) {
        if (pos == nums.size() - 1) {
            res.push_back(nums);
            return;
        }
        
        sort(nums.begin() + pos, nums.end());
        for (int i = pos; i < nums.size(); i++) {
            if (i != pos && nums[i - 1] == nums[i])
                continue;
            swap(nums[i], nums[pos]);
            dfs(nums, pos + 1, res);
            swap(nums[i], nums[pos]);
        }
    }
};