package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jefeferreira/password-generator/generator"
)

func main() {
	length := flag.Int("length", 16, "Comprimento da senha")
	uppercase := flag.Bool("uppercase", true, "Incluir letras maiúsculas")
	lowercase := flag.Bool("lowercase", true, "Incluir letras minúsculas")
	numbers := flag.Bool("numbers", true, "Incluir números")
	symbols := flag.Bool("symbols", true, "Incluir caracteres especiais")
	count := flag.Int("count", 1, "Número de senhas a gerar")
	flag.Parse()

	cfg := generator.Config{
		Length:         *length,
		IncludeUpper:   *uppercase,
		IncludeLower:   *lowercase,
		IncludeNumbers: *numbers,
		IncludeSymbols: *symbols,
	}

	for i := 0; i < *count; i++ {
		password, err := generator.Generate(cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao gerar senha: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Senha %d: %s\n", i+1, password)
	}
}
