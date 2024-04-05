package exception

import "fmt"

type builtIn struct {
	Reason string
	Status int
	Err    error
}

func newBuiltIn(reason string, status int, errs ...error) builtIn {
	var err error = nil
	if len(errs) > 0 {
		err = errs[0]
	}

	return builtIn{
		Reason: reason,
		Status: status,
		Err:    err,
	}
}

func (b *builtIn) Error() string {
	if b.Err != nil {
		return fmt.Sprintf("예외가 발생한 이유는 다음과 같습니다. \n%s\n세부정보: %s", b.Reason, b.Err.Error())
	}
	return fmt.Sprintf("예외가 발생한 이유는 다음과 같습니다. \n%s", b.Reason)
}
