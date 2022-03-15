// gen for home toolset
package job

import (
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _DemoJobSingle *DemoJob

func GetAllProvider() []interface{} {
	return []interface{}{
		NewDemoJob(),
	}
}

func NewDemoJob() *DemoJob {
	if _DemoJobSingle == nil {
		DemoJob := &DemoJob{}
		app.AfterProvider(DemoJob, "")
		_DemoJobSingle = DemoJob
	}
	return _DemoJobSingle
}
