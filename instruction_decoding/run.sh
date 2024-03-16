echo "Compiling asm files..."
(cd examples && ../../nasm/nasm simple.asm)
(cd examples && ../../nasm/nasm complex.asm)

echo "Running the application..."
go run .