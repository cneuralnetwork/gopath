# Packages and modules

A module names a collection of packages. Exported identifiers begin with an uppercase letter. internal directories limit who can import a package.

:::spec
Go 1.22 and newer use the behavior described here. Run the example with `go run .` in a module.
:::

## Start with a small program

```go
go mod init example.com/myapp
```

Read the program from the declarations toward the expressions. Go rewards explicit data flow: you can usually see where a value comes from, where it changes, and what a function returns. When an operation can fail, check its error at the point where you have enough context to explain it. This makes code easier to debug than a chain of silent conversions or exceptions.

## Make the behavior visible

Try changing one input at a time and predict the output before you run the program. Print intermediate values while learning; remove temporary output when the test captures the behavior. If a snippet uses a library package, read its `go doc` entry as well as the example. The documentation tells you which edge cases the function promises to handle.

:::gotcha
A common mistake is to assume that Go will infer what you meant from another language's syntax. Check types, return values and zero values explicitly. A compiler error is useful feedback about the contract you wrote.
:::

## Practice the idiom

The exercises below begin with a compiling implementation. They deliberately return an incomplete result. Run it first, read the output, then make one small change and test again. Visible tests are examples of the contract; hidden tests cover edge cases. If a test fails, inspect the function's input and expected output before changing code.

:::idiom
Keep the smallest implementation that makes the behavior clear. Use gofmt to format it. Prefer a direct expression over clever state when the task is simple.
:::

## Why Go is shaped this way

Go favors tools and conventions shared by every project. The same formatter, test command and package rules work as a program grows. That consistency makes code written by another developer easier to read. You do not need an elaborate abstraction before the concrete behavior is understood. Start with a small function, name the data it needs, and let a test state its observable result.

As you move to the next concept, connect this one to previous work: Which values are copied? Which operation may fail? Where should cleanup occur? Which behavior belongs in a test? Those questions matter more than memorizing a syntax fragment. Change the code, run it locally, and use a failing test as a precise question about what the program does.
