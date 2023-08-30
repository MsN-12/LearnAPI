package main

func main() {
	app := App{}
	err := app.Initialise(DBUser, DBPassword, DBName)
	if err != nil {
		return
	}
	app.Run("localhost:1234")
}
