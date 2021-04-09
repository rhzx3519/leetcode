class Solution(object):
    def reversePairs(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        self.aux = [0]*len(nums)

        def merge_sort(nums, l, r):
            if l < r:
                mid = l + (r - l)/2
                merge_sort(nums, l, mid)
                merge_sort(nums, mid+1, r)
                self.ans += find(nums, l, r)    # 在排好序的两个子数组间，求翻转对数量
                merge(nums, l, mid, r)
                
        
        def merge(nums, l, mid, r):
            i, j = l, mid+1
            aux = self.aux
            for k in range(l, r+1):
                aux[k] = nums[k]
            for k in range(l, r+1):
                if i > mid:
                    nums[k] = aux[j]
                    j += 1
                elif j > r:
                    nums[k] = aux[i]
                    i += 1
                elif aux[i] <= aux[j]:
                    nums[k] = aux[i]
                    i += 1
                else:
                    nums[k] = aux[j]
                    j += 1

        def find(nums, l, r):
            ans = 0
            mid = l + (r - l)/2
            j = mid + 1
            for i in range(l, mid+1):
                while j <= r and nums[i] > 2*nums[j]:
                    ans += (mid - i + 1)
                    j += 1
            return ans

        self.ans = 0
        merge_sort(nums, 0, len(nums)-1)
        # print nums
        
        return self.ans
        
            