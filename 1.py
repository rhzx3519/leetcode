# def gcd(a, b):
#     return a if b==0 else gcd(b, a%b)

# print gcd(2, 6)    
# print 5%1

# n = 3

# for i in range(n-1, -1, -1):
#     for j in range(1, n-i):
#     	print i, j



def dullStack(a):
	n = len(a)
	idx = [0]*(n+1)
	idx[0] = -1
	st = []
	p = 0
	for i in range(n):
		while idx[p] != -1 and a[idx[p]] >= a[i]:
			p -= 1
		idx[p+1] = i
	return idx

# 寻找
def foo(a):
	n = len(a)
	res = [-1]*n
	st = []
	for i in range(n):
		while st and a[st[-1]] < a[i]:
			res[st.pop()] = i
		st.append(i)
	return res
		


if __name__ == '__main__':
	a = [1, 3, 6, 4, 5, 7, 1]
	# a = [1, 4, 3, 2]
	print foo(a)