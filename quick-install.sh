#!/usr/bin/env bash

# This script installs maven-toolbox.
#
# Quick install: `curl http://bit.ly/maven-toolbox | bash`
#
# Acknowledgments: Taken from https://getmic.ro/

set -e
set -u
set -o pipefail

function githubLatestTag {
    finalUrl=$(curl "https://github.com/$1/releases/latest" -s -L -I -o /dev/null -w '%{url_effective}')
    echo "${finalUrl##*v}"
}

UNKNOWN_OS_MSG=<<-'EOM'
/=====================================\
|      COULD NOT DETECT PLATFORM      |
\=====================================/

Uh oh! We couldn't automatically detect your operating system. You can file a
bug here: https://github.com/harshzalavadiya/maven-toolbox

To continue with installation, please choose from one of the following values:

- linux64
- osx
- win64
EOM


platform=''
machine=$(uname -m)

if [[ "$OSTYPE" == "linux"* ]]; then
  if [[ "$machine" == "arm"* || "$machine" == "aarch"* ]]; then
    platform='linux-arm'
  elif [[ "$machine" == *"86" ]]; then
    platform='linux32'
  elif [[ "$machine" == *"64" ]]; then
    platform='linux64'
  fi
elif [[ "$OSTYPE" == "darwin"* ]]; then
  platform='osx'
elif [[ "$OSTYPE" == "freebsd"* ]]; then
  if [[ "$machine" == *"64" ]]; then
    platform='freebsd64'
  elif [[ "$machine" == *"86" ]]; then
    platform='freebsd32'
  fi
elif [[ "$OSTYPE" == "openbsd"* ]]; then
  if [[ "$machine" == *"64" ]]; then
    platform='openbsd64'
  elif [[ "$machine" == *"86" ]]; then
    platform='openbsd32'
  fi
elif [[ "$OSTYPE" == "netbsd"* ]]; then
  if [[ "$machine" == *"64" ]]; then
    platform='netbsd64'
  elif [[ "$machine" == *"86" ]]; then
    platform='netbsd32'
  fi
fi

if test "x$platform" = "x"; then
  cat <<EOM
Uh oh! We couldn't automatically detect your operating system. You can file a
bug here: https://github.com/harshzalavadiya/maven-toolbox

To continue with installation, please choose from one of the following values:

- linux64
- osx
- win64
EOM
  read -rp "> " platform
else
  echo "Detected platform: $platform"
fi

TAG=$(githubLatestTag harshzalavadiya/maven-toolbox)

echo "Downloading https://github.com/harshzalavadiya/maven-toolbox/releases/download/v$TAG/maven-toolbox-$platform.tar.xz"
curl -L "https://github.com/harshzalavadiya/maven-toolbox/releases/download/v$TAG/maven-toolbox-$platform.tar.xz" > maven-toolbox.tar.xz

tar xf maven-toolbox.tar.xz "maven-toolbox-$platform"

mkdir -p @ci
mv "maven-toolbox-$platform" ./@ci/maven-toolbox
chmod +x ./@ci/maven-toolbox

rm maven-toolbox.tar.xz

echo "✨ maven-toolbox-$TAG Downloaded"