## Understanding Loops in Python: `for`, `while`, and Loop `else`

Loops allow you to execute a block of code repeatedly, either a fixed number of times or while a condition is true. Python has two main types of loops: `for` and `while`. Both can optionally include an `else` clause.

### `for` Loop

The `for` loop is used to iterate over a sequence such as a list, tuple, string, or range:

```python
fruits = ["apple", "banana", "cherry"]

for fruit in fruits:
    print(fruit)
```

**Output:**

```sh
apple
banana
cherry
```

#### Using `range()` with `for`

You can use the `range()` function to loop a specific number of times:

```python
for i in range(5):
    print(i)
```

**Output:**

```sh
0
1
2
3
4
```

### `while` Loop

The `while` loop runs as long as a condition is `True`:

```python
x = 0

while x < 5:
    print(x)
    x += 1
```

**Output:**

```sh
0
1
2
3
4
```

### `else` Clause with Loops

Python allows an optional `else` block with loops. The `else` block executes **only if the loop completes normally**, without a `break`:

```python
for i in range(3):
    print(i)
else:
    print("Loop completed without break")
```

**Output:**

```sh
0
1
2
Loop completed without break
```

```python
x = 0
while x < 3:
    print(x)
    x += 1
else:
    print("While loop finished without break")
```

**Output:**

```sh
0
1
2
While loop finished without break
```

### Using `break` in Loops

The `break` statement exits the loop immediately:

```python
for i in range(5):
    if i == 3:
        break
    print(i)
```

**Output:**

```sh
0
1
2
```

Notice that the `else` block will **not** run if a `break` occurs.

### Using `continue` in Loops

The `continue` statement skips the rest of the current iteration and moves to the next iteration:

```python
for i in range(5):
    if i % 2 == 0:
        continue
    print(i)
```

**Output:**

```sh
1
3
```

### Nested Loops

Loops can be nested inside each other:

```python
for i in range(1, 4):
    for j in range(1, 4):
        print(f"{i} x {j} = {i*j}")
```

**Output:**

```sh
1 x 1 = 1
1 x 2 = 2
1 x 3 = 3
2 x 1 = 2
2 x 2 = 4
2 x 3 = 6
3 x 1 = 3
3 x 2 = 6
3 x 3 = 9
```

### Conclusion

Python’s loops (`for` and `while`) are versatile and allow you to repeat actions efficiently. Combining loops with `else`, `break`, and `continue` gives fine-grained control over iteration. Nested loops enable handling complex scenarios such as matrices or multi-level data structures.
