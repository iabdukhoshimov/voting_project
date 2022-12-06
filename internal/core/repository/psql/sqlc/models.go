// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package sqlc

import (
	"time"

	null "gopkg.in/guregu/null.v4"
)

type Board struct {
	// uuid.new is used
	ID              string      `json:"id"`
	Title           null.String `json:"title"`
	Icon            null.String `json:"icon"`
	TotalAmount     null.Int    `json:"total_amount"`
	AcceptStartDate time.Time   `json:"accept_start_date"`
	AcceptEndDate   time.Time   `json:"accept_end_date"`
	ReviewStartDate time.Time   `json:"review_start_date"`
	ReviewEndDate   time.Time   `json:"review_end_date"`
	VotingStartDate time.Time   `json:"voting_start_date"`
	VotingEndDate   time.Time   `json:"voting_end_date"`
	CreatedAt       null.Time   `json:"created_at"`
	UpdatedAt       null.Time   `json:"updated_at"`
}

type Chronology struct {
	// uuid.new is used
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

type District struct {
	ID       int32  `json:"id"`
	Title    string `json:"title"`
	RegionID int32  `json:"region_id"`
}

type DistrictBudget struct {
	BoardID    string    `json:"board_id"`
	DistrictID int32     `json:"district_id"`
	Budget     int32     `json:"budget"`
	UpdatedBy  string    `json:"updated_by"`
	CreatedAt  null.Time `json:"created_at"`
	UpdatedAt  null.Time `json:"updated_at"`
}

type Initiative struct {
	// uuid.new is used
	ID              string      `json:"id"`
	Title           string      `json:"title"`
	Images          []string    `json:"images"`
	Description     null.String `json:"description"`
	Author          string      `json:"author"`
	BoardID         string      `json:"board_id"`
	VoteCount       null.Int    `json:"vote_count"`
	Status          int32       `json:"status"`
	RequestedAmount null.Int    `json:"requested_amount"`
	GrantedAmount   null.Int    `json:"granted_amount"`
	RegionID        int32       `json:"region_id"`
	DistrictID      int32       `json:"district_id"`
	QuarterID       int32       `json:"quarter_id"`
	CreatedAt       null.Time   `json:"created_at"`
	UpdatedAt       null.Time   `json:"updated_at"`
}

type InitiativeChronology struct {
	InitiativeID string    `json:"initiative_id"`
	ChronologyID string    `json:"chronology_id"`
	UpdatedBy    string    `json:"updated_by"`
	CreatedAt    null.Time `json:"created_at"`
	UpdatedAt    null.Time `json:"updated_at"`
}

type Quarter struct {
	ID         int32  `json:"id"`
	Title      string `json:"title"`
	DistrictID int32  `json:"district_id"`
}

type Region struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

type Status struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

type SubInitiative struct {
	// uuid.new is used
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	VoteCount       null.Int  `json:"vote_count"`
	InitiativeID    string    `json:"initiative_id"`
	RequestedAmount null.Int  `json:"requested_amount"`
	GrantedAmount   null.Int  `json:"granted_amount"`
	CreatedAt       null.Time `json:"created_at"`
	UpdatedAt       null.Time `json:"updated_at"`
}

type User struct {
	// uuid.new is used
	ID          string      `json:"id"`
	Fullname    string      `json:"fullname"`
	PhoneNumber string      `json:"phone_number"`
	Role        null.String `json:"role"`
	Username    null.String `json:"username"`
	Password    null.String `json:"password"`
	Status      null.Int    `json:"status"`
	RegionID    int32       `json:"region_id"`
	DistrictID  int32       `json:"district_id"`
	CreatedAt   null.Time   `json:"created_at"`
	UpdatedAt   null.Time   `json:"updated_at"`
}

type Vote struct {
	PhoneNumber  string    `json:"phone_number"`
	InitiativeID string    `json:"initiative_id"`
	BoardID      string    `json:"board_id"`
	CreatedAt    null.Time `json:"created_at"`
	UpdatedAt    null.Time `json:"updated_at"`
}