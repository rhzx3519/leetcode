class Solution {
public:
    vector<vector<string>> partition(string s) {
        vector<vector<string>> res;
        int ls = s.size();
        vector<vector<bool>> is_pal(ls, vector<bool>(ls, false));
        for (int i = 0; i < ls; i++)
            is_pal[i][i] = true;
        for (int i = 0; i < ls-1; i++) 
            is_pal[i][i+1] = (s[i] == s[i+1]);
        
        for (int i = 2; i < ls; i++) 
            for (int j = 0; i+j < ls; j++)
                is_pal[j][i+j] = is_pal[j+1][i+j-1]&&(s[j]==s[i+j]);
        
        vector<string> v;
        dfs(0, s, v, res, is_pal);
        return res;
    }
    
    void dfs(int idx, string &s, vector<string> &v, vector<vector<string>> &res,  vector<vector<bool>> &is_pal) {
        int ls = s.size();
        if (idx == ls) {
            res.push_back(v);
            return;
        }
        
        for (int i = idx; i < ls; i++) {
            if (!is_pal[idx][i]) continue;
            v.push_back(s.substr(idx, i-idx+1));
            dfs(i+1, s, v, res, is_pal);
            v.pop_back();
        }
            
    }
};