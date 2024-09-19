package billingRepository

import (
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	"github.com/muhammadariyanto/billing-engine/internal/repository"
)

type billingRepository struct {
	data map[string]*billingModel.Billing
}

func New() repository.IBillingRepository {
	return &billingRepository{
		data: make(map[string]*billingModel.Billing),
	}
}
