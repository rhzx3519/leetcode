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

// 简单版本，使用hash判断是否重复
class Solution {
public:
    vector<string> findRepeatedDnaSequences(string s) {
        vector<string> res;
        unordered_map<string, int> um;
        int len = 10, n = s.size();
        for (int i = 0; i <= n-len; i++) {
            string str = s.substr(i, len);
            if (um.count(str) == 0)
                um[str] = 1;
            else
                ++um[str];
            if (um[str]==2)
                res.push_back(str);
        }
        return res;
    }
};