"""
Create a function 'top_n' that takes a list of numbers and an integer 'n' as arguments and returns the 'n' largest numbers in the list as a tuple.
Arguments:
- numbers: a list of numbers
- n: an integer representing the number of largest numbers to return
Returns:
- a tuple with the 'n' largest numbers in the list
Error cases:
if 'n' is greater than the length of the list, return the list throw a ValueError with the message 'n is greater than the length of the list'.
"""
def top_n(numbers, n):
    if n > len(numbers):
        raise ValueError('n is greater than the length of the list')
    return tuple(sorted(numbers, reverse=True)[:n])