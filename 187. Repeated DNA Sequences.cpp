class Solution {
public:
    vector<string> findRepeatedDnaSequences(string s) {
        vector<string> res;
        int key = 0, len = 10, ls = s.size();
        for (int i = 0; i < ls && i < len - 1; i++) 
            key = (key << 3) | (s[i]&0x7);
            
        unordered_map<int, int> um;
        for (int i = len - 1; i < ls; i++) {
            key = ((key << 3) & 0x3fffffff) | (s[i]&0x7);
            int cnt = ++um[key];
            if (cnt == 2)
                res.push_back(s.substr(i-len+1, len));
        }
        
        return res;
    }
};