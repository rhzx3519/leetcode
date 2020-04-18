class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> um;
        for (auto &i : nums)
            um[i]++;
        
        vector<int> res;
        priority_queue<pair<int, int>> pq;
        for (auto it = um.begin(); it != um.end(); it++) {
            pq.push(make_pair(it->second, it->first));
            if (pq.size() > um.size() - k) {
                res.push_back(pq.top().second);
                pq.pop();
            }
        }
        
        return res;
    }
};