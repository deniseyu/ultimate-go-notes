# Arrays

## Semantics

Array sizes have to be known at compile time, so you can't use a variable
to set the size of the array!

Value semantics mean we operate on our own copy of data. A "for i, item := range items"
loop will iterate over a copy of the data. A "for i := range items" will pass around a
reference to the original data.

## Slices

"Slices are the most important data structure in Go!"

We've already seen built-in types (strings, ints) and custom types (user-defined structs).

Reference types: Slices, maps, functions are all reference types because they are
data structures with pointers! Strings resemble reference types, but they're not because the
zero value is empty.

The "make" call will make a new slice with three words:

1. Address of backing array
2. Length: total # of accessible elements from that pointer position
3. Capacity: max # of elements in backing array. Slices have backing arrays. For future growth!

### Appending Slices

When you instantiate an empty slice, the first word actually points to a global
empty backing array. There is a slight performance hit associated with this, so
it's better to start with a nil slice when you don't know:

"var mySlice []string" over "mySlice := []string{}"

Appending is a value-semantic API which does not mutate the underlying data. If the backing
array isn't long enough to have another element added, the compiler will copy over the data
into a 2x longer backing array under the hood, and the address will update. The original
backing array will get swept up eventually.

#### Memory Leaks in Go

In C-like languages, you have to manage memory yourself and tell the compiler when to GC.
In Go, a memory leak is when we maintain a reference to a heap value and the reference doesn't
go away.

1. Are you creating extra goroutines that should be terminating, but are not? This is the
classic memory leak scenario in Go.
2. Maps are commonly used as caches, but how do you clean up unused keys? If you don't, the
map will grow indefinitely. You should delete keys based on time or event.
3. Append calls! If the value being appended to is different from the new, then we leave
references hanging around.
4. Not calling close()

---

Every time Go performs the copy over to a new array, it uses some extra resources. But if
we *do* know about what we might need upfront, we can declare a size for the backing array and
allocate those resources upfront, reducing the number of copy operations.

The backing array lives on the heap, so accessing it is expensive.

## Slices of Slices

When we take a subset of a slice using the syntax

mySlice[1:3]

we share the same backing array, but the length and capacity are shorter.

Side effects:
* If the sub-slice changes a value, then the original slice will also be modified
* In cases where a slice references other slices, it's possible to get into a situation
where backing arrays get updated for some references but not others

To get around this, we can try to use append to force a new backing array to be created.
However, if capacity is bigger than length, this won't happen. There is syntax to force "copy-on-write":

slice2 := slice1[2:4:4]

The third parameter stipulates that length and capacity are going to be the same!

You can also call "copy".

## Strings & Slices

Strings are UTF-8 based. They can be ranged over! But not "character"-by-character, rather
code point by code point, because different characters have different byte sizes. Chinese
characters take up more space than English!

The rune length tells us how many bytes fit between code points.

## Range

It's possible to change the value of the underlying data from inside of a "range" call. As
long as we're using value semantics, that's okay, because we're iterating using a copy.
Pointer semantics are where you get in trouble.

## Maps

When you iterate over a map, the entries will appear randomly. You can import the "sort"
package and call 'sort.Strings(keys)' to alphabetize the keys!

Maps do their best to keep data contiguous, but not as well as slices and arrays.


