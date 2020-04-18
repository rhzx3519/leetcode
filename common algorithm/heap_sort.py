def heapify(a, idx, size):
    if idx >= size:
        return
    max_idx = idx
    l = 2*idx
    r = 2*idx + 1
    if l < size:
        if a[l] > a[idx]:
            max_idx = l
        
    if r < size:                                                                                                         
        if a[r] > a[idx]:
            max_idx = r
    
    if max_idx != idx:
        a[max_idx], a[idx] = a[idx], a[max_idx]
        heapify(a, max_idx, size)


def one_heap(a, size):
    for i in range(size-1, -1, -1):
        heapify(a, i, size)
    # print a

def hp(a):
    n = len(a)
    for i, _ in enumerate(a):
        one_heap(a, n-i)

if __name__ == '__main__':
    pass