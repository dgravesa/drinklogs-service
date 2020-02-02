package auth

// PermissionsReader is an interface for retrieving a user's permissions.
type PermissionsReader interface {
	UserPermissions(uid string) UserPermissions
}

// UserPermissions contains a user's read and write permissions for drink log data.
type UserPermissions struct {
	read, write map[string]bool
}

// HasRead returns true if current permissions include read; otherwise, false.
func (p *UserPermissions) HasRead(uid string) bool {
	_, isReadable := p.read[uid]
	return isReadable
}

// HasWrite returns true if current permissions include write; otherwise, false.
func (p *UserPermissions) HasWrite(uid string) bool {
	_, isWriteable := p.write[uid]
	return isWriteable
}
