class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        string t;
        if (strs.empty()) return t;
        
        int size = strs[0].size();
        t = strs[0];
        for (int i = 1; i < strs.size(); i++)
        {
            int j;
            for (j = 0; j < min(strs[i].size(), t.size()); j++)
            {
                if (strs[i][j] != t[j]) break;
            }
            t = t.substr(0, j);
        }
        
        return t;
    }
};