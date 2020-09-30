// YOU SHOULD NEVER HAVE FILES NAMED THINGS Like "utils", "helpers", etc (see:
// https://dave.cheney.net/2019/01/08/avoid-package-names-like-base-util-or-common). This was done simply for the demo
package main

func getPlaceValue(n int) string {
	if n > 999999999 {
		return "billions"
	} else if n > 99999999 {
		return "hundred millions"
	} else if n > 9999999 {
		return "ten millions"
	} else if n > 999999 {
		return "millions"
	} else if n > 99999 {
		return "hundred thousands"
	} else if n > 9999 {
		return "ten thousands"
	} else if n > 999 {
		return "thousands"
	} else if n > 99 {
		return "hundreds"
	} else if n > 9 {
		return "tens"
	} else if n >= 0 {
		return "ones"
	}

	return "oops"
}

func reverseString(s string) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
