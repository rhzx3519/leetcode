class Solution {
public:
    int longestSubstring(string s, int k) {
        unordered_map<char, int> um;
        string t;
        for (char c: s)
            um[c]++;
        
        int i = 0;
        while (i < s.size()) {
            if (um[s[i]] < k)
                break;
            i++;
        }
        if (i==s.size()) return s.size();

        int left = longestSubstring(s.substr(0, i), k);
        while (i < s.size() && um[s[i]] < k)
            i++;
        int right = longestSubstring(s.substr(i), k);
        return max(left, right);
    }
};