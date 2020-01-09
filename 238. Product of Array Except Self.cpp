class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        vector<int> res;
        int n = nums.size();
        if (n == 0) return res;
        res.resize(n);

        int i, product = 1;
        for (i = n-1; i >= 0; i--) {
            res[i] = product;
            product *= nums[i];
        }
        
        product = 1;
        for (i = 0; i < n; i++) {
            res[i] *= product;
            product *= nums[i];
        }
        
        return res;
    }
};