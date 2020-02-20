package gotwilio

import "testing"

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
