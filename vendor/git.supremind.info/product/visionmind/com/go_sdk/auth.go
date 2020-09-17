package go_sdk

type IAuthClient interface {
	CheckToken(token string) (err error)
	GetSysOrgTree(token string) (root *SysOrgNode, err error)
}

type SysOrgNode struct {
	ID       int           `json:"id"`
	Code     string        `json:"code,omitempty"`
	Name     string        `json:"string"`
	Children []*SysOrgNode `json:"children"`
}
