class Solution {
public:
    int leastInterval(vector<char>& tasks, int n) {
        map<char, int> mp;
        int count = 0;
        for (char &task : tasks) {
            mp[task]++;
            count = max(count, mp[task]);
        }
        int res = (count-1)*(n+1);
        for (auto it = mp.begin(); it != mp.end(); it++) {
            if (it->second == count)
                res++;
        }
        return max(int(tasks.size()), res);
    }
};

/**
 * 等于数量最多的任务乘以时间间隔 公式就是 (count-1) * (n+1) + 数量等于count的任务数目
*/