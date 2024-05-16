# Lambda Functions

Anonymous function.
Pure, does not keep any state.
```python
def identity(x):
    return x
```
turns into

```python
lambda x: x
```

The lambda function is composed of 3 parts:
- lambda keyword
- bound variable: x   # An argument to a function
- body: x

As opposed to bound is a free variable - _scope_!

## Save a lambda in a variable
addOne = lambda x: x + 1

## You can have multiple arguments
fullName = lambda first, last: f'{first.title()} {last.title()}'

## Higher order functions
Lambdas are often used with higher-order functions, which _take
one or more functions as arguments or return one or more
functions_.

## Arguments
(lambda x,y,z: x+y+z)(1,2,3)
### You can set default values
(lambda x,y,z=3: x+y+z)(1,2)
### Pass a bunch of args in
`(lambda *args: sum(args))(1,2,3)`
### Pass in a bunch of kwargs
`(lambda **kwargs: sum(kwargs.values()))(one=1, two=2, three=3)`

## Closure!

## Common Uses
Used with built-in functions
- map()
- filter()
- functools.reduce()

**Key Functions**
Key functions are _higher order functions_ that take a parameter
key as a named argument.
- sort()
- sorted(), min(), max()
- nlargest() and nsmallest(): **in the heapq module**

Sorting
ids = ['id1', 'id2', 'id30', 'id3', 'id22', 'id100']
sorted(ids)) # Lexicographic sort
sortedIds = sorted(ids, key=lambda x: int(x[2:])) # Integer sort

***I don't really know how sort works***

[Sorting](./python-sorting.md)
