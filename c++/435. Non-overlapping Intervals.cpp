/*
	根据结束时间升序排序，删除重叠的区间
*/
/**
 * Definition for an interval.
 * struct Interval {
 *     int start;
 *     int end;
 *     Interval() : start(0), end(0) {}
 *     Interval(int s, int e) : start(s), end(e) {}
 * };
 */
class Solution {
public:
    static bool cmp(const Interval& i1, const Interval& i2)
    {   
        if (i1.end != i2.end)
            return i1.end < i2.end;
        return i1.start < i2.start;
    }
    
    int eraseOverlapIntervals(vector<Interval>& intervals) {
        vector<Interval>& a = intervals;
        if (a.empty()) return 0;
        
        sort(a.begin(), a.end(), cmp);
        int n = a.size();
        for (int i = 0; i < a.size()-1;)
        {
            if (a[i].end > a[i+1].start)
                a.erase(a.begin() + i + 1);
            else
                i++;
        }
        
        return n - a.size();
    }
};