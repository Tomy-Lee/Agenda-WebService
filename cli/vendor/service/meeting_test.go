package service_test

// import (
// 	"go/build"
// 	"runtime"
// 	"service"
// 	"testing"

// 	"entity"
// )

// func init() {
// 	if runtime.GOOS == "windows" {
// 		entity.MeetingModel.Init(build.Default.GOPATH + "\\tmp\\test_meeting_data.json")
// 		entity.UserModel.Init(build.Default.GOPATH + "\\tmp\\test_user_data.json")
// 		entity.CurSessionModel.Init(build.Default.GOPATH + "\\tmp\\test_curUser.json")
// 	} else {
// 		entity.MeetingModel.Init("/tmp/test_meeting_data.json")
// 		entity.UserModel.Init("/tmp/test_user_data.json")
// 		entity.CurSessionModel.Init("/tmp/test_curUser.json")
// 	}
// }

// func TestMeetingModel(t *testing.T) {
// 	model := entity.MeetingModel
// 	newMeeting := entity.Meeting{
// 		Title:         "shadowsocks",
// 		Speecher:      "test1",
// 		Participators: []string{"test2", "test3", "test3"},
// 		StartTime:     "2017-10-10/12:00:00",
// 		EndTime:       "2017-10-10/13:00:00",
// 	}
// 	model.AddMeeting(&newMeeting)
// 	foundMeetings := model.FindBy(func(oneMeeting *entity.Meeting) bool {
// 		return oneMeeting.Title == "shadowsocks"
// 	})
// 	if len(foundMeetings) != 1 {
// 		t.Errorf("Didn't find the test meeting")
// 	}
// 	foundMeeting := model.FindByTitle("shadowsocks")
// 	if foundMeeting == nil {
// 		t.Errorf("Didn't find the test meeting")
// 	}

// 	model.DeleteParticipatorFromMeeting(foundMeeting, "test2")
// 	foundMeeting = model.FindByTitle("shadowsocks")
// 	if len(foundMeeting.Participators) != 2 {
// 		t.Errorf("Error from deleteParticipator test")
// 	}

// 	model.AddParticipatorToMeeting(foundMeeting, "test2")
// 	foundMeeting = model.FindByTitle("shadowsocks")
// 	if len(foundMeeting.Participators) != 3 {
// 		t.Errorf("Error from addParticipator test")
// 	}
// }

// func TestMeetingService(t *testing.T) {
// 	var err error
// 	err = service.Register("test1", "test1", "email@email.com", "12345678912")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = service.Register("test2", "test2", "email@email.com", "12345678912")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = service.Register("test3", "test3", "email@email.com", "12345678912")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if err := service.Login("test1", "test1"); err != nil {
// 		t.Fatal(err)
// 	}

// 	err = service.AddMeeting("shadowsocksTest", []string{"test2", "test3"}, "2017-01-01/12:00:00", "2017-01-01/12:00:00")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	err = service.DeleteParticipatorFromMeeting("shadowsocksTest", []string{"test2"})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	err = service.AddParticipatorToMeeting("shadowsocksTest", []string{"test2"})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	err = service.DeleteAllMeeting()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	results, err := service.FindMeetingByTime("2017-01-01/12:00:00", "2017-01-01/12:00:00")
// 	if err != nil {
// 		t.Fatal(err)
// 	} else if len(results) != 0 {
// 		t.Errorf("Error with FindAll")
// 	}
// }
