class Solution {
public:
    vector<string> restoreIpAddresses(string s) {
        vector<string> res;
        for (int i = 1; i <= 3; i++)
            dfs(0, 0, i, "", s, res);
        return res;
    }
    
    void dfs(int cnt, int start, int end, string str, string &s, vector<string> &res) {
        if (end > s.size() || cnt > 3) return;
        string sub = s.substr(start, end - start);
        if (sub.size() > 1 && sub[0] == '0') return;
        if (atoi(sub.c_str()) > 255) return;
        str += sub;
        if (end != s.size())
            str += '.';
        else {
            if (cnt == 3) {
                res.push_back(str);
                return;
            }
        }
        
        for (int i = 1; i <= 3; i++)
            dfs(cnt+1, end, end + i, str, s, res);
    }
    
};