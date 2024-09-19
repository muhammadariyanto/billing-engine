package customerModel

import "time"

type Customer struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	IsDelinquent bool `json:"-"`
}
