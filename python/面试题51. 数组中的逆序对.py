class Solution(object):
    def reversePairs(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return self.sort(nums, 0, len(nums)-1)

        
    def sort(self, nums, l, r):
        res = 0
        if l>=r: return 0
        mid = (l+r)>>1
        res += self.sort(nums, l, mid)
        res += self.sort(nums, mid+1, r)
        res += self.merge(nums, l, r)
        return res

    def merge(self, a, l, r):
        res = 0
        mid = l + (r-l)/2
        b = [0]*(r-l+1)
        i = l
        j = mid+1
        k = 0
        while k<=r-l+1:
            if i<=mid and j<=r:
                if a[i] <= a[j]:
                    b[k] = a[i]
                    i += 1
                else:
                    res += mid -i + 1   # 数组间的逆序对数目
                    b[k] = a[j]
                    j += 1
            elif i<=mid:
                b[k] = a[i]
                i += 1
            elif j<=r:
                b[k] = a[j]
                j += 1

            k += 1

        a[l:r+1] = b
        return res
