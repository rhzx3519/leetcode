/**
 * Definition for undirected graph.
 * struct UndirectedGraphNode {
 *     int label;
 *     vector<UndirectedGraphNode *> neighbors;
 *     UndirectedGraphNode(int x) : label(x) {};
 * };
 */
class Solution {
public:
    UndirectedGraphNode *cloneGraph(UndirectedGraphNode *node) {
        if (!node) return nullptr;
        UndirectedGraphNode* res = new UndirectedGraphNode(node->label);;
        queue<UndirectedGraphNode*> que;
        que.push(node);
        unordered_map<UndirectedGraphNode*, UndirectedGraphNode*> mp;
        mp[node] = res;
        
        while (!que.empty()) {
            auto t = que.front();
            que.pop();
            for (auto &k : t->neighbors) {
                 UndirectedGraphNode* r;
                if (!mp.count(k)) {
                    r = new UndirectedGraphNode(k->label);
                    mp[k] = r;
                    que.push(k);
                } else
                    r = mp[k];
                mp[t]->neighbors.push_back(r);
            }
        }
        return res;
    }
};