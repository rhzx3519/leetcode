class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        unordered_set<int> us(nums1.begin(), nums1.end());
        vector<int> res;
        for (int n : nums2) {
            if (us.count(n)) {
                res.push_back(n);
                us.erase(n);
            }
        }
        us.clear();
        
        return res;
    }
};