class Solution(object):
    def jump(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        i = pos = 0
        n = len(nums)
        step = 0
        if n<=1: return 0
        while True:
            new_pos = pos
            while i <= pos:
                new_pos = max(new_pos, i + nums[i])
                i += 1
            if pos==new_pos: break
            pos = new_pos
            print i, pos
            step += 1
            if pos >=n-1: break
            
        return step

class Solution(object):
    def jump(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        if n <= 1:
            return 0
        max_dis = step = reach = 0
        for i in range(n):
            max_dis = max(max_dis, i + nums[i])
            if max_dis >= n-1:
                step += 1
                break
            if i == reach: # 当i遍历到最远距离的时候step+1
                step += 1
                reach  = max_dis
            
        return step

# Dijkstra算法的变形——将第一步第二步...第N步可以到达的节点依
# 次放入map_arrive中。或者说是对图执行广度优先遍历（图的广度
# 优先遍历可以找到某节点到所有节点最少步数，图的深度优先遍历可以
# 找到无环图的拓扑结构）
# 将第一步到达的节点放到map_arrive中
# 如果目的节点不在map_arrive中，试图从map_arrive再往前走一步
# 将可以到达的节点再次放入到map_arrive中。一直如此执行下去。                

if __name__ == '__main__':
    nums = [2,3,1,1,4]
    su = Solution()
    su.jump(nums)        
