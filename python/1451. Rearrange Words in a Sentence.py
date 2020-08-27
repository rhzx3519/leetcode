class Solution(object):
    def arrangeWords(self, text):
        """
        :type text: str
        :rtype: str
        """
        return ' '.join(sorted(text.lower().split(' '), key=len)).capitalize()
