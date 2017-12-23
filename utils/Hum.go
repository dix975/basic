package utils

func reverse(lst []interface{}) chan interface{} {

	ret := make(chan interface{})
	go func() {
		for i, _ := range lst {
			ret <- lst[len(lst)-1-i]
		}
		close(ret)
	}()
	return ret
}
