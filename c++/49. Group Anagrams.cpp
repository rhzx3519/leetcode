class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        map<string, vector<string>> mp;
        for(auto &str : strs) {
            string s = str;
            sort(str.begin(), str.end());
            mp[str].push_back(s);
        }
        vector<vector<string>> res;
        for (map<string, vector<string>>::iterator iter = mp.begin(); iter != mp.end(); iter++)
            res.push_back(iter->second);
        return res;
    }
};