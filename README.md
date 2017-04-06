# gopostal

[![Build Status](https://travis-ci.org/openvenues/gopostal.svg?branch=master)](https://travis-ci.org/openvenues/gopostal)

Go/cgo interface to [libpostal](https://github.com/openvenues/libpostal), a C library for fast international street address parsing and normalization.

## Usage

To expand address strings into normalized forms suitable for geocoder queries:

```go
package main

import (
    "fmt"
    expand "github.com/openvenues/gopostal/expand"
)

func main() {
    expansions := expand.ExpandAddress("Quatre-vingt-douze Ave des Ave des Champs-Élysées")

    for i := 0; i < len(expansions); i++ {
        fmt.Println(expansions[i])
    }
}
```

To parse addresses into components:

```go
package main

import (
    "fmt"
    parser "github.com/openvenues/gopostal/parser"
)

func main() {
    parsed := parser.ParseAddress("781 Franklin Ave Crown Heights Brooklyn NY 11216 USA")
    fmt.Println(parsed)
}
```

## Prerequisites

Before using the Go bindings, you must install the libpostal C library. Make sure you have the following prerequisites:

**On Ubuntu/Debian**
```
sudo apt-get install curl autoconf automake libtool pkg-config
```

**On CentOS/RHEL**
```
sudo yum install curl autoconf automake libtool pkgconfig
```

**On Mac OSX**
```
sudo brew install curl autoconf automake libtool pkg-config
```

**Installing libpostal**

```
git clone https://github.com/openvenues/libpostal
cd libpostal
./bootstrap.sh
./configure --datadir=[...some dir with a few GB of space...]
make
sudo make install

# On Linux it's probably a good idea to run
sudo ldconfig
```

## Installation

For expansions:

```
go get github.com/openvenues/gopostal/expand
```

For parsing:
```
go get github.com/openvenues/gopostal/parser
```

## Tests

```
go test github.com/openvenues/gopostal/...
```
