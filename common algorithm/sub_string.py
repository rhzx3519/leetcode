#!usr/bin/env python  
#-*- coding:utf-8 _*-  

# 给定一个字符串，求它的所有字串

class Solution(object):

	def subStr(self, s):

		def dfs(idx):
			if idx==len(s):
				return

			a = []
			for i in range(idx, len(s)):
				a.append(s[i])
				res.append(a[:])
			dfs(idx+1)

		res = []
		dfs(0)
		print res

if __name__ == '__main__':
	su = Solution()
	su.subStr('123')

