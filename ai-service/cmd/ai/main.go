package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Konstantin-Korolyov/go-support-assistant/proto/support"
)

type server struct {
	pb.UnimplementedAIServiceServer
}

func (s *server) ClassifyMessage(ctx context.Context, req *pb.ClassifyRequest) (*pb.ClassifyResponse, error) {
	// Заглушка: возвращаем тестовую категорию
	return &pb.ClassifyResponse{
		Category:   "general",
		Confidence: 0.8,
		ExtractedEntities: map[string]string{
			"order_id": "12345",
		},
	}, nil
}

func (s *server) GenerateReply(ctx context.Context, req *pb.GenerateReplyRequest) (*pb.GenerateReplyResponse, error) {
	// Заглушка
	return &pb.GenerateReplyResponse{
		ReplyText:     "Это автоматический ответ. Скоро с вами свяжется оператор.",
		Confidence:    0.6,
		RequiresHuman: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAIServiceServer(s, &server{})

	log.Println("AI Service gRPC server starting on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
