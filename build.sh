#!/bin/bash

TARGET=$1

if [ -n ${TARGET} ]; then
  rustup target add ${TARGET}
fi

rm -rf ./lodepng.h
rm -rf ./libimagequant.h
rm -rf ./libimagequant.a
rm -rf ./lib/lodepng.o
rm -rf ./lib/libimagequant-main
rm -rf ./liblodepng.a

rm -rf ./__MACOSX

cd ./lib

unzip ./libimagequant-main.zip

cd ./libimagequant-main/imagequant-sys

if [ -n ${TARGET} ]; then
  cargo build --release --target ${TARGET}
else
  cargo build --release
fi

cd ..

if [ $? -ne '0' ]; then
  echo 'Build libimagequant Failed'
  exit
else
  echo 'Build libimagequant Success'
fi

cp ./imagequant-sys/libimagequant.h ../../libimagequant.h

if [ -n ${TARGET} ]; then
cp ./target/${TARGET}/release/libimagequant_sys.a ../../libimagequant.a
else
cp ./target/release/libimagequant_sys.a ../../libimagequant.a
fi

cd ..

rm -rf ./libimagequant-main

gcc -c lodepng.c
ar rcs lodepng.a lodepng.o
mv lodepng.a ../liblodepng.a

rm -rf ./lodepng.o

cp ./lodepng.h ../lodepng.h

cd ..