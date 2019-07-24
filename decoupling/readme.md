# Decoupling

We want to try to minimize cascading changes: this is where decoupling comes in.

Too many developers start with behaviour and decoupling; but that's wrong! You should
start with solving the concrete data transformation problem, then decouple.

Think of allowing data to have behaviour.

Method: A function that has a receiver. Go convention is to have methods after type.

A piece of data should have behaviour by exception, not by rule. This is the opposite
of OOP, where data is encapsulated inside objects. That's not how Go should be written.
Functions should be your first choice until they're not reasonable anymore.

"If the method has to mutate the data, I'll use a pointer! Otherwise I'll use the value."
BK says this is wrong. We have to take a step back and think about the semantics.

-- Don't ever use pointers as the returned value of initializers --

## Semantic consistency is everything!

There are three classes of data:

1. Built-in types

  You ought to be using *value* types always. Everybody can get their own copy
  of the built-in types. Strings are designed to be mutable.
  Fields in a struct should be treated as values too.
  The only exception is marshaling data into a struct where a field might be nil.

2. Reference types - slices and maps

  Value semantics too! Other than decoding and unmarshaling, everybody should get
  a copy of the slice value. They're all pointer types anyway, so passing by
  reference would be passing the address of an address.

3. Struct types

  Things get more interesting with custom-defined types. Pointer semantics should
  be the exception, value should be default.

### How do you decide whether to choose a value or pointer?

Are you trying to slightly modify a value, or are you returning a brand new version?

Slightly modification --> pointer
Example: A user's email changes, but everything else stays the same.

Brand new entity --> value
Example: Add 5 minutes to a time object; a new time is returned

When unsure... it's better to use pointer semantics because not everything
can be copied.

Within functions, oftentimes pointers will be used. That's fine, because values
are ultimately returned, so consumers of the API will deal in value semantics.

### Exceptions

Unmarshaling and decoding always need pointers. In the Time library, you can see
that for all the Unmarshal functions, they switch to pointer semantics.

---------

Assume that it's dangerous to make copies of data after it's been mutated!

## Function/Method Variables

"Methods are made up!" Methods give syntactic sugar that a piece of data has behaviour.
Under the hood, the compiler passes the receiver as the first parameter to a function.
It will automatically adjust to the pointer if it's a pointer function.

Functions in Go are typed values! That means we can pass a function by its name anywhere.

When should data have behaviour, and when should it not?

BK advises to avoid setters and getters.

When we decouple, we incur a "cost of allocation" when data gets copied. But if it
prevents cascading changes from blowing up the codebase it's worth it.

## Interfaces & Polymorphism

Polymorphism: You write a program and it behaves differently based on the data it
operates on. Tom Kurtz

BK: code changes it behaviour based on the concrete data it's based on.

When we want to process lots of different types of data with a single piece
of code, polymorphism comes into play.

Interfaces are not "real": they define a contract of behaviour.

Interface types are "valueless". When a function takes an interface type, it means we
need to pass a concrete type that implements that interface.

## Method Sets

Constants only exist at compile time, not runtime; therefore they have no addresses!
They can't be used as pointer receivers.

When a method is defined using the pointer receiver, it can be called on the pointer or
the value. This is not true for value receivers, because values have no addresses.
If the address is passed, it can be dereferenced into the value.

The concrete value stored in an interface is not addressable. You have to explicitly
pass the pointer if you want to share data and change stuff.

## Embedding

A type can be embedded in another type, which allows for "inner set promotion": the
methods on the inner type can be called directly on the outer type.

When the outer type implements the method as well, it will override the inner type's.
But until we call something, it's not promoted.

This is not for type reuse; rather for composition.

## Exporting

Go's equivalent of encapsulation.

The basic unit of compilation is a package. Every package lives in its own folder.
Each folder should contain a file called <folder>.go.

Import packages by their folder name. Import paths are always relative to your $GOPATH.

You can theoretically export a "New" function from a package that returns an instance of
a private custom type. The compiler has access to it, but it's considered dirty.

When exporting types, the names must be upcased, and struct fields that are lowercase
cannot be set or accessed from outside.

Oftentimes, types will be unexported, but fields are exported. This is done because of
marshaling and unmarshaling data. Only fields that are exported will be marshaled.
