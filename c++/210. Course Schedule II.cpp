class Solution {
public:
    vector<int> findOrder(int numCourses, vector<pair<int, int>>& prerequisites) {
        int nv = numCourses, ne = prerequisites.size();
        vector<pair<int, int>>& e = prerequisites;
        queue<int> que;
        vector<vector<int>> v(nv);
        vector<int> in(nv, 0);
        
        int i, j, x , y;
        for (i = 0; i < ne; i++) {
            x = e[i].first;
            y = e[i].second;
            v[y].push_back(x);
            in[x]++;
        }
        
        for (i = 0; i < nv; i++) {
            if (in[i] == 0)
                que.push(i);
        }
        
        vector<int> res;
        int ec = ne;
        while (!que.empty()) {
            x = que.front();
            que.pop();
            res.push_back(x);
            
            while (!v[x].empty()) {
                y = v[x].back();
                v[x].pop_back();
                --in[y];
                --ec;
                if (in[y] == 0)
                    que.push(y);
            }
        }
        
        if (ec == 0)
            return res;
        res.clear();
        return res;
    }
};