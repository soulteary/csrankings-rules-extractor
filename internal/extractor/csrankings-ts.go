package extractor

import (
	"os"
	"strings"

	"github.com/dop251/goja"
)

func GetCSRankingsTS(filename string) (string, error) {
	buf, _ := os.ReadFile(filename)
	vm := goja.New()

	csrankingsTS, _ := os.ReadFile("data/emeryberger/CSrankings/csrankings.ts")
	vm.Set("code", string(csrankingsTS))

	v, err := vm.RunString(strings.Join([]string{string(buf), "GetCSRankingsConfig(code)"}, "\n"))
	if err != nil {
		return "", err
	}

	// TODO parse to json
	return v.Export().(string), nil
}
