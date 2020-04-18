class Solution {
public:
    int totalHammingDistance(vector<int>& nums) {
        int res = 0;
        int n = nums.size();
        for (int i = 0; i <= 30; i++) {
            int n1 = 0;
            int t = 0;
            int remain = n;
            for (int j = 0; j < n; j++) {
                n1 += nums[j]&1;
                nums[j] >>= 1;
                t = n1 * (n-n1);
                if (nums[j] == 0)
                    remain--;
            }
            res += t;
            if (remain==0) break;
            
        }
        return res;
    }
};


//  任意两个数字之间的某一位的Hamming distance等于两个numOne * (n - numZero)