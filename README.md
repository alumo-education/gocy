```
  __ _  ___   __ _ _   _
 / _` |/ _ \ / _` | | | |
| (_| | (_) | (_| | |_| |
 \__, |\___/ \__, |\__,_|
 |___/          |_|
```
[![Go Report Card](https://goreportcard.com/badge/github.com/alumo-education/gocy)](https://goreportcard.com/report/github.com/alumo-education/gocy)

`gocy` is an expressive Cypher builder

* [Basics](#basics)
* [Expressions](#expressions)
    * [Complex Example](#complex-example)
* [Querying](#querying)
    * [Dataset](#dataset)
        * [Prepared Statements](#dataset_prepared)
    * [Database](#database)
    * [Transactions](#transactions)
* [Logging](#logging)
* [Adapters](#adapters)
* [Contributions](#contributions)
* [Changelog](https://github.com/doug-martin/goqu/tree/master/HISTORY.md)

This library was built with the following goals:

* Make the generation of Cypher easy and enjoyable
* Provide a DSL that accounts for the common Cypher expressions, NOT every nuance of the language.
* Allow users to use Cypher when desired
* Provide a simple query API for querying nodes

## Features

TODO


## Installation

```sh
go get -u github.com/alumo-education/gocy
```


<a name="basics"></a>
## Basics

TODO

<a name="querying"></a>
## Querying

TODO

<a name="logging"></a>
## Logging

TODO

<a name="adapters"></a>
## Adapters

TODO

<a name="contributions"></a>
## Contributions

I am always welcoming contributions of any type. Please open an issue or create a PR if you find an issue with any of the following.

* An issue with Documentation
* You found the documentation lacking in some way

If you have an issue with the package please include the following

* A description of the problem
* A short example of how to reproduce (if applicable)

Without those basics it can be difficult to reproduce your issue locally. You may be asked for more information but that is a good starting point.

### New Features

New features and/or enhancements are great and I encourage you to either submit a PR or create an issue. In both cases include the following as the need/requirement may not be readily apparent.

1. The use case
2. A short example

If you are issuing a PR also also include the following

1. Tests - otherwise the PR will not be merged
2. Documentation - otherwise the PR will not be merged
3. Examples - [If applicable] see example_test.go for examples

If you find an issue you want to work on please comment on it letting other people know you are looking at it and I will assign the issue to you.

If want to work on an issue but dont know where to start just leave a comment and I'll be more than happy to point you in the right direction.

### Running tests

TODO

## License

`gocy` is released under the [MIT License](http://www.opensource.org/licenses/MIT).
