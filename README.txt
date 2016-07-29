Based on https://github.com/bradfitz/iter

Now supports more types.

Added empty interfaces, reflection and an allocation for good measure.

Because sometimes you want to iterate over a *testing.B without having to type `b.N` or the length of a buffered channel without having to type `len(ch)`.

See https://godoc.org/github.com/voutasaurus/iter
