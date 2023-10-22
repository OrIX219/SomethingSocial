package posts

import "time"

type PostFilter struct {
	Author   *string
	Limit    *int
	DateFrom *time.Time
	DateTo   *time.Time
	Sort     *string
	Vote     *string
}
