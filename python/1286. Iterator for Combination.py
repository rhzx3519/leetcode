characters = 'abc'
combinationLength = 2
arr = []

def dfs(idx, a):
    if idx==len(characters):
        if len(a)==combinationLength:
            arr.append(''.join(a))
        return

    a.append(characters[idx])
    dfs(idx + 1, a)
    a.pop()
    dfs(idx + 1, a)

dfs(0, [])

print arr