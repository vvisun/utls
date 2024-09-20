package timeutil

type StateVarType int
type SaveMode int32

const (
	SaveOneDay           SaveMode = 1
	SaveOneWeekOfMonday  SaveMode = 2 //每周一凌晨更新
	SaveOneWeekOfTuesDay SaveMode = 3 //每周二凌晨更新
	SaveForever          SaveMode = 4
)
