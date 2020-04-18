class Solution {
public:
    vector<int> findRedundantConnection(vector<vector<int>>& edges) {
        int n = edges.size();
        vector<int> pre(n+1);
        for (int i = 1; i <= n; i++)
            pre[i] = i;

        for (auto &edge : edges) {
            int x = edge[0], y = edge[1];
            int px = find(x, pre);
            int py = find(y, pre);
            if (px != py)
                pre[py] = px;
            else
                return {x, y};
        }
        return {};            
    }

    int find(int x, vector<int> &pre) { // 并查集查询父节点
        if (pre[x] != x)
            pre[x] = find(pre[x], pre);
        return pre[x];
    }
};

// 并查集，查询第一个祖先节点相同的两个节点之间的边