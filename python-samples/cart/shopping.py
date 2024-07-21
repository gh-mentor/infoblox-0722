from io import StringIO
import sys


class Item:
    def __init__(self, name, price):
        self.name = name
        self.price = price

class Cart:
    def __init__(self):
        self.items = []

    def add_item(self, item):
        self.items.append(item)

class Cashier:
    def checkout(self, cart):
        for index, item in enumerate(cart.items):
            print(f"Item #{index + 1} {item.name} ${item.price:5.2f}")