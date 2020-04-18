def gcd(a, b):
    return a if b==0 else gcd(b, a%b)

print gcd(2, 6)    
print 5%1

n = 3

for i in range(n-1, -1, -1):
    for j in range(1, n-i):
    	print i, j