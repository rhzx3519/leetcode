class Solution {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        if (nums.empty()) return false;
        
        unordered_map<int, int> um;
        vector<int> &a = nums;
        int ls = nums.size();
        if (ls == 1) return false;
        
        for (int i = 0; i < ls && i < k + 1; i++) {
            ++um[a[i]];
            if (um[a[i]] > 1)
                return true;
        }
        
        for (int i = k + 1; i < ls; i++) {
            um.erase(a[i - k - 1]);
            ++um[a[i]];
            if (um[a[i]] > 1)
                return true;
        }
        
        return false;
    }
};