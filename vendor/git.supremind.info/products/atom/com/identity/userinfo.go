package identity

type UserIdentity struct {
	Username      string   `json:"username,omitempty"`
	Fullname      string   `json:"fullname,omitempty"`
	Email         string   `json:"email,omitempty"`
	UpstreamToken string   `json:"upstream_token,omitempty"`
	SuperUser     bool     `json:"super_user,omitempty"`
	Groups        []string `json:"groups,omitempty"`
}
