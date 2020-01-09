/*
	1. 根据身高从大到小排序，如果身高相同，根据k值从小到大排序
	2. 根据k值插入到结果列表的k位置
*/
class Solution {
public:
    static bool cmp(const pair<int, int>& p1, const pair<int, int>& p2)
    {
        if (p1.first != p2.first) return p1.first > p2.first;
        return p1.second < p2.second;
    }
    
    vector<pair<int, int>> reconstructQueue(vector<pair<int, int>>& people) {
        sort(people.begin(), people.end(), cmp);
        vector<pair<int, int>> res;
        for (int i = 0; i < people.size(); i++)
        {
            // printf("(%d, %d), ", people[i].first, people[i].second);
            int pos = people[i].second;
            res.insert(res.begin() + pos, people[i]);
        }
        return res;
    }
};