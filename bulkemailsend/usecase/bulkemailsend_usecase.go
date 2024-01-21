package usecase

import (
	"time"

	"github.com/reinhardjs/sayakaya/domain"
)

type bulkEmailSendUsecase struct {
	bulkEmailSendRepo domain.BulkEmailSendRepository
	contextTimeout    time.Duration
}

// NewBulkEmailSendUsecase will create new an bulkEmailSendUsecase object representation of domain.BulkEmailSendUsecase interface
func NewBulkEmailSendUsecase(repo domain.BulkEmailSendRepository) domain.BulkEmailSendUsecase {
	return &bulkEmailSendUsecase{
		bulkEmailSendRepo: repo,
	}
}

func (a *bulkEmailSendUsecase) BulkSend(bulkEmailSend *domain.BulkEmailSend) (err error) {

	err = a.bulkEmailSendRepo.BulkSend(bulkEmailSend)
	if err != nil {
		return err
	}

	return
}
