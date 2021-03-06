gotogether
==========

To run go code concurrently.

[![CircleCI](https://circleci.com/gh/caiguanhao/gotogether.svg?style=svg)](https://circleci.com/gh/caiguanhao/gotogether)

```go
gotogether.Parallel{
	func() {
		time.Sleep(100 * time.Millisecond)
	},
	func() {
		time.Sleep(300 * time.Millisecond)
	},
	func() {
		time.Sleep(200 * time.Millisecond)
	},
}.Run()

gotogether.Queue{
	Concurrency: 5,
	AddJob: func(jobs *chan interface{}) {
	},
	DoJob: func(job *interface{}) {
	},
}.Run()
```

See [docs](https://godoc.org/github.com/caiguanhao/gotogether) for usage and examples.

LICENSE: MIT

Copyright (C) 2016 Cai Guanhao (Choi Goon-ho)
