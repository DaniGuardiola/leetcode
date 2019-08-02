/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
/*
My solution uses the algorithm I learned with problem
#7 in which an integer is reversed by "popping and pushing
digits", which is achieved by using modulo and sum of 10
and by shifting in decimal base.

Then, I compare the reversed decimal with the copy I made
of the original one (because the reversing algorithm mutates
the original one).

EDIT: the official solution points to something I though about
but that I didn't end up implementing, which is reversing just
half of the digits, which makes the algorithm faster and avoids
potential overflow problems.

I achieved this by doing what the solution states:
"Since we divided the number by 10, and multiplied the reversed
number by 10, when the original number is less than the reversed
number, it means we've processed half of the number digits."

Additionally, they also point to a special case in addition to
negative numbers in which the number will never be a palindrome:
numbers ending with zero, because an integer will never start with
zero (with the exception of zero itself). I implemented this as well.
*/
import "fmt"
func isPalindrome(x int) bool {
	// negative or integers ending in zero cannot be palindromes
    if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}
	rev := 0
	for x > rev { // reverse half of the integer
		pop := x % 10
		x /= 10
		rev = rev * 10 + pop
	}
	// in case the original integer was odd, we can take care
	// of the middle digit (which does not matter in a palindrome)
	// by also comparing the original integer to the right-shifted
	// reversed half of the integer
	
	if x == rev || x == rev / 10 {
		return true
	}
	return false
}

