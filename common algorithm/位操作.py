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


//集合A、B，元素c --> int A,B   c = 0 ~ 31
//A中插入c
A |= (1<<c)
//A中去除c
A &= ~(1<<c)
A ^= (1<<c)
//A B 合并
A | B
//判断B是不是A的子集
return (A&B) == B
//判断c在不在A里
return A & (1<<c)
//lowbit
return x & (-x);
//枚举A的全部子集
for(int i = A; i; i = (i-1) & A)
{
    //do something
}