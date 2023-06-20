package service

import (
	"context"
	"github.com/ecosafety/lec.letter.initial.service/internal/gateway/document_generator"

	"github.com/ecosafety/lec.letter.initial.service/internal/storage/memory/letter_memory"
	"github.com/ecosafety/lec.letter.initial.service/internal/storage/memory/objective_memory"
	"github.com/ecosafety/lec.letter.initial.service/internal/storage/memory/purpose_memory"

	"github.com/ecosafety/lec.letter.initial.service/internal/storage/pg_letter"
	"github.com/ecosafety/lec.letter.initial.service/internal/storage/pg_objective"
	"github.com/ecosafety/lec.letter.initial.service/internal/storage/pg_purpose"

	"log"
)

type Service struct {
	letterStorage     *pg_letter.Storage
	letterMemCache    *letter_memory.Cache
	objectiveStorage  *pg_objective.Storage
	objectiveMemCache *objective_memory.Cache
	purposeStorage    *pg_purpose.Storage
	purposeMemCache   *purpose_memory.Cache
	generatorC        *document_generator.Client
}

func NewService(ctx context.Context, letterStorage *pg_letter.Storage, letterMemCache *letter_memory.Cache,
	objectiveStorage *pg_objective.Storage, objectiveMemCache *objective_memory.Cache,
	purposeStorage *pg_purpose.Storage, purposeMemCache *purpose_memory.Cache,
	generatorC *document_generator.Client) *Service {

	letters, err := letterStorage.GetList(ctx)
	if err != nil {
		log.Fatalf("couldn't create service: %s", err.Error())
	}
	letterMemCache.SetData(letters)

	objectives, err := objectiveStorage.GetList(ctx)
	if err != nil {
		log.Fatalf("couldn't create service: %s", err.Error())
	}
	objectiveMemCache.SetData(objectives)

	purposes, err := purposeStorage.GetList(ctx)
	if err != nil {
		log.Fatalf("couldn't create service: %s", err.Error())
	}
	purposeMemCache.SetData(purposes)

	return &Service{
		letterStorage:     letterStorage,
		letterMemCache:    letterMemCache,
		objectiveStorage:  objectiveStorage,
		objectiveMemCache: objectiveMemCache,
		purposeStorage:    purposeStorage,
		purposeMemCache:   purposeMemCache,
		generatorC:        generatorC,
	}
}
