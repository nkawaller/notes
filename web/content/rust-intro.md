# Ownership in Rust

Ownership is what allows rust to make memory safety guarantees without needing
a garbage collector.

## Ownership Defined

- a set of rules
- that defines
- how a Rust program manages memory

## Basics

All programs (in any language) have to manage they way they use memory

- python uses a garbage collector
- in c you must explicitly allocate and free memory
- rust uses ownership rules

### Stack and Heap Review

All data stored on the stack must have a known, fixed size. Data with an
unknown size at compile time must be stored on the heap. The heap is less
organized——here's how it works

- request a certain amount of space
- the memory allocator finds an empty spot that's big enough
- and returns a pointer (the address of that location)

Because this pointer is a known, fixed size, it can be stored on the stack. If
you want to access the actual data, you must follow the pointer. This is a
slower process.

When your code calls a function, the values passed into the function and the
function's local variables get pushed onto the stack. When the function is over
those values get popped off.

The main purpose of ownership is to manage heap data.

## The Rules

- Each value in Rust has an owner
- There can only be one owner at a time
- When the owner goes out of scope, the value will be dropped

## The String Type

All the previous types we covered were of a known size——they can be stored on
the stack. The String type is more complex and requires use of the heap.

When we call `String::from()` it requests the memory it needs. Now in most
garbage collected languages, the GC keeps track of and cleans up memory that
isn't being used anymore. In rust, the memory is automatically returned once
the variable that owns it goes out of scope. When a variable goes out of scope,
Rust calls a special fn called `drop`.

### Two Examples

```
let x = 5;
let y = x;
```

In this code, we bind the value `5` to `x`; then make a copy of the value in
`x` and bind it to `y`. Integers are a simple type (pushed onto the stack). But
it's different with the `String` type:

```
let s1 = String::from("hello");
let s2 = s1;
```

This doesn't work the same way. Instead, what we get is a pointer to store on
the stack; and it points us to memory on the heap. This pointer is made up of
three parts

- ptr: address of the heap location
- len: how much memory (in bytes) the contents of the String is currently using
- capacity: total amount the String has received from the allocator

In the above code example, s2 is a copied pointer to the same data on the heap.
Also, once we assign s2, Rust considers s1 no loger valid. Rust calls this a
_move._ You could say tht s1 was _moved_ to s2.

> other languages call this copying of the pointer only a shallow copy. It's
> only called a move in rust because of the extra invalidation step. Also, rust
> never automatically makes deep copies of your data

### Clone

If you _do_ want to make a deep copy of the heap data, (instead of a shallow
copy of just the stack data), you can call `clone().` This is a more expensive
operation.

### Copy

Fixed size types such as integers get a special annotation called `Copy` and
the following code would work:

```
let x = 5;
let y = x;
println!("{x}");
```

Here, x doesn't get invalidated.

## Ownership and Functions

The mechanics of passing a value to a function are similar to those when
assigning a variable.

```
let s = String::from("hello");

takes_ownership(s);

let x = 5;

makes_copy(x);

```

## Return Values and Scope

A function can take ownership and transfer ownership when it returns the value.
However all this taking and returning ownership all the time can get a bit
tedious. What if we want to let a function use a value but not take ownership?
Lucky for us, Rust enables this through the use of references.

## References

A reference is like a pointer——it's an address we can follow to access the data
stored at that address. However, unlike a pointer, a reference is guaranteed to
point to a valid value of a particular type for the life of that reference.

Here's how you use them:

```
fn calculate_length(s: &String) -> usize {}

calculate_length(&s1);
```

References allow you to refer to some value without taking ownership of it.

Just as in C, `*` is the dereference operator.

References are immutable by default just like variables. But you can change
them by making the references mutable:

```
&mut
```

_Important:_ if a reference is mutable, you can't have any other references to
that value; only one at a time. This prevents multiple references to the same
data at the same time which can lead to _data races._

Data Race:

- two or more pointers access the same data at the same time
- at least one of the pointers is being used to write to the data
- there's no mechanism being used to synchronize access to the data

Also, you can't have both immutable and mutable references to the same data at
the same time. You _can_ however have many immutable references to the same
data at the same time.

### Dangling References

It's easy to create a dangling pointer by accident:

- a pointer references a location in memory
- that memory is given to someone else (through `free()`)
- the pointer still points to that location (but the data has changed!)

## The Slice Type

Slices let you reference a contiguous sequence of elements in a collection
rather than the whole collection. It's a reference to _part_ of a string. The
way you indicate this type in function signatures is like this:

