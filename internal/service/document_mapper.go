package service

import (
	"context"

	gateway "github.com/ecosafety/lec.letter.initial.service/internal/pb/document_generator_api"
	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"

	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"
	"github.com/google/uuid"
)

func mapDocumentListToReq(ctx context.Context, req []*domain.Document) ([]*gateway.DocumentData, error) {
	var docs []*gateway.DocumentData
	for _, doc := range req {
		mDoc, mErr := mapDocumentToReq(ctx, doc)
		if mErr != nil {
			return nil, mErr
		}

		docs = append(docs, mDoc)
	}

	return docs, nil
}

func mapDocumentToReq(_ context.Context, req *domain.Document) (*gateway.DocumentData, error) {
	return &gateway.DocumentData{
		OrderNumber: int32(req.OrderNumber),
		Title:       req.Title,
	}, nil
}

func mapDocumentListToDomain(ctx context.Context, req []*pb.DocumentData) ([]*domain.Document, error) {
	var docs []*domain.Document
	for _, doc := range req {
		mDoc, mErr := mapDocumentToDomain(ctx, doc)
		if mErr != nil {
			return nil, mErr
		}

		docs = append(docs, mDoc)
	}

	return docs, nil
}

func mapDocumentToDomain(_ context.Context, req *pb.DocumentData) (*domain.Document, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	return &domain.Document{
		Id:          id,
		OrderNumber: int(req.GetOrderNumber()),
		Title:       req.GetTitle(),
	}, nil
}
