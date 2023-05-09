#!/bin/bash

rm -rf ./lodepng.h
rm -rf ./lodepng.c
rm -rf ./libimagequant.h
rm -rf ./libimagequant.a

rm -rf ./__MACOSX

cd ./lib

unzip ./libimagequant-main.zip

cd ./libimagequant-main/imagequant-sys
make static

cd ..

cargo build --release

if [ $? -ne '0' ]; then
  echo 'Build libimagequant Failed'
  exit
else
  echo 'Build libimagequant Success'
fi

cp ./imagequant-sys/libimagequant.h ../../libimagequant.h
cp ./target/release/libimagequant_sys.a ../../libimagequant.a

cd ..

rm -rf ./libimagequant-main

cp ./lodepng.c ../lodepng.c
cp ./lodepng.h ../lodepng.h

cd ..