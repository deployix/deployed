package v1

import (
	actionableVersionV1 "github.com/deployix/deployed/pkg/actionableVersion/v1"
	historyV1 "github.com/deployix/deployed/pkg/history/v1"
	utilsV1 "github.com/deployix/deployed/pkg/utils/v1"
)

type Channel struct {
	Description       string                                `yaml:"description,omitempty"`
	ActionableVersion actionableVersionV1.ActionableVersion `yaml:"actionableVersion"`
	History           []historyV1.History
}

// AppendActionableVersion adds actionable version to history
func (c *Channel) AppendActionableVersion(v actionableVersionV1.ActionableVersion) {
	history := historyV1.History{
		Version: v.Version,
		Date:    utilsV1.GetCurrentDateTimeAsString(utilsV1.DateTimeLayoutFromTypeName(v.DateTime)),
	}

	c.History = append(c.History, history)
}
