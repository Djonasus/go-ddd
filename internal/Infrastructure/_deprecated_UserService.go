

//DEPRECATED

// import (
// 	entity "ddd/internal/Entity"
// 	repository "ddd/internal/Repository"
// 	"log"
// )

// // Сервис для работы с репозиторием из контроллера
// type UserService struct {
// 	userRepository repository.UserRepository
// }

// func NewUserService(ur repository.UserRepository) *UserService {
// 	return &UserService{
// 		userRepository: ur,
// 	}
// }

// func (us *UserService) Create(login, password string) error {
// 	u, err := entity.NewUser(login, password)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	err = us.userRepository.Create(u)
// 	log.Println(err)
// 	return err
// }

// // func (us *UserService) Delete(login, password string) error {

// // }
