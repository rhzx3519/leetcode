#!usr/bin/env python  
#-*- coding:utf-8 _*- 

class Solution(object):
    def searchInsert(self, nums, target):
        left,right = 0,len(nums)-1
        while left<=right:
            mid = (left+right)//2
            if target == nums[mid]:
                while mid<len(nums)-1:
                    if nums[mid]!=nums[mid+1]:
                        break
                    mid+=1
                return mid+1
            elif target < nums[mid]:
                right = mid - 1 
            else:
                left = mid + 1 
        
        return left

    def biSearch(self, que, val):
        l, r = 0, len(que)-1
        while l <= r:
            mid = l + (r-l)//2
            key = que[mid]
            if key==val:
                if mid==len(que)-1 or key!=que[mid+1]:
                    return mid + 1
                else:
                    l = mid + 1
            elif key > val:
                r = mid -1
            else:
                l = mid + 1
        return l

so=Solution()
s1=[1,3,5,6]
s2=[1,3,5,5,5,6]
s3 = [2]
target=5
print(so.searchInsert(s1,target))
print(so.searchInsert(s2,target))
print(so.biSearch(s3,target))
