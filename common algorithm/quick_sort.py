#!usr/bin/env python  
#-*- coding:utf-8 _*-  


# 一趟分割，可用于链表的快速排序
def partition(a, l, r):
    i = j = l
    key = a[l]
    while i<=r:
        if a[i] < key:
            j += 1
            a[i], a[j] = a[j], a[i]
        i += 1
        print a, j
    a[l], a[j] = a[j], a[l]
    print a
    return j


# def partition(a, l, h):     
#     t = a[l]
#     while l<h:
#         while l<h and a[h] >= t:
#             h -= 1
#         a[l] = a[h]
#         while l<h and a[l] <= t:
#             l += 1
#         a[h] = a[l]
#     a[l] = t
#     return l

# recusion
def qsort(a, l, h):
    if l < h:
        t = partition(a, l, h)
        qsort(a, l, t-1)
        qsort(a, t+1, h)

# iteration
def iter_qsort(a):
    st = []
    l, h = 0, len(a)-1
    st.append(l)
    st.append(h)
    while len(st)>=2:
        h = st.pop(-1)
        l = st.pop(-1)
        t = partition(a, l, h)
        if l < t:
            st.append(l)
            st.append(t-1)
        if t < h:
            st.append(t+1)
            st.append(h)


if __name__ == '__main__':
    a = [3,2,1,3,3,2,5,2,3,1,6,]
    # qsort(a, 0, len(a)-1)
    partition(a, 0, len(a)-1)
    # print a