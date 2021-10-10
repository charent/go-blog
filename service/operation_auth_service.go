package service

type OperationAuth struct {

}

func (o *OperationAuth) OperationCheck(userId int, operation string) (val bool){
	val = false

	// Mysql auto increase 的第一个id为1
	if userId == 1 {
		val = true
		return
	}

	if operation == ""{
		val = false
	}
	return
}