package qrcode

import (
	"io"

	"github.com/skip2/go-qrcode"
)

type RecoveryLevel = qrcode.RecoveryLevel

const (
	Low RecoveryLevel = iota
	Medium
	High
	Highest
)

func Generate(content string, level RecoveryLevel, writer io.Writer) error {
	code, err := qrcode.New(content, level)
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte(code.ToSmallString(false)))
	return err
}
