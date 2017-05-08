class Solution {
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        vector<int> res;
        unordered_map<int, int> um;
        for (auto i : nums1)
            um[i]++;
        for (auto i : nums2) {
            if (um.find(i) != um.end() && --um[i] >= 0) res.push_back(i);
        }
        return res;
    }
};