set -e

function doBuild {
  cp ~/code/src/github.com/runningwild/glop/gos/darwin/lib/libglop.so darwin/lib/
  go build --tags $1 .
  rm -rf $1.app
  mkdir -p $1.app/Contents/MacOS
  mkdir -p $1.app/Contents/lib
  mv glomple $1.app/Contents/MacOS/$1
  cp darwin/lib/* $1.app/Contents/lib/
  cp -r data/* $1.app/Contents/
}

doBuild 'example'
open example.app
