"""
Create a container class 'Stack' that has the following methods:
- 'push' that adds an element to the stack
- 'pop' that removes the top element from the stack
- 'peek' that returns the top element from the stack
- 'is_empty' that returns True if the stack is empty, False otherwise
Details:
- The stack should be implemented using a list with a capacity value provided as an argument to the constructor.
- Fild names should use the prefix 'm_' to indicate that they are private.
Error cases:
- If the stack is full and 'push' is called, raise a ValueError with the message 'Stack is full'.
"""
class Stack:
    def __init__(self, capacity):
        self.m_capacity = capacity
        self.m_stack = []

    def push(self, element):
        if len(self.m_stack) == self.m_capacity:
            raise ValueError('Stack is full')
        self.m_stack.append(element)

    def pop(self):
        return self.m_stack.pop()

    def peek(self):
        return self.m_stack[-1]

    def is_empty(self):
        return not self.m_stack