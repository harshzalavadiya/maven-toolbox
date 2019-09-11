# Maven Toolbox 🧰

Power packed utilities for maven ⚡

![demo](https://media.giphy.com/media/VCPVjSPn5FVVLqHmC7/giphy.gif)

## Features

- Update `~/.m2/settings.xml` with artifactory credintials
- Find and Update `hibernate.cfg.xml`
- Update dynamically generated swagger/openapi SDK for artificatory publishing
- Update properties file

## How to use

- Download platform specific binary release from [here](https://github.com/harshzalavadiya/maven-toolbox/releases)
- Copy executable to root directory
- Setup environment variables (for example `export MTPROP_FOO=bar`)
- Run executable with specific arguements from below

## Documentation

```txt
usage: main [<flags>] <command> [<args> ...]

Commands:
  help [<command>...]
    Show help.

  configure-m2
    Update '~/.m2/settings.xml'

  configure-sdk
    Update 'target/sdk/pom.xml'

  configure-hibernate
    Finds and updates 'hibernate.cfg.xml'

  configure-properties <file> [<prefix>]
    Update properties file
      Args:
      <file>    relative properties file path
      <prefix>  environment variable prefix (prefix will be stripped with underscore, case insensitive)
```

## Building

```sh
go build -ldflags "-s -w"
```
