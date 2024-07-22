package wordgame

/*
Create a function 'isPalindrome that takes a string 'word' and returns a boolean value indicating whether the string is a word that is written the same forwards and backwards.
Arguments:
	- 'word': a string that needs to be checked for palindrome.
Return:
	- A boolean value indicating whether the string is a palindrome.
Example:
	"racecar" is a palindrome
	"bob" is a palindrome
	"noon" is a palindrome
	"hello" is not a palindrome
	"world" is not a palindrome
*/
func IsPalindrome(word string) bool{

	for i := range word{
		l := len(word)
		if word[i] != word[l - i - 1]{
			return false
		}
	}
	return true
}