package auth

// PermissionsReader is an interface for retrieving a user's permissions.
type PermissionsReader interface {
	UserPermissions(uid uint64) UserPermissions
}

// UserPermissions contains a user's read and write permissions for drink log data.
type UserPermissions struct {
	read, write map[uint64]bool
}

// HasRead returns true if current permissions include read; otherwise, false.
func (p *UserPermissions) HasRead(uid uint64) bool {
	_, isReadable := p.read[uid]
	return isReadable
}

// HasWrite returns true if current permissions include write; otherwise, false.
func (p *UserPermissions) HasWrite(uid uint64) bool {
	_, isWriteable := p.write[uid]
	return isWriteable
}
