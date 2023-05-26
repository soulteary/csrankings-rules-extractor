package extractor

import (
	"context"
	"time"

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
