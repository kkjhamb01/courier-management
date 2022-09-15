package api

import (
	"context"

	pb "github.com/kkjhamb01/courier-management/party/proto"
)

func (s *grpcServer) Upload(ctx context.Context, in *pb.UploadDocumentRequest) (*pb.UploadDocumentResponse, error) {
	return s.service.Upload(ctx, in)
}

func (s *grpcServer) GetDocumentsOfUser(ctx context.Context, in *pb.GetDocumentsOfUserRequest) (*pb.GetDocumentsOfUserResponse, error) {
	return s.service.GetDocumentsOfUser(ctx, in)
}

func (s *grpcServer) GetDocument(ctx context.Context, in *pb.GetDocumentRequest) (*pb.GetDocumentResponse, error) {
	return s.service.GetDocument(ctx, in)
}

func (s *grpcServer) Download(ctx context.Context, in *pb.DownloadRequest) (*pb.DownloadResponse, error) {
	return s.service.Download(ctx, in)
}

func (s *grpcServer) DirectDownload(ctx context.Context, in *pb.DirectDownloadRequest) (*pb.DirectDownloadResponse, error) {
	return s.service.DirectDownload(ctx, in)
}
