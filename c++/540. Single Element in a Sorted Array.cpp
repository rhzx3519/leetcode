class Solution {
public:
    int singleNonDuplicate(vector<int>& nums) {
        int n = nums.size();
        if (n==1) return nums[0];
        for (int i = 0; i < n/2; i++) {
            if (nums[2*i] != nums[2*i+1]) {
                return nums[2*i];
            }

        }
        return nums.back();
    }
};