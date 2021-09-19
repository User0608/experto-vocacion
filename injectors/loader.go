package injectors

import (
	"log"
	"sync"

	"github.com/user0608/expertos/configs"
	"github.com/user0608/expertos/database"
	"github.com/user0608/expertos/handlers"
	"github.com/user0608/expertos/repository"
	"github.com/user0608/expertos/services"
	"gorm.io/gorm"
)

var ones sync.Once

func init() {
	ones.Do(func() {
		conf, err := configs.LoadDBConfigs("db_config.json")
		if err != nil {
			log.Fatalln("Err db configs,", err.Error())
		}
		log.Println("Configuraciones de base de datos cargadas!")
		db = database.GetDBConnextion(conf)
		initRepository(db)
		initServices()
		initHandlers()
	})
}
func initRepository(gdb *gorm.DB) {
	loginRepository = repository.NewLoginRepository(gdb)
	estudianteRepository = repository.NewEstudianteRepository(gdb)
	casmRepository = repository.NewCASMRepository(gdb)
	bergerRepository = repository.NewBergerRepository(gdb)
	heaRepository = repository.NewHEARepository(gdb)
	testRepository = repository.NewTestRepository(gdb)
}
func initServices() {
	loginService = services.NewLoginService(loginRepository)
	estudianteService = services.NewEstudianteRepository(estudianteRepository)
	questionService = services.NewQuestionService(casmRepository, bergerRepository, heaRepository)
	testService = services.NewTestService(testRepository)
}
func initHandlers() {
	loginHandler = handlers.NewLogginHandler(loginService)
	estudianteHandler = handlers.NewEstudianteService(estudianteService)
	questionHandler = handlers.NewQuestionHandler(questionService)
	testHandler = handlers.NewTestHandler(testService)
}
