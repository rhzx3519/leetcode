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
        int i =0, size = intervals.size();
        
        sort(intervals.begin(), intervals.end(), comp);
        while (i < size - 1) {
            if (intervals[i].start <= intervals[i+1].start && intervals[i].end >= intervals[i+1].end) {
                intervals.erase(intervals.begin() + i + 1);
                size--;
            }
            else if (intervals[i].end >= intervals[i+1].start && intervals[i].end < intervals[i+1].end) {
                intervals[i].end = intervals[i+1].end;
                intervals.erase(intervals.begin() + i + 1);
                size--;
            }else
                i++;
        }
        return intervals;
    }
    
    static bool comp(const Interval &ele1, const Interval &ele2) {
        return ele1.start < ele2.start;
    }
};