package request

import "time"

type PeriodRequest struct {
	Period          uint      `json:"period"`
	Stage1Starttime time.Time `json:"stage1_startTime"`
	Stage1Endtime   time.Time `json:"stage1_endTime"`
	Stage2Starttime time.Time `json:"stage2_startTime"`
	Stage2Endtime   time.Time `json:"stage2_endTime"`
	Stage3Starttime time.Time `json:"stage3_startTime"`
	Stage3Endtime   time.Time `json:"stage3_endTime"`
	Stage4Starttime time.Time `json:"stage4_startTime"`
	Stage4Endtime   time.Time `json:"stage4_endTime"`
	Stage5Starttime time.Time `json:"stage5_startTime"`
	Stage5Endtime   time.Time `json:"stage5_endTime"`
}
