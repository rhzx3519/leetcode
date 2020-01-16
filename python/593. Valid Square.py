#!/usr/bin/env python
# -*- coding: utf-8 -*-

class Solution(object):
    def validSquare(self, p1, p2, p3, p4):
        """
        :type p1: List[int]
        :type p2: List[int]
        :type p3: List[int]
        :type p4: List[int]
        :rtype: bool
        """
        a = []
        points = [p1, p2, p3, p4]
        for i in xrange(len(points)):
        	for j in xrange(i+1, len(points)):
        		d = self.dis(points[i], points[j])
        		if d==0:
        			return False
        		if len(a)<2 and d not in a:
        			a.append(d)
        		elif d not in a:
        			return False
        return True

    def dis(self, p1, p2):
    	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])


if __name__ == '__main__':
	p1 = [1,0]
	p2 = [0,1]
	p3 = [0,-1]
	p4 = [-1,0]

	s = Solution()
	print s.validSquare(p1, p2, p3, p4)