class Solution {
public:
    string findLongestWord(string s, vector<string>& d) {
        sort(d.begin(), d.end(), [](const string &s1, const string &s2){
            if (s1.size() != s2.size())
                return s1.size() > s2.size();
            return s1 < s2;
        });

        for (int i = 0; i < d.size(); i++) {
            if (is_subsequence(d[i], s))
                return d[i];
        }
        return "";
    }

    bool is_subsequence(string &a, string &b) {
        auto i = a.begin();
        auto j = b.begin();
        while (i!=a.end() && j!=b.end()) {
            if (*i==*j)
                i++;
            j++;
        }
        return i==a.end();
    }
};