// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"gopkg.in/guregu/null.v3/zero"
)

type Author struct {
	ID   int64       `json:"id"`
	Name string      `json:"name"`
	Bio  zero.String `json:"bio"`
}