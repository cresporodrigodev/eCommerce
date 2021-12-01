package command

type RegisterNewUserCommand struct {
	Alias      string
	FirstName  string
	SecondName string
	Email      string
	Admin      string
	Password   string
}

func NewRegisterNewUserCommand(alias, firstName, secondName, email, admin, password string) RegisterNewUserCommand {
	return RegisterNewUserCommand{Alias: alias, FirstName: firstName, SecondName: secondName, Email: email, Admin: admin, Password: password}
}

func (cmd RegisterNewUserCommand) CommandId() string {
	const id = "registerNewUser"
	return id
}
