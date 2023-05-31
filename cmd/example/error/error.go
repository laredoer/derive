package err

// 系统正在处理中，请稍后再试。
//
// #[i18n(code = 2501506, zh-HK = "系統正在處理中，請稍後再試", zh-CN = "系统正在处理中，请稍后再试", en = "System is processing, please try again later")]
type SystemIsProcessing int32

// RewardIsInProgress 奖励正在发放中，无需再次领取
//
// #[i18n(code = 400, default="en",zh-HK = "獎勵正在發放中，無需再次領取", zh-CN = "奖励正在发放中，无需再次领取", en = "Reward is in progress")]
type RewardIsInProgress int32
