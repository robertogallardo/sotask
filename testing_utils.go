//go:build windows
// +build windows

package sotask

func createTestTask(taskSvc TaskService) RegisteredTask {
	newTaskDef := taskSvc.NewTaskDefinition()
	newTaskDef.AddAction(ExecAction{
		Path: "cmd.exe",
		Args: "/c timeout $(Arg0)",
	})
	newTaskDef.Settings.MultipleInstances = TASK_INSTANCES_PARALLEL

	task, _, err := taskSvc.CreateTask("\\Sotask\\TestTask", newTaskDef, true)
	if err != nil {
		panic(err)
	}

	return task
}
