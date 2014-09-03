package gopinba

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInSlice(t *testing.T) {

	slice := []string{"a", "b", "c", "d", "e"}

	position, exist := inSlice("c", slice)

	assert.True(t, exist)
	assert.Equal(t, 2, position)

	position, exist = inSlice("z", slice)

	assert.False(t, exist)
	assert.Equal(t, position, -1)
}

func BenchmarkPinba(b *testing.B) {

	pinba := New(&Options{})

	var tmr *timer

	for i := 0; i < b.N; i++ {

		request := pinba.Request()

		for i := 0; i < 10; i++ {

			tmr = request.TimerStart(&Tags{"group": "go", "operation": "run"})

			request.TimerStop(tmr)
		}
	}
}
