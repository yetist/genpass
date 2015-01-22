#!/bin/bash
eval $(grep "PkgName.*=" *.go|tr -d [:space:])
eval $(grep "PkgVersion.*=" *.go|tr -d [:space:])
NV=${PkgName}-${PkgVersion}
olddir=$(realpath $PWD)

if [ ! -f $PkgName ]; then
	godep go build
fi

if [ -f $PkgName ]; then
	TMP=`mktemp -d`
	DEST=${TMP}/$NV
	mkdir -p ${DEST}
	install -m755 genpass ${DEST}/genpass
	cp -r templates ${DEST}
	cp -r locale ${DEST}
	cp -r static ${DEST}
	cd ${TMP}
	tar -zcf $olddir/$NV.tar.gz *
	echo "$NV.tar.gz" is ready
	rm -rf ${TMP}
	cd $olddir
fi
