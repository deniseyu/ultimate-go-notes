# Design Guidelines

Somewhere along the line, we forgot:

* Codebases don't have to be huge
* More abstractions are not necessarily better
* We rely too much on virtualization, and forget about the hardware
* Every decision has a cost

Core Go programming values:

Quality
Efficiency
Simplicity

Mental models matter for codebases! The average experienced developer can only hold
10000 lines of code worth of mental models at a time. Go enables you to write less code,
and hopefully produce less complexity, so maybe our mental models can be better-managed.

## Productivity vs Performance

Historically, productivity has been prioritized over performance by engineering teams.
We shouldn't worry about shaving the last few nanoseconds off an algorithm, but
performance does matter. Software keeps undoing the hard-earned gains of hardware!

With Go, you can get huge performance gains by learning a little bit about how it works.

## Correctness vs Performance

We should always optimize for correctness, not performance. Lack of performance will come from:

* Latency (network I/O, etc)
* Memory allocation & garbage collection
* Data access
* Algorithmic efficiency

Go should take care of the first three for you!

As engineers we should worry about correctness primarily, over performance. We have to be comfortable
reading and writing clear and simple code to optimize for correctness.

## Code Reviews

Go's priorities:

1. Integrity: being serious about reliability
  a. Every read and write has to be accurate, consistent, and efficient
  b. Data transformation have to be accurate, consistent, and efficient
  c. Error Handling
    - 92% of critical failures caused by poor error handling
2. Readability
  a. The average developer on your team should be able to build a full mental model of your codebase
     If you're above average, your obligation is to not produce clever code!
3. Simplicity: Hiding complexity, but not cost
