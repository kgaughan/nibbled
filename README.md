# nibbled

nibbled is a [Gopher][] daemon.

This is yet another COVID-19 project, reviving an old project I'd intended on starting when I first started using Go in anger.

It has partial [gophermap][] support, but only supports a subset of [Gophernicus][]'s functionality. It supports `#`, `.`, and `*`.

## Running under Docker

By default, nibbled will serve from `/srv/gopher`. The most straightforward way to get up and running is to mount a local directory at that path:

```sh
docker run -d -p 70:70 -v /path/to/your/gopherhole:/srv/gopher ghcr.io/kgaughan/nibbled:latest
```

[Gopher]: https://tools.ietf.org/html/rfc1436
[Gopher URI scheme]: https://tools.ietf.org/html/rfc4266
[GopherII]: https://www.ietf.org/archive/id/draft-matavka-gopher-ii-03.txt
[Gophernicus]: https://github.com/gophernicus/gophernicus/
[gophermap]: https://github.com/gophernicus/gophernicus/blob/master/README.gophermap
[goth]: https://gitlab.com/commonshost/goth
[Gemini]: https://gemini.circumlunar.space/
