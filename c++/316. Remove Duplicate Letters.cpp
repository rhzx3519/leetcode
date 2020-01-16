class Solution {
public:
    string removeDuplicateLetters(string s) {
        static const int DICT_SIZE = 26;
        int ls = s.size();
        string res;
        if (ls < 2) return s;
        
        int c1[DICT_SIZE] = {0}, c2[DICT_SIZE] = {0}; // c1记录了源字符串中出现的字符数量，c2记录了结果字符串中出现的字符数量
        int i;
        for (i = 0; i < ls; i++)
            ++c1[s[i] - 'a'];
            
        i = 0; int k;
        for (i = 0; i < ls; i++) {
            k = s[i] - 'a';
            if (c2[k] > 0) {
                c1[k]--;
                continue;
            }
            
            while (res.size() > 0 && res.back() >= s[i] && c1[res.back() - 'a'] > 0) { 
                --c2[res.back() - 'a'];
                res.pop_back();
            }
            res.push_back(s[i]);
            
            c1[k]--;
            c2[k]++;
        }
        
        return res;
    }
};