package response

import (
	database "booking-konstruksi/database/migration"
	"time"
)

type TimeLine struct {
	ID           int        `json:"id"`
	KonstruksiID int        `json:"konstruksi_id"`
	Name         string     `json:"name"`
	Date         *time.Time `json:"date"`
	Description  *string    `json:"description"`
}

func TimeLineResponse(timeLine *database.TimeLine) TimeLine {
	return TimeLine{
		ID:           timeLine.ID,
		KonstruksiID: timeLine.KonstruksiID,
		Name:         timeLine.Name,
		Date:         timeLine.Date,
		Description:  timeLine.Description,
	}
}
