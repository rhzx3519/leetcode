class MyCalendarTwo {
public:
    MyCalendarTwo() {
        
    }
    
    bool book(int start, int end) {
        for (auto &c : overlaps) {
            if (start  < c[1] && end > c[0]) {
                return false;
            }
        }

        for (auto &c : calendar) {
            if (start  < c[1] && end > c[0]) {
                overlaps.push_back({max(start, c[0]), min(end, c[1])});
            }
        }        

        calendar.push_back({start, end});
        return true;
    }

    vector<vector<int>> calendar;
    vector<vector<int>> overlaps;  // overlap记录重叠的区间
};

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * MyCalendarTwo* obj = new MyCalendarTwo();
 * bool param_1 = obj->book(start,end);
 */