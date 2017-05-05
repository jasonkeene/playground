
# Preface

- Read on the train, no notes.

# Chapter 1 - Introduction

- Read on the train, no notes. Basic introduction to concurrency concepts.

# Chapter 2 - Thread Safety

- Frameworks impose thread safety requirements as they spin up threads for
  you.
- Stateless objects are thread safe because all their state is temporary and
  stored on the thread's stack and is (typically) not accessible from other
  threads.
- If req/resp handler needs to remember stuff from one request to another that
  is when thread safety of the handler needs to be considered.  ++operator is
  not atomic, it is composed of load, add and store steps aka read-modify-write.
- Race conditions are where the correctness of an algorithm is dependent on
  lucky timing in the ordering of actions taken by each thread.
- The most common race condition is check-then-act.
- Not all data races are race conditions and not all race conditions are data
  races. check-then-act is when you make an observation about the system then
  take an action, however by the time you take the action the observation could
  have become invalid.
- Lazy initialization is where you postpone initing an object until it is
  needed yet ensure that it is only initialized once. This commonly uses
  check-then-act.
- Read-modify-write is a state transition from state A->B. This requires a
  thread to have exclusive access to reading and writing the state. If another
  thread reads the state before the first thread is done modifying it will have
  an invalid understanding of the state. Additionally, if a thread writes while
  the first thread is writing it can overwrite the first thread's write.
- Atomic operations are indivisible. Thread B must wait for thread A to
  complete operations on the state before it can continue.
- check-then-act and read-modify-write are both compound actions that need to
  be atomic to ensure thread safety.
- Java has `java.util.concurrent.atomic` that gives you types that can be used
  to encapsulate values and operate on them atomically to ensure thread
  safety.
- Adding more than one atomic state variable does not guarantee thread safety
  if the operations on those atomic variables need to be atomic themselves.
- Synchronized blocks and methods guard critical sections of code. You can
  pass it an instance or class and it will add a lock to that instance/class.
  These locks are called intrinsic locks or monitor locks.
- Intrinsic locks are mutexes.
- Overly broad locks can cause performance issues.
- Intrinsic locks are reentrant, meaning that if thread A already holds a lock
  if it tries to acquire it again, it will succeed and not deadlock.
- For a given shared mutable variable, all access (reads and writes) must be
  guarded by the same lock to ensure thread safety.
- It is recommended to only use a single lock to guard any given variable.
- Invariants that involve multiple variables must be guarded by the same lock
  so that operations on them are atomic.
- There is a balance between granularity and coarseness. If you are too course
  in your locking you have less concurrency, if you are too granular you pay
  the cost of lock overhead.

## Discussion

- How useful is it to make a type partially thread safe? For instance, if a
  diode is only used with one writer, one reader, it could be thread safe
  under those conditions but not when there are multiple writers or multiple
  readers. Another example is Eno's channel buffer used in his gRPC streaming
  demo.
- Reentrant locks seem costly. They seem to be useful for intrinsic locks with
  inheritance. Is there a use in Go? Can you even implement such a thing
  considering Go's lack of support for thread locals?
- Is there a situation where you would want to use more than one lock to guard
  a particular variable?

# Chapter 3 - Sharing Objects

- synchronized keyword is not only about guarding critical sections but also
  involves memory visibility.
- writes from one thread are not guaranteed to be visible to another thread,
  either in a timely manner or at all.
- Reordering of writes can occur as well if the critical section is not
  synchronized.
- When sharing variables without synchronization you may see some variables be
  stale and others not. Stale reads are when writes from one thread are not
  visible to the reader thread.
- out-of-thin-air safety is where a thread reads a value that was previously
  written by another thread. It is guaranteed to not be a random value.
- 64-bit value load/store operations are not atomic and as a result, you may
  read the high 32 bits from one write but the low 32 bits from another write.
  This means 64-bit values need to be marked as volatile or use synchronization
  when reading and writing.
- the volatile keyword may be used to ensure visibility of a variable. This
  prevents the compiler and runtime from reordering memory operations
  pertaining to this variable. All writes to memory before the write to the
  volatile memory are visible to the thread after reading from that variable.
