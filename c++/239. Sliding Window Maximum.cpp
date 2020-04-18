class Solution {
public:
<<<<<<< HEAD
	vector<int> maxSlidingWindow(vector<int>& nums, int k) {
		vector<int> res;
		deque<int> que;
		int i = 0, j = 0;
		int n = nums.size();
		for (int i = 0; i < n; i++) {
			while (!que.empty() && nums[que.back()] < nums[i])
				que.pop_back();
			que.push_back(i);

			if (i - que.front() == k)
				que.pop_front();

			if (i >= k - 1)
				res.push_back(nums[que.front()]);

		}
		return res;
	}
};

//利用一个 双端队列，在队列中存储元素在数组中的位置， 并且维持队列的严格递减, ，也就说维持队首元素是 **最大的 **，
//当遍历到一个新元素时, 如果队列里有比当前元素小的，就将其移除队列，以保证队列的递减。当队列元素位置之差大于 k，就将队首元素移除
=======
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        vector<int> res;
        int n = nums.size();
        if (n == 0) return res;
        
        k = min(n, k);
        map<int, int> window;
        for (int i = 0; i < k; i++)
            ++window[nums[i]];
        res.push_back(window.rbegin()->first);
        
        for (int i = k; i < n; i++) {
            --window[nums[i-k]];
            if (window[nums[i-k]] == 0)
                window.erase(nums[i-k]);
            ++window[nums[i]];
            res.push_back(window.rbegin()->first);
        }
        window.clear();
        
        return res;
    }
};
>>>>>>> 837bde34d06bb470ffcea088785f591f742870d7
