

/*
 * @lc app=leetcode id=7 lang=rust
 *
 * [7] Reverse Integer
 */
impl Solution {
    // to be honest, I cheated and looked up the solution
    // which consists on "popping and pushing digits" in a
    // brilliant way using modulo and shifting the integers
    // around in decimal base:
    pub fn reverse(mut x: i32) -> i32 {
        let mut rev = 0; // initialize empty reversed int
        while x != 0 { // if x is 0 we're finished
            let pop = x % 10; // pop last digit with modulo 10, genius!
            if int32_overflows(rev, pop) {
                return 0; // return 0 if pushing the value would overflow
            }
            x /= 10; // divide by 10 to right shift in decimal base
            rev = rev * 10 + pop; // multiply by 10 to left shift in decimal base
        }
        rev // return once done
    }
}

fn int32_overflows(rev: i32, pop: i32) -> bool {
    // this video helped me understand how this overflow check works:
    // https://www.youtube.com/watch?v=ThdOYKsFSK0

    // basically, this function will check if the next digit that
    // will be "popped" into the current reversed int will overflow
    // its 32bit signed integer memory

    // this is achieved by doing the following check:
    // - if rev is already larger than the right-shifted max i32 value,
    //   it will overflow no matter the digit being pushed
    // - if rev is equal to the right-shifted max i32 value,
    //   it will overflow if the digit being pushed is larger than
    //   the last digit of the max i32 value

    // same logic (but inverted) applies to the minimum

    // pretty cool! :D
    let max: i32 = i32::max_value() / 10; // right shift in decimal base
    if rev > max || (rev == max && pop > 7) {
        return true;
    }
    let min: i32 = i32::min_value() / 10; // right shift in decimal base
    if rev < min || (rev == min && pop < -8) {
        return true;
    }
    false
}

