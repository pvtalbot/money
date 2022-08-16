package queries

type GetAllErrors struct{}

type GetAllErrorsHandler struct{}

type Error struct {
	Code        string
	Translation string
}

func (h GetAllErrorsHandler) Handle(q GetAllErrors) ([]*Error, error) {
	return []*Error{
		{
			Code:        "createuser-1",
			Translation: "An account with this username already exists",
		},
	}, nil
}
