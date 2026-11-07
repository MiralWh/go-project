package works

import (
	"errors"
	"math"
	"testing"
)

// -------------------------------------------------------------------------
func TestDivide(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr error
	}{
		// позитив
		{name: "положительные", a: 10, b: 2, want: 5},
		{name: "отрицательное a", a: -7, b: 4, want: -1.75},
		{name: "отрицательное b", a: 7, b: -4, want: -1.75},
		{name: "два отрицательных", a: -10, b: -2, want: 5},
		{name: "ноль в числителе", a: 0, b: 5, want: 0},
		{name: "дробный результат", a: 1, b: 3, want: 1.0 / 3.0},
		{name: "большие числа", a: 1e12, b: 1e6, want: 1e6},

		// негатив
		{name: "b == 0", a: 5, b: 0, wantErr: ErrorDividedByZero},
		{name: "a и b == 0", a: 0, b: 0, wantErr: ErrorDividedByZero},
		{name: "b == -0.0", a: 5, b: math.Copysign(0, -1), wantErr: ErrorDividedByZero},
	}

	for _, tc := range tests {

		tc := tc // для старых версий (почитал в сети что иногда нужно)

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := divide(tc.a, tc.b)

			if tc.wantErr != nil {

				if err == nil {
					t.Fatalf("divide(%v, %v): ожидалась ошибка %v, получен nil", tc.a, tc.b, tc.wantErr)
				}
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("divide(%v, %v): ошибка %v, ожидалась %v", tc.a, tc.b, err, tc.wantErr)
				}
				if got != 0 {
					t.Errorf("divide(%v, %v): при ошибке результат должен быть 0, получен %v", tc.a, tc.b, got)
				}
				return
			}

			if err != nil {
				t.Fatalf("Divide(%v, %v): неожиданная ошибка: %v", tc.a, tc.b, err)
			}

			if !approxEqual(got, tc.want, 1e-12) {
				t.Errorf("Divide(%v, %v) = %v, ожидалось %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// -------------------------------------------------------------------------

func TestDivide_ErrorWrapping(t *testing.T) {
	t.Parallel()

	_, err := divide(1, 0)
	if err == nil {
		t.Fatal("ожидалась ошибка, получен nil")
	}
	if !errors.Is(err, ErrorDividedByZero) {
		t.Errorf("errors.Is(err, ErrDivisionByZero) = false; err = %v", err)
	}
	if msg := err.Error(); msg == "" || !contains(msg, "ошибка деления на ноль") /*contains(msg, "lsadfjllsdflksfdlkkl;fsdklf") - тут сфейлится*/ {
		t.Errorf("сообщение об ошибке не содержит контекста: %q", msg)
	}
}

// -------------------------------------------------------------------------

// сравнение с округлением
func approxEqual(a, b, eps float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) {
		return math.IsNaN(a) && math.IsNaN(b)
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) {
		return a == b
	}
	return math.Abs(a-b) <= eps
}

func contains(s, substr string) bool {
	return len(substr) == 0 || (len(s) >= len(substr) && indexOf(s, substr) >= 0)
}

func indexOf(s, substr string) int {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
