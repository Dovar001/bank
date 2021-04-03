package types



//Money  представляет собой денежную сумму в минимальных единицах(центы, копейки, и т.д)
type Money int64

//Category  представляет сабой категорию в которой был совершен платёж
type Category string

//Status представляет собой статус платежа
type Status string

//Предопределлённые статуси платежей
const (
	StatusOK Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"		
)


//Payment представляет  информацию о платеже  
type Payment struct {
	ID int
	Amount Money
	Category Category
	Status Status 
}