package main

import "testing"

func Test_structures(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			structures()
		})
	}
}

func Test_accessingStructureMembers(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessingStructureMembers()
		})
	}
}

func Test_structuresAsFunctionArguments(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			structuresAsFunctionArguments()
		})
	}
}

func Test_printBook(t *testing.T) {
	type args struct {
		book Books
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printBook(tt.args.book)
		})
	}
}

func Test_printBookbyPointers(t *testing.T) {
	type args struct {
		book *Books
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printBookbyPointers(tt.args.book)
		})
	}
}
