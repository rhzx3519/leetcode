class Solution {
public:
    int maxProduct(vector<string>& words) {
        int lw = words.size();
        vector<pair<int, int>> v;
        for (auto &w : words)
            v.push_back(make_pair(get_mask(w), w.size()));
        
        int res = 0;
        for (int i = 0; i < lw; i++) {
            for (int j = i + 1; j < lw; j++) {
                if ((v[i].first & v[j].first) == 0)
                    res = max(res, v[i].second*v[j].second);
            }
        }
        
        return res;
    }
    
    int get_mask(string &w) {
        int res = 0;
        for (auto &ch : w) 
            res |= 1<<(ch-'a');
        return res;
    }
};