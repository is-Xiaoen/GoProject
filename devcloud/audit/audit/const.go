package audit

const (
	META_AUDIT_KEY = "audit"
)

func Enable(v bool) (string, bool) {
	return META_AUDIT_KEY, v
}
