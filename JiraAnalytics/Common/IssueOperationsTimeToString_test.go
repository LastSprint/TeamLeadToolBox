package Common

import (
	"gitlab.com/surfstudio/infrastructure/spa/spa-backend-jira-packages/models"
	"testing"
)

func TestGroupByAssigneeWorksRight(t *testing.T) {
	// Arrange

	issues := []models.IssueEntity{
		{
			Fields:    models.IssueFieldsEntity{
				Assignee:  models.UserReferenceEntity{
					Name:        "1",
				},
			},
		},
		{
			Fields:    models.IssueFieldsEntity{
				Assignee:  models.UserReferenceEntity{
					Name:        "4",
				},
			},
		},
		{
			Fields:    models.IssueFieldsEntity{
				Assignee:  models.UserReferenceEntity{
					Name:        "2",
				},
			},
		},
		{
			Fields:    models.IssueFieldsEntity{
				Assignee:  models.UserReferenceEntity{
					Name:        "4",
				},
			},
		},
		{
			Fields:    models.IssueFieldsEntity{
				Assignee:  models.UserReferenceEntity{
					Name:        "1",
				},
			},
		},
	}

	// Act

	result := GroupByAssignee(issues)

	// Assert

	if len(result) != 3 {
		t.FailNow()
	}

	if len(result["1"]) != 2 {
		t.FailNow()
	}

	if len(result["4"]) != 2 {
		t.FailNow()
	}

	if len(result["2"]) != 1 {
		t.FailNow()
	}
}

func TestGroupByAssigneeExcludeEmptyNames(t *testing.T) {
	// Arrange

	issues := []models.IssueEntity{
		{
			Fields:    models.IssueFieldsEntity{
				Assignee:  models.UserReferenceEntity{
					Name:        "",
				},
			},
		},
		{
			Fields:    models.IssueFieldsEntity{
				Assignee:  models.UserReferenceEntity{
					Name:        "1",
				},
			},
		},
		{
			Fields:    models.IssueFieldsEntity{
				Assignee:  models.UserReferenceEntity{
					Name:        "2",
				},
			},
		},
	}

	// Act

	result := GroupByAssignee(issues)

	// Assert

	if len(result) != 2 {
		t.FailNow()
	}

	if len(result["1"]) != 1 {
		t.FailNow()
	}

	if len(result["2"]) != 1 {
		t.FailNow()
	}
}

func TestTimeToStringReturnRightValueFor1h23m(t *testing.T) {
	// Arrange

	seconds := 1 * 60 * 60 + 23 * 60

	// Act

	result := TimeToStringView(seconds)

	// Assert

	if result != "1 h 23 m" {
		t.FailNow()
	}
}

func TestTimeToStringReturnRightValueFor23mAnd1Sec(t *testing.T) {
	// Arrange

	seconds := 23 * 60 + 1

	// Act

	result := TimeToStringView(seconds)

	// Assert

	if result != "23 m" {
		t.FailNow()
	}
}

func TestTimeToStringReturnRightValueFor1Sec(t *testing.T) {
	// Arrange

	seconds := 1

	// Act

	result := TimeToStringView(seconds)

	// Assert

	if result != "1 sec" {
		t.Log(result)
		t.FailNow()
	}
}