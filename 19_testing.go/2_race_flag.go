// race flag can be false negatives. It can not find race conditions.
// we can put logic inside the loop to increase change to find race condition.

func TestingZmthing(t *testing.T) {
	for i := 0; i < 100; i++ {
		// Actual logic
	}
}