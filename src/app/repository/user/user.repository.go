package user

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/event"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindOne(id string, result *user.User) error {
	return r.db.First(&result, "id = ?", id).Error
}

func (r *Repository) FindByStudentID(sid string, result *user.User) error {
	return r.db.First(&result, "student_id = ?", sid).Error
}

func (r *Repository) CreateOrUpdate(result *user.User) error {
	if r.db.Where("id = ?", result.ID.String()).Updates(&result).RowsAffected == 0 {
		return r.db.Create(&result).Error
	}
	return nil
}

func (r *Repository) Verify(studentId string) error {
	return r.db.Model(&user.User{}).First(&user.User{}, "student_id = ?", studentId).Update("is_verify", true).Error
}

func (r *Repository) Create(in *user.User) error {
	return r.db.Create(&in).Error
}

func (r *Repository) Update(id string, result *user.User) error {
	return r.db.Where(id, "id = ?", id).Updates(&result).First(&result, "id = ?", id).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&user.User{}).Error
}

func (r *Repository) VerifyEstamp(uId string, eId string) bool {
	//Check if this user has this estamp
	var thisUser user.User
	r.db.First(&thisUser, "id = ?", uId)

	for _, event := range thisUser.Events {
		if event.ID.String() == eId {
			return true
		}
	}
	return false
}

func (r *Repository) ConfirmEstamp(uId string, eId string) error {
	//Check if this user has this estamp
	var thisUser user.User
	result := r.db.First(&thisUser, "id = ?", uId)
	if result.Error != nil {
		return result.Error
	}

	var thisEvent event.Event
	result = r.db.First(&thisEvent, "id = ?", eId)
	if result.Error != nil {
		return result.Error
	}

	thisUser.Events = append(thisUser.Events, &thisEvent)
	return nil
}

func (r *Repository) FindUserEstamp(uId string, foundList []string, unfoundList []string) error {
	//Find all estamps this user has and hasn't
	var thisUser user.User
	result := r.db.First(&thisUser, "id = ?", uId)
	if result.Error != nil {
		return result.Error
	}

	var allEvent []*event.Event
	result = r.db.Model(&event.Event{}).Find(allEvent)
	if result.Error != nil {
		return result.Error
	}

	flag := make(map[string]bool)
	for _, e := range thisUser.Events {
		flag[e.ID.String()] = true
	}

	for _, e := range allEvent {
		if flag[e.ID.String()] == true {
			foundList = append(foundList, e.ID.String())
		} else {
			unfoundList = append(unfoundList, e.ID.String())
		}
	}
	return nil
}
