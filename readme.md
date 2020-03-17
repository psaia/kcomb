# kcomb

Compute (and stream/generate) every possible combination within N sets of data using an implementation of [n choose k](https://en.wikipedia.org/wiki/Binomial_coefficient).

This could be useful when doing something like generating every possible template string given a set of 
values for each variable in the template. This works even when the number of variables per template is unknown.

When working with larger sets, this can easily result in millions of results. The `CombineGenerator` func will allow you to create a pipeline and efficently iterate without eating much CPU or memory.

See [tests](kcomb_test.go) and the [demo](demo/main.go) for usage and benchmarking.


```
go get -u github.com/psaia/kcomb
```

![data stream](demo/fruit.gif)
