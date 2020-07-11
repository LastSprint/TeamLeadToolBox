package Common

import (
	"gitlab.com/surfstudio/infrastructure/spa/spa-backend-com-packages/utils"
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

	return utils.JoinByCharacter(result, " and ", "")
}