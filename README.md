Some examples to use Golang in the best practices:

- GoRoutines
- Error Handling
- Interfaces
- CAP theorem
- Consistency
- Resilience

Goroutines VS Mutex VS Channels

| Feature                  | Goroutines (`go func()`) | `sync.Mutex`                  | Channels (`chan`)             |
|--------------------------|--------------------------|-------------------------------|-------------------------------|
| **Purpose**              | Enables concurrency      | Prevents race conditions      | Facilitates goroutine communication |
| **Creates Concurrency?** | ✅ Yes                    | ❌ No (only protects shared data) | ✅ Yes (by managing goroutine interaction) |
| **Synchronization?**     | ❌ No built-in sync       | ✅ Yes                         | ✅ Yes                         |
| **Blocking Behavior?**   | Non-blocking             | Blocks other goroutines       | Blocks when channel is full/empty |
| **Best For?**            | Running multiple tasks   | Protecting shared memory      | Passing data between goroutines |
