struct Gap {
public:
    Gap(vector<int>& _inter, int _idx)
    {
        inter =_inter;
        idx = _idx;
    }

public:
    vector<int> inter;
    int idx;
    int right = -1;
};

class Solution {
public:
    vector<int> findRightInterval(vector<vector<int>>& intervals) {
        vector<int> res;
        vector<Gap> gaps;
        for (int i = 0; i < intervals.size(); i++) {
            gaps.push_back(Gap(intervals[i], i));
        }
            
        sort(gaps.begin(), gaps.end(), [](const Gap &g1, const Gap &g2){
            if (g1.inter[0] == g2.inter[0]) return g1.inter[1] < g2.inter[1];
            return g1.inter[0] < g2.inter[0];
        });

        int n = intervals.size();
        for (int i = 0; i < n; i++) {
            int j = i+1;
            for (; j < n; j++) {
                if (gaps[j].inter[0] >= gaps[i].inter[1]){
                    gaps[i].right = gaps[j].idx;
                    break;
                }
                    
            }
        }
        res.resize(n, -1);
        for (int i = 0; i < n; i++)
            res[gaps[i].idx] = gaps[i].right;
            
        return res;
    }
};

// 傻一点方法