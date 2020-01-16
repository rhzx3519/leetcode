class Solution {

public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        dfs(nums, 0, nums.size() - 1);
        return ans;
    }
    
    void dfs(vector<int> nums, int left, int right) {
        if (left == right)
            ans.push_back(nums);

        for (int i = left; i <= right; i++) {
            if (i != left && nums[left] == nums[i]) continue; 
            swap(nums[left], nums[i]);
            dfs(nums, left + 1, right);
        }
    }

private:
    vector<vector<int>> ans;

};