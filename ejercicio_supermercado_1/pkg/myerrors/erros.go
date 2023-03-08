package myerrors


type ServerError struct{
	Message string
}

func (s ServerError) Error() string{
	return s.Message
}

type ResourseNotFound struct{
	Message string
}

func (e ResourseNotFound) Error() string{
	return e.Message
}

type DuplicatedError struct{
	Message string
}

func (e DuplicatedError) Error() string{
	return e.Message
}



