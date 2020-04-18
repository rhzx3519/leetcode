#!usr/bin/env python  
#-*- coding:utf-8 _*- 

# KMP算法是一种字符串匹配算法，可以在 O(n+m) 的时间复杂度内实现两个字符串的匹配。
# 如何更好的理解和掌握 KMP 算法? - 阮行止的回答 - 知乎
# https://www.zhihu.com/question/21923021/answer/1032665486

# P 的 next 数组定义为：next[i] 表示 P[0] ~ P[i] 这一个子串，使得 前k个字符恰等于后k个字符 的最大的k.

class KMPSerach(object):
	"""docstring for KMPSerach"""
	def __init__(self):
		super(KMPSerach, self).__init__()

	def getNxt(self, p, x):
		'''获取 next[x]
		'''
		for i in range(x, 0, -1): 		# 从 x 到 1 枚举
			if p[0:i] == p[x-i+1:x+1]:
				return i 				# 如果前缀等于后缀则返回
		return 0

	def buildNxt(self, p):
		nxt = [0]	# next[0] 必然为0
		x = 1			# 从next[1]开始求
		now = 0	
		while x < len(p):
			if p[now]==p[x]: # 如果p[now]==p[x]，向右扩展一位
				now += 1
				x += 1
				nxt.append(now)
			elif now:
				now = nxt[now-1]	# 缩小now，改成nxt[now-1]
			else:
				nxt.append(0)		# now已经为0，无法再缩，故next[x] = 0
				x += 1
		return nxt

	def search(self, s, p):
		nxt = [self.getNxt(p, x) for x in range(len(p))]
		# nxt = self.buildNxt(p)
		print nxt
		tar = 0		# tar: 主串中将要匹配的位置
		pos = 0		# pos: 模式串中将要匹配的位置
		while tar < len(s):
			if s[tar]==p[pos]:	# 若两个字符相等，则tar,pos各前进一步
				tar += 1
				pos += 1
			elif pos:			# 失配了。若pos!=0，则依据nxt数组移动标尺
				pos = nxt[pos-1]
			else:
				tar += 1		# p[0]失配，直接移动tar

			if pos==len(p):			# 匹配成功
				print (tar - pos)	# 输出主串上的起始位置
				pos = nxt[pos-1]	# 移动标尺

if __name__ == '__main__':
	s = 'abcdabaabacewda'
	p = 'abaabac'
	kmpSearch = KMPSerach()
	kmpSearch.search(s, p)
