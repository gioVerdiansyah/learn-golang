package my_modules

var User = struct {
	Name  string
	Email string
}{}

func init() {
	User.Name = "Verdi"
	User.Email = "verdi@gmail.com"

	//fmt.Println("my_modules imported!")
}
