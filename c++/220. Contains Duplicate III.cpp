class Solution {
public:
     bool containsNearbyAlmostDuplicate(vector<int>& nums, int k, int t) {
        set<int64_t> window; // set is ordered automatically 
        for (int i = 0; i < nums.size(); i++) {
            if (i > k) window.erase((int64_t)nums[i-k-1]); // keep the set contains nums i j at most k
            // |x - nums[i]| <= t  ==> -t <= x - nums[i] <= t;
            auto pos = window.lower_bound((int64_t)nums[i] - (int64_t)t); // x-nums[i] >= -t ==> x >= nums[i]-t 
            // x - nums[i] <= t ==> |x - nums[i]| <= t    
            if (pos != window.end() && *pos - (int64_t)nums[i] <= t) return true;
            window.insert((int64_t)nums[i]);
        }
        return false;
    }
};