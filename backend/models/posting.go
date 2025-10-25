package models

type Posting struct {
	GormModel
	Photo   string
	Caption string
}
