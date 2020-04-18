class Codec:
    def __init__(self):
        self.cache = {}
        self.rev_cache = {}
        self.idx = 0
        
    def base64(self, num):
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

    def encode(self, longUrl):
        """Encodes a URL to a shortened URL.
        
        :type longUrl: str
        :rtype: str
        """
        if not self.cache.has_key(longUrl):
            self.cache[longUrl] = self.base64(self.idx)
            self.rev_cache[self.cache[longUrl]] = longUrl
            self.idx += 1
        return self.cache[longUrl]


    def decode(self, shortUrl):
        """Decodes a shortened URL to its original URL.
        
        :type shortUrl: str
        :rtype: str
        """
        if shortUrl in self.rev_cache:
            return self.rev_cache[shortUrl]
        return ''
        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.decode(codec.encode(url))