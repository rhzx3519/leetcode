class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        int res = 0, l, r, len;
        map<int, int> mp;
        for (auto &i : nums) {
            if (!mp.count(i)) {
                l = mp.count(i-1)?mp[i-1] : 0;
                r = mp.count(i+1)?mp[i+1] : 0;
                len = l + r + 1;
                res = max(res, len);
                mp[i-l] = len;
                mp[i+r] = len;
                mp[i] = len;
            }
        }
        
        return res;
    }
};