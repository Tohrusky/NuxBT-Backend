package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteCountBinary(t *testing.T) {
	type args struct {
		b uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "B",
			args: args{b: 103},
			want: "103 B",
		},
		{
			name: "KB",
			args: args{b: 1024},
			want: "1.0 kB",
		},
		{
			name: "MB",
			args: args{b: 1024 * 1024 * 3},
			want: "3.1 MB",
		},
		{
			name: "GB",
			args: args{b: 1024 * 1024 * 1024 * 3},
			want: "3.2 GB",
		},
		{
			name: "TB",
			args: args{b: 1024 * 1024 * 1024 * 1024 * 3},
			want: "3.3 TB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ByteCountBinary(tt.args.b), "ByteCountBinary(%v)", tt.args.b)
		})
	}
}
