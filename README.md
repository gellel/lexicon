# Hashtable
Hashtable is a [Go](https://github.com/golang/go) package that provides a generic hashtable with extended functionality. It abstracts common map operations, such as adding, deleting, iterating, and more, making it easier to work with maps in Go.

![Hashtable]()

[![PkgGoDev](https://pkg.go.dev/badge/github.com/lindsaygelle/hashtable)](https://pkg.go.dev/github.com/lindsaygelle/hashtable)
[![Go Report Card](https://goreportcard.com/badge/github.com/lindsaygelle/hashtable)](https://goreportcard.com/report/github.com/lindsaygelle/hashtable)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/lindsaygelle/hashtable)](https://github.com/lindsaygelle/hashtable/releases)
[![GitHub](https://img.shields.io/github/license/lindsaygelle/hashtable)](LICENSE.txt)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)

## Features

## Installation
You can install it in your Go project using `go get`:

```sh
go get github.com/lindsaygelle/hashtable
```

## Usage
Import the package into your Go code:

```Go
import (
	"github.com/lindsaygelle/hashtable"
)
```

## Methods
Provided methods for `&hashtable.Hashtable[K]V`.


## Examples


## Docker
A [Dockerfile](./Dockerfile) is provided for individuals that prefer containerized development.

### Building
Building the Docker container:
```sh
docker build . -t hashtable
```

### Running
Developing and running Go within the Docker container:
```sh
docker run -it --rm --name hashtable hashtable
```

## Docker Compose
A [docker-compose](./docker-compose.yml) file has also been included for convenience:
### Running
Running the compose file.
```sh
docker-compose up -d
```

## Contributing
We warmly welcome contributions to Hashtable. Whether you have innovative ideas, bug reports, or enhancements in mind, please share them with us by submitting GitHub issues or creating pull requests. For substantial contributions, it's a good practice to start a discussion by creating an issue to ensure alignment with the project's goals and direction. Refer to the [CONTRIBUTING](./CONTRIBUTING.md) file for comprehensive details.

## Branching
For a smooth collaboration experience, we have established branch naming conventions and guidelines. Please consult the [BRANCH_NAMING_CONVENTION](./BRANCH_NAMING_CONVENTION.md) document for comprehensive information and best practices.

## License
Hashtable is released under the MIT License, granting you the freedom to use, modify, and distribute the code within this repository in accordance with the terms of the license. For additional information, please review the [LICENSE](./LICENSE) file.

## Security
If you discover a security vulnerability within this project, please consult the [SECURITY](./SECURITY.md) document for information and next steps.

## Code Of Conduct
This project has adopted the [Amazon Open Source Code of Conduct](https://aws.github.io/code-of-conduct). For additional information, please review the [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md) file.

## Acknowledgements
