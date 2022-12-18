// While using a worker pooling-like pattern, meaning spinning a fixed number of
// goroutines that need to send data to a shared channel. In that case, we can tie
// the channel size to the number of goroutines created.

//  When using channels for rate-limiting problems. For example, if we need to
// enforce resource utilization by bounding the number of requests, we should set
// up the channel size according to the limit.

// If we are outside of these cases, using a different channel size should be done cau-
// tiously. It’s pretty common to see a codebase using magic numbers for setting a chan-
// nel size:
// ch := make(chan int, 40)
// Why 40? What’s the rationale? Why not 50 or even 1000? Setting such a value should
// be done for a good reason. Perhaps it was decided following a benchmark or perfor-
// mance tests. In many cases, it’s probably a good idea to comment on the rationale for
// such a value.

// Let’s bear in mind that deciding about an accurate queue size isn’t an easy prob-
// lem. First, it’s a balance between CPU and memory. The smaller the value, the more
// CPU contention we can face. But the bigger the value, the more memory will need to
// be allocated.