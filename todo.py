import random

with open ('todo.txt', 'r') as f:
    items = f.read().splitlines()

f.close()
print (random.choice(items))
