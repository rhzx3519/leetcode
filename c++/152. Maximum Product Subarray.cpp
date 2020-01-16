class Solution {
public:
    /**
     * @param nums: a vector of integers
     * @return: an integer
     */
    int maxProduct(vector<int>& nums) {
        // write your code here
        int pos_max = nums[0];
        int neg_max = nums[0];
        int ret = nums[0];
        for (int i = 1; i < nums.size(); i++) {
            int tmp_pos_max = pos_max;
            int tmp_neg_max = neg_max;
            pos_max = max(nums[i], max(nums[i]*tmp_pos_max, nums[i]*tmp_neg_max));
            neg_max = min(nums[i], min(nums[i]*tmp_pos_max, nums[i]*tmp_neg_max));
            ret = max(pos_max, ret);
        }
        
        return ret;
    }
};