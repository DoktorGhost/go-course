package main

import (
	"bytes"
	"os"
	"testing"
)

func TestOperateTV(t *testing.T) {
	tests := []struct {
		name           string
		tv             TVer
		expectedOutput string
	}{
		{
			name: "Samsung TV",
			tv: &Samsunger{
				status: true,
				model:  "Samsung XL-100500",
			},
			expectedOutput: "Model: Samsung XL-100500\nIs TV on? true\nTurned off. Is TV on? false\nTurned on. Is TV on? true\n",
		},
		{
			name: "LG TV",
			tv: &LGer{
				status: true,
				model:  "LG SmartTV-3000",
			},
			expectedOutput: "Model: LG SmartTV-3000\nIs TV on? true\nTurned off. Is TV on? false\nTurned on. Is TV on? true\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем pipe для перехвата вывода
			r, w, _ := os.Pipe()
			originalStdout := os.Stdout // Сохраняем оригинальный stdout
			os.Stdout = w               // Подменяем stdout на pipe

			// Выполняем функцию
			operateTV(tt.tv)

			// Закрываем писатель и восстанавливаем оригинальный stdout
			w.Close()
			os.Stdout = originalStdout

			// Читаем из pipe
			var buf bytes.Buffer
			buf.ReadFrom(r)
			got := buf.String()

			if got != tt.expectedOutput {
				t.Errorf("operateTV() = %q; want %q", got, tt.expectedOutput)
			}
		})
	}
}

func TestLGHub(t *testing.T) {
	tv := &LGer{}
	expected := "LG Hub"
	got := tv.LGHub()
	if got != expected {
		t.Errorf("LGHub() = %q; want %q", got, expected)
	}
}

func TestSamsungHub(t *testing.T) {
	tv := &Samsunger{}
	expected := "Samsung Hub"
	got := tv.SamsungHub()
	if got != expected {
		t.Errorf("SamsungHub() = %q; want %q", got, expected)
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)
	expected := "Model: Samsung XL-100500\nIs TV on? true\nTurned off. Is TV on? false\n" +
		"Turned on. Is TV on? true\n\nModel: LG SmartTV-3000\nIs TV on? true\nTurned off. Is TV on? false\nTurned on. Is TV on? true\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
