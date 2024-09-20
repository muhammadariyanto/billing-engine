package customerModel

type CreateCustomerRequest struct {
	Name string `json:"name" validate:"required"`
}
