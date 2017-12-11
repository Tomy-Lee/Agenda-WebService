package entities

// User .
type User struct {
	Username string `json:"username" gorm:"primary_key;column:username"`
	Password string `json:"password" gorm:"column:password"`
	Email    string `json:"email" gorm:"column:email"`
	Phone    string `json:"phone" gorm:"column:phone"`
}

// TableName .
func (*User) TableName() string {
	return "user"
}

// UserService .
type UserService struct{}

// UserServ .
var UserServ = UserService{}

func init() {
	addServ(&UserServ)
}

func (*UserService) load() {
	u := &User{}
	if !gormDb.HasTable(u) {
		gormDb.CreateTable(u)
	}
}

// Add .
func (*UserService) Add(u *User) {
	tx := gormDb.Begin()
	checkErr(tx.Error)

	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}

	tx.Commit()
}

// Delete .
func (*UserService) Delete(u *User) {
	tx := gormDb.Begin()
	checkErr(tx.Error)

	if err := tx.Delete(u).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}

	tx.Commit()
}

// FindAll .
func (*UserService) FindAll() []User {
	ulist := make([]User, 0, 0)
	checkErr(gormDb.Find(&ulist).Error)
	return ulist
}

// FindByUsername .
func (*UserService) FindByUsername(username string) *User {
	users := make([]User, 0, 0)
	checkErr(gormDb.Where([]string{username}).Find(&users).Error)
	if len(users) == 0 {
		return nil
	}
	return &users[0]
}
