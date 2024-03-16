package main

import "os"

//const inputPath = "examples/simple"

const inputPath = "examples/complex"

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, stat.Size())
	_, err = f.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func main() {
	data, err := readFile(inputPath)
	if err != nil {
		println(err.Error())
		return
	}
	parser := NewAsmParser(data)
	asm := parser.Parse()
	println(asm)
}