```
&str
```

## Structs

Structs give you a way to group data (you can do this with tuples). Think of a
struct as an object's data attributes in OOP. They (along with enums) are the
building blocks for creating new types in your program.

They're like tuples in that both hold multiple related values. But with structs
you name each piece of data. This makes them more flexible because you don't
need to rely on the order of the data.

Say we had a `User` struct; we could define a `build_user` function that would
instantiate a new user.

> When I see the syntax for tuple structs: `struct Color(i32, i32, i32)` I think
> of building types from the ground up. In this case, we're building a type that
> is built upon the i32 primitive type.

### Example Program

In the calculate area of rectangle program, we have a function that performs
the calculation:

```
fn area(width: u32, height: u32) -> u32 {
```

The problem with it is it's not as clear as it could be. It's supposed to find
the area of one rectangle, but this function takes in two parameters. And it's
not clear that the parameters are related.

We could refactor with tuples, and that would be better in a way. But because
tuples don't name their arguments, it's still not entirely clear what's going
on.

_Structs add meaning by labeling the data._ Being able to pass one "Rectangle"
into the area function, and then being able to label what the two numbers mean
(width and height) makes the program a lot clearer——the function signature now
says exactly what you mean.

### Traits

Ever see this error???

```
error[E0277]: `Rectangle` doesn't imple
```

How can we print our Rectangle all at once?

One thing we can do is add this to our struct:

```
#[derive(Debug)]
```

### Methods

Methods are similar in many ways to functions. You declare them with the `fn`
keyword. However they are defined within the context of a struct, and their
first parameter is always `self.`

