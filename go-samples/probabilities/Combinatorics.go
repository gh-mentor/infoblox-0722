package mathfuncs

/*
This package provides operations that can be used in counting and arranging objects in probabilistic mathematics.
*/

/*
Create a 'Combinatorics' struct that will provide bound functions for calculating combinatorial operations. Add a map 'memo' with a cap of 50 to store the results of previously calculated factorials.
*/
type Combinatorics struct {
	memo map[uint]uint
}

/*
Create a function 'NewCombinatorics()' that initializes a 'Combinatorics' object and returns it.	
*/
func NewCombinatorics() Combinatorics {
	return Combinatorics{memo: make(map[uint]uint, 50)}
}

/*
Create a bound function 'Factorial()' that calculates the factorial of a number 'n' and returns the result. It uses memoization to store the results of previously calculated factorials.
Arguments:
 - 'n': the number for which the factorial is to be calculated.
Return:
 - The factorial of 'n'.
Implementation:
- If the Combinatorics object does not exist in the 'memo' map, create it.
- If the factorial of 'n' is not in the 'memo' map, calculate it and store it in the 'memo' map.
- Return the factorial of 'n'.
*/
func (c Combinatorics) Factorial(n uint) uint {
	if c.memo == nil {
		c.memo = make(map[uint]uint, 50)
	}
	if _, ok := c.memo[n]; !ok {
		if n == 0 {
			c.memo[n] = 1
		} else {
			c.memo[n] = n * c.Factorial(n-1)
		}
	}
	return c.memo[n]
}


/*
Create a bound function 'Permutations()'.
Arguments:
- takes two unsigned integer parameters 'n' and 'r'.
- returns the number of permutations of 'n' items taken 'r' at a time.
Definition:
- Permutations are the number of ways to arrange 'r' items from a set of 'n' items.
Implementation:
	return Factorial(n) / Factorial(n - r);
*/
func (c Combinatorics) Permutations(n, r uint) uint {
	return c.Factorial(n) / c.Factorial(n-r)
}
