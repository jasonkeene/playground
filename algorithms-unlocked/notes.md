
## Chapter 1

Two desired properties of an algorithm:

- Correctness
- Resource Efficienty

Correctness is getting the right answer. Sometimes it is acceptable to have an
algorithm that gives wrong answers, such as the prime number check that is
part of RSA which has an error rate of 2**50. Additionally, there is a class
of algorithms called approximation algorithms that can give appoximate
answers. An example of this is finding the fastest route for GPS navigation.

Resources efficiency primarily pertains to:

- Execution Time
- Memory Space

but can also pertain to things like:

- Network I/O
- Disk I/O
- Use of Entropy

Since measuring execution time depends on factors outside the algorithm, such
as execution speed of the machine, we evaluate algorithms based on the size of
their inputs and how increasing that size affects the rate of growth in
execution time.

## Chapter 2

Sentinel linear search is an optimization that prevents you from checking if
you are out of bounds of the array.

In the definition of big theta notation he says:

> The idea is that if we have two functions f(n) and g(n) we say that f(n) is
> Θ(g(n)) if f(n) is within a constant factor of g(n) for sufficiently large
> n.

This was a bit confusing at first but what he is saying is:

- g(n) is just some expression parametrized by n, say n**2
- f(n) is a function, in our situation it is a function that maps n to the
  execution time, say something like 5n**2 + 3n + 20

Big theta allows for an upper and lower bound that is of the form provided.
The actual upper and lower bound functions can be different but must have the
same form (same degree).

Big O notation states the character of the upper bound only. Big omega
notation states the lower bound only. Therefore if a f(n) is Θ(g(n)) then by
definition it is also O(g(n)) and Ω(g(n)).

Loop invariants are important for demonstrating that the loop does the right
thing. They consist of:

- Initialization: Statement that is true before the first iteration
- Maintenance: If it is true before an iteration of the loop, it remains true
  before the next iteration.
- Termination: The loop terminates, and when it does, the loop invariant,
  along with the reason that loop terminated, gives us a useful property.

For recursive algorithms to work you need a base case and the recursive call
must be a smaller set of the problem that will eventually terminate with the
base case.

## Chapter 3

Binary search works in O(lg n) time but only on sorted arrays.  You need to
sort the array before using binary search. If you are only doing a few
searches then linear searching makes more sense.

There are four sorting algorithms covered in this chapter:

- Selection Sort
- Insertion Sort
- Merge Sort
- Quick Sort

The sort key is the information in the value of the element that is used to
sort the array. Satellite data is information associated with the sort key
that needs to travel with the key when it is moved.

Sorting in algorithms differs from what one might consider sorting such as
placing clothes into different groups. This is called bucketing or binning.

For binary search, to find the time complexity, ask how many iterations of the
loop do we need to repeatedly halve a subarray to get to size 1. This is the
same as the amount of times we would have to double the size of the array from
1 to reach n. This can be expressed as:

    2**x = n  or  log2(n) = x

This demonstrates that the loop is dependent on the input n by log2(n). Since
the rest of the algorithm is O(1) things makes binary search O(log2(n)).

It is possible to beat Θ(lg n) worst-case time for searching, but only if we
organize data in more elaborate ways and make certain assumptions about the
keys.

Rearranging the elements of the array is called permuting the array.

The sumation of an arithmetic series is:

    n(a1 + an)/2

With selection sort, since the algorithm swaps in the outer loop there is only
Θ(n) swaps. This might be useful if swapping elements is costly.

Insertion sort's best case is Θ(n) if the array is already sorted. Its worst
case is Θ(n**2) if it is reverse sorted.

Insertion sort is great if the data is already almost sorted as it can
approach Θ(n) time.

Merge sort runs in Θ(n lg n) time in all cases. However, it takes up more
space as it has to copy the elements of array into temporary memory in order
to merge them. Merge sort uses the divide and conquer strategy. Divide and
conquer splits a task into sub tasks, recursively solves those sub tasks, and
then joins the results.

Quick sort uses divide and conquer as well. It is Θ(n**2) in the worst case
but in the average case it is Θ(n lg n). Quick sort is sorted in place so it
uses less memory than merge sort and has better constant factors than merge
sort. Quick sort's form is:

- Choose the right most element as the pivot point.
- Reorder the elements so that all elements that are lower than the value at
  the pivot point are to its left, all elements that are greater than the
  value at the pivot point are to its right.
- Recursively apply sort to 0..pivot-1 and pivot+1..n-1.

Quick sort's worst case is when it is already sorted or is reverse sorted. In
those situations I will take Θ(n**2). In order to get Θ(n lg n), before we
partition we can simply swap the last element with a random element. This will
make getting the worst case quite difficult since it would rely on being realy
unlucky in your random swaps.

## Chapter 4

Counting sort and Radix sort run in Θ(n) time by "bending the rules" so to
speak. Both of these algoritms work by not comparing keys directly to one
another. Comparison sorting is where the elements of the array are compared
with each other.

For counting sort the keys need to be in a range and are numerically indexed,
ie 0-500. The algorithm then counts the frequency of each of these values and
uses this information to rearrange the values into a new array. This means it
does not sort in place. This sort is stable and is Θ(m+n) and m is typically
constant.

Radix sort uses a stable sort (such as a counting sort) to sort larger keys
from right to left. The time complexity for this is dependent on the amount of
digits in the keys that are being sorted, however this is typically constant.

## Chapter 5

One method for getting a linear ordering from a DAG is to start with an
node that has nothing pointing to it, remove it from the graph and place it
in the list. Continue this process until there is no nodes remaining. Placing
nodes in order like this is called topological sorting. There can be multiple
different sortings of a DAG that are correct.

There are a few representations for the edges in a graph:

1. adjacency matrix: A two dimensional array where each row and column
   represents a vertex id. If there is a 1 at matrix[u][v] that represents an
   edge is present. This takes O(n**2) memory.
2. unorderd list: This is just an unordered list of (u, v) doubles.
3. adjacenty list: This is a hybrid of the other two representations. It is an
   array of lists.

Lists used here can either be array backed or a linked list.

Topological sort takes O(n + m) time where n is the amount of vertices and m
is the amount of edges.

A PERT chart is a project management tool that can be represented as a DAG.
The critical path of a PERT chart is the path that taked the most time to
complete. This represents the minimum time to complete the project as a whole.
If we negate the time it takes to complete each task then we can use a sortest
path algorithm to solve for the critical path.

One way to compute the shortest path is to compute the smallest weight between
a starting vertex and all other vertices. If you also record the previous
vertex that for each vertex along the shortest path you can then construct a
single shortest path from u to v. To do this, first topological sort the DAG,
then relax each edge leaving each vertex in the toplological order.
