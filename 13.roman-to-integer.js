/*
 * @lc app=leetcode id=13 lang=javascript
 *
 * [13] Roman to Integer
 */

// @lc code=start
/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function (s) {
  let lastC = false
  let total = 0
  ;[...s].forEach(c => {
    switch (c) {
      case 'I':
        total += 1
        break
      case 'V':
        total += lastC === 'I' ? 3 : 5
        break
      case 'X':
        total += lastC === 'I' ? 8 : 10
        break
      case 'L':
        total += lastC === 'X' ? 30 : 50
        break
      case 'C':
        total += lastC === 'X' ? 80 : 100
        break
      case 'D':
        total += lastC === 'C' ? 300 : 500
        break
      case 'M':
        total += lastC === 'C' ? 800 : 1000
        break
    }
    lastC = c
  })
  return total
}
// @lc code=end
