#
# @lc app=leetcode id=1 lang=python3
#
# [1] Two Sum
#

# [2, 8, 3, 56, 12, 45]\n68
class Solution:
    def twoSum(self, nums, target):
        """
        My initial solution had a loop inside a loop -I believe that
        made its running time O(n^2)- to compare every number in the
        list with each one of the remaining ones (with higher index).

        I checked the top solution submitted by Google and I found their
        brilliant solution which is to use a hash map to store each of the
        numbers that would be needed for each index in order to sum up to
        the target, allowing a nice and linear O(n) running time because you
        only need to loop once to see if the number at the current index is
        a key in the hash map (and, therefore, it is a solution alongside the
        index stored as a value).

        I tested it and it seemed to be broken because of the + 1, so I
        removed it and it runs just fine now. Cool one.

        EDIT: using enumerate as seen in a comment for the top solution.
        Looks nicer and it's probably more efficient bc of less list lookups.

        EDIT 2: fixed return type - List -> Tuple

        EDIT 3: it might actually be faster without the typing lib or just
        types altogether, let's see

        EDIT 4: what the hell runtime went up, let me test this one sec
        """
        map = {}
        for i, num in enumerate(nums):
            if num not in map:
                map[target - num] = i
            else:
                return map[num], i
        return -i, -i

