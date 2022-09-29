package sherbet

import (
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/v12"
)

// BaseResponse is struct
type BaseResponse struct {
	StatusCode int
	Body       *iris.Map
	Error      *error
}

// returnResponseGet response get to the client
func returnResponseGet(ctx iris.Context, baseResponse *BaseResponse) {
	ctx.JSON(*baseResponse.Body)
}

// returnResponseError response error to the client
func returnResponseError(ctx iris.Context, baseResponse *BaseResponse) {
	ctx.StopExecution()
	ctx.StatusCode(baseResponse.StatusCode)
	ctx.WriteString((*baseResponse.Error).Error())
}

// ReturnResponse response get to the client
func ReturnResponse(ctx iris.Context, baseResponse *BaseResponse) {
	if baseResponse.StatusCode == http.StatusOK {
		returnResponseGet(ctx, baseResponse)
	} else {
		returnResponseError(ctx, baseResponse)
	}
}

// BuildResponseGet return a get response
func BuildResponseGet(body *iris.Map) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusOK,
		Body:       body,
	}
}

// BuildResponseUnexpectedWrong return a unexpected wrong response
func BuildResponseUnexpectedWrong(err *error) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusBadRequest,
		// ErrorType:  "A unexpected wrong occurred",
		Error: err,
	}
}

// BuildResponseParametersWrong return a parameters wrong response
func BuildResponseParametersWrong(err *error) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusBadRequest,
		Body:       nil,
		// ErrorType:  "Parameters is wrong",
		Error: err,
	}
}

// BuildResponseExecuteSQLWrong return a execute sql wrong response
func BuildResponseExecuteSQLWrong(err *error) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusBadRequest,
		Body:       nil,
		// ErrorType:  "Wrong response when execute sql",
		Error: err,
	}
}

// BuildResponseBuildSQLWrong return a build sql wrong response
func BuildResponseBuildSQLWrong(err *error) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusBadRequest,
		Body:       nil,
		// ErrorType:  "Wrong response when build sql",
		Error: err,
	}
}

// CombineUpdateSetMap combine two array to a map for update's set method
func CombineUpdateSetMap(keys []string, values []interface{}) map[string]interface{} {
	var result = make(map[string]interface{}, len(keys))

	for key, val := range keys {
		result[val] = values[key]
	}

	return result
}

// RetrieveList is func
func RetrieveList(
	database *sqlx.DB,
	response *BaseResponse,
	table *string,
	fields *[]string,
	where *interface{},
	orderBy *[]string,
	limit *uint64,
	offset *uint64,
	data interface{},
) {
	if sql, arguments, err := squirrel.StatementBuilder.
		Select((*fields)...).
		From(*table).
		Where(*where).
		Limit(*limit).
		Offset(*offset).
		OrderBy((*orderBy)...).
		ToSql(); err != nil {
		response = BuildResponseBuildSQLWrong(&err)
	} else {
		if err := database.Select(data, sql, arguments...); err != nil {
			response = BuildResponseExecuteSQLWrong(&err)
		} else {
		}
	}
}

// RetrieveTotal is func
func RetrieveTotal(
	database *sqlx.DB,
	response *BaseResponse,
	table *string,
	field *string,
	where *interface{},
	data interface{},
) {
	if sql, arguments, err := squirrel.StatementBuilder.
		Select(*field).
		From(*table).
		Where(*where).
		Limit(1).
		ToSql(); err != nil {
		response = BuildResponseBuildSQLWrong(&err)
	} else {
		if err = database.Get(data, sql, arguments...); err != nil {
			response = BuildResponseExecuteSQLWrong(&err)
		} else {
		}
	}
}

// RetrieveDetail is func
func RetrieveDetail(
	database *sqlx.DB,
	response *BaseResponse,
	table *string,
	fields *[]string,
	where *interface{},
	data interface{},
) {
	if sql, arguments, err := squirrel.StatementBuilder.
		Select((*fields)...).
		From(*table).
		Where(*where).
		ToSql(); err != nil {
		response = BuildResponseBuildSQLWrong(&err)
	} else {
		err := database.Get(data, sql, arguments...)
		if err != nil {
			response = BuildResponseExecuteSQLWrong(&err)
		} else {
		}
	}
}
