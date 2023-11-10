package keyinp

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	inputStr, er := reader.ReadString('\n')
	if er != nil {
		return 0, er
	}
	inputStr = strings.TrimSpace(inputStr)
	inputStr = strings.Replace(inputStr, ",", ".", 1)
	inputFloat, er := strconv.ParseFloat(inputStr, 64)
	if er != nil {
		return 0, er
	}
	return inputFloat, nil
}
