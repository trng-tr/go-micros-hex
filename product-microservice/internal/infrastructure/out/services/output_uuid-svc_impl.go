package services

import (
	"fmt"

	"github.com/google/uuid"
)

// OutUuidGeneratorServiceImpl implemnt OutUuidGeneratorService
type OutUuidGeneratorServiceImpl struct {
}

// NewOutGenerateSkuServiceImpl construtor func
func NewOutUuidGeneratorServiceImpl() *OutUuidGeneratorServiceImpl {
	return &OutUuidGeneratorServiceImpl{}
}

// GenerateProductSku implement interface OutUuidGeneratorService
func (o *OutUuidGeneratorServiceImpl) GenerateUuid() string {
	var uu string = uuid.NewString()
	return fmt.Sprintf("%s-%s", uu[:4], uu[4:8])
}
