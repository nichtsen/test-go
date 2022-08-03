# test-go
# gate
gate determines which moudule will handle which message

# skeleton
skeleton dedicated to regitsers `ChanRPC`

# activity
Activities: (type, id) -> Activity

## task pass
type UserTaskPassActivity struct {
        *UserBaseActivity
        // 当前等级
        Level            uint32
        // 当前经验
        Exp              uint32
        // 是否购买198元特权
        PassBought       bool
        // ??
        PassRewardGotten bool
        // 任务的完成状态
        Tasks            map[string]*UserActivityTask
        // 标识每级的普通奖励和解锁奖励是否领取
        StatusOfRewards  map[uint32]*LevelRewardStatus
}


