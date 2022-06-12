package controller

import (
	"sync"

	"github.com/goldencoderam/so-p2_processes/src/model"
	"github.com/goldencoderam/so-p2_processes/src/object"
)

type mainController struct {
	Processor *model.Processor
}

var lock = &sync.Mutex{}

var mainControllerInstance *mainController

func GetMainControllerInstance() *mainController {
	if mainControllerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if mainControllerInstance == nil {
			mainControllerInstance = &mainController{
				Processor: &model.Processor{},
			}
		}
	}
	return mainControllerInstance
}

func (m *mainController) AddProcessToProcessor(process *object.Process) {
    m.Processor.AddProcessToReadyList(process)
}
