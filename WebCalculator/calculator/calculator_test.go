package calculator

import "testing"

func TestCalculatorResult(t *testing.T) {
	type args struct {
		exp string
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "positive case",
			args: args{
				exp: "23 + 25 - 20 * 9 / 6",
			},
			want:    42,
			wantErr: false,
		},
		{
			name: "Float case",
			args: args{
				exp: "23 - 20 * 7 / 2",
			},
			want:    10.5,
			wantErr: false,
		},
		{
			name: "error check case",
			args: args{
				exp: "12 / 0",
			},
			want:    -999,
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculatorResult(tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculatorResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculatorResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
