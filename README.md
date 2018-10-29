# swagger-version-manager

_If you're using swagger-codegen I'd recommend giving OpenAPI Codegen a try. I've written `openapi-version-manager` as well ([link](https://github.com/place1/openapi-version-manager))_

## What's this?
`swagger-version-manager` is a tiny CLI tool to help you manage different versions of `swagger-codegen` on your local
machine.

```bash
$ swagger-version-manager --help
NAME:
   swagger-version-manager - A new cli application

USAGE:
   swagger-version-manager-darwin-amd64 [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     current  show the current swagger codegen version
     list     list available swagger codegen versions
     use      use the specified swagger-codegen-version
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

**list available swagger versions**
```bash
$ swagger-version-manager list
3.0.0-rc1
3.0.0-rc0
2.3.1
2.3.0
2.2.3
2.2.2
2.2.1
2.2.0
2.1.6
2.1.5
```

**use a specific swagger-codegen version**
```bash
$ swagger-version-manager use 2.2.3
downloading http://search.maven.org/remotecontent?filepath=io/swagger/swagger-codegen-cli/2.2.3/swagger-codegen-cli-2.2.3.jar
 3.25 MiB / 13.19 MiB [===================>                                                                     ]  25% 00m01
```
```bash
$ swagger-codegen version
2.2.3
```

## Installation

### MacOS
```bash
curl -Lo swagger-version-manager https://github.com/Place1/swagger-version-manager/releases/download/v1.0.0/swagger-version-manager-darwin-amd64
chmod +x ./swagger-version-manager
mv ./swagger-version-manager /usr/local/bin/
swagger-version-manager --help
```

### Linux
```bash
curl -Lo swagger-version-manager https://github.com/Place1/swagger-version-manager/releases/download/v1.0.0/swagger-version-manager-linux-amd64
chmod +x ./swagger-version-manager
mv ./swagger-version-manager /usr/local/bin/
swagger-version-manager --help
```

### Windows
1. Download the windows binary from the release page
2. rename it to `swagger-version-manager.exe`
3. execute it using powershell
    - `swagger-version-manager` required admin privileges on windows because it writes the `swagger-codegen`
      executable to `C:\Windows\System32\swagger-codegen`. I'm not familiar with windows but i'd love
      a PR that changes this behaviour so that admin privileges are not required.

## OpenAPI
If you found this tool useful, you may also be interested in the OpenAPI version: https://github.com/place1/openapi-version-manager
