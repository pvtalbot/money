package errors

type DuplicateEntityError struct{}

func (e DuplicateEntityError) Error() string {
	return "duplicate entity already exists"
}
