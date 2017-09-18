class Solution {
public:
    string convert(string s, int numRows) {
        int n = numRows, k = 0;
        int offset = 1;
        vector<string> a(n, string());
        for (int i = 0; i < s.size(); i++)
        {
            a[k].push_back(s[i]);
            if (k == 0)
                offset = 1;
            else if (k == n-1)
                offset = -1;
            k += offset;
        }
        
        string res;
        for (auto &str : a)
        {
            if (!str.empty())
                res += str;
        }
            
        return res;
    }
};