package extractor

import (
	"context"
	"strings"
	"time"

	"github.com/dop251/goja"
	pb "github.com/soulteary/csrankings-rules-extractor/internal/extractor/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetPythonAST(addr string, code string) (string, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	c := pb.NewConverterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.PythonAST(ctx, &pb.ConvertRequest{Code: code})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}

func ParsePyResponse(data string, js []byte) (string, error) {
	vm := goja.New()
	vm.Set("data", data)

	v, err := vm.RunString(string(js))
	if err != nil {
		return "", err
	}
	return v.Export().(string), nil
}

func GetPyConfig(addr string, buf []byte, js []byte) (string, error) {
	// trick: Comment out the code that is irrelevant to the parsing configuration to avoid errors in AST parsing (lower versions of Python)
	code := strings.ReplaceAll(string(buf), `if (pvmatcher := TECSCounterColon.match(pages)):`, "#")
	python, err := GetPythonAST(addr, code)
	if err != nil {
		return "", err
	}

	result, err := ParsePyResponse(python, js)
	if err != nil {
		return "", err
	}
	return result, nil
}
