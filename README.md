# Maven Toolbox 🧰

Power packed utilities for maven ⚡

## Features

- Update `~/.m2/settings.xml` with artifactory credintials
- Find and Update `hibernate.cfg.xml`
- Update dynamically generated swagger/openapi SDK for artificatory publishing
- Update properties file

```txt
usage: main [<flags>] <command> [<args> ...]

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --version  Show application version.

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

## Demo

![demo](https://media.giphy.com/media/VCPVjSPn5FVVLqHmC7/giphy.gif)

## Building

```sh
go build -ldflags "-s -w"
```