- If you can confine an object to a single thread at any given time, it can be
  used without synchronization. For example, a connection pool that is used to
  acquire a connection can be used by multiple threads since the pool is safe.
  Once a connection is acquired it is confined to a given thread. As a result,
  the connection does not need to be thread-safe.
- Ad-hoc thread confinement is where it is left up to the implementation to
  not share references.
- Stack and thread local confinement are also alternatives.
- Immutable objects are always thread-safe.
- If you do not properly publish shared state, using volatile or
  synchronization, you will have a bad time.
- If the object is immutable then no synchronization is needed in order to
  publish.

## Discussion

- In Go, are operations in the atomic package guaranteed to be visible by
  other goroutines?
- Using immutable holder objects for variables that are related by an
  invariant along with volatile for visibility is interesting since it
  requires no synchronization.

# Chapter 4 - Composing Objects

- Encapsulation makes analyzing concurrent programs easier. With public state,
  you have to worry about how the entire program might access the state vs the
  type's methods.
- Instance confinement is where data is encapsulated and manipulated
  concurrently through a set of methods that serialize access to that state.
- Monitor pattern is just marking all methods synchronized and calling it a
  day.
- Delegating thread safety is where you hand of thread safety to fields that
  are thread safe. This does not work if you have invariants that relate to
  multiple, thread safe fields.
- Client side locking is just using the same lock of the type you are
  extending.
- Composition is similar to composition in Go where you embed other types and
  delegate methods to those types.
- Document thread safety guarantees for users and synchronization policy for
  maintainers.
- Java documentation, at least at the time of writing, is not great when it
  comes to documenting thread safety guarantees.
- Latches, barriers, semaphores, and blocking queues are types of
  synchronizers.
- Latches block all threads until the terminal state is reached in which case
  all threads unblock.
- CountDownLatch is like a WaitGroup.
- Semaphores control how many things can run at once. It is backed by an int
  counter that is typically inited to be a certain value. This value can grow
  and shrink. If it goes < 0 threads will block until it grows >= 0. This is
  useful for implementing pools.
- Barriers are useful for releasing workers to work in parallel and then
  joining them back up.

## Discussion

- Intrinsic locks are public and exposed to the outside world. That seems
  really messed up.

# Chapter 5 - Building Blocks

- Delegation can be powerful. Use built-in thread safe classes to delegate
  thread safety where possible.
- `Collections.synchronizedXxx` synchronize every public method.
- Compound actions still need to use client-side synchronization.
- Iteration can have timing issues as you call `.size` first and then n `.get`
  calls after. If another thread modifies the map after `.size` was called you
  get an exception or don't get all the data.
- Java5 has an iteration syntax that will throw an exception of modifications
  are made to the collection during iteration.
- You can clone collections for safe iteration since the copy will be thread
  confined. 
- Gotta watch out for hidden state access via `toString`, `hashCode`, and
  `equal` methods.
- `ConcurrentHashMap` uses lock striping vs a single lock for the entire map.
- Iteration on `ConcurrentHashMap` does not require holding a client lock as
  it is weakly consistent. This allows for writes while iterating that may or
  may not show up in the results of the iteration.
- Work stealing is where each worker thread has its own deque. When it is done
  with all the items in its queue it can read from the tail of another
  thread's deque. This minimizes contention compared to a single work queue for
  all workers.

# Chapter 6 - Task Execution

- Choosing appropriate task size and isolation boundaries is critical.
- Creating a thread per task will allocate stacks for each thread and can cause
  paging and/or out of memory errors. Go doesn't suffer this problem quite as
  bad since it mux's goroutines onto a pool of OS threads and has lightweight
  stacks for each goroutine.
- Alternatively, you can have a pool of threads that read off a queue that you
  can put work on from a single thread that accepts connections.
- An execution policy should be defined at deploy time and establish:
  - How many tasks can execute concurrently?
  - How many tasks may be queued waiting to execute?
  - What actions to take before and after executing a task?
  - How to reject tasks?
  - What order should tasks execute in (FIFO, LIFO, priority)?
  - What threads execute what tasks?
- homogeneous workloads that can independently execute allow for high
  parallelism.

## Discussion

- It seems Java has put much more thought into lifecycle management of
  concurrent actors.

# Chapter 7 - Cancellation and Shutdown

