class Solution {
public:
    vector<int> majorityElement(vector<int>& nums) {
        vector<int> res;
        int n = nums.size();
        if (n == 0) return res;
        
        int n1, n2;
        int c1 = 0, c2 = 0;
        
        for (auto &i : nums) {
            if (i == n1)
                c1++;
            else if (i == n2)
                c2++;
            else if (c1 == 0) {
                n1 = i;
                c1 = 1;
            } else if (c2 == 0) {
                n2 = i;
                c2 = 1;
            } else {
                c1--;
                c2--;
            }
        }

        c1 = c2 = 0;
        for (auto &i : nums) {
            if (i==n1)
                c1++;
            else if (i == n2)
                c2++;
        }
        
        if (c1 > n/3.f)
            res.push_back(n1);
        if (c2 > n/3.f)
            res.push_back(n2);
        
        return res;
    }
};