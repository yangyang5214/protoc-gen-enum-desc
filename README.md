# protoc-gen-enum-desc

gen enum description method(extended field comment)

### enum-desc

proto enum does not support enums similar to other languages, and provides declarations in key-value format.

In some cases, the proto enum declares the key, but needs to feedback the value (or field description or field comment) at the same time.

This plug-in generates corresponding methods by parsing enum annotations.

### usage

- proto

```shell
enum TaskStatus{
  CREATED = 0; //已创建
  RUNNING = 1; //运行中
  STOPPED = 2; //已停止
  FAILED = 3; //失败
  SUCCESS = 4; //成功
}
```

- protoc run

```shell
cd examples
protoc --proto_path=.  --go_out=paths=source_relative:. --enum-desc_out=paths=source_relative:.  task_status.proto
```

- gen TaskStatus_desc

```shell
package task

var (
	TaskStatus_desc = map[TaskStatus]string{
		TaskStatus_CREATED: "已创建",
		TaskStatus_RUNNING: "运行中",
		TaskStatus_STOPPED: "已停止",
		TaskStatus_FAILED:  "失败",
		TaskStatus_SUCCESS: "成功",
	}
)

func (x TaskStatus) Desc() string {
	return TaskStatus_desc[x]
}
```

- 使用

```shell
func TestName(t *testing.T) {
	t.Log(TaskStatus_RUNNING.Desc())
}
```
