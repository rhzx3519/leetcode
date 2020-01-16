class Solution {
public:
    int minPatches(vector<int>& nums, int n) {
        int64_t miss = 1;
        int res = 0, i = 0;
        while (miss <= n) {
            if (i < nums.size() && nums[i] <= miss)
                miss += nums[i++];
            else {
                miss += miss;
                res++;
            }
        }
        
        return res;
    }
};