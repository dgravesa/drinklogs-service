package auth

// VerifyReadAccess returns true if accessor has read access to owner's data.
func VerifyReadAccess(accessorID, ownerID string) bool {
	// TODO implement
	return accessorID == ownerID
}

// VerifyWriteAccess returns true if accessor has write access to owner's data.
func VerifyWriteAccess(accessorID, ownerID string) bool {
	// TODO implement
	return accessorID == ownerID
}
