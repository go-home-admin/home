// 代码由home-admin生成, 不需要编辑它

package job

import (
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
)

var DemoJobSingle *DemoJob

func NewDemoJobProvider() *DemoJob {
	DemoJob := &DemoJob{}
	return DemoJob
}

func InitializeNewDemoJobProvider() *DemoJob {
	if DemoJobSingle == nil {
		DemoJobSingle = NewDemoJobProvider()

		home_constraint.AfterProvider(DemoJobSingle)
	}

	return DemoJobSingle
}
