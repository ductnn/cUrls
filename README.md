# cUrls

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/ductnn/cUrls/pulls)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**cUrls** is a simple tool crawl urls from domain using [colly](https://github.com/gocolly/colly)
library.

## Installation

First, install [golang](https://go.dev/doc/install)

Then, clone from soure code and install:

```sh
git clone https://github.com/ductnn/cUrls.git
cd cUrls
go get
```

## Usage

Run command:

```
go run curls.go > sub.txt
# Enter domain you want to crawl.
# Example
http://httpbin.org/
```

Check results in file *sub.txt*:

```txt
Visiting http://httpbin.org/
Link found: "\n        \n            \n            \n            \n        \n    " -> https://github.com/requests/httpbin
Link found: "the developer - Website" -> https://kennethreitz.org
Link found: "Send email to the developer" -> mailto:me@kennethreitz.org
Link found: "Flasgger" -> https://github.com/rochacbruno/flasgger
Link found: "HTML form" -> /forms/post
```

So done !!! =))))

### Show your support
Give a ⭐ if you like this application ❤️

## Contribution
Contributions are more than welcome in this project!

## License
The MIT License (MIT). Please see [LICENSE](license) for more information.
