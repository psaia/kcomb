# kcomb

kcomb implements a variation of [n choose k](https://en.wikipedia.org/wiki/Binomial_coefficient) in order
to compute every possible distinct combination of values within a series of columns.

This could be useful when doing something like generating every possible template string given a set of 
values for each variable in the template.

```
go get -u github.com/psaia/kcomb
```

See [tests](kcomb_test.go) and the [demo](demo/main.go) for usage and benchmarking.

![data stream](demo/fruit.gif)
