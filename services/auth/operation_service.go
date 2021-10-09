package auth

func operationCheck(roleId int, operation string) (val bool){
	val = false

	// Mysql auto increase 的第一个id为1
	if roleId == 1 {
		val = true
		return
	}

	if operation == ""{
		val = false
	}
	return
}