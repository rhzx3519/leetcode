#!usr/bin/env python  
#-*- coding:utf-8 _*-  

# 01背包矩阵遍历剩余承重量时倒序，完全背包遍历时正序。


matrix = [[0, 0, 0],
		[0, 0, 0],
		[0, 0, 0],]
n = len(matrix)

# 斜向遍历
print '---------------------------'
for l in range(1, n):
	for i in range(n-l):
		j = i+l
		print (i, j)

print '---------------------------'
# i从n-1往0遍历，j从0往n遍历，斜三角形上方
for i in range(n-1, -1, -1):
    for j in range(i+1, n):		
    	print (i, j)


print '---------------------------'
# i从n-1往0遍历，j从0往n遍历，斜三角形下方
for i in range(n-1, -1, -1):
    for j in range(n-i, n):		
    	print (i, j)    	
