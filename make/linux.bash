# Run from magnus root directory.
USER=runningwild
PACKAGE=glomple

cd $GOPATH/src/github.com/$USER/$PACKAGE

if go build . ; then
		rm -rf bin
		mkdir bin

		cp $PACKAGE bin/base
		cp ../glop/gos/linux/lib/libglop.so bin/libglop.so
		echo "LD_LIBRARY_PATH=$LD_LIBRARY_PATH:. ./base" > bin/$PACKAGE
		chmod 776 bin/$PACKAGE

		cd bin
		./$PACKAGE
fi
