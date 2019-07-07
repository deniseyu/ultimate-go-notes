# Language Syntax

## Variables

Types tell you what a bit pattern represents!

They tell us:

* Size of the data
* What it represents

Types can change size based on architecture, especially integers. It's better to use the int type
unless you really need a specific integer size.

Zero values are built-in integrity feature. Using a var assignment always gives you the zero value.

Go does not cast; but you can convert for a memory cost. You could use the Unsafe package if you really wanted to cast,
but the idiomatic Go approach is to convert. Another built-in integrity feature!

## Struct types

Go doesn't have "objects"! But you can define types, then create values from those type declarations.

When you declare a struct with various fields, Go attempts to use a concept called alignment on word boundaries.
Go prefers to not let values cross over word boundaries, so it tries to fit values within boundaries instead.
This means that we'll usually end up with padding.

If you have a 1-byte value, it'll always fit inside a boundary. The Go compiler has to know the size of a value
to figure out where it can fit.

```
0    1    2    3    4    5    6    7    8
|----|----|----|----|----|----|----|----|
| fl |    |  int16  |       int32       |
|----|----|----|----|----|----|----|----|
      ^ padding
```

The int16 won't start on address 1. It has to start on 2! If you have a small data type like a bool followed by
a big data type like int64, the bool will be padded out by 7 bytes!

One possible micro-optimization is order your struct fields from largest to smallest.

What if you have two structs with identical fields and therefore identical memory layouts?
You can't assign one value of type A to a value of type B, even if it might theoretically be okay!
This is *implicit conversion*, and it's risky -- historically it has caused more problems than good.

If you really need to do this assignment, then you have to explicitly convert.

When a type is named, you have to explicitly convert. But if the value being converted is a literal type
(an anonymous struct), then it's okay to convert it to a named type!

## Pointers Semantics: Passing Values

Everything in Go is passed by value as we cross programmatic boundaries!

Threads are paths of execution from the OS's point of view.

Three areas of memory:

- Data segment (Global vars, etc)
- Stack:
  Every thread is given a stack. Usually 1 MB of memory for each thread. The choice of the stack as the data
  structure is enforced by hardware design.
    - Every Goroutine also has its own stack. It's 2K in size (used to be 4K!), which is much smaller than the
    OS stack.. so that multiple Goroutines can be scheduled in one thread!
    - A Goroutine only has direct access to memory for the frame that it's operating on. If we need to
    execute a data transformation in a goroutine, it has to fit in that frame.
    - This sort of sandboxes the goroutine: it prevents the goroutine from interfering with memory elsewhere
- Heap

When we pass things around using values, goroutines are operating in the background!

## Pointer Semantics: Sharing Data

Passing by reference means we store an address, then pass a copy of that address.

Pointers are literal types, unnamed types. You can't just pass "\*" as a parameter type because it doesn't have
enough information about what data type is underneath. In order to modify a value in memory, we need to know
what type it is.

Memory below the active frame doesn't have integrity and should be treated as invalid.
