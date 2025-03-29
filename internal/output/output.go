package output

import "github.com/u2d-man/polyfeed/internal/core"

type Output interface {
	Send(articles []core.Article) error
}
