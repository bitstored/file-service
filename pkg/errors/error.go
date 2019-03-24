package errors

type Kind int

// Error represents errors for the service.
type Error struct {
	Kind Kind
	Err  error
}

const (
	KindOther        Kind = iota // Unclassified error.
	KindUnauthorized             // Unauthorized
	KindNotFound                 // entity not found
	KindExisting                 // already exists
)

func NewError(k Kind, e error) *Error {
	return &Error{Kind: k, Err: e}
}

func (e *Error) Error() string {
	switch e.Kind {
	case KindOther:
		return "other error"
	case KindNotFound:
		return "entity not found"
	case KindUnauthorized:
		return "not authorized for this action"
	case KindExisting:
		return "entity exists"
	default:
		return "something is wrong"
	}
}