- Similar to Go, Java provides no mechanism to safely force a thread to stop.
- Each thread has an interrupted status that can be set from the outside.
  Blocking processes typically read from this value and throw an exception.
- There are forms of blocking that do not throw interrupted exception such as
  synchronous IO and intrinsic locks.
- A poison pill is a sentinel value that is put in a queue to signal teardown
  once it is reached.
- If you are a thread pool and are calling untrusted code, it is best to catch
  all exceptions.
- JVM will fire shutdown hooks and possibly finalizers on graceful shutdown.
- Daemon threads are stuff like the GC that the JVM will preemptively abort on
  shutdown.

## Discussion

- Seems like context with cancel func and deadline/timeout handle most of the
  shutdown situations explained in this chapter.

# Chapter 8 - Applying Thread Pools

- If you are using the executor framework your tasks should be:
  - Independent of one another, not depend on results, timing, or side effects
    of other tasks.
  - Should not use thread local state or otherwise exploit thread confinement.
    You should be able to swap out a single threaded executor with a thread
    pool without issues.
  - Take a long time.
- If you have tasks that block waiting for results of other tasks this can
  cause starvation deadlock.
- Pools can be sized by `N = CPUs * Utilization * (1 + Wait/Compute)`
- Using a pool avoids unbounded thread creation but you can still run out of
  memory if tasks are produced at a higher rate than they are processed. You
  can either throw messages on the ground or rate limit input in some way.
- Before running into memory consumption issues, however, the more tasks
  waiting in the queue the high the latency per task.
- Bounding the pool or queue size with tasks that block on other tasks can
  cause starvation deadlocks.
- Saturation policies describe behavior when the pool's queue is full:
  - Abort: Throws an exception notifying the caller.
  - Caller Runs: Do not discard or throw an exception but run the task given
    in the caller's thread.
  - Discard: Silently discard the new task given.
  - Discard Oldest: Silently discard the next task to run to provide room in
    the queue for the new task.
  - Block: Block the thread writing into the queue until there is room.
- Recursive algorithms can be made parallel if there are tasks that have no
  dependencies on intermediate results. For example computing a result for
  each node in the tree. You can walk the tree sequentially, then do each
  computation in a task with a shared output queue, then collect the results.

## Discussion

- Caller Runs policy is interesting. Not relevant to our domain but
  interesting none the less.

# Chapter 9 - GUI Applications

- GUI frameworks are single threaded for a reason, accessing state from
  multiple threads with events triggered by hardware and events triggered by
  the application can lead to ordering issues in lock acquisition and thus
  deadlocks.
- Long running processes can be handled in worker threads that then write
  events back into the event thread. All GUI state is thread confined to the
  event thread.

# Chapter 10 - Avoiding Liveness Hazards

- Lock-ordering deadlocks can be fixed by having all threads that need the
  same locks acquire them in the same order.
- Calling alien methods while holding a lock risks getting into a deadlock as
  the alien method can then try to acquire a lock.
- Only making open calls to alien methods lowers deadlock risk and makes
  analysis easier.
- Modifying thread priorities is platform dependent and can cause starvation
  of lower priority threads.
- Livelock is where a thread is active but can not make progress. This can
  occur when multiple threads change their state in response to each other.
  For instance, two people walking down a hall both attempt to move out of each
  other's way but then are in each other's way again.

# Chapter 11 - Performance and Scalability

- "How fast" measurements are performance, service time, latency.
- "How much" measurements are scalability, capacity, throughput.
- When making performance decisions:
  - What does it mean to be "faster"?
  - Under what conditions will this approach actually be faster? Light/heavy
    load? Large/small datasets? Do you have measurements?
  - How often are these conditions likely to arise? Do you have measurements?
  - How often are conditions different.
  - What hidden costs, such as maintenance risk, are there?
- Schedulers will run a given task for a minimum amount of time to mitigate
  context switches and increase throughput.
- Increased lock contention increases context switches and serialization
  costs.
- Lock contention probability is determined by how frequent the lock is
  requested and how long the lock is held for once acquired.
- Lock splitting and striping are methods of providing higher granularity in
  situations where contention is high.
- Avoid hot fields where possible. Striping, volatile, or atomics can help
  with mitigating the cost of hot fields.
