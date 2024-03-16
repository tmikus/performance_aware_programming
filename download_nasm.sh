NASM_VERSION="2.16.01"
NASM_FILENAME="nasm-$NASM_VERSION-macosx.zip"
curl -o "$NASM_FILENAME" https://www.nasm.us/pub/nasm/releasebuilds/$NASM_VERSION/macosx/$NASM_FILENAME
unzip $NASM_FILENAME
rm $NASM_FILENAME
mv "nasm-$NASM_VERSION" nasm
