package pstring

import "github.com/gosimple/slug"

// ToSlug ...
func ToSlug(s string) string {
	return slug.Make(s)
}
