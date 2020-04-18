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
    vector<Interval> merge(vector<Interval>& intervals) {
        vector<Interval> res;
        vector<Interval>& a= intervals;
        int n = a.size();
        
        sort(a.begin(), a.end(), [&](const Interval& a, const Interval& b){
            return a.start < b.start;    
        });
        
        for (int i = 0; i < n - 1;)
        {
            if (a[i].end >= a[i+1].start)
            {
                a[i].start = min(a[i].start, a[i+1].start);
                a[i].end = max(a[i].end, a[i+1].end);
                a.erase(a.begin()+i+1);
                n--;
            }
            else
                i++;
        }
        return a;
    }
};