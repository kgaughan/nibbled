# gribbled

gribbled is meant to be a fun exercise so that I can play around with Go
properly for the first time.

The nice thing about Gopher as a protocol for an exercise like this is that
it's a relatively simple protocol and easy to implement. It will also require
getting familiar with the many parts of the language and its standard libraries
so that it can be implemented, making it an ideal candidate for implementation.

The design is intended to be similar to nibbled, another Gopher client I'm
planning on writing as a counterpoint. The repository for nibbled is here:

    https://github.com/kgaughan/nibbled

The single biggest difference between the two from the outside will be that the
Python daemon stores its configuration in a configuration file, whereas the Go
daemon uses command line arguments for the same. This is maninly because (a) I
didn't want to deal with extra dependencies and (b) I didn't want to write my
own config parser. I could've used the built-in JSON or XML parsers, but
neither of those are particularly writer-friendly formats.
