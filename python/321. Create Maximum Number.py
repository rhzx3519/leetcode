#!usr/bin/env python  
#-*- coding:utf-8 _*- 

class Solution(object):
    def maxNumber(self, nums1, nums2, k):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :type k: int
        :rtype: List[int]
        """
        def pick_max(nums, k):
            '''
                保持相对顺序，取nums中k个元素组成最大数字
            '''
            st = []
            drop = len(nums) - k
            for num in nums:
                while drop and st and st[-1] < num:
                    st.pop()
                    drop -= 1
                st.append(num)
            return st[:k]
        
        def merge(A, B):
            '''
            和并A B数组
            '''
            ans = []
            while A or B:
                bigger = A if A > B else B
                ans.append(bigger.pop(0))
            return ans
        
        return max(merge(pick_max(nums1, i), pick_max(nums2, k - i)) for i in range(k+1) if i <= len(nums1) and k - i <= len(nums2))

def pick_max(nums, k):
    '''
        保持相对顺序，取nums中k个元素组成最大数字
    '''
    st = []
    drop = len(nums) - k
    for num in nums:
        while drop and st and st[-1] < num:
            st.pop()
            drop -= 1
        st.append(num)
    return st[:k]


if __name__ == '__main__':
    print pick_max([3, 4, 6, 5], 3)