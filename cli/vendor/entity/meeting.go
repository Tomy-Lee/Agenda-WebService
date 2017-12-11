package entity

// Meeting model for one meeting
type Meeting struct {
	Title         string   `json:"tile"`
	Speecher      string   `json:"speecher"`
	Participators []string `json:"participators"`
	StartTime     string   `json:"startTime"`
	EndTime       string   `json:"endTime"`
}

type meetingDb struct {
	Data []Meeting `json:"data"`
}

type meetingModel struct {
	storage
	meetings map[string]*Meeting
}

var (
	// MeetingModel for meetings
	MeetingModel meetingModel
)

func init() {
	// addModel(&MeetingModel, "meeting_data")
}

// Init initialize a meeting model
func (model *meetingModel) Init(path string) {
	logger.Println("[meetingmodel] initializing")
	model.path = path
	model.meetings = make(map[string]*Meeting)

	model.load()
	logger.Println("[meetingmodel] initialized")
}

// Addmeeting add a new meeting to database
func (model *meetingModel) AddMeeting(meeting *Meeting) {
	logger.Println("[meetingmodel] try adding new meeting", meeting.Title)
	model.meetings[meeting.Title] = meeting
	model.dump()
	logger.Println("[meetingmodel] added new meeting", meeting.Title)
}

// AddParticipatorToMeeting Add a Participator To Meeting
func (model *meetingModel) AddParticipatorToMeeting(meeting *Meeting, participator string) {
	logger.Println("[meetingmodel] try adding a participator to meeting", meeting.Title)
	curMeetingParticipators := model.meetings[meeting.Title].Participators
	model.meetings[meeting.Title].Participators = append(curMeetingParticipators, participator)
	model.dump()
	logger.Println("[meetingmodel] added a participator to meeting", meeting.Title)
}

// FindMeetingCondition filter function to query meeting
type FindMeetingCondition func(*Meeting) bool

// FindBy find meetingList with provided condition
func (model *meetingModel) FindBy(condition FindMeetingCondition) []Meeting {
	result := []Meeting{}
	for _, meeting := range model.meetings {
		if condition(meeting) {
			result = append(result, *meeting)
		}
	}
	return result
}

// FindByTitle find meeting by meetingname
func (model *meetingModel) FindByTitle(meetingname string) *Meeting {
	return model.meetings[meetingname]
}

func (model *meetingModel) DeleteMeeting(meeting *Meeting) {
	logger.Println("[meetingmodel] try deleting a meeting", meeting.Title)
	delete(model.meetings, meeting.Title)
	model.dump()
	logger.Println("[meetingmodel] deleted a meeting", meeting.Title)
}

func (model *meetingModel) DeleteParticipatorFromMeeting(meeting *Meeting, participator string) {
	logger.Println("[meetingmodel] try deleting a participator from meeting", meeting.Title)
	curMeetingParticipators := model.meetings[meeting.Title].Participators
	for i, p := range curMeetingParticipators {
		if p == participator {
			curMeetingParticipators = append(curMeetingParticipators[:i], curMeetingParticipators[i+1:]...)
			break
		}
	}
	model.meetings[meeting.Title].Participators = curMeetingParticipators
	model.dump()
	logger.Println("[meetingmodel] deleted a participator from meeting", meeting.Title)
}

func (model *meetingModel) load() {
	var meetingDb meetingDb
	model.storage.load(&meetingDb)
	for index, meeting := range meetingDb.Data {
		model.meetings[meeting.Title] = &meetingDb.Data[index]
	}
}

func (model *meetingModel) dump() {
	var meetingDb meetingDb
	for _, meeting := range model.meetings {
		meetingDb.Data = append(meetingDb.Data, *meeting)
	}
	model.storage.dump(&meetingDb)
}
