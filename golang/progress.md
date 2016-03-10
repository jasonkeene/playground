## [Documentation](http://golang.org/doc/)

 - [x] [How to write Go code](http://golang.org/doc/code.html)
 - [ ] [Language Specification](http://golang.org/ref/spec)
 - [ ] [Writing Web Applications](http://golang.org/doc/articles/wiki/)
 - [ ] [Effective Go](http://golang.org/doc/effective_go.html)
 - [ ] [Command go](http://golang.org/cmd/go/)
 - [ ] [Package testing](http://golang.org/pkg/testing/)
 - [ ] [Package json](http://golang.org/pkg/encoding/json/)


## Blogs

#### [The Go Blog](http://blog.golang.org/)

 - [ ] [Go's Declaration Syntax](http://blog.golang.org/gos-declaration-syntax)
 - [ ] [Defer, Panic, and Recover](http://blog.golang.org/defer-panic-and-recover)
 - [ ] [Go Slices: usage and internals](http://blog.golang.org/go-slices-usage-and-internals)
 - [ ] [Strings, bytes, runes and characters in Go](http://blog.golang.org/strings)
 - [ ] [Concurrency is not parallelism](http://blog.golang.org/concurrency-is-not-parallelism)
 - [x] [JSON and Go](http://blog.golang.org/json-and-go)
 - [ ] [Error handling and Go](http://blog.golang.org/error-handling-and-go)
 - [x] [Defer, Panic, and Recover](http://blog.golang.org/defer-panic-and-recover)

#### [Dave Cheney](http://dave.cheney.net/category/golang)

<!--
// use this function to pull links out of Dave's blog
(function () {
    var out = "";
    var elements = document.querySelectorAll('.post .entry-title a');
    for (var i = 0; i < elements.length; i++) {
        var e = elements[i];
        out += " - [ ] [" + e.text + "](" + e.href + ")\n";
    }
    console.log(out);
})();
-->

 - [ ] [Investigate your package dependencies with prdeps and Unix](http://dave.cheney.net/2016/02/29/investigate-your-package-dependencies-with-prdeps-and-unix)
 - [ ] [Unhelpful abstractions](http://dave.cheney.net/2016/02/06/unhelpful-abstractions)
 - [ ] [cgo is not Go](http://dave.cheney.net/2016/01/18/cgo-is-not-go)
 - [ ] [Are Go maps sensitive to data races ?](http://dave.cheney.net/2015/12/07/are-go-maps-sensitive-to-data-races)
 - [ ] [A whirlwind tour of Go’s runtime environment variables](http://dave.cheney.net/2015/11/29/a-whirlwind-tour-of-gos-runtime-environment-variables)
 - [ ] [Wednesday pop quiz: spot the race](http://dave.cheney.net/2015/11/18/wednesday-pop-quiz-spot-the-race)
 - [ ] [The Legacy of Go](http://dave.cheney.net/2015/11/15/the-legacy-of-go)
 - [ ] [Let’s talk about logging](http://dave.cheney.net/2015/11/05/lets-talk-about-logging)
 - [ ] [Bootstrapping Go 1.5 on non Intel platforms](http://dave.cheney.net/2015/10/16/bootstrapping-go-1-5-on-non-intel-platforms)
 - [ ] [Padding is hard](http://dave.cheney.net/2015/10/09/padding-is-hard)
 - [ ] [Building Go 1.5 on the Raspberry Pi](http://dave.cheney.net/2015/09/04/building-go-1-5-on-the-raspberry-pi)
 - [x] [Cross compilation with Go 1.5](http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5)
 - [x] [Using gb as my every day build tool](http://dave.cheney.net/2015/08/19/using-gb-as-my-every-day-build-tool)
 - [x] [Performance without the event loop](http://dave.cheney.net/2015/08/08/performance-without-the-event-loop)
 - [x] [Why Go and Rust are not competitors](http://dave.cheney.net/2015/07/02/why-go-and-rust-are-not-competitors)
 - [x] [gb, a project based build tool for the Go programming language](http://dave.cheney.net/2015/06/09/gb-a-project-based-build-tool-for-the-go-programming-language)
 - [x] [Friday pop quiz: the smallest buffer](http://dave.cheney.net/2015/06/05/friday-pop-quiz-the-smallest-buffer)
 - [x] [Hear me speak about Go performance at OSCON](http://dave.cheney.net/2015/05/31/hear-me-speak-about-go-performance-at-oscon)
 - [x] [Struct composition with Go](http://dave.cheney.net/2015/05/22/struct-composition-with-go)
 - [x] [Introducing gb, a project based build tool for the Go programming language](http://dave.cheney.net/2015/05/12/introducing-gb)
 - [ ] [2015 is going to be the year of Go](http://dave.cheney.net/2015/03/28/2015-is-going-to-be-the-year-of-go)
 - [ ] [A parable about practice](http://dave.cheney.net/2015/03/26/a-parable-about-practice)
 - [ ] [Simplicity and collaboration](http://dave.cheney.net/2015/03/08/simplicity-and-collaboration)
 - [ ] [Cross compilation just got a whole lot better in Go 1.5](http://dave.cheney.net/2015/03/03/cross-compilation-just-got-a-whole-lot-better-in-go-1-5)
 - [ ] [Unofficial Go 1.4.2 tarballs now available](http://dave.cheney.net/2015/02/26/unofficial-go-1-4-2-tarballs-now-available)
 - [ ] [Lost in translation](http://dave.cheney.net/2015/02/25/lost-in-translation)
 - [ ] [Thanks Brainman](http://dave.cheney.net/2015/02/13/thanks-brainman)
 - [ ] [Errors and Exceptions, redux](http://dave.cheney.net/2015/01/26/errors-and-exceptions-redux)
 - [ ] [Inspecting errors](http://dave.cheney.net/2014/12/24/inspecting-errors)
 - [ ] [Unofficial Go 1.4 tarballs now available](http://dave.cheney.net/2014/12/14/unofficial-go-1-4-tarballs-now-available)
 - [ ] [Friday pop quiz: the size of things](http://dave.cheney.net/2014/12/12/friday-pop-quiz-the-size-of-things)
 - [ ] [Minimum one liner followup](http://dave.cheney.net/2014/12/05/minimum-one-liner-followup)
 - [ ] [Friday pop quiz: minimum one liner](http://dave.cheney.net/2014/12/05/friday-pop-quiz-minimum-one-liner)
 - [ ] [Five suggestions for setting up a Go project](http://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
 - [ ] [Visualising dependencies](http://dave.cheney.net/2014/11/21/visualising-dependencies)
 - [ ] [Error handling vs. exceptions redux](http://dave.cheney.net/2014/11/04/error-handling-vs-exceptions-redux)
 - [ ] [Go, frameworks, and Ludditry](http://dave.cheney.net/2014/10/26/go-frameworks-and-ludditry)
 - [ ] [Simple profiling package moved, updated](http://dave.cheney.net/2014/10/22/simple-profiling-package-moved-updated)
 - [ ] [Functional options for friendly APIs](http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
 - [ ] [That trailing comma](http://dave.cheney.net/2014/10/04/that-trailing-comma)
 - [ ] [Unofficial Go 1.3.3 ARM tarballs now available](http://dave.cheney.net/2014/10/03/unofficial-go-1-3-3-arm-tarballs-now-available)
 - [ ] [Using // +build to switch between debug and release builds](http://dave.cheney.net/2014/09/28/using-build-to-switch-between-debug-and-release)
 - [ ] [go list, your Swiss army knife](http://dave.cheney.net/2014/09/14/go-list-your-swiss-army-knife)
 - [ ] [How to install multiple versions of Go](http://dave.cheney.net/2014/09/13/how-to-install-multiple-versions)
 - [ ] [Go’s runtime C to Go rewrite, by the numbers](http://dave.cheney.net/2014/09/01/gos-runtime-c-to-go-rewrite-by-the-numbers)
 - [ ] [Go has both make and new functions, what gives ?](http://dave.cheney.net/2014/08/17/go-has-both-make-and-new-functions-what-gives)
 - [ ] [Tinyterm: A silly terminal emulator written in Go](http://dave.cheney.net/2014/08/03/tinyterm-a-silly-terminal-emulator-written-in-go)
 - [ ] [Visualising the Go garbage collector](http://dave.cheney.net/2014/07/11/visualising-the-go-garbage-collector)
 - [ ] [Unofficial Go 1.3 ARM tarballs now available](http://dave.cheney.net/2014/07/08/unofficial-go-1-3-arm-tarballs-now-available)
 - [ ] [Ice cream makers and data races](http://dave.cheney.net/2014/06/27/ice-cream-makers-and-data-races)
 - [ ] [Five things that make Go fast](http://dave.cheney.net/2014/06/07/five-things-that-make-go-fast)
 - [ ] [What does go build build ?](http://dave.cheney.net/2014/06/04/what-does-go-build-build)
 - [ ] [On declaring variables](http://dave.cheney.net/2014/05/24/on-declaring-variables)
 - [ ] [Go 1.3 linker improvements](http://dave.cheney.net/2014/05/22/go-1-3-linker-improvements)
 - [ ] [Accidental method value](http://dave.cheney.net/2014/05/19/accidental-method-value)
 - [ ] [Unofficial Go 1.2.2 and 1.3beta1 tarballs for ARM now available](http://dave.cheney.net/2014/05/09/unofficial-go-1-2-2-and-1-3beta1-tarballs-for-arm-now-available)
 - [ ] [term: low level serial with a high level interface](http://dave.cheney.net/2014/05/08/term-low-level-serial-with-a-high-level-interface)
 - [ ] [autobench-next updated for Go 1.3](http://dave.cheney.net/2014/05/04/autobench-next-updated-for-go-1-3)
 - [ ] [How to install multiple versions of Go](http://dave.cheney.net/2014/04/20/how-to-install-multiple-versions-of-go)
 - [ ] [Associative commentary follow up](http://dave.cheney.net/2014/03/30/associative-commentary-follow-up)
 - [ ] [Associative commentary](http://dave.cheney.net/2014/03/28/associative-commentary)
 - [ ] [The empty struct](http://dave.cheney.net/2014/03/25/the-empty-struct)
 - [ ] [Thoughts on Go package management six months on](http://dave.cheney.net/2014/03/22/thoughts-on-go-package-management-six-months-on)
 - [ ] [Channel Axioms](http://dave.cheney.net/2014/03/19/channel-axioms)
 - [ ] [Pointers in Go](http://dave.cheney.net/2014/03/17/pointers-in-go)
 - [ ] [pdp11 presentation slides](http://dave.cheney.net/2014/03/08/pdp11-presentation-slides)
 - [ ] [Using go test, build and install](http://dave.cheney.net/2014/01/21/using-go-test-build-and-install)
 - [ ] [Unofficial Go 1.2 tarballs for ARM now available](http://dave.cheney.net/2013/12/29/unofficial-go-1-2-tarballs-for-arm-now-available)
 - [ ] [Go 1.2 performance improvements](http://dave.cheney.net/2013/12/02/go-12-performance-improvements)
 - [ ] [A Go client for Joyent Manta](http://dave.cheney.net/2013/11/21/a-go-client-for-joyent-manta)
 - [ ] [Benchmarking Go 1.2rc5 vs gccgo](http://dave.cheney.net/2013/11/19/benchmarking-go-1-2rc5-vs-gccgo)
 - [ ] [Evaluation order oddity](http://dave.cheney.net/2013/11/15/evaluation-order-oddity)
 - [ ] [More simple test coverage in Go 1.2](http://dave.cheney.net/2013/11/14/more-simple-test-coverage-in-go-1-2)
 - [ ] [Stupid Go declaration tricks](http://dave.cheney.net/2013/11/14/stupid-go-declaration-tricks)
 - [ ] [Subcommand handling in Go](http://dave.cheney.net/2013/11/07/subcommand-handling-in-go)
 - [ ] [Calling for autobench contributions for Go 1.2](http://dave.cheney.net/2013/11/04/calling-for-autobench-contributions-for-go-1-2)
 - [ ] [How does the go build command work ?](http://dave.cheney.net/2013/10/15/how-does-the-go-build-command-work)
 - [ ] [How to use conditional compilation with the go build tool](http://dave.cheney.net/2013/10/12/how-to-use-conditional-compilation-with-the-go-build-tool)
 - [ ] [Why I think Go package management is important](http://dave.cheney.net/2013/10/10/why-i-think-go-package-management-is-important)
 - [ ] [Simple test coverage with Go 1.2](http://dave.cheney.net/2013/10/07/simple-test-coverage-with-go-1-2)
 - [ ] [#golang tweet popularity](http://dave.cheney.net/2013/09/21/golang-tweet-popularity)
 - [ ] [Release candidate 1 tarballs for ARM now available](http://dave.cheney.net/2013/09/21/release-candidate-1-tarballs-for-arm-now-available)
 - [ ] [Unofficial Go 1.1.2 tarballs for ARM now available](http://dave.cheney.net/2013/09/19/unofficial-go-1-1-2-tarballs-for-arm-now-available)
 - [ ] [How to include C code in your Go package](http://dave.cheney.net/2013/09/07/how-to-include-c-code-in-your-go-package)
 - [ ] [Introducing autobench-next](http://dave.cheney.net/2013/08/26/introducing-autobench-next)
 - [ ] [Go 1.1 on the Cubieboard 2](http://dave.cheney.net/2013/08/06/go-1-1-on-the-cubieboard-2)
 - [ ] [An introduction to cross compilation with Go 1.1](http://dave.cheney.net/2013/07/09/an-introduction-to-cross-compilation-with-go-1-1)
 - [ ] [Introducing profile, super simple profiling for Go programs](http://dave.cheney.net/2013/07/07/introducing-profile-super-simple-profiling-for-go-programs)
 - [ ] [Unofficial Go 1.1.1 tarballs for ARM now available](http://dave.cheney.net/2013/07/02/unofficial-go-1-1-1-tarballs-for-arm-now-available)
 - [ ] [How to write benchmarks in Go](http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
 - [ ] [Stress test your Go packages](http://dave.cheney.net/2013/06/19/stress-test-your-go-packages)
 - [ ] [How to install Go 1.1 on CentOS 5.9](http://dave.cheney.net/2013/06/18/how-to-install-go-1-1-on-centos-5)
 - [ ] [You don’t need to set GOROOT, really](http://dave.cheney.net/2013/06/14/you-dont-need-to-set-goroot-really)
 - [ ] [Writing table driven tests in Go](http://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go)
 - [ ] [How Go uses Go to build itself](http://dave.cheney.net/2013/06/04/how-go-uses-go-to-build-itself)
 - [ ] [Why is a Goroutine’s stack infinite ?](http://dave.cheney.net/2013/06/02/why-is-a-goroutines-stack-infinite)
 - [ ] [Go 1.1 performance improvements, part 3](http://dave.cheney.net/2013/05/28/go-11-performance-improvements-part-3)
 - [ ] [Go 1.1 performance improvements, part 2](http://dave.cheney.net/2013/05/25/go-11-performance-improvements-part-2)
 - [ ] [Go 1.1 performance improvements](http://dave.cheney.net/2013/05/21/go-11-performance-improvements)
 - [ ] [Go 1.1 tarballs for linux/arm](http://dave.cheney.net/2013/05/20/go-11-tarballs-for-linux-arm)
 - [ ] [Go and Juju at Canonical slides posted](http://dave.cheney.net/2013/05/11/go-and-juju-at-canonical-slides-posted)
 - [ ] [Curious Channels](http://dave.cheney.net/2013/04/30/curious-channels)
 - [ ] [What is the zero value, and why is it useful ?](http://dave.cheney.net/2013/01/19/what-is-the-zero-value-and-why-is-it-useful)
 - [ ] [Go, the language for emulators](http://dave.cheney.net/2013/01/09/go-the-language-for-emulators)
 - [ ] [Testing Go on the Raspberry Pi running FreeBSD](http://dave.cheney.net/2012/12/31/testing-go-on-the-raspberry-pi-running-freebsd)
 - [ ] [The Go Programming Language (2009)](http://dave.cheney.net/2012/11/18/the-go-programming-language-2009)
 - [ ] [Notes on exploring the compiler flags in the Go compiler suite](http://dave.cheney.net/2012/10/07/notes-on-exploring-the-compiler-flags-in-the-go-compiler-suite)
 - [ ] [Mikio Hara’s ipv4 package](http://dave.cheney.net/2012/09/27/mikio-haras-ipv4-package)
 - [ ] [Installing Go on the Raspberry Pi](http://dave.cheney.net/2012/09/25/installing-go-on-the-raspberry-pi)
 - [ ] [An introduction to cross compilation with Go](http://dave.cheney.net/2012/09/08/an-introduction-to-cross-compilation-with-go)
 - [ ] [Another go at the Next Big Language](http://dave.cheney.net/2012/09/03/another-go-at-the-next-big-language)
 - [ ] [August Go meetup slides](http://dave.cheney.net/2012/08/19/august-go-meetup-slides)
 - [ ] [How the Go language improves expressiveness without sacrificing runtime performance](http://dave.cheney.net/2012/02/11/how-the-go-language-improves-expressiveness-without-sacrificing-runtime-performance)
 - [ ] [Introducing gmx, runtime instrumentation for Go applications](http://dave.cheney.net/2012/02/05/introducing-gmx-runtime-instrumentation-for-go-applications)
 - [ ] [Why Go gets exceptions right](http://dave.cheney.net/2012/01/18/why-go-gets-exceptions-right)
 - [ ] [Three new SSH client features in Go weekly.2011-11-18](http://dave.cheney.net/2011/11/21/three-new-ssh-client-features-in-go-weekly-2011-11-18)
 - [ ] [Scratching my own itch, or how to publish multicast DNS records in Go](http://dave.cheney.net/2011/10/15/scratching-my-own-itch-or-how-to-publish-multicast-dns-records-in-go)
 - [ ] [Simple extended attribute support for Go](http://dave.cheney.net/2011/07/31/simple-extended-attribute-support-for-go)
 - [ ] [Using Clang 2.9 to build Google Go](http://dave.cheney.net/2011/06/07/using-clang-2-9-to-build-google-go)
 - [ ] [Netgear Stora as an ARM development platform](http://dave.cheney.net/2011/03/31/netgear-stora-as-an-arm-development-platform)
 - [ ] [Using Multicast UDP in Go](http://dave.cheney.net/2011/02/19/using-multicast-udp-in-go)
 - [ ] [ack! and Go source files](http://dave.cheney.net/2011/02/11/ack-and-go-source-files)
 - [ ] [How to run godoc under launchd on OS X](http://dave.cheney.net/2010/11/21/how-to-run-godoc-under-launchd-on-os-x)
 - [ ] [How to dial remote SSL/TLS services in Go](http://dave.cheney.net/2010/10/05/how-to-dial-remote-ssltls-services-in-go)

#### Others

 - [ ] [How to use interfaces in Go](http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)


## Wiki

 - [ ] [Learn Concurrency](https://github.com/golang/go/wiki/LearnConcurrency)


## [Talks](https://talks.golang.org/)

#### 2007

 - [ ] [Advanced Topics in Programming Languages: Concurrency/message passing Newsqueak](https://www.youtube.com/watch?v=hB05UFqOtFA)

#### 2009

 - [x] [The Go Programming Language](https://www.youtube.com/watch?v=rKnDgT73v8s)

#### 2010

 - [x] [Another Go at Language Design](https://www.youtube.com/watch?v=7VcArS4Wpqk)
 - [x] [Public Static Void](https://www.youtube.com/watch?v=5kj5ApnhPAE)
 - [ ] [Origins of Go Concurrency Style](https://www.youtube.com/watch?v=3DtUzH3zoFo)

#### 2011

 - [x] [Rob Pike on Google Go: Concurrency, Type System, Memory Management and GC](http://www.infoq.com/interviews/pike-google-go)
 - [ ] [Rob Pike on Parallelism and Concurrency in Programming Languages](http://www.infoq.com/interviews/pike-concurrency)
 - [ ] [Another Go at Language Design](https://www.youtube.com/watch?v=aIgyp5nvdqc)
 - [x] [Writing Web Apps in Go](https://www.youtube.com/watch?v=-i0hat7pdpk)
 - [ ] [Practical Go Programming](https://www.youtube.com/watch?v=2-pPAvqyluI)

#### 2012

 - [ ] [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs) ([slides](http://talks.golang.org/2012/concurrency.slide))
 - [ ] [Go: a simple programming environment](http://vimeo.com/53221558) ([slides](http://talks.golang.org/2012/simple.slide))
 - [ ] [Concurrency Is Not Parallelism](http://vimeo.com/49718712)
 - [x] [Going Go](https://www.youtube.com/watch?v=on5DeUyWDqI)
 - [x] [The path to Go 1](https://www.youtube.com/watch?v=bj9T2c2Xk_s)
 - [x] [Meet the Go Team](https://www.youtube.com/watch?v=sln-gJaURzk)

#### 2013

 - [ ] [Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw) ([slides](http://talks.golang.org/2013/advconc.slide))
 - [ ] [Go at Google](http://www.infoq.com/presentations/Go-Google)
 - [x] [Fireside Chat with the Go Team](https://www.youtube.com/watch?v=p9VUCp98ay4)

#### 2014

 - [x] [Toward Go 1.3, and beyond](https://www.youtube.com/watch?v=mQ4hwLgSvUs)
 - [x] [Opening Keynote by Rob Pike](https://www.youtube.com/watch?v=VoS7DsT1rdM)
 - [x] [Implementing a bignum calculator](https://www.youtube.com/watch?v=PXoG0WX0r_E)

#### 2015

 - [x] [GopherFest 2015](https://www.youtube.com/watch?v=Fx304EfqtMo)
 - [ ] [Go and the Modern Enterprise](https://www.youtube.com/watch?v=iFR_7AKkJFU)
 - [x] [Gophercon - Prometheus: Designing and Implementing a Modern Monitoring Solution in Go](https://www.youtube.com/watch?v=1V7eJ0jN8-E)
 - [x] [GopherCon - Uptime: Building Resilient Services with Go](https://www.youtube.com/watch?v=PyBJQA4clfc)


## Screencasts

 - [x] [Writing, building, installing, and testing Go code](https://www.youtube.com/watch?v=XCsL89YtqCs)


## [Code Walks](http://golang.org/doc/codewalk/)

 - [ ] [Share Memory By Communicating](http://golang.org/doc/codewalk/sharemem/)
 - [ ] [First-Class Functions in Go](http://golang.org/doc/codewalk/functions/)
