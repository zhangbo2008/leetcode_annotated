#   Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:
#
 #   '.' Matches any single character.​​​​
  #  '*' Matches zero or more of the preceding element.
   # The matching should cover the entire input string (not partial).
#写一个匹配.和*的正则函数
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        def helper(s, i, p, j): #动态规划,前i个s字符和p的前j个字符.返回是否匹配.
            if j == -1:
                return i == -1
            if i == -1:
                if p[j] != '*':
                    return False
                return helper(s, i, p, j-2)  # -1表示空字符串.
            if p[j] == '*':
                if p[j-1] == '.' or p[j-1] == s[i]:
                    if helper(s, i-1, p, j):
                        return True
                return helper(s, i, p, j-2)# *当空白用,因为j-1匹配不上,*当他乘以0
            if p[j] == '.' or p[j] == s[i]:
                return helper(s, i-1, p, j-1) # 就是.的使用.等于任意字符,所以等价于i-1跟j-1
            return False

        return helper(s, len(s)-1, p, len(p)-1)