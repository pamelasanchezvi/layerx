package handlers
import (
	"github.com/mesos/mesos-go/mesosproto"
	"github.com/layer-x/layerx-mesos-tpi_v2/mesos_master_api/mesos_data"
	"github.com/layer-x/layerx-commons/lxerrors"
"github.com/layer-x/layerx-commons/lxlog"
	"github.com/Sirupsen/logrus"
	"github.com/layer-x/layerx-mesos-tpi_v2/framework_manager"
	"github.com/layer-x/layerx-core_v2/lxtypes"
	"github.com/pborman/uuid"
	"github.com/layer-x/layerx-core_v2/layerx_tpi"
)


func HandleSubscribeRequest(tpi *layerx_tpi.LayerXTpi, frameworkManager framework_manager.FrameworkManager, frameworkUpid *mesos_data.UPID, call *mesosproto.Call_Subscribe) error {
	frameworkInfo := call.GetFrameworkInfo()
	frameworkName := frameworkInfo.GetName()
	frameworkId := frameworkInfo.GetId().GetValue()
	if frameworkId == "" {
		frameworkId = frameworkName+uuid.New()
	}

	taskProvider := &lxtypes.TaskProvider{
		Id: frameworkId,
		Source: frameworkUpid.String(),
	}
	err := tpi.RegisterTaskProvider(taskProvider)
	if err != nil {
		err = lxerrors.New("registering framework as new task provider with layer x", err)
		lxlog.Errorf(logrus.Fields{
			"error": err.Error(),
			"frameworkName": frameworkName,
			"frameworkId": frameworkId,
			"tpi": tpi,
		}, "handling subscribe call request")
		return err
	}

	err = frameworkManager.NotifyFrameworkRegistered(frameworkName, frameworkId, frameworkUpid)
	if err != nil {
		err = lxerrors.New("sending framework registered message to framework", err)
		lxlog.Errorf(logrus.Fields{
			"error": err.Error(),
			"frameworkName": frameworkName,
			"frameworkId": frameworkId,
			"frameworkUpid": frameworkUpid.String(),
		}, "handling subscribe call request")
		return err
	}
	return nil
}
