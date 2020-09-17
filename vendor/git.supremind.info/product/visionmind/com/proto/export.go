package proto

import (
	"errors"
	"fmt"
)

type ExportType int

const (
	ExportTypeNone           ExportType = iota
	ExportType_General_2And2            // 通用 2+2
	ExportType_General_4In1             // 通用 4合1
	ExportType_General_4                // 通用 4张单独
	ExportType_Wuhan_4In1               // 武汉 4合1
	ExportType_Taicang_4In1             //太仓 4合1
	// TODO，后续增加导出类型在这边添加
	ExportTypeEnd
)

func CheckExportType(event *EventMsg, exportType int) (err error) {
	if event == nil {
		err = errors.New("event is nil")
		return
	}
	if exportType <= int(ExportTypeNone) || exportType >= int(ExportTypeEnd) {
		err = fmt.Errorf("exportType in invalid, %d", exportType)
		return
	}

	switch exportType {
	case int(ExportType_General_4):
		// 4 张独立图片导出，不要求是 3 张取证图
		if exportType == int(ExportType_General_4) && len(event.Snapshot) < 1 {
			err = fmt.Errorf("snapshot len < 1 for exportType %d", exportType)
			return
		}
	default:
		if len(event.Snapshot) < 3 {
			err = fmt.Errorf("snapshot len < 3 for exportType %d", exportType)
			return
		}
	}

	return nil
}
