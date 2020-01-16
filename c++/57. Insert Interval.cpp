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
    vector<Interval> insert(vector<Interval>& intervals, Interval newInterval) {
        int i = 0, lap = 0;
        while (i < intervals.size()) {
            if (newInterval.end < intervals[i].start) break;
            else if (newInterval.start > intervals[i].end) {}
            else {
                newInterval.start = min(newInterval.start, intervals[i].start);
                newInterval.end = max(newInterval.end, intervals[i].end);
                lap++;
            }
            i++;
        }

        if (lap > 0) 
            intervals.erase(intervals.begin() + i - lap, intervals.begin() + i);
        intervals.insert(intervals.begin() + i - lap, newInterval);
        return intervals;
    }
};