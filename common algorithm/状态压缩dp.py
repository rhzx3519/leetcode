#!usr/bin/env python  
#-*- coding:utf-8 _*-

def foo(arr):
    n = len(arr)
    dp = [0]*(1<<n)
    for i in range(1<<n):
        for j in range(n):
            if i&(1<<j):
                dp[i] += arr[j]
    print dp
    return dp


if __name__ == '__main__':
    arr = [1, 2]
    foo(arr)

# 给定一个整型数组arr, 随机选取arr中的元素相加(每个元素可以选择1次或者0次)，返回和的所有可能情况, 
# 例：arr = [1, 2], 返回[0, 1, 2, 3]