"""
Create a function 'low_high' that takes a list of numbers as an argument and returns a tuple with the lowest and highest numbers in the list.
Arguments:
- numbers: a list of numbers
Returns:
- a tuple with the lowest and highest numbers in the list
Error cases:
- If the list is empty, return a tuple with two None values.

"""
def low_high(numbers):
    if not numbers:
        return None, None
    return min(numbers), max(numbers)