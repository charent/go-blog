package auth

func operationCheck(roleId int, operation string) (val bool){
	val = false
	if roleId == 0 {
		val = true
		return
	}
	if operation == ""{
		val = false
	}
	return
}