/*
 * Vetstoria Booking API
 *
 * Booking API for 3rd party consumption
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

type AvailabilityResponseSlotsDto struct {
	ScheduleId int32 `json:"scheduleId,omitempty"`
	Slot time.Time `json:"slot,omitempty"`
}
