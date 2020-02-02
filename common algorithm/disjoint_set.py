lookup = []

def find(x):
	if x!=lookup[x]:
		lookup[x] = find(lookup[x])
	return lookup[x]

def join(x, y):
	px = find(x)
	py = find(y)
	if px!=py:
		lookup[py] = px