package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		level    string
		expected LogLevel
	}{
		{"1", FATAL}, {"FATAL", FATAL},
		{"2", ERROR}, {"ERROR", ERROR},
		{"3", WARN}, {"WARN", WARN},
		{"4", INFO}, {"INFO", INFO},
		{"5", DEBUG}, {"DEBUG", DEBUG},
		{"6", TRACE}, {"TRACE", TRACE},
		{"invalid", INFO},
	}
	for _, tt := range tests {
		l := NewLogger(tt.level, &bytes.Buffer{})
		if l.level != tt.expected {
			t.Errorf("NewLogger(%q) = %v, want %v", tt.level, l.level, tt.expected)
		}
	}
}

func TestLoggerLevels(t *testing.T) {
	buf := &bytes.Buffer{}
	l := NewLogger("TRACE", buf)
	
	l.Fatalf("fatal")
	l.Errorf("error")
	l.Warnf("warn")
	l.Infof("info")
	l.Debugf("debug")
	l.Tracef("trace")
	
	out := buf.String()
	if !strings.Contains(out, "[FATAL]") || !strings.Contains(out, "[TRACE]") {
		t.Error("Expected all log levels to be written")
	}
}

func TestLoggerFiltering(t *testing.T) {
	buf := &bytes.Buffer{}
	l := NewLogger("ERROR", buf)
	l.Infof("should not appear")
	if buf.Len() > 0 {
		t.Error("INFO should be filtered at ERROR level")
	}
}

func TestLoggerNoColor(t *testing.T) {
	buf := &bytes.Buffer{}
	l := NewLogger("INFO", buf)
	
	// Test without color code
	l.Log(INFO, "TEST", "message")
	if !strings.Contains(buf.String(), "message") {
		t.Error("Expected message to be logged without color")
	}
	
	// Test with color code
	buf.Reset()
	l.Log(INFO, "TEST", "colored", "\x1b[31m")
	if !strings.Contains(buf.String(), "colored") || !strings.Contains(buf.String(), "\x1b[31m") {
		t.Error("Expected colored message to be logged")
	}
	
	// Test filtered level
	buf.Reset()
	l.Log(DEBUG, "DEBUG", "filtered")
	if buf.Len() > 0 {
		t.Error("DEBUG should be filtered at INFO level")
	}
}
