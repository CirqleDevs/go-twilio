package gotwilio

import "testing"

// check REALLY carefully that I mapped everything correctly
func TestSmsStatus_String(t *testing.T) {

	if SmsStatusQueued.String() != "queued" {
		t.Error("bad queued value")
	}

	if SmsStatusFailed.String() != "failed" {
		t.Error("bad failed")
	}

	if SmsStatusSent.String() != "sent" {
		t.Error("bad sent")
	}

	if SmsStatusDelivered.String() != "delivered" {
		t.Error("bad delivered")
	}

	if SmsStatusUndelivered.String() != "undelivered" {
		t.Error("bad undelivered")
	}
}

// check REALLY carefully that I mapped everything correctly
func TestVoiceStatus_String(t *testing.T) {

	if VoiceStatusQueued.String() != "queued" {
		t.Error("bad queued")
	}

	if VoiceStatusRinging.String() != "ringing" {
		t.Error("bad ringing")
	}

	if VoiceStatusInProgress.String() != "in-progress" {
		t.Error("bad in-progress")
	}

	if VoiceStatusCompleted.String() != "completed" {
		t.Error("bad completed")
	}

	if VoiceStatusBusy.String() != "busy" {
		t.Error("bad busy")
	}

	if VoiceStatusFailed.String() != "failed" {
		t.Error("bad failed")
	}

	if VoiceStatusNoAnswer.String() != "no-answer" {
		t.Error("bad no-answer")
	}
}
