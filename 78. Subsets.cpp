class Solution {
public:
    vector<vector<int>> subsets(vector<int>& nums) {
        vector<vector<int>> res;
        // vector<int> v;
        // int n = nums.size();
        // dfs(0, n, 0, v, nums, res);
        // return res;
        iterater(nums, res);
        return res;
    }    
    
    void dfs(int idx, int n, int k, vector<int>& v, vector<int>& nums, vector<vector<int>>& res)
    {
        if (idx == n)
        {
            res.push_back(v);
            return;
        }

        v.push_back(nums[idx]);
        dfs(idx+1, n, k, v, nums, res);
        v.pop_back();
        dfs(idx+1, n, k, v, nums, res);
        
    }
    
    void iterater(vector<int>& nums, vector<vector<int>>& res)
    {
        res.push_back(vector<int>());
        int l, r;
        for (int i = 0; i < nums.size(); i++)
        {
            l = 0; r = res.size();
            if (l == r) continue;
            for (int j = l; j < r; j++)
            {
                vector<int> v;
                for (int k = 0; k < res[j].size(); k++)
                    v.push_back(res[j][k]);
                v.push_back(nums[i]);
                res.push_back(v);
            }
        }
    }
};