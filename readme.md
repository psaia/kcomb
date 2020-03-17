# kcomb

![data stream](demo/fruit.gif)

kcomb implements a variation of [n choose k](https://en.wikipedia.org/wiki/Binomial_coefficient) in order
to compute every possible distinct combination of values with respect to column (or set) of data belong to.

This could be useful when doing something like generating every possible template string given a set of 
values for each variable in the template.

```
go get -u github.com/psaia/kcomb
```

See [tests](kcomb_test.go) for usage and benchmarking.
