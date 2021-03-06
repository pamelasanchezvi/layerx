package framework_api_handlers

import (
	"github.com/emc-advanced-dev/layerx/layerx-core/layerx_rpi_client"
	"github.com/emc-advanced-dev/pkg/errors"
	"github.com/mesos/mesos-go/mesosproto"
)

func HandleStatusUpdate(lxRpi *layerx_rpi_client.LayerXRpi, status *mesosproto.TaskStatus) error {
	statusString := "task " + status.GetTaskId().GetValue() + " in state" + status.GetState().String()
	err := lxRpi.SubmitStatusUpdate(status)
	if err != nil {
		return errors.New("failed to submit status {"+statusString+"} to layerx core", err)
	}
	return nil
}
