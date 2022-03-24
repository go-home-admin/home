package demo

import (
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/bootstrap/utils"
	"testing"
)

func TestOrmTableName_First(t *testing.T) {
	providers.NewApp()

	first, _ := NewOrmTableName().First()

	utils.Dump(first)
}
