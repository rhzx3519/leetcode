
import collections
class Solution(object):

    def twoSum(self, arr, s):
        ans = []
        mem = collections.defaultdict(int)
        for num in arr:
            mem[num] += 1

        for num in arr:
            t = s - num
            if t not in mem:
                continue 
            c1 = mem[t]
            c2 = mem[num]
            if t != num:
                if c1>0 and c2>0:
                    mem[t] -= 1
                    mem[num] -= 1
            else:
                if c1 >= 2:
                    mem[t] -= 2

        # print mem
        for k, v in mem.iteritems():
            if v > 0:
                ans.extend([k]*v)

        print ans
        return ans

if __name__ == '__main__':
    # arr = [1, 2, 9, 9, 17, 17]
    arr = [1,2,9,17,17]
    # arr = [1,2,9,17,1]
    s = 18
    su = Solution()
    su.twoSum(arr, s)