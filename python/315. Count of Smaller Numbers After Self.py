class Node(object):
    def __init__(self, val):
        self.val = val
        self.left = self.right = None
        self.left_count = 0

class Solution(object):
    def countSmaller(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        n = len(nums)
        counts = [0]*n
        root = None

        def build(node, num, idx):
            if not node:
                return Node(num)

            if num<=node.val:
                node.left_count += 1
                node.left = build(node.left, num, idx)
            else:
                counts[idx] += node.left_count + 1
                node.right = build(node.right, num, idx)
            return node
            
        for i in range(n-1, -1, -1):
            root = build(root, nums[i], i)

        return counts

class Solution(object):
    def countSmaller(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        n = len(nums)
        res = [0]*n
        a = [i for i in range(n)]

        def merge_sort(arr):
            if len(arr) <= 1:
                return arr
            mid = len(arr)//2
            left = merge_sort(arr[:mid])
            right = merge_sort(arr[mid:])
            return merge(left, right)
        
        def merge(left, right):
            tmp = []
            i = j = 0
            while i<len(left) or j<len(right):
                if j==len(right) or i < len(left) and nums[left[i]] <= nums[right[j]]:
                    tmp.append(left[i])
                    res[left[i]] += j
                    i += 1
                else:
                    tmp.append(right[j])
                    j += 1
            return tmp

        merge_sort(a)
        return res

# 思路： 归并排序，前有序数组和后有序数组归并时，前有序数组出列时，后有序数组中有j个比它小的数
# time:O(NlogN), space:O(N)        

# 思路：从后往前遍历构造二叉查找树，数的每个节点都记录了左子树节点数量，每次寻找插入位置时，都会加上左子树节点数
# time：O(N*logN), space: O(N)         

if __name__ == '__main__':
    nums = [5,2,6,1]
    su = Solution()
    print su.countSmaller(nums)
