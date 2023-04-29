package i18n

type ErrorCode int32

const (
	// TeamFull is returned when a team is full.
	// #[i18n(zh-HK = "红包已被领取", zh-CN = "红包已被领取", en = "Red packet has been received")]
	// #[clone]
	TeamFull ErrorCode = 4001
	// RedPacketHasBeenReceived is returned when a red packet has been received.
	// #[i18n(zh-HK = "红包已被领取", zh-CN = "红包已被领取", en = "Red packet has been received")]
	RedPacketHasBeenReceived ErrorCode = 401
)

func (ErrorCode) Code() int {
	return 1
}

func bar() {
	TeamFull.Code()
}
