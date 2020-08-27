class Solution(object):
    def permuteUnique(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums.sort()
        res = []
        n = len(nums)

        def dfs(idx, end, nums):
            if idx==end:
                res.append(nums)
                return
            for i in range(idx, end+1):
                if i!=idx and nums[i]==nums[idx]: continue
                nums[i], nums[idx] = nums[idx], nums[i]
                dfs(idx+1, end, nums[:])
                
        
        dfs(0, n-1, nums)
        return res


class Solution(object):
    def permuteUnique(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = []
        if not nums:
            return res
        
        def dfs(nums, tmp):
            if not nums:
                res.append(tmp)
                return
            
            mem = set()
            for i in range(len(nums)):
                if nums[i] not in mem:
                    mem.add(nums[i])
                    dfs(nums[:i] + nums[i+1:], tmp + [nums[i]])
        dfs(nums, [])
        return res        

if __name__ == '__main__':
    # nums = [[0,0,0,1,9],[0,0,0,9,1],[0,0,1,0,9],[0,0,1,9,0],[0,0,9,1,0],[0,0,9,0,1],[0,1,0,0,9],[0,1,0,9,0],[0,1,9,0,0],[0,9,0,1,0],[0,9,0,0,1],[0,9,1,0,0],[0,9,0,1,0],[0,9,0,0,1],[1,0,0,0,9],[1,0,0,9,0],[1,0,9,0,0],[1,9,0,0,0],[9,0,0,1,0],[9,0,0,0,1],[9,0,1,0,0],[9,0,0,1,0],[9,0,0,0,1],[9,1,0,0,0],[9,0,0,1,0],[9,0,0,0,1],[9,0,1,0,0],[9,0,0,1,0],[9,0,0,0,1]]
    nums = [1,2,3]
    su = Solution()
    print su.permuteUnique(nums)