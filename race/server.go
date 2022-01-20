package race

import (
	"context"
	"crypto/md5"

	b64 "encoding/base64"

	pb "redbelly.grpcrace/proto"
)

type Server struct {
	pb.UnimplementedChecksumServiceServer
}

func (s Server) Checksum(ctx context.Context, in *pb.CheckRequest) (*pb.CheckResponse, error) {
	result := md5.Sum([]byte(in.GetPayload()))
	return &pb.CheckResponse{Sum: b64.StdEncoding.EncodeToString(result[:]), Size: int32(len(in.GetPayload()))}, nil
}
