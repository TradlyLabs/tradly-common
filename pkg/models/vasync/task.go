package vasync

import "fmt"

type automationCompatibleType string

const (
	ACT_COMMON automationCompatibleType = "common"
	ACT_PAGING automationCompatibleType = "paging"
)

func AutomationCompatiblePattern(typ automationCompatibleType) string {
	return fmt.Sprintf("automation_comptible:%s", typ)
}

type walletHandleType string

const (
	WT_COMMON walletHandleType = "common"
)

func WalletHandlePattern(typ walletHandleType) string {
	return fmt.Sprintf("wallet:%s", typ)
}

func (t automationCompatibleType) String() string {
	return string(t)
}
