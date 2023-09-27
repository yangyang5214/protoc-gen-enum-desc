package task

import "testing"

func TestName(t *testing.T) {
	t.Log(TaskStatus_RUNNING.Desc())
}
