package initialize

func Run() {
	LoadConfig()

	InitDb()

	r := InitRouter()

	r.Run(":5000")
}
