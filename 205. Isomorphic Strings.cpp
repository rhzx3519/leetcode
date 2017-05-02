class Solution {
public:
    bool isIsomorphic(string s, string t) {
        unordered_map<char,char> ums, umt;
        int ls = s.size(), lt = t.size();
        if (ls != lt) return false;
        
        for (int i = 0; i < ls; i++) {
            if (ums.find(s[i]) == ums.end() && umt.find(t[i]) == umt.end()) {
                ums[s[i]] = t[i];
                umt[t[i]] = s[i];
            }
            if (ums[s[i]] != t[i] || umt[t[i]] != s[i])
                return false;
        }
        
        return true;
    }
};