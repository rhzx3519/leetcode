class Solution(object):
    
	def gcd(self, a, b):
		if b:
			return self.gcd(b, a%b)
		else:
			return a

	def lcm(self, a, b):
		return a*b/self.gcd(a, b)


	def fractionAddition(self, expression):
		"""
		:type expression: str
		:rtype: str
		"""
		a = []
		b = []
		i = 0
		sign = 1
		if expression[0]=='-':
			sign = -1
			i += 1
		n = len(expression)
		while i < n:
			t = 0
			while i<n and (expression[i]>='0' and expression[i]<='9'):
				t = t*10 + int(expression[i])
				i += 1
			a.append(sign*t) # get numerator
			i += 1 		# skip /

			t = 0
			while i<n and (expression[i]>='0' and expression[i]<='9'):
				t = t*10 + int(expression[i])
				i += 1
			b.append(t) # get denominator
			if i < n:
				sign = 1 if expression[i]=='+' else -1
			i += 1

		n = len(a)
		for i in xrange(n-1):
			m = a[i]*b[i+1] + a[i+1]*b[i]
			n = b[i] * b[i+1]
			t = self.gcd(m, n)
			a[i+1] = m/t
			b[i+1] = n/t

		return '{}/{}'.format(a[-1], b[-1])