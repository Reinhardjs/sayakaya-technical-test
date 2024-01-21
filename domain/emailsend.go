package domain

type BulkEmailSend struct {
	Recipients []string
	Subject    string
	Message    string
}

type BulkEmailSendUsecase interface {
	BulkSend(bulkEmailSend *BulkEmailSend) error
}

type BulkEmailSendRepository interface {
	BulkSend(bulkEmailSend *BulkEmailSend) error
}
