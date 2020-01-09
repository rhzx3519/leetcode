class Solution {
public:
    string frequencySort(string s) {
        unordered_map<char, int> um;
        set<char> st;
        for (char ch : s)
        {
            st.insert(ch);
            um[ch]++;
        }
        vector<char> a(st.begin(), st.end());
        sort(a.begin(), a.end(), [&um](const char& c1, const char& c2){
            return um[c1] > um[c2];
        });
        
        string res;
        for (int i = 0; i < a.size(); i++) {
            string t(um[a[i]], a[i]);
            res.append(t);
        }
        return res;
    }
};