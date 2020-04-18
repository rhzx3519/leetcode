unordered_map<string, int> mp;    

struct cmp {
    bool operator()(string s1, string s2) {
        if (mp[s1] != mp[s2]) return mp[s1] > mp[s2];
        return s1 < s2;
    }
};

class Solution {
public:
    vector<string> topKFrequent(vector<string>& words, int k) {
        mp.clear();
        priority_queue<string, vector<string>, cmp> que;
        
        for (string &word : words)
            mp[word]++;
        for (auto it = mp.begin(); it != mp.end(); it++) {
            que.push(it->first);
            if (que.size() > k)
                que.pop();
        }

        vector<string> res;
        while (!que.empty())
        {
            res.push_back(que.top());
            que.pop();
        }
        reverse(res.begin(), res.end());
        return res;
    }

};