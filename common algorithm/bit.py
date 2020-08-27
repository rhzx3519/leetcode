#!usr/bin/env python  
#-*- coding:utf-8 _*-  

# 清除最低位的1
def foo(n):
	n &= (n -1)  

def countBinaryOne(n):
	c = 0
	while n:
		n &= (n-1)
		c += 1
	return c

if __name__ == '__main__':
	print countBinaryOne(0b010010)