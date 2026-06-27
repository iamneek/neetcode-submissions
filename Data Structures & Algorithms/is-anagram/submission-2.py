import collections

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        counter1 = collections.Counter(s)
        counter2 = collections.Counter(t)
        return counter1 == counter2
        # count1 = {}
        # count2 = {}

        # for e in s:
        #     if e in count1:
        #         count1[e] += 1
        #         continue
        #     count1[e] = 1

        # for e in t:
        #     if e in count2:
        #         count2[e] += 1
        #         continue
        #     count2[e] = 1

        # return count1==count2
        
        