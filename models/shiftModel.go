// So actually idea in shift is to make them fair for employee's and for employer
// Employee starts work and click button, on panel, that's saying "Start shift"
// You need to input expected work time e.g. 8 hours
// After shift click button End shift
//
// You need to input actual work time e.g. 7 hours 
// And you will see how much you worked and how much you should work
// If you worked more than you should, you will get extra money
// If you worked less than you should, you will get less money
// If you worked exactly as you should, you will get normal money
//
// So, you need to create 2 functions
// 1. StartShift
// 2. EndShift
//
// StartShift should create new shift with start time and expected work time
// EndShift should update shift with end time and worked time
//
// You need to create 2 models
// 1. Shift
// 2. ShiftUpdate
//
// Shift should have fields:
// 1. ID
// 2. Start
// 3. ShouldWork
// 4. End
// 5. Worked
//
// ShiftUpdate should have fields:
// 1. Start
// 2. ShouldWork
// 3. End
// 4. Worked

// format of shift date should be e.g "2022-01-02 15:04:05"

package models

import (
	"time"

	"github.com/google/uuid"
)

// Shift struct is a model of shift data fields
type Shift struct {
	Base
	Employee uuid.UUID `json:"user_id" binding:"required"`
	Start *time.Time `json:"shift_start" sql:"DEFAULT:current_timestamp"` 
	ShouldWork *time.Time `json:"work_time" binding:"required"`

	End *time.Time `json:"shift_end"`
	Worked *time.Time `json:"worked"`
	
	StartDescription string `json:"start_description"`  // additional info about work/being late etc.
	EndDescription string `json:"end_description"`  // additional info about work/being late etc.
}