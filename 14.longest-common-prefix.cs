/*
 * @lc app=leetcode id=14 lang=csharp
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
public class Solution
{
    public string LongestCommonPrefix(string[] strs)
    {
        if (strs.Length == 0) return "";
        StringBuilder prefixBuilder = new StringBuilder();
        int i = 0;
        bool done = false;
        while (!done)
        {
            char? c = null;
            foreach (var str in strs)
            {
                if (str.Length <= i) { done = true; break; }
                if (!c.HasValue) c = str[i];
                else if (c != str[i]) { done = true; break; }
            }
            if (!done) prefixBuilder.Append(c);
            i++;
        }
        return prefixBuilder.ToString();
    }
}
// @lc code=end