- Object pooling to minimize allocations is mostly a bad idea with Java.

# Chapter 12 - Testing Concurrent Programs

- There are two categories of tests for concurrent types. Tests for safety and
  tests for liveness.
- Testing concurrent software is hard.
- Performance tests:
  - Collect data, draw graphs.
  - Measure for latency and throughput.
  - Disable optimizations such as dead code elimination as they will sometimes
    eliminate your benchmark code.

## Discussion

- Perhaps we should run our unit tests on a timer as well to help expose bugs
  due to interleavings of goroutines/threads that are not common.
- Running more goroutines than procs and more procs than CPU cores along with
  calling `runtime.Gosched` in tests loops can generate more interleavings.
- Pausing between benchmark runs to allow for GC to fully clean up.

# Chapter 13 - Explicit Locks

- Java provides a ReentrantLock that can be used to explicitly lock.
- Use `finally` to ensure unlock occurs even when an exception happens.
- Polling and timed locks are useful for deadlock avoidance when lock
  acquisition ordering can not be guaranteed.
- Explicit locks are interruptable.
- Hand over hand locking is where you must acquire another lock to release
  your current lock.
- ReentrantLock is faster than intrinsic locks pre-Java 6.
- ReentrantLocks can be fair on unfair. Fair being FIFO for acquisition.
  Fairness costs a lot.
- ReadWriteLock allows for multiple readers at the same time, similar to Go's
  RWLock. They should only be used in read-heavy situations.

# Chapter 14 - Building Custom Synchronizers

- State dependence is where threads are blocked until a certain state exists.
- All objects in Java have an intrinsic condition queue.
- Conditional locking cooperates with the runtime to put threads to sleep and
  wake them up when the state changes.
- Wait can unblock even if notify or notifyAll are not called by another
  thread.
- You may use notify instead of notifyAll in situations where you have a
  single condition predicate.
- Intrinsic condition queues should be encapsulated and not exposed, just like
  intrinsic locks.
- Use multiple Conditions where you have multiple condition predicates.

## Discussion

- Intrinsic with synchronized keyword makes using conditional mutexes easier
  than in Go.

# Chapter 15 - Atomic Variables and Nonblocking Synchronization

- Lock/wait free data structures are immune to liveness issues.
- Putting threads to sleep is costly.
- Priority inversion is where a high priority thread is blocked on a lock held
  by a lower priority thread. Its effective priority is limited by the other
  thread.
- CAS, compare and swap (or compare and set) allows you to atomically read
  from a piece of memory and at a later time atomically write to that memory.
  If the value has changed in the meantime it will fail and allow you to try
  again. This is lock-free but not wait-free.
- JVM will use a spin lock if CAS type instructions are not supported by the
  target hardware.
- Locks can perform better under ultra-high contention scenarios.
- It is a great idea to not share data across cores at all if you can.
- The ABA problem is an issue with CAS where some other thread has changed the
  value from A to B and back to A. This is invalid for some algorithms. To
  mitigate this you can read/write a tuple of (value, version). Every write will
  increment the version and thus invalidate other CAS operations.

## Discussion

- The JVM adaptively determines if a thread should spin or be suspended based
  on how long the lock has previously been held for. That's cool!
- It's weird that they implemented Counters with CAS and not atomic increment.

# Chapter 16 - The Java Memory Model

- A memory model is the conditions under which a write in one core is visible
  to another core.
- ISAs will have their own memory model along with instructions called
  barriers or fences to get additional memory coordination guarantees.
- Sequential consistency is where every read of a variable will see the last
  write to that value from any processor as if there is a sequential ordering
  to all instructions. This is not reality.
- Actions by one thread might appear to execute in different orders from the
  perspective of different threads.
- Happens-before relationship prevents reordering.
- It is possible to piggyback on other synchronization that is forcing a
  happens-before ordering.
- Unsafe-publication can occur if you don't have a happens-before relationship
  between thread B accessing a reference to an object thread A published.
- static initialization happens before a class can be used by threads and is
  guarded by a lock. As a result, the statically initialized variable is
  visible to all threads.

## Discussion

- Piggybacking seems very cool and also very dangerous.
- I wonder what visibility guarantees are given to `func init() {}` or `var
  foo = InitFoo()`.
