class Solution {
public:
    vector<vector<int>> subsetsWithDup(vector<int>& nums) {
        vector<vector<int>> res;
        vector<int> tmp;
        res.push_back(tmp);
        sort(nums.begin(), nums.end());
        int start, end, p = 0;
        
        for (int i = 0; i < nums.size(); i++) {
            start = (i!=0 && nums[i] == nums[i-1])?p:0;
            end = res.size();
            if (start == end) continue;
            
            for (int j = start; j < end; j++) {
                vector<int> v;
                for (int k = 0; k < res[j].size(); k++)
                    v.push_back(res[j][k]);
                v.push_back(nums[i]);
                res.push_back(v);
            }
            
            p = end;
        }
        
        return res;
    }
};