package main

import ("fmt"
"math/rand/v2"
"os"
"io"
"os/exec"
)

func main() {
	fmt.Print(fuzzer(1000, 32, 32))
	generateFile()
	//execute()
}


// cria o arquivo de entradas
func generateFile(){
	f, err := os.Create("testFile.txt")
	if err != nil {
		fmt.Println("Erro")
	}
	defer f.Close()
	fmt.Println(f.Name())
	
	f.Write([]byte(fuzzer(100, 32, 32)))
}

// Gera a saída para o fuzzer
func fuzzer(maxLength int, charStart int, charRange int) string {
	var stringLength = rand.IntN(maxLength)
	var out = ""
	
	//converte um número aleatório em caractere
	for i := 0; i < stringLength; i++ {
		out += string(rune(rand.IntN(charRange) + charStart))
	}
	return out
}


// Não funciona ainda
func execute(){
	
	program := "bc"
	filePath := "input.txt" // Assuming `input.txt` is the file path defined previously
	
	
	// Run the subprocess
	cmd := exec.Command(program, filePath)
	cmd.Stdin = nil // equivalent to subprocess.DEVNULL in Python
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe for Cmd", err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error creating StderrPipe for Cmd", err)
		return
	}
	
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting Cmd", err)
		return
	}
	
	outBytes, err := io.ReadAll(stdout)
	if err != nil {
		fmt.Println("Error reading stdout", err)
		return
	}
	
	errBytes, err := io.ReadAll(stderr)
	if err != nil {
		fmt.Println("Error reading stderr", err)
		return
	}
	
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for Cmd", err)
		return
	}
	
	fmt.Println("Output:", string(outBytes))
	fmt.Println("Errors:", string(errBytes))
}