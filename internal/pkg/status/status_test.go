package status_test

import (
	"encoding/json"
	"testing"

	"github.com/pthethanh/robusta/internal/pkg/status"
)

func TestStatusMarshal(t *testing.T) {
	s := status.New(200, 200, "ok")
	b, err := json.Marshal(s)
	if err != nil {
		t.Error(err)
		return
	}
	wants := `{"code":200,"message":"ok","status":200}`
	if string(b) != wants {
		t.Errorf("got %s, wants %s", string(b), wants)
	}
}

func TestStatusUnmarshal(t *testing.T) {
	msg := `{"code":200,"message":"ok","status":200}`
	var s status.Status
	if err := json.Unmarshal([]byte(msg), &s); err != nil {
		t.Error(err)
		return
	}
	wants := status.New(200, 200, "ok")
	if s.Code() != wants.Code() || s.Message() != wants.Message() || s.Status() != wants.Status() {
		t.Errorf("got %v, wants %v", s, wants)
	}
}
