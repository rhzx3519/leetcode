#!usr/bin/env python  
#-*- coding:utf-8 _*- 

class Solution(object):
    def minimumJumps(self, forbidden, a, b, x):
        """
        :type forbidden: List[int]
        :type a: int
        :type b: int
        :type x: int
        :rtype: int
        """
        if x==0: return 0
        forbidden = set(forbidden)
        vis = [0]*6001
        vis[0] = 1
        que = [(0, 0)]
        step =0
        while que:

            sz = len(que)
            while sz:
                sz -= 1
                i, prev = que.pop(0)
                # print i
                
                if i==x: return step
      
                if i+a < 4001 and (i+a) not in forbidden and not vis[i+a]:
                    vis[i] = 1
                    que.append((i+a, 1))
                if prev!=-1 and i-b>0 and (i-b) not in forbidden and not vis[i-b]:
                    que.append((i-b, -1))

            step += 1
        return -1


if __name__ == '__main__':
    forbidden = [1,6,2,14,5,17,4]
    a = 16
    b = 9
    x = 7

    forbidden = [162,118,178,152,167,100,40,74,199,186,26,73,200,127,30,124,193,84,184,36,103,149,153,9,54,154,133,95,45,198,79,157,64,122,59,71,48,177,82,35,14,176,16,108,111,6,168,31,134,164,136,72,98]
    a = 29
    b = 98
    x = 80
    su = Solution()
    print su.minimumJumps(forbidden, a, b, x)