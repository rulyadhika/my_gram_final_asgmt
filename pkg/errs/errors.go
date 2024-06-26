package errs

// not found err
type NotFoundError struct {
	msg string
}

func NewNotFoundError(err string) error {
	return &NotFoundError{msg: err}
}

func (n *NotFoundError) Error() string {
	return n.msg
}

// not found err

// bad request err
type BadRequestError struct {
	msg string
}

func NewBadRequestError(err string) error {
	return &BadRequestError{msg: err}
}

func (b *BadRequestError) Error() string {
	return b.msg
}

// bad request err

// unprocessable entity err
type UnprocessableEntityError struct {
	msg string
}

func NewUnprocessableEntityError(err string) error {
	return &UnprocessableEntityError{msg: err}
}

func (n *UnprocessableEntityError) Error() string {
	return n.msg
}

// unprocessable entity err

// internal server err
type InternalServerError struct {
	msg string
}

func NewInternalServerError(err string) error {
	return &InternalServerError{msg: err}
}

func (n *InternalServerError) Error() string {
	return n.msg
}

// internal server err

// conflict err
type ConflictError struct {
	msg string
}

func NewConflictError(err string) error {
	return &ConflictError{msg: err}
}

func (n *ConflictError) Error() string {
	return n.msg
}

// conflict err

// unauthorized err
type UnauthorizedError struct {
	msg string
}

func NewUnauthorizedError(err string) error {
	return &UnauthorizedError{msg: err}
}

func (n *UnauthorizedError) Error() string {
	return n.msg
}

// unauthorized err

// forbidden err
type ForbiddenError struct {
	msg string
}

func NewForbiddenError(err string) error {
	return &ForbiddenError{msg: err}
}

func (n *ForbiddenError) Error() string {
	return n.msg
}

// forbidden err
