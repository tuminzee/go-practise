# Memory Allocation of Arrays in Go

## Arrays in Go

In Go, an array is a fixed-size sequence of elements of the same type. The size of the array is part of its type. When you create an array, memory is allocated for all its elements.

**Example:**

```go
var a [5]int
a[0] = 1
a[1] = 2
a[2] = 3
a[3] = 4
a[4] = 5

fmt.Println(a)  // Output: [1 2 3 4 5]
```

In this example, `a` is an array of 5 integers. Memory is allocated for 5 integers.

## Slices in Go

Slices are more flexible and powerful than arrays. They are a lightweight abstraction over an array. A slice is a descriptor of an array segment and consists of a pointer to the array, the length of the segment, and its capacity.

**Example:**

```go
a := [5]int{1, 2, 3, 4, 5}
b := a[1:4]

fmt.Println(b)  // Output: [2 3 4]
```

Here, `b` is a slice that refers to elements from index 1 to 3 of array `a`.

## Shallow Copy

A shallow copy in Go can be demonstrated using slices. When you assign one slice to another, you are copying the slice header, but both slices point to the same underlying array.

**Example:**

```go
a := []int{1, 2, 3, 4, 5}
b := a
b[0] = 10

fmt.Println(a)  // Output: [10 2 3 4 5]
fmt.Println(b)  // Output: [10 2 3 4 5]
```

Here, modifying `b` affects `a` because both slices share the same underlying array.

## Deep Copy

A deep copy in Go involves copying the data from one slice to another, ensuring that modifications to the new slice do not affect the original slice.

**Example:**

```go
a := []int{1, 2, 3, 4, 5}
b := make([]int, len(a))
copy(b, a)
b[0] = 10

fmt.Println(a)  // Output: [1 2 3 4 5]
fmt.Println(b)  // Output: [10 2 3 4 5]
```

Here, `b` is a deep copy of `a`. Modifying `b` does not affect `a` because `b` has its own copy of the underlying array.

## The `:=` Operator

The `:=` operator performs a shallow copy when used with slices
when you have a slice and you assign it to another slice, it performs a shallow copy, not a deep copy. This means that both slices will reference the same underlying array. Any modifications made to the elements of one slice will be reflected in the other slice.

Here is an example to illustrate this:

```go
a := []int{1, 2, 3, 4, 5}
b := a
b[0] = 10

fmt.Println(a)  // Output: [10 2 3 4 5]
fmt.Println(b)  // Output: [10 2 3 4 5]
```

In this example:
- `a` is a slice that references an array containing `[1, 2, 3, 4, 5]`.
- `b` is assigned to `a`, so `b` now references the same underlying array as `a`.
- When we modify `b[0]`, the change is reflected in `a` because both `a` and `b` reference the same array.


## Summary

- **Assignment (`=`):** Creates a new reference to the same object (slice) or a copy of the value (array).
- **Shallow Copy:** Copies the slice header, but both slices point to the same underlying array.
- **Deep Copy:** Creates a completely independent copy of the data.
- **The `:=` Operator:** Shallow copy.