# Maven Toolbox 🧰

Power packed utilities for maven ⚡

## Features

- Built for CI-CD 💫
- No Dependency 😇
- Cross Platform 🐿️

## How to use

- Download platform specific binary release from [here](https://github.com/harshzalavadiya/maven-toolbox/releases)
- Copy executable to root directory
- Setup environment variables (for example `export MTPROP_FOO=bar`)
- Run executable with specific arguements from below

## Actions

#### `configure-m2`

Update `~/.m2/settings.xml` with artifactory credintials

```sh
export ARTIFACTORY_ENDPOINT=admin
export ARTIFACTORY_USERNAME=password
export ARTIFACTORY_URL=http://localhost:8081/artifactory
./maven-toolbox configure-m2
```

#### `configure-sdk`

Update dynamically generated swagger/openapi SDK project for artificatory publishing

```sh
export ARTIFACTORY_URL=http://localhost:8081/artifactory
./maven-toolbox configure-sdk
```

#### `configure-hibernate`

find and update database configuration `hibernate.cfg.xml`

```sh
export DB_USERNAME="postgres"
export DB_PASSWORD="postgres123"
export DB_URL="jdbc:postgresql://localhost:5432/db"
./maven-toolbox configure-hibernate
```

#### `configure-properties`

Update properties file

```sh
export MTPROP_FOO="xyz" # `MTPROP_` is a default prefix
./maven-toolbox configure-properties path/to/your.properties
```

<details><summary>Example</summary>
<p>

![demo](https://media.giphy.com/media/VCPVjSPn5FVVLqHmC7/giphy.gif)

</p>
</details>

## Building

```sh
go build -ldflags "-s -w"
```
