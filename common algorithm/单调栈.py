#!usr/bin/env python  
#-*- coding:utf-8 _*-  

# 单调递减序列，寻找右边第一个比i大的数据的下标
def foo(a):
	n = len(a)
	res = [-1]*n
	st = []
	for i in range(n):
		while st and a[st[-1]] < a[i]:
			res[st.pop()] = i
		st.append(i)
	return res

def bar(a):
	n = len(a)
	res = [-1]*n
	st = []
	for i in range(n-1, -1, -1):
		while st and a[st[-1]] > a[i]:
			res[st.pop()] = i
		st.append(i)
	return res

if __name__ == '__main__':
	a = [1, 3, 6, 4, 5, 7, 1]
	print bar(a)