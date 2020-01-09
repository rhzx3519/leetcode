class Solution {
public:
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