We define them within implementation blocks (`impl`), and everything within the
block will be associated with that struct. (All functions within an impl block
are called _associated functions._

The main reason to use methods over functions is for organization. It's easier
if everythin that can act on a Rectangle is within one `impl` block instead of
scattered all over the place.

You can define a method that has the same name as one of the fields on the
struct——and often these methods just return the value that's in the field with
the same name. These methods are often called _getters._ The advantage to using
getters is if you make the field private and the method public, then you enable
read-only access to that field as part of the type's public API.

We can have associated functions that don't have self as their first parameter
and therefore are not methods. This pattern is often used for constructors that
return a new instance of the struct. The convention is to name these `new().`
To call this function we use the `::` syntax.

## Enums

While structs give you a way of grouping together related fields and data,
enums give you a way of saying a value is one of a possible set of values. An
example of this is with IP addresses. Currently all IPs are either version 4 or
version 6. So we can enumerate all possible variants:

```
enum IpAddrKind {
    V4,
    V6,
}
```

Now we have a custom data type we can use elsewhere in our code. And you can
put data directly into enum variants:

```
enum IpAddrKind {
    V4(String),
    V6(String),
}
```

You can define methods on enums too! Just put it in an `impl` block.

### Option

This enum provided by the standard library encodes the scenario where a value
could be something, or it could be nothing.

Many languages implement the concept of null——NULL is a value that represents
the absence of a value. Rust doesn't do this, but it does have the Option enum
which allows you to encode this concept:

```
enum Option<T> {
    None,
    Some(T),
}
```

> <T> is the generic type parameter.

In order to have a value that can possibly be null, you must explicitly opt in
by making the type of that value `Option<T>`. Then when you use that value you
are required to explicitly handle the case when the value is null.

### Control Flow with Match

Match seems very similar to an if statement, but there's one big difference. An
if statement needs a boolean conditon to be met——_but with match the condition
can be any type_.

match is very useful with Option<T> —— we can create an arm to execute if there
is no value (None), and another to execute when there is a value.

Rust devs combine match with enums often. A common pattern is match against an
enum; bind a variable to the data inside; execute code based on it.

Use the underscore to match "all other cases"; combine that with an empty tuple
to "do nothing in all the other cases:

```
_ => ()
```

#### if let

this expression behaves the same way as `match`. Use in the case where you're
trying to match one pattern, and ignore all other values. It's like match
shorthand.

## Packages and Crates

rustc == the rust compiler
A _crate_ is the smallest amount of code the compiler considers at a time. And
crates can contain modules. _Library crates_ don't contain a main() function
and don't compile to an executable. Rust devs often refer to these when they
say "crate"——you could also say "library."

A _package_ is a bundle of one or more crates that provides a set of
functionality.

### Steps in creating a package

1. run `cargo new my-project`
2. Cargo creates a few things, including a Cargo.toml file——this file is what
   gives us a package.
3. Cargo follows the convention that `src/main.rs` is the crate root of a
   binary crate. If your package is just a library crate, your crate root will
   be `src/lib.rs` instead. (You can have both too!).

### Paths, Scope, Privacy

The `use` keyword brings a path into scope.
The `pub` keyword makes items public.

[how modules work]: https://doc.rust-lang.org/book/ch07-02-defining-modules-to-control-scope-and-privacy.html

### Modules

_Modules_ let us organize code within a crate for readability and easy reuse.
Code within a module is private by default.

By using modules you can group related definitions together. Programmers using
this code can navigate the codebase based on these groups, and when they need
to add new functionality, they know where the code should go.

The module tree looks a lot like a filesystem directory tree. You can organize
your code in this same way.

```
crate
 └── front_of_house
     ├── hosting
     │   ├── add_to_waitlist
     │   └── seat_at_table
     └── serving
         ├── take_order
         ├── serve_order
         └── take_payment
```

### Referring to Items in the Module Tree

Two ways:

- absolute path (starts with `crate` keyword)
- relative path (uses `self` or `super`)

### The `use` keyword

`use` provides a shortcut to bring paths into scope.

#### The Standard Library

The standard library ships with the rust language, so you don't need to add it
to Cargo.toml, but you do need to use the `use` keyword to access it.

#### Nested Paths

If you're brining a bunch of items in from the same crate or module, it's more
concise to use this syntax:

```
use std::{cmp::Ordering, io};
```

Use the glob operator to bring in all public items:

```
use std::collections::*;
```

This is most often used when testing, to bring everything into the tests module

### Separating Modules into Different Files

As modules get large, you might want to move their definitions to a separate
file for better organization.

[example]: https://doc.rust-lang.org/book/ch07-05-separating-modules-into-different-files.html

## Collections

Unlike array and tuple types, the data these collections point to is stored on
the heap. The three you'll use most often are:

- vector (resizable array)
- string (collection of characters)
- hash map (kev value store)

### Vectors

Create a new vector like this:

```
let v: Vec<i32> = Vec::new();
```

Here, we're adding the type annotation because it's initialized without any
values, and Rust can't infer what we intend to use it for. If you do initialize
it with values, you don't need the type annotation:

```
let v = vec![1, 2, 3];
```

Vectors can only store elements of the same type. If you need to store a
collection of different typed items, use an enum.

### Strings

Strings are stored as a collection of bytes.

We should quickly remind ourselves that string slices (`str`, `&str`) are also
referred to as strings, but here we're talking about a growable, mutable, owned
object.

String literals are slices:

```
let data = "hello";
```

To make it a String type we can use `.to_string()`

```
let s = data.to_string();
```

You can concatenate strings using the `+` operator, but that can get a bit
unwieldly. As things get more complex, use the `format!` macro:

```
let s = format!("{s1}-{s2}-{s3}");
```

Note! Rust doesn't support accesing Strings by index; no s[0]. This is a
complex topic - look into it more.

### Hash Maps

They use the SipHash hash function.

## Panic! and Error Handling

In C if you try to access a array past it's alloted memory you get unspecified
behavior——this is called a _buffer overread_ and can lead to security issues.
Rust prevents this by throwing a panic! instead.

A _backtrace_ is a list of all the functions that have been called to get to
the failure point.

### Result

Use the Result enum for errors that aren't so serious that we need to unwind
all the way back out. It has two variants, `Ok` and `Err`——we can use a match
expression to handle each case. And we can drill deeper and handle different
error cases using match as well. Use ErrorKind for this.

Match expressions are primitive——one higher level way to handle this that is a
bit more readable is to use closures.

Other ways to do this are using `unwrap` or `expect`. Expect is the common
convention in production-quality code among Rust devs.

There's also the `?` operator we can use.

### When to use panic! or Result

In prototyping and testing start off with unwrap and expect——these make it
clear where your failures are, and you can always go back to them when you're
ready to make your code more robust.

In general, use panic! when your code could end up in a bad state:

- invalid or missing values are passes to your code
- something unexpected as opposed to something that happens occasionally
- when the code further down the line depends on not being in a bad state

However if a failure is expected, it's more appropriate to return a Result.

## Generic Types, Traits, and Lifetimes

Generics are a tool for handling teh duplication of concepts (controlling
complexity via abstraction?)

Generics:

- replace specific types with...
- a placeholder that represents multiple types

### Basic way to avoid duplication

- Write code to fine the largets number in a list
- Then say you want to find the largest number in a different list
- You "could" copy and paste the code
- But a better way would be to
  - create an abstraction
  - by defining a function
  - that operates on any list of integers

Here's a basic formula:

1. Identify duplicate code
2. Extract the duplicate code into the body of the function and specify the
   inputs and return values of that code in the function signature
3. Update the two instances of duplicated code to call the function instead

We can use this same process with generics. Just like a function can operate on
"any list", generics allow code to operate on abstract types——like operating on
a slice of `i32` or a slice of `char` values.

### Generic Data Types

The basic syntax for a function is:

```
fn largest<t>(list: &[T]) -> &T {}
```

You could also create struct using generics. For instance a `Point` struct
where x and y were either ints or floats.

```
struct Point<T> {...
```

or, if x is a different type from y:

```
struct Point<T, U> {...
```

## Traits

Traits define functionality a particular type has and can share with other
types. They are similar to the concept of `interfaces` in other languages.

## Lifetimes

Lifetimes are another kind of generic. They ensure that references are valid as
long as we need them to be.

Every reference in rust has a lifetime. Most of the time they are implicit and
inferred. However there are times when rust explicitly needs you to annotate
them. For example:

```
fn main() {
    let r;

    {
        let x = 5;
        r = &x;
    }

    println!("r: {}", r);
}
```

This code won't compile because:

- we set the value of r as a reference to x
- but by the time we try to use r, x has gone out of scope (it's memory was
  deallocated)

### The Borrow Checker

The borrow checker compares scopes to determine whether all borrows are
valid. Take this code for example:

```
fn longest(x: &str, y: &str) -> &str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}
```

This will give us a lifetime error. Which variable (x or y) will be returned?
Which variable needs to live longer? We don't know if the `if` or `else` block
will be executed. So we'll apply a generic lifetime that could handle either.

Here's how we'd annotate the above function:

What this says is, "the returned reference will be valid as long as both the
parameters are valid." Lifetime annotations don't change the length of
lifetimes; they just describe the relationships.

We're specifying that the borrow checker should reject any values that don't
adhere to these constraints. They go into the function signature because they
are part of the "contract" established.

> Ultimately, lifetime syntax is about connecting the lifetimes of various
> parameters and return values of functions.

All in all, lifetimes prevent dangling references.

## Testing

Correctness is important, and the type system shoulders a large part of the
burden in proving it, but it can't do it all.

### Basic Test Flow

1. Set up any needed data or state
2. Run the code you want to test
3. Assert the results are what you expect

This line: `use super::*;` means "make anything we define in the outer module
available to this tests module.

### Test Organization

There are two broad categories of tests:

1. Unit (one module at a time)
2. Integration (external, the same as any user (or other code) would)

`#[cfg(test)]` —— only compile and run this code with `cargo test`

Rust does allow you to test private functions.

> tests are just rust code, and the tests module is just another module

Integration tests can only call functions that are part of your library's
public API. These kind of tests go here:

```
adder
├── Cargo.lock
├── Cargo.toml
├── src
│   └── lib.rs
└── tests
    └── integration_test.rs
```

You can add more than one file in this dir, and it's good practice to group
tests together by the functionality they cover.

## Iterators

In the iterator pattern, the iterator is responsible for the logic of
iterating over each item and determining when the sequence has
finished.

In Rust, iterators are lazy, meaning they have no effect until you
call methods that consume the iterator to use it up.

Say we have a vector:

```
let v1 = vec![1, 2, 3];
```

Now we can create an iterator like this:

```
let v1_iter = v1.iter();
```

This doesn't do anything yet; we're just storing the iterator in the
variable `v1_iter`

> when you use a for loop, under the hood you're creating and
> consuming an iterator

In C and JavaScript you need to explicitly start a variable at index
0 and increment that variable to get values from a vector. Iterators
do this for you.

Iterator adaptors are methods defined on the iterator trait that don't
consume the iterator, but instead produce different iterators by
changing some aspect of the original iterator. `map()` is one example.
The `map()` method returns a new iterator that produces modified
items.

> map() takes a closure;
> so does filter()

In our minigrep I/O project, we were able to use an iterator to avoid
having to check the length of the args (the iterator handles that for
us).

We were also able to change our `search()` function. Instead of an
intermediate, mutable state variable, we could handle it in a more
functional style.

Most Rust programmers gravitate towards using iterators over loops.
Then instead of fiddling with lower-level looping mechanics, you can
focus on the high-level objective of the loop.

> iterators get compiled down to the same low level code that loops do
> so performance is comparable

## Smart Pointers

References are the most basic type of pointer. Smart pointers add more
functionality.

- References only borrow data
- Smart pointers own the data they point to

### Cons List

This is a recursive data type. It's a data structure that comes from
Lisp, and it's made up of nested pairs. It's Lisp's version of a
linked list (pairs of pairs).

### Interior Mutability

This takes advantage of unsafe Rust. Unsafe code is code that's
manually checked rather than relying on the compiler to do it.

## Is Rust OOP?

In OOP, an object is something that bundles data, and methods that act
on that data into a single entity. In this sense yes, rust is oop:

- structs and enums have data
- impl blocks provide methods on structs and enums

Encapsulation is another key idea in OOP——hiding the implementation
details and forcing interaction via public API. Rust offers this
through the use of `pub`. Everything else is private by default.

Inheritance as a Type System (interesting!) —— Rust doesn't provide
inheritance; a struct can't inherit from another struct.

Polymorphism —— you can substitute multiple objects for each other at
runtime if they share certain characteristics. Or... "code that can
work with data of multiple types."

We can use trait objects to implement polymorphism in Rust instead of
inheritance. The specific purpose of traits is to allow abstraction
across common behavior.

## Fearles Concurrency

Concurrent programming is where different parts of a program execute
independently.

Parallel programming is where different parts of a program execute at
the same time.

Ownership and type systems are tools that Rust uses to manage both
memory safety and concurrency problems.

Message-passing concurrency is where channels send messages between
threads.

Shared-state concurrency is where multiple threads have access to some
piece of data.

### Using Threads to Run Code Simultaneously

In most operating systems, an executed program's code is run in a
process. The OS manages multiple processes at once.

Within a single program, you can also have independent parts that run
simultaneously. The features that run these intependent parts are
called threads.

> Example: A web server could have multiple threads so it could
> respond to more than one request at a time.

Rust uses a 1:1 model of thread implementation. In this model, a
program uses one OS thread per one language thread.

When using the main thread along with spawned threads, all spawned
threads are shut down when the main thread completes. To ensure all
spawned threads complete, we can save the return value of the spawned
thread in a variable, and call the `.join()` method at the end. This
will make sure all spawned threads finish befor main exits.

Calling `.join()` blocks the thread currently running until the thread
represented by the handle terminates.

### move

There's a closure passed into thread::spawn, and if we want our
spawned thread to use any data from our main thread we need that
closure to capture the values it needs.

By adding the `move` keyword to our closure, we force the closure to
take ownership of the values it's using.

### Message Passing

Message passing is used to ensure safe concurrency. Here's a slogan
from the Go language:

> Do not communicate by sharing memory; instead, share memory by
> communicating

To accomplish message-passing, Rust provides an implementation of
channels. A channel is a general programming concept by which data is
sent from one thread to another.

Think of a channel like a stream or a river. Put a rubber duck in it
and it'll travel downstream to the end of the river.

A channel has two parts:

- transmitter (upstream)
- receiver (downstream)

One part of your code calls methods on the transmitter to send data,
and another part checks the receiving end for arriving messages.

Use `mpsc` to use channels. It stand for:

> multiple producer, single consumer

### Shared-State Concurrency

Message passing isn't the only game in town. You can also achieve
concurrency through multiple threads accessing the same shared data.

This is blowing off the first part of the Go slogan:

> do not communicate by sharing memory...

Rust allows this, and a mutex is one of the more common concurrency
primitives for shared memory.

> This is in SICP!!!

Mutex stands for "mutual exclusion"

The way it works is it allows only one thread to access some data at
any given time. It does this through a lock.

Mutexes are tricky to get right (which is why people love channels).
But Rust's type system and ownership rules make it so you can't get
them wrong!

When you put an i32 in a mutex, you must use `.lock()` to access the
i32 inside. The mutex will prevent you from accessing it otherwise.

Mutexes return smart pointers so you need to dereference them.

If you try to use a mutex with multiple threads you'll need to enable
multiple ownership using Arc<T>. It's like the type Rc<T> but the A
stands for atomic. It's an atomically reference counted type. They are
safe to share across threads.

## Patterns and Matching

The program matches values against the patterns to determine whether
it has the correct shape of the data to continue running a particular
piece of code.

Overall, match expressions are very useful in distinguishing between
different kinds of data.

The compiler checks match expressions and makes sure they cover every
possible case——this isn't the case with `if let` expressions, and it's
an advantage to using `match`.

### Refutability

Patterns come in two forms:

- refutable: patterns that can fail to match some possible value
- irrefutable: patterns that will match for any possible value
