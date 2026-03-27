package service

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

const codeCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

type CodeGenerator interface {
	GenerateCode(length int) (string, error)
}

type codeGenerator struct{}

func NewCodeGenerator() CodeGenerator {
	return &codeGenerator{}
}

func (g *codeGenerator) GenerateCode(length int) (string, error) {
	var b bytes.Buffer
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(codeCharset))))
		if err != nil {
			return "", err
		}
		b.WriteByte(codeCharset[idx.Int64()])
	}
	return b.String(), nil
}
