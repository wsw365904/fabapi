/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package allcontroller

import (
	"github.com/wsw365904/fabapi/core/fabsdk"
	"github.com/wsw365904/fabapi/core/fabsdk/chaincodeimpl"
	"github.com/wsw365904/fabapi/core/fabsdk/channelimpl"
	"github.com/wsw365904/fabapi/core/fabsdk/fabsdkimpl"
	"github.com/wsw365904/fabapi/core/fabsdk/ledgerimpl"
	"github.com/wsw365904/fabapi/core/fabsdk/models"
	"github.com/wsw365904/fabapi/internal/app/controller/chaincodecontroller"
	"github.com/wsw365904/fabapi/internal/app/controller/channelcontroller"
	"github.com/wsw365904/fabapi/internal/app/controller/ledgercontroller"
	"github.com/wsw365904/fabapi/internal/app/controller/logcontroller"
	"github.com/wsw365904/fabapi/internal/app/service/chaincodeserviceimpl"
	"github.com/wsw365904/fabapi/internal/app/service/channelserviceimpl"
	"github.com/wsw365904/fabapi/internal/app/service/ledgerserviceimpl"
	"github.com/wsw365904/fabapi/internal/app/service/logserviceimpl"
	"github.com/wsw365904/fabapi/internal/pkg/config/fabconfig"

	"github.com/wsw365904/wswlog/wlogging"
)

var logger = wlogging.MustGetLoggerWithoutName()

type Controller struct {
	ChaincodeController *chaincodecontroller.ChaincodeController
	ChannelController   *channelcontroller.ChannelController
	LedgerController    *ledgercontroller.LedgerController
	LogController       *logcontroller.LogController
	FabSdk              fabsdk.FabSdk
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewController(chaincodeController *chaincodecontroller.ChaincodeController,
	channelController *channelcontroller.ChannelController,
	ledgerController *ledgercontroller.LedgerController,
	logController *logcontroller.LogController,
	fabSdk fabsdk.FabSdk) *Controller {
	return &Controller{
		ChaincodeController: chaincodeController,
		ChannelController:   channelController,
		LedgerController:    ledgerController,
		LogController:       logController,
		FabSdk:              fabSdk,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func New() (*Controller, *models.Other, error) {
	logger.Debug("NewController enter")
	fabSdk, sdk, err := fabsdkimpl.NewFabSdk()
	if err != nil {
		logger.Error("SetupFabricSDK err:", err)
		return nil, nil, err
	}
	other, err := fabconfig.NewConfig()
	if err != nil {
		logger.Error("LoadConfig err:", err)
		return nil, nil, err
	}
	return NewController(
		chaincodecontroller.NewChaincodeController(chaincodeserviceimpl.NewChaincodeService(chaincodeimpl.NewChaincodeOp(other, sdk))),
		channelcontroller.NewChannelController(channelserviceimpl.NewChannelService(channelimpl.NewChannelOp(other, sdk))),
		ledgercontroller.NewLedgerController(ledgerserviceimpl.NewLedgerService(ledgerimpl.NewLedgerOp(other, sdk))),
		logcontroller.NewLogController(logserviceimpl.NewLogService()),
		fabSdk), other, nil
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *Controller) Close() {
	logger.Debug("Controller enter close")
	if c == nil {
		return
	}
	c.ChaincodeController.Close()
	c.ChannelController.Close()
	c.LedgerController.Close()
	c.LogController.Close()
	c.FabSdk.Close()
	c = nil
}
