package entities

// Session .
type Session struct {
	Openid   string `gorm:"primary_key;AUTO_INCREMENT;column:openid"`
	Username string `gorm:"column:username"`
}

// TableName .
func (*Session) TableName() string {
	return "session"
}

// SessionService .
type SessionService struct{}

// SessionServ .
var SessionServ = SessionService{}

func init() {
	addServ(&SessionServ)
}

func (*SessionService) load() {
	s := &Session{}
	if !gormDb.HasTable(s) {
		gormDb.CreateTable(s)
	}
}

// Add .
func (*SessionService) Add(s *Session) {
	tx := gormDb.Begin()
	checkErr(tx.Error)

	if err := tx.Create(s).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}

	tx.Commit()
}

// Delete .
func (*SessionService) Delete(openid string) {
	s := &Session{Openid: openid}
	tx := gormDb.Begin()
	checkErr(tx.Error)

	if err := tx.Delete(s).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}

	tx.Commit()
}

// FindByOpenid .
func (*SessionService) FindByOpenid(openid string) *Session {
	sessions := make([]Session, 0, 0)
	checkErr(gormDb.Where([]string{openid}).Find(&sessions).Error)
	if len(sessions) == 0 {
		return nil
	}
	return &sessions[0]
}
