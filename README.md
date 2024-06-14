# Extism Go PDK Plugin

See more documentation at https://github.com/extism/go-pdk and
[join us on Discord](https://extism.org/discord) for more help.

This repo demonstrates an unusual interaction (bug?) between the
[Extism JavaScript SDK] and the [Extism Go PDK].

[Extism JavaScript SDK]: https://extism.github.io/js-sdk/
[Extism Go PDK]: https://pkg.go.dev/github.com/extism/go-pdk

## Build

First, build the plugin:

```bash
$ ./build.sh
```

Next, set up a Python local HTTP server:

```bash
$ ./scripts/python-server.sh
```

Finally, open your browser to: http://localhost:8080/examples/count-vowels/
and open the dev tools console, and you will see this:

```
Using simulated Extism SDK...
{"Count":7,"Total":7,"Vowels":"aeiouAEIOU"}
{"Count":7,"Total":14,"Vowels":"aeiouAEIOU"}
Using official Extism JavaScript SDK...
{"Count":10,"Total":10,"Vowels":"aeiouAEIOU"}
{"Count":8,"Total":8,"Vowels":"aeiouAEIOU"}
```

Note that the last line should say `"Total":18`, not `"Total":8`.

Also note that this problem cannot be duplicated using the [MoonBit PDK]
with the Extism JavaScript SDK, so there appears to be some odd interaction
between the Go PDK and JavaScript SDK.

[MoonBit PDK]: https://github.com/extism/moonbit-pdk
