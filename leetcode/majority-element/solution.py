from collections import Counter


class Solution:
    def majorityElement(self, num):
        return Counter(num).most_common(1)[0][0]
