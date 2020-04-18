class MyCalendar {
public:
    MyCalendar() {
        
    }
    
    bool book(int start, int end) {
        for (auto &ca : calendar) {
            if (start < ca[1] && ca[0] < end)
                return false;
        }
        calendar.push_back({start, end});
        return true;
    }

    vector<vector<int>> calendar;
};

/**
 * Your MyCalendar object will be instantiated and called as such:
 * MyCalendar* obj = new MyCalendar();
 * bool param_1 = obj->book(start,end);
 */