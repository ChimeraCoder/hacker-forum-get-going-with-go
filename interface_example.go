type Account interface {
    Deposit(int) error
}

type CheckingAccount struct {
    Balance int 
        superSecretId int64
}

func (destination CheckingAccount) Deposit(amount int) error {
    destination.Balance += amount
    return nil
}

// We want to donate money to any account type (Checking, Savings)
// Donations require a tax receipt
func Donate(recipient Account, amount int) error {
    err := recipient.Deposit(amount)
    if err != nil{
        return err
    }
    IssueTaxReceipt(amount)
}
