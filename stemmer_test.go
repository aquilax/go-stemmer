package stemmer

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestStem(t *testing.T) {
	reader := strings.NewReader(`ката ==> к 2`)
	rules, _ := LoadRulesStream(reader, 1)
	word := Stem("кокошката", rules)
	want := "кокошк"
	if word != want {
		t.Errorf("Stem() = %v, want %v", word, want)
	}
}

func TestLoadRulesStream(t *testing.T) {
	type args struct {
		reader       io.Reader
		stemBoundary int
	}
	tests := []struct {
		name    string
		args    args
		want    Rules
		wantErr bool
	}{
		{
			"Loads rules as expected",
			args{
				strings.NewReader(`ката ==> к 2`),
				1,
			},
			Rules{"ката": "к"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadRulesStream(tt.args.reader, tt.args.stemBoundary)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadRulesStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadRulesStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadRulesIntegration(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping integration test in CI environment")
	}
	type args struct {
		fileName     string
		stemBoundary int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"context1 loads",
			args{
				"stem_rules_context_1.utf8.txt",
				1,
			},
			5033,
			false,
		},
		{
			"context2 loads",
			args{
				"stem_rules_context_2.utf8.txt",
				1,
			},
			22199,
			false,
		},
		{
			"context3 loads",
			args{
				"stem_rules_context_3.utf8.txt",
				1,
			},
			56797,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadRules(tt.args.fileName, tt.args.stemBoundary)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadRules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("LoadRules() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestStemIntegration(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping integration test in CI environment")
	}
	rules, _ := LoadRules("stem_rules_context_1.utf8.txt", 1)
	word := Stem("кокошката", rules)
	want := "кокошк"
	if word != want {
		t.Errorf("Stem() = %v, want %v", word, want)
	}
}
