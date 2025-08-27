package endpoint

type ACCESS_MODE uint8

const (
	ACCESS_MODE_READ = iota
	ACCESS_MODE_READ_WRITE
)

const (
	META_REQUIRED_AUTH_KEY      = "required_auth"
	META_REQUIRED_CODE_KEY      = "required_code"
	META_REQUIRED_PERM_KEY      = "required_perm"
	META_REQUIRED_ROLE_KEY      = "required_role"
	META_REQUIRED_AUDIT_KEY     = "required_audit"
	META_REQUIRED_NAMESPACE_KEY = "required_namespace"
	META_RESOURCE_KEY           = "resource"
	META_ACTION_KEY             = "action"
)
