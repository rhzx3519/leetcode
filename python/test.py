#!usr/bin/env python  
#-*- coding:utf-8 _*-  
from operator import itemgetter

def base64(num):
    a = []
    while num:
        a.append(num % 62)
        num /= 62
    a.reverse()
    b = []
    for i in range(len(a)):
        if a[i]<26:
            ch = chr(ord('a') + a[i])
        elif a[i]<52:
            ch = chr(ord('A') + a[i]-26)
        else:
            ch = chr(ord('0') + a[i]-52)
        b.append(ch)
    return ''.join(b)

print base64(62)