#!usr/bin/env python  
#-*- coding:utf-8 _*-  


def merge(a, l, r):
    mid = l + (r-l)/2
    b = [0]*(r-l+1)
    i = l
    j = mid+1
    k = 0
    while i <= mid and j <= r:
        if a[i] < a[j]:
            b[k] = a[i]
            i += 1
        else:
            b[k] = a[j]
            j += 1
        k += 1

    while i <= mid:
        b[k] = a[i]
        i += 1
        k += 1

    while j <= r:
        b[k] = a[j]
        j += 1
        k += 1

    i = l
    while i <= r:
        a[i] = b[i-l]
        i += 1

def sort(a, l , r):
    if l < r:
        mid = l + (r-l)/2
        sort(a, l, mid)
        sort(a, mid+1, r)
        merge(a, l, r)