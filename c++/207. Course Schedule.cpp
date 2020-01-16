class Solution {
public:
    bool canFinish(int numCourses, vector<pair<int, int>>& prerequisites) {
        int nv = numCourses, ne = prerequisites.size();
        vector<pair<int, int>> &e = prerequisites;
        vector<vector<int>> v(nv);
        vector<int> ind(nv, 0);
        queue<int> que;
        
        for (int i = 0; i < ne; i++) {
            int a = e[i].first, b = e[i].second;
            ind[a]++;
            v[b].push_back(a);
        }
        
        for (int i = 0; i < nv; i++) {
            if (ind[i] == 0)
                que.push(i); // 入度为0的结点进队列
        }
        
        int nc = ne;
        while (!que.empty()) {
            int a = que.front();
            que.pop();
            while(!v[a].empty()) {
                int b = v[a].back();
                v[a].pop_back();
                --ind[b];
                if (ind[b] == 0)
                    que.push(b);
                nc--;
            }
        }
        
        return (nc==0);
    }
};