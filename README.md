
![google translate logo](https://resize.indiatvnews.com/en/resize/newbucket/1200_-/2019/12/google-translate-1575354629.jpg)

# GoogletransX

[![language](https://img.shields.io/badge/language-Golang-blue)](https://golang.org/)
[![Documentation](https://godoc.org/github.com//yuriizinets/googletransx?status.svg)](https://godoc.org/github.com/mind1949/googletrans)
[![Go Report Card](https://goreportcard.com/badge/github.com/yuriizinets/googletransx)](https://goreportcard.com/report/github.com/mind1949/googletrans)

This is fork of mind1949/googletrans with extended features:

* BulkTranslate with goroutines processing
* HTML translate support
* Built-in API server for integrations
* Dockerized server

## Installation

```bash
$ go get -u github.com/yuriizinets/googletransx
...
```

## Usage

For usage doc of basic operations, check original [repo](https://github.com/mind1949/googletrans)  
Here will be mentioned only extended features

HTML Mime-Type

```go
package main

import (
    "github.com/yuriizinets/googletransx"
)

func main() {
    input := "<div>Test<span>translate</span></div>"
    result, err := googletransx.Translate(googletransx.TranslateParams{
        Text: input,
        Src: "auto",
        Dest: "zh-CN",
        MimeType: "text/html",
    })
    fmt.Println(result) // Translated{} struct with translated html
}
```

BulkTranslate (HTML is supported too)

```go
package main

import (
    "github.com/yuriizinets/googletransx"
)

func main() {
    inputs := []string{"example", "test"}
    params := []googletransx.TranslateParams{}
    for _, input := range inputs {
        params = append(params, googletransx.TranslateParams{
            Text: input,
            Src: "auto",
            Dest: "zh-CN",
        })
    }
    results, err := googletransx.BulkTranslate(params)
    if err != nil {
        panic(err)
    }
    fmt.Println(results) // Results are []Translated
}
```

API Server (native)

```bash
$ go run github.com/yuriizinets/googletransx/server.runtime
Starting server at :25021
```

API Server (docker)

```bash
$ docker pull docker.pkg.github.com/yuriizinets/googletransx/googletransx:latest
$ docker run -p 25021:25021 googletransx
Starting server at :25021
```

## Known issues

* `socket: too many open files` too many connections while making BulkTranslate. Can be temporary fixed with ulimit -Sn 100000
* Pronunciation is not working
* Somethings not working? Just create new issue
