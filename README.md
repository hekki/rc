# rc
Command line tool for radix conversion.

# install
```sh
$ go build
$ mv rc /usr/local/bin/
```

# usage
```
Usage:
  rc [num] [flags]

Flags:
  -b, --binary   binary flag
  -h, --help     help for rc
  -e, --hex      hex flag
```

* Decimal source
```sh
$ rc 254
Decimal: 254
Binary: 11111110
Hex: fe
```

* Binary source
```sh
$ rc -b 10101010
Decimal: 170
Binary: 10101010
Hex: aa
```

* Hex source
```sh
$ rc -e ab
Decimal: 171
Binary: 10101011
Hex: ab
```
