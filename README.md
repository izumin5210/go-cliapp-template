# go-cliapp-template
[![Build Status](https://travis-ci.org/izumin5210/go-cliapp-template.svg?branch=master)](https://travis-ci.org/izumin5210/go-cliapp-template)
[![Coverage Status](https://coveralls.io/repos/github/izumin5210/go-cliapp-template/badge.svg?branch=master)](https://coveralls.io/github/izumin5210/go-cliapp-template?branch=master)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)][license]

A basic cli application template with Golang.

## Usage

1. Copy this repository
2. Rewrite the following sections in `Makefile` as your application configurations

```make
NAME := awesome-cli
VERSION := 0.1.0
```

3. If you want to use Travis CI, you need to set `COVERALLS_TOKEN` in settings tab on travis-ci.org.

## Acknowledgments
- [Go でツール書くときの Makefile 晒す - Qiita](http://qiita.com/dtan4/items/8c417b629b6b2033a541) (Japanese)
- [みんなのGo言語【現場で使える実践テクニック】](http://amzn.to/2hsvvB3) (Japanese)
- [tcnksm/ghr: Upload multiple artifacts to GitHub Release in parallel](https://github.com/tcnksm/ghr)

## License
Licensed under [MIT License][license].

[license]: https://izumin.mit-license.org/2016
