package bookdepository

import "luminnovel/internal/entity"

type Product struct {
	entity.ProductBookDepository
	Volume string
}
