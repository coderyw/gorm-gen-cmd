// Package dao
// @Author: yinwei
// @File: Querier
// @Version: 1.0.0
// @Date: 2023/11/22 13:28

package generate

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id uint64) (gen.T, error) // GetByID query data by id and return it as *struct*
}
