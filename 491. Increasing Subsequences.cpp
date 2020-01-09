class Solution {
public:
    vector<vector<int>> findSubsequences(vector<int>& nums) {
        
        set<vector<int>> s;
        vector<int> v;
        dfs(0, v, nums, s);
        vector<vector<int>> res(s.begin(), s.end());
        return res;
    }

    void dfs(int idx, vector<int> &v, vector<int> &nums,  set<vector<int>> &s) {
        if (v.size()>=2) {
            s.insert(v);
        }
        if (idx == nums.size()) {
            return;
        }
        for (int i = idx; i < nums.size(); i++) {
            if (v.empty() || v.back() <= nums[i]) {
                v.push_back(nums[i]);
                dfs(i+1, v, nums, s);
                v.pop_back();
            }
        }
    }
};