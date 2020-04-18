#include <queue>

struct comparator{
    bool operator()(const vector<int>& p1, const vector<int> &p2) {
        return p1[0]*p1[0] + p1[1]*p1[1] < p2[0]*p2[0] + p2[1]*p2[1];
    }
};

class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int K) {
        priority_queue<vector<int>, vector<vector<int>>, comparator> que;
        for (int i = 0; i < points.size(); i++) {
            que.push(points[i]);
            if(que.size()>K) {
                que.pop();
            }
        }
        vector<vector<int>> ans;
        while (!que.empty()) {
            ans.push_back(que.top());
            que.pop();
        }
        return ans;
    }
};