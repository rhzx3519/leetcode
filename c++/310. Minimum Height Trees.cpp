class Solution {
public:
    vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges) {
        vector<int> t;
        for (int i = 0; i < n; i++)
            t.push_back(i); 

        if (n<=2) return t;
        vector<vector<int>> v(n);
        vector<int> ind(n, 0);
        for (int i = 0; i < edges.size(); i++) {
            int a = edges[i][0];
            int b = edges[i][1];
            v[a].push_back(b);
            v[b].push_back(a);
            ind[a]++;
            ind[b]++;
        }

        vector<int> leaves;
        for (int i = 0; i < n; i++) {
            if (ind[i]==1) {
                leaves.push_back(i);
            }
        }

        while (n > 2) {
            n -= leaves.size();
            vector<int> new_leaves;
            for (int leaf : leaves) {
                for (int i = 0; i < v[leaf].size(); i++) {
                    int b = v[leaf][i];
                    ind[b]--;
                    if (ind[b] == 1)
                        new_leaves.push_back(b);
                }
            }
            leaves = new_leaves;
        }

        return leaves;
    }

};