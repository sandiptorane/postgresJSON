package validate

import (
	"os"
	"testing"
)

func TestIsJSON(t *testing.T) {
	var test = []struct {
		Name string
		Data string
		expected bool
	}{
		{
			Name :"Fail",
			Data: `customer`,
			expected: false,
		},
		{
			Name: "success",
			Data: `{"name":"sandip"}`,
			expected: true,
		},
		{
			Name: "fail",
			Data:`{"name":sandip}`,
			expected: false,
		},
		{
			Name: "Fail",
			Data : `"hello"`,
			expected: false,
		},
		{
			Name: "success with int",
			Data: `{"name":12}`,
			expected: true,
		},
		{
			Name: "fail on wrong input",
			Data: `{"user":,"pass":"123"}`,
			expected: false,
		},
		{
			Name: "success with multiple inputs",
			Data: `[{"user":"sandip","pass":"pass123"},
					{"user":"pankj","pass":"pass123"}]`,
			expected: true,
		},
	}

	for i := range test{
		tc:= test[i]
		t.Run(tc.Name,func(t *testing.T){
			actual := IsJSON(tc.Data)
			if tc.expected!=actual{
				t.Error("expected:",tc.expected,"got:",actual)
				os.Exit(1)
			}
		})
	}

}

func TestUnmarshal(t *testing.T) {
	data := `"user"`
	Unmarshal(data)
}