class Solution {
public:
	bool search(vector<int>& nums, int target) {
		if (nums.empty()) return false;
		int i;
		for (i = 0; i < nums.size() - 1; i++) {
			if (nums[i] > nums[i + 1])
				break;
		}
		int l, r, len = nums.size();
		if (target > nums[i]) return false;
		if (i + 1 < len && target < nums[i + 1]) return false;
		if (target <= nums[i] && target >= nums[0]) {
			l = 0; r = i;
		}
		else if (target >= nums[i + 1] && target <= nums[len - 1]) {
			l = i + 1; r = len - 1;
		}
	
		while (l <= r) {
			int mid = (l + r) >> 1;
			if (nums[mid] == target)
				return true;
			else if (nums[mid] > target)
				r = mid - 1;
			else
				l = mid + 1;
		}

		return false;
	}
};