/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
func isPalindrome(x int) bool {
    if x < 0 { // negative integers cannot be a palindrome
		return false
	}
	// my solution uses the algorithm I learned with problem
	// #7 in which an integer is reversed by "popping and pushing
	// digits", which is achieved by using modulo and sum of 10
	// and by shifting in decimal base

	// then, I compare the reversed decimal with the copy I made
	// of the original one (because the reversing algorithm mutates
	// the original one)
	xCopy := x
	rev := 0
	for x > 0 {
		pop := x % 10
		x /= 10
		rev = rev * 10 + pop
	}
	if xCopy == rev {
		return true
	}
	return false
}

