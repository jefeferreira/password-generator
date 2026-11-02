package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const (
	UppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	NumberChars    = "0123456789"
	SymbolChars    = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

type Config struct {
	Length         int
	IncludeUpper   bool
	IncludeLower   bool
	IncludeNumbers bool
	IncludeSymbols bool
}

func Generate(cfg Config) (string, error) {
	if cfg.Length < 4 {
		return "", errors.New("comprimento minímo é 4 caracteres")
	}

	var charPool strings.Builder
	requiredChars := []byte{}

	if cfg.IncludeUpper {
		charPool.WriteString(UppercaseChars)
		ch, err := randomSecureChar(UppercaseChars)
		if err != nil {
			return "", err
		}
		requiredChars = append(requiredChars, ch)
	}
	
	if cfg.IncludeLower {
		charPool.WriteString(LowercaseChars)
		ch, err := randomSecureChar(LowercaseChars)
		if err != nil {
			return "", err
		}
		requiredChars = append(requiredChars, ch)
	}
	
	if cfg.IncludeNumbers {
		charPool.WriteString(NumberChars)
		ch, err := randomSecureChar(NumberChars)
		if err != nil {
			return "", err
		}
		requiredChars = append(requiredChars, ch)
	}
	
	if cfg.IncludeSymbols {
		charPool.WriteString(SymbolChars)
		ch, err := randomSecureChar(SymbolChars)
		if err != nil {
			return "", err
		}
		requiredChars = append(requiredChars, ch)
	}

	if charPool.Len() == 0 {
		return "", errors.New("pelo menos um tipo de caractere deve ser incluído")
	}

	password := make([]byte, cfg.Length)
	for i, ch := range requiredChars {
		if i >= cfg.Length {
			break
		}
		password[i] = ch
	}

	pool := charPool.String()
	for i := len(requiredChars); i < cfg.Length; i++ {
		ch, err := randomSecureChar(pool)
		if err != nil {
			return "", err
		}
		password[i] = ch
	}

	if err := cryptoShuffle(password); err != nil {
		return "", err
	}

	return string(password), nil
}

func randomSecureChar(s string) (byte, error) {
	if len(s) == 0 {
		return 0, errors.New("string vazia")
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(s))))
	if err != nil {
		return 0, err
	}
	return s[n.Int64()], nil
}

func cryptoShuffle(data []byte) error {
	for i := len(data) - 1; i > 0; i-- {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}
		j := int(n.Int64())
		data[i], data[j] = data[j], data[i]
	}
	return nil
}
