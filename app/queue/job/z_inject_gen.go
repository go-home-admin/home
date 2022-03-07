// gen for home toolset
package job

var _DemoJobSingle *DemoJob

func GetAllProvider() []interface{} {
	return []interface{}{
		NewDemoJob(),
	}
}

func NewDemoJob() *DemoJob {
	if _DemoJobSingle == nil {
		DemoJob := &DemoJob{}
		_DemoJobSingle = DemoJob
	}
	return _DemoJobSingle
}
