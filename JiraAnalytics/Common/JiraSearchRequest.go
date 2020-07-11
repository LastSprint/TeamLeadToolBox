package Common

import (
	jsrv "github.com/LastSprint/JiraGoIssues/services"
	"strconv"
)

type JiraSearchRequest struct {
	Wrapped jsrv.RequestConvertible
	Sprint string
}

func (req JiraSearchRequest) GetAdditionFields() []jsrv.JiraField {
	return req.Wrapped.GetAdditionFields()
}

func (req JiraSearchRequest) GetUseOnlyAdditionalFields() bool {
	return req.Wrapped.GetUseOnlyAdditionalFields()
}

// MakeJiraRequest конвертирует структуру в строку JQL запроса.
func (req JiraSearchRequest) MakeJiraRequest() string {

	result := []string{}

	if len(req.Sprint) != 0 {

		sprintVal := req.Sprint

		if _, err := strconv.Atoi(req.Sprint); err != nil {
			sprintVal = `"` + sprintVal +`"`
		}

		result = append(result, "Sprint = " + sprintVal)
	}

	wrappedStr := req.Wrapped.MakeJiraRequest()

	if len(wrappedStr) != 0 {
		result = append(result, wrappedStr)
	}

	return joinByCharacter(result, " and ", "")
}

// JoinByCharacter объединяет массив строк через символ `character`
// и может "окружить" каждый элемент массива значением`surroundBy`
func joinByCharacter(arr []string, delim string, surroundBy string) string {
	result := ""

	arrLen := len(arr)

	for i, item := range arr {

		result += surroundBy + item + surroundBy

		if i < arrLen-1 {
			result += delim
		}
	}

	return result
}
