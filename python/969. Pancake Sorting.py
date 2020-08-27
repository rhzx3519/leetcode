class Solution(object):
    def pancakeSort(self, A):
        """
        :type A: List[int]
        :rtype: List[int]
        """
        ans = []
        n = len(A)
        for i in range(n, 0, -1):
            idx = A.index(i)
            if idx == i-1:
                continue
            ans.append(idx + 1)
            A = A[:idx+1][::-1] + A[idx+1:] # 反转[0, idx]之间的元素，将A[idx]这个最大值放到数组头
            A = A[:i][::-1] + A[i:] # 反转[0:i-1], 将A[0]放到A[i-1], 最大值在数组尾
            ans.append(i)
            print A
        return ans