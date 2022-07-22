package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_NOT_EXIST_ACTIONS       = 10011
	ERROR_CHECK_EXIST_ACTION_FAIL = 10012
	ERROR_ADD_ACTION_FAIL         = 10013
	ERROR_DELETE_ACTION_FAIL      = 10014
	ERROR_EDIT_ACTION_FAIL        = 10015
	ERROR_COUNT_ACTIONS_FAIL      = 10016
	ERROR_GET_ACTIONS_FAIL        = 10017
	ERROR_GET_ACTION_FAIL         = 10018

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)