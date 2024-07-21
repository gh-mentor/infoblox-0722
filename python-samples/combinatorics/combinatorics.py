"""
This file contains functions for combinatorics.
"""

"""
Create a function 'factorial' that returns the factorial of a number.
The factorial of a number is the product of all positive integers less than or equal to the number.
Arguments:
n -- an integer
Return:
an integer
Examples:
factorial(-1) => None
factorial(0) => 1
factorial(1) => 1
factorial(2) => 2
factorial(3) => 6
"""
def factorial(n):
    if n == 0 or n == 1:
        return 1
    return n * factorial(n - 1)

"""
Create a function 'permutations' that returns the number of permutations of k elements from a set of n elements.
The number of permutations of k elements from a set of n elements is the number of ways to arrange k elements from a set of n elements.
Arguments:
n -- an integer
k -- an integer
Return:
an integer
Examples:
permutations(0, 0) => 1
permutations(1, 0) => 1
permutations(1, 1) => 1
permutations(2, 1) => 2
permutations(2, 2) => 2
permutations(3, 1) => 3
permutations(3, 2) => 6
"""
def permutations(n, k):
    return factorial(n) // factorial(n - k)







