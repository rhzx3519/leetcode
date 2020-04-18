class Solution {
public:
    int findMaxLength(vector<int>& nums) {
        int res = 0;
        int n = nums.size();
        int sum = 0;
        unordered_map<int, int> um;
        um[0] = -1;
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i]==1)
                sum += 1;
            else
                sum += -1;
            if (um.count(sum) == 0)
                um[sum] = i;
            else {
                int last = um[sum];
                res = max(res, i - last);
            }
        }

        return res;
    }
};