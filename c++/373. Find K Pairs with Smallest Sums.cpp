struct compare{
    bool operator() (pair<int, int> a, pair<int, int> b) {
        return a.first + a.second < b.first + b.second;
    }
};

class Solution {
public:
    vector<pair<int, int>> kSmallestPairs(vector<int>& nums1, vector<int>& nums2, int k) {
        priority_queue<pair<int,int>, vector<pair<int,int>>, compare> pq;
        for (int i = 0; i < nums1.size(); i++) {
            for (int j = 0; j < nums2.size(); j++) {
                pq.push(make_pair(nums1[i], nums2[j]));
                if (pq.size() > k)
                    pq.pop();
            }
        }
        
        vector<pair<int,int>> res;
        while(!pq.empty()) {
            res.push_back(pq.top());
            pq.pop();
        }
        
        return res;
    }
};