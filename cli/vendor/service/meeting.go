package service

import (
	"entity"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02/15:04:05"

func validateNewMeeting(meeting *entity.Meeting) error {
	if len(meeting.Title) == 0 {
		return fmt.Errorf("Title should not be empty")
	}
	if entity.MeetingModel.FindByTitle(meeting.Title) != nil {
		return fmt.Errorf("Title '%s' already exists", meeting.Title)
	}
	if len(meeting.Participators) == 0 {
		return fmt.Errorf("Meeting must have participator")
	}
	for _, participatorName := range meeting.Participators {
		if entity.UserModel.FindByUsername(participatorName) == nil {
			return fmt.Errorf("Participator %s is not existed", participatorName)
		}
		if participatorName == meeting.Speecher {
			return fmt.Errorf("You can't be the participator")
		}
	}

	err := validateNewTimeInterval(meeting.StartTime, meeting.EndTime)
	if err != nil {
		return err
	}

	if flag, err := validateNewMeetingTime(meeting); !flag {
		return err
	}

	// ...TODO
	return nil
}

func validateNewTimeInterval(startTime string, endTime string) error {
	if len(startTime) == 0 {
		return fmt.Errorf("Empty start time")
	}
	// Start Time format must be XXXX-XX-XX/XX:XX
	_, err := time.Parse(timeFormat, startTime)

	if err != nil {
		return fmt.Errorf("Illegal StartTime format")
	}

	if len(endTime) == 0 {
		return fmt.Errorf("Empty end time")
	}
	// End Time format must be XXXX-XX-XX/XX:XX
	_, err = time.Parse(timeFormat, endTime)

	if err != nil {
		return fmt.Errorf("Illegal EndTime format")
	}

	if startTime > endTime {
		return fmt.Errorf("StartTime is after the EndTime")
	}
	return nil
}

func validateNewMeetingTime(newMeeting *entity.Meeting) (bool, error) {
	// add all meetings belong to the speecher to the results
	results := entity.MeetingModel.FindBy(func(m *entity.Meeting) bool {
		// check from speecher field
		if m.Speecher == newMeeting.Speecher {
			return true
		}
		// check from participator field
		for _, participator := range m.Participators {
			if participator == newMeeting.Speecher {
				return true
			}
		}
		return false
	})

	// add all meetings belong to the participators to the results
	for _, participatorName := range newMeeting.Participators {
		results = append(results, entity.MeetingModel.FindBy(func(m *entity.Meeting) bool {
			// check from speecher field
			if m.Speecher == participatorName {
				return true
			}
			// check from participator field
			for _, participator := range m.Participators {
				if participator == participatorName {
					return true
				}
			}
			return false
		})...)
	}

	// check the time
	return validateFreeTime(newMeeting.StartTime, newMeeting.EndTime, results)
}

func validateFreeTime(startTime string, endTime string, meetings []entity.Meeting) (bool, error) {
	for _, oldMeeting := range meetings {
		if endTime > oldMeeting.StartTime && oldMeeting.EndTime > startTime {
			return false, fmt.Errorf("Time conflited with the meeting %s", oldMeeting.Title)
		}
	}
	return true, nil
}

// AddMeeting add a meeting
func AddMeeting(title string, participatorName []string, startTime string, endTime string) (err error) {
	// Login first
	if err := checkIfLoggedin(); err != nil {
		return err
	}
	speecherName := entity.CurSessionModel.GetCurUser()
	newMeeting := &entity.Meeting{
		Title:         title,
		Speecher:      speecherName,
		Participators: participatorName,
		StartTime:     startTime,
		EndTime:       endTime,
	}
	// validate if format is ok
	err = validateNewMeeting(newMeeting)
	if err != nil {
		return
	}
	entity.MeetingModel.AddMeeting(newMeeting)
	return
}

// FindMeetingByTime find the meeting according to the time interval
func FindMeetingByTime(startTime string, endTime string) ([]entity.Meeting, error) {
	if err := checkIfLoggedin(); err != nil {
		return nil, err
	}
	err := validateNewTimeInterval(startTime, endTime)
	if err != nil {
		return nil, err
	}
	results := entity.MeetingModel.FindBy(func(m *entity.Meeting) bool {
		if endTime > m.StartTime || m.EndTime > startTime {
			return true
		}
		return false
	})
	return results, nil
}

// DeleteAllMeeting delete all meeting of the speecher
func DeleteAllMeeting() error {
	if err := checkIfLoggedin(); err != nil {
		return err
	}
	curUser := entity.CurSessionModel.GetCurUser()
	results := entity.MeetingModel.FindBy(func(m *entity.Meeting) bool {
		if m.Speecher == curUser {
			return true
		}
		return false
	})
	for _, curMeeting := range results {
		entity.MeetingModel.DeleteMeeting(&curMeeting)
	}
	return nil

}

// DeleteFromMeeting delete a meeting
func DeleteFromMeeting(title string) error {
	if err := checkIfLoggedin(); err != nil {
		return err
	}
	meeting := entity.MeetingModel.FindByTitle(title)
	if meeting.Speecher != entity.CurSessionModel.GetCurUser() {
		// delete participator from a meeting
		entity.MeetingModel.DeleteParticipatorFromMeeting(meeting, entity.CurSessionModel.GetCurUser())
		// check if the participator is 0
		if len(entity.MeetingModel.FindByTitle(title).Participators) == 0 {
			entity.MeetingModel.DeleteMeeting(meeting)
		}
		return nil
	}
	// delete the whole meeting
	entity.MeetingModel.DeleteMeeting(meeting)
	return nil
}

// AddParticipatorToMeeting add participator to a meeting
func AddParticipatorToMeeting(title string, participatorNames []string) error {
	if err := checkIfLoggedin(); err != nil {
		return err
	}

	curMeeting := entity.MeetingModel.FindByTitle(title)
	if curMeeting == nil {
		return fmt.Errorf("Meeting %s doesn't exist", title)
	}

	for _, participatorName := range participatorNames {
		participator := entity.UserModel.FindByUsername(participatorName)
		if participator == nil {
			return fmt.Errorf("Participator %s doesn't exist", participatorName)
		}

		results := entity.MeetingModel.FindBy(func(m *entity.Meeting) bool {
			// check from speecher field
			if m.Speecher == participatorName {
				return true
			}
			// check from participator field
			for _, participator := range m.Participators {
				if participator == participatorName {
					return true
				}
			}
			return false
		})

		_, err := validateFreeTime(curMeeting.StartTime, curMeeting.EndTime, results)
		if err != nil {
			return err
		}

	}

	for _, participatorName := range participatorNames {
		entity.MeetingModel.AddParticipatorToMeeting(curMeeting, participatorName)
	}

	return nil
}

// DeleteParticipatorFromMeeting delete a participator from the meeting
func DeleteParticipatorFromMeeting(title string, participatorNames []string) error {
	if err := checkIfLoggedin(); err != nil {
		return err
	}
	curMeeting := entity.MeetingModel.FindByTitle(title)
	if curMeeting == nil {
		return fmt.Errorf("Meeting %s doesn't exist", title)
	}

	for _, participatorName := range participatorNames {
		participator := entity.UserModel.FindByUsername(participatorName)
		if participator == nil {
			return fmt.Errorf("Participator %s doesn't exist", participatorName)
		}

		flag := false
		for _, curParticipator := range curMeeting.Participators {
			if curParticipator == participatorName {
				entity.MeetingModel.DeleteParticipatorFromMeeting(curMeeting, curParticipator)
				flag = true
				// check if the participator is 0
				if len(entity.MeetingModel.FindByTitle(title).Participators) == 0 {
					entity.MeetingModel.DeleteMeeting(curMeeting)
				}
				break
			}
		}
		if !flag {
			return fmt.Errorf("Participator %s is not in the meeting %s", participatorName, title)
		}
	}
	return nil
}
