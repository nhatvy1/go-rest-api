package initialize

func Run() {
	LoadConfig()

	db := InitDb()

	r := InitRouter(db)

	r.Run(":5000")
}
