class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int res = 0;
        vector<int> dict(256, -1);
        int i = 0, start = 0;
        for (; i < s.size(); i++) {
            start = max(start, dict[s[i]] + 1);
            dict[s[i]] = i;
            res = max(res, i - start + 1);
        }
        
        return res;
    }
};