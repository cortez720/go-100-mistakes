package main

import "context"

type watcher struct{} // Can contain connections to db.

func NewWatcher() watcher {
	return watcher{}
}

func (watcher) watchV1() {}

func (watcher) watchV2(context.Context) {} // Context-aware function.

func listing1() {
	watcher := NewWatcher()
	go watcher.watchV1() // Retain gourutine in memory.
}

func listing2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Cancelling the watch function. But we can retain some connections opened.

	watcher := NewWatcher()
	go watcher.watchV2(ctx)
}

func (watcher) close() {} // Gracefully closes all connections.

func listing3() {
	watcher := NewWatcher()
	defer watcher.close() // Shutdown all db connections.

	go watcher.watchV1()

}

// We must have to plan when we should stop the gourutine.
