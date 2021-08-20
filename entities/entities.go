package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

type Product struct {
	ProductID    uuid.UUID `json:"product_id" gorm:"primary_key; unique;type:uuid; column:id_request;default:uuid_generate_v4()"`
	ProductName  string    `json:"product_name" validate:"required,numeric"`
	Price        float64   `json:"price" validate:"required,numeric"`
	CounterOrder int32     `json:"counter_order"`
}

type Member struct {
	MemberID   uuid.UUID `json:"member_id" gorm:"primary_key; unique;type:uuid; column:id_request;default:uuid_generate_v4()"`
	MemberName string    `json:"member_name" validate:"required"`
}

type Payment struct {
	PaymentID     uuid.UUID `json:"payment_id" gorm:"primary_key; unique;type:uuid; column:id_request;default:uuid_generate_v4()"`
	ProductID     string    `json:"product_id" validate:"required"`
	ProductName   string    `json:"product_name"`
	MemberID      string    `json:"member_id" validate:"required"`
	MemberName    string    `json:"member_name"`
	Qty           int32     `json:"qty" validate:"required,numeric"`
	Price         float64   `json:"price"`
	FullPrice     float64   `json:"full_price"`
	StatusPayment string    `json:"status_payment"`
}

type Paid struct {
	PaymentID     string `json:"payment_id" validate:"required"`
	MemberID      string `json:"member_id" validate:"required"`
	StatusPayment string `json:"status_payment" validate:"required"`
}

type TemporalMemberRequest struct {
	Times        time.Time
	TypesTimes   string
	Task         string
	TaskType     string
	Data         *Member
	WorkflowName string
	Retry        int32
	Attempt      int32
	Interval     time.Duration
	MaxInterval  time.Duration
}

type TemporalOrderRequest struct {
	Times        time.Time
	TypesTimes   string
	Task         string
	TaskType     string
	Data         *Payment
	WorkflowName string
	Retry        int32
	Attempt      int32
	Interval     time.Duration
	MaxInterval  time.Duration
}

type TemporalPaymentRequest struct {
	Times        time.Time
	TypesTimes   string
	Task         string
	TaskType     string
	Data         *Paid
	WorkflowName string
	Retry        int32
	Attempt      int32
	Interval     time.Duration
	MaxInterval  time.Duration
}
