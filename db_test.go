package main

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/SAP/go-hdb/driver"
)

func Test_GetHanaVersion_OK(t *testing.T) {
	/*first, create the mock DB*/
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	version := string("200.123.11")
	rows := sqlmock.NewRows([]string{"VERSION"}).AddRow(version)
	mock.ExpectQuery("SELECT VERSION FROM \"SYS\".\"M_DATABASE\"").WillReturnRows(rows)

	g := gsConn{Conn: db, Connected: true}

	res, err := g.GetHanaVersion()
	if err != nil {
		t.Error("query failed")
	}
	if res != version {
		t.Errorf("db returned %s but expected %s\n", res, version)
	}
}

func Test_GetHanaVersion_Query_Error(t *testing.T) {
	/*first, create the mock DB*/
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT VERSION FROM \"SYS\".\"M_DATABASE\"").WillReturnError(fmt.Errorf("some error"))

	g := gsConn{Conn: db, Connected: true}

	_, err = g.GetHanaVersion()
	if err == nil {
		t.Error("mocksql sent error but function didn not report it")
	}
}

func Test_gsConn_UserHasMonRole_OK(t *testing.T) {
	type args struct {
		user string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.

		{"Good 01", gsConn{}, args{"David"}, true, false},
		{"Good 02", gsConn{}, args{"1234"}, true, false},
		{"Good 02", gsConn{}, args{"CAPS_USER"}, true, false},
	}

	for _, tt := range tests {
		/*configure mock db*/
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow("1")
		q1 := fmt.Sprintf("SELECT COUNT(GRANTEE) FROM \"SYS\".\"GRANTED_ROLES\" WHERE GRANTEE = '%s' AND ROLE_NAME = 'MONITORING'", strings.ToUpper(tt.args.user))
		mock.ExpectQuery(q1).WillReturnRows(rows)
		tt.g.Conn = db

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.UserHasMonRole(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("gsConn.UserHasMonRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gsConn.UserHasMonRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gsConn_UserHasMonRole_No_Role(t *testing.T) {
	type args struct {
		user string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.

		{"No Role 01", gsConn{}, args{"David"}, false, false},
		{"No Role 02", gsConn{}, args{"1234"}, false, false},
		{"No Role 02", gsConn{}, args{"CAPS_USER"}, false, false},
	}

	for _, tt := range tests {
		//configure mock db
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow("0")
		q1 := fmt.Sprintf("SELECT COUNT(GRANTEE) FROM \"SYS\".\"GRANTED_ROLES\" WHERE GRANTEE = '%s' AND ROLE_NAME = 'MONITORING'", strings.ToUpper(tt.args.user))
		mock.ExpectQuery(q1).WillReturnRows(rows)
		tt.g.Conn = db

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.UserHasMonRole(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("gsConn.UserHasMonRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gsConn.UserHasMonRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gsConn_UserHasMonRole_DB_Error(t *testing.T) {
	type args struct {
		user string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.

		{"DB Error 01", gsConn{}, args{"David"}, false, true},
		{"DB Error 02", gsConn{}, args{"1234"}, false, true},
		{"DB Error 03", gsConn{}, args{"CAPS_USER"}, false, true},
	}

	for _, tt := range tests {
		//configure mock db
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		q1 := fmt.Sprintf("SELECT COUNT(GRANTEE) FROM \"SYS\".\"GRANTED_ROLES\" WHERE GRANTEE = '%s' AND ROLE_NAME = 'MONITORING'", strings.ToUpper(tt.args.user))
		mock.ExpectQuery(q1).WillReturnError(fmt.Errorf("some error"))
		tt.g.Conn = db

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.UserHasMonRole(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("gsConn.UserHasMonRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gsConn.UserHasMonRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gsConn_CheckSchema_Exists(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		want    bool
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS"}, true, false},
		{"Good 02", gsConn{}, args{"sTRANG__Naame"}, true, false},
		{"Good 03", gsConn{}, args{"12333"}, true, false},
	}
	for _, tt := range tests {
		//configure mock db
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow("1")
		q1 := fmt.Sprintf("SELECT COUNT(SCHEMA_NAME) FROM \"PUBLIC\".\"SCHEMAS\" WHERE SCHEMA_NAME = '%s'", tt.args.schema)
		mock.ExpectQuery(q1).WillReturnRows(rows)
		tt.g.Conn = db

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.CheckSchema(tt.args.schema)
			if (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CheckSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gsConn.CheckSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gsConn_CheckSchema_Not_Exist(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		want    bool
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS"}, false, false},
		{"Good 02", gsConn{}, args{"sTRANG__Naame"}, false, false},
		{"Good 03", gsConn{}, args{"12333"}, false, false},
	}
	for _, tt := range tests {
		//configure mock db
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow("0")
		q1 := fmt.Sprintf("SELECT COUNT(SCHEMA_NAME) FROM \"PUBLIC\".\"SCHEMAS\" WHERE SCHEMA_NAME = '%s'", tt.args.schema)
		mock.ExpectQuery(q1).WillReturnRows(rows)
		tt.g.Conn = db

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.CheckSchema(tt.args.schema)
			if (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CheckSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gsConn.CheckSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gsConn_CheckSchema_DB_Error(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		want    bool
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS"}, false, true},
		{"Good 02", gsConn{}, args{"sTRANG__Naame"}, false, true},
		{"Good 03", gsConn{}, args{"12333"}, false, true},
	}
	for _, tt := range tests {
		//configure mock db
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		q1 := fmt.Sprintf("SELECT COUNT(SCHEMA_NAME) FROM \"PUBLIC\".\"SCHEMAS\" WHERE SCHEMA_NAME = '%s'", tt.args.schema)
		mock.ExpectQuery(q1).WillReturnError(fmt.Errorf("some error"))
		tt.g.Conn = db

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.CheckSchema(tt.args.schema)
			if (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CheckSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gsConn.CheckSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gsConn_dropSchema_OK(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS"}, false},
		{"Good 02", gsConn{}, args{"GtrShop"}, false},
		{"Good 03", gsConn{}, args{"122!"}, false},
		{"Good 04", gsConn{}, args{"_just_lowers"}, false},
	}
	for _, tt := range tests {
		/*configure mock db*/
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		t.Run(tt.name, func(t *testing.T) {
			q1 := fmt.Sprintf("DROP SCHEMA \"%s\" CASCADE", tt.args.schema)
			mock.ExpectExec(q1).WillReturnResult(sqlmock.NewResult(0, 0))
			tt.g.Conn = db

			if err := tt.g.DropSchema(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.dropSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_DropSchema_DB_Error(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS"}, true},
		{"Good 02", gsConn{}, args{"GtrShop"}, true},
		{"Good 03", gsConn{}, args{"122!"}, true},
		{"Good 04", gsConn{}, args{"_just_lowers"}, true},
	}
	for _, tt := range tests {
		/*configure mock db*/
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		t.Run(tt.name, func(t *testing.T) {
			q1 := fmt.Sprintf("DROP SCHEMA \"%s\" CASCADE", tt.args.schema)
			mock.ExpectExec(q1).WillReturnError(fmt.Errorf("some error"))
			tt.g.Conn = db

			if err := tt.g.DropSchema(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.DropSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_TransactExecRows(t *testing.T) {
	type args struct {
		statements []string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{GetMasterDataStatements("test")}, false},
		{"Good 02", gsConn{}, args{GetSchemaStatements("test")}, false},
	}
	for _, tt := range tests {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		mock.ExpectBegin()
		for _, statement := range tt.args.statements {
			mock.ExpectExec(statement).WillReturnResult(sqlmock.NewResult(0, 0))
		}
		mock.ExpectCommit()

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.TransactExecRows(tt.args.statements); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.TransactExecRows() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_TransactExec_DB_Error_Rollback_OK(t *testing.T) {
	type args struct {
		statements []string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Rollback 01", gsConn{}, args{GetMasterDataStatements("test")}, true},
		{"Rollback 02", gsConn{}, args{GetSchemaStatements("test")}, true},
	}
	for _, tt := range tests {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		mock.ExpectBegin()
		mock.ExpectExec(tt.args.statements[0]).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.TransactExecRows(tt.args.statements); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.TransactExecRows() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_TransactExec_DB_Error_Rollback_Fail(t *testing.T) {
	type args struct {
		statements []string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Rollback Fail 01", gsConn{}, args{GetMasterDataStatements("test")}, true},
		{"Rollback Fail 02", gsConn{}, args{GetSchemaStatements("test")}, true},
	}
	for _, tt := range tests {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		mock.ExpectBegin()
		mock.ExpectExec(tt.args.statements[0]).WillReturnError(fmt.Errorf("EXEC FAILED"))
		mock.ExpectRollback().WillReturnError(fmt.Errorf("ROLLBACK FAILED"))

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.TransactExecRows(tt.args.statements); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.TransactExecRows() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_TransactExec_Trnx_Start_Fail(t *testing.T) {
	type args struct {
		statements []string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Trnx Start Fail 01", gsConn{}, args{GetMasterDataStatements("test")}, true},
		{"Trnx Start Fail  02", gsConn{}, args{GetSchemaStatements("test")}, true},
	}
	for _, tt := range tests {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		mock.ExpectBegin().WillReturnError(fmt.Errorf("Failed to start transaction"))

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.TransactExecRows(tt.args.statements); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.TransactExecRows() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_TransactExec_Commit_Fail(t *testing.T) {
	type args struct {
		statements []string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Commit Fail 01", gsConn{}, args{GetMasterDataStatements("test")}, true},
		{"Commit Fail 02", gsConn{}, args{GetSchemaStatements("test")}, true},
	}
	for _, tt := range tests {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		mock.ExpectBegin()
		for _, statement := range tt.args.statements {
			mock.ExpectExec(statement).WillReturnResult(sqlmock.NewResult(0, 0))
		}
		mock.ExpectCommit().WillReturnError(fmt.Errorf("Commit Failed"))

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.TransactExecRows(tt.args.statements); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.TransactExecRows() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_CreateSchema_Good(t *testing.T) {
	type args struct {
		schema string
		drop   bool
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS", true}, false},
		{"Good 02", gsConn{}, args{"gs", true}, false},
		{"Good 03", gsConn{}, args{"MySchema", true}, false},
		{"Good 04", gsConn{}, args{"012345", true}, false},
		{"Good 05", gsConn{}, args{"!!", true}, false},
	}
	for _, tt := range tests {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		rows := sqlmock.NewRows([]string{"COUNT(SCHEMA_NAME)"}).AddRow("1")
		q1 := fmt.Sprintf("SELECT COUNT(SCHEMA_NAME) FROM \"PUBLIC\".\"SCHEMAS\" WHERE SCHEMA_NAME = '%s'", tt.args.schema)
		q2 := fmt.Sprintf("DROP SCHEMA \"%s\" CASCADE", tt.args.schema)
		mock.ExpectQuery(q1).WillReturnRows(rows)
		mock.ExpectExec(q2).WillReturnResult(sqlmock.NewResult(0, 0))

		mock.ExpectBegin()
		for _, statement := range GetSchemaStatements(tt.args.schema) {
			mock.ExpectExec(statement).WillReturnResult(sqlmock.NewResult(0, 0))
		}
		mock.ExpectCommit()

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.CreateSchema(tt.args.schema, tt.args.drop); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CreateSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_CreateSchema_Error(t *testing.T) {
	type args struct {
		schema string
		drop   bool
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS", true}, true},
		{"Good 02", gsConn{}, args{"gs", true}, true},
		{"Good 03", gsConn{}, args{"MySchema", true}, true},
		{"Good 04", gsConn{}, args{"012345", true}, true},
		{"Good 05", gsConn{}, args{"!!", true}, true},
	}
	for _, tt := range tests {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		mock.ExpectBegin()
		statements := GetSchemaStatements(tt.args.schema)
		mock.ExpectExec(statements[0]).WillReturnError(fmt.Errorf("INSERT ERROR"))

		mock.ExpectRollback()

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.CreateSchema(tt.args.schema, tt.args.drop); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CreateSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_CreateMasterData_OK(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS"}, false},
		{"Good 02", gsConn{}, args{"gs"}, false},
		{"Good 03", gsConn{}, args{"MySchema"}, false},
		{"Good 04", gsConn{}, args{"012345"}, false},
		{"Good 05", gsConn{}, args{"!!"}, false},
	}
	for _, tt := range tests {

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		mock.ExpectBegin()
		for _, statement := range GetMasterDataStatements(tt.args.schema) {
			mock.ExpectExec(statement).WillReturnResult(sqlmock.NewResult(0, 0))
		}
		mock.ExpectCommit()

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.CreateMasterData(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CreateMasterData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_CreateMasterData_Error(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS"}, true},
		{"Good 02", gsConn{}, args{"gs"}, true},
		{"Good 03", gsConn{}, args{"MySchema"}, true},
		{"Good 04", gsConn{}, args{"012345"}, true},
		{"Good 05", gsConn{}, args{"!!"}, true},
	}
	for _, tt := range tests {

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		mock.ExpectBegin()
		statements := GetMasterDataStatements(tt.args.schema)
		mock.ExpectExec(statements[0]).WillReturnError(fmt.Errorf("INSERT ERROR"))
		mock.ExpectRollback()

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.CreateMasterData(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CreateMasterData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_InsertProducts_OK(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS"}, false},
		{"Good 02", gsConn{}, args{"gs"}, false},
		{"Good 03", gsConn{}, args{"MySchema"}, false},
		{"Good 04", gsConn{}, args{"012345"}, false},
		{"Good 05", gsConn{}, args{"!!"}, false},
	}
	for _, tt := range tests {

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		mock.ExpectBegin()
		for _, statement := range GetProductStatements(tt.args.schema) {
			mock.ExpectExec(statement).WillReturnResult(sqlmock.NewResult(0, 0))
		}
		mock.ExpectCommit()

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.InsertProducts(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.InsertProducts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_InsertProducts_DB_Error(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS"}, true},
		{"Good 02", gsConn{}, args{"gs"}, true},
		{"Good 03", gsConn{}, args{"MySchema"}, true},
		{"Good 04", gsConn{}, args{"012345"}, true},
		{"Good 05", gsConn{}, args{"!!"}, true},
	}
	for _, tt := range tests {

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		mock.ExpectBegin()
		statements := GetProductStatements(tt.args.schema)
		mock.ExpectExec(statements[0]).WillReturnError(fmt.Errorf("INSERT ERROR"))
		mock.ExpectRollback()

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.InsertProducts(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.InsertProducts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_CreateSchema_Schema_Not_Present(t *testing.T) {
	type args struct {
		schema string
		drop   bool
	}
	tests := []struct {
		name    string
		g       *gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", &gsConn{}, args{"GS", true}, false},
		{"Good 02", &gsConn{}, args{"gs", true}, false},
		{"Good 03", &gsConn{}, args{"MySchema", true}, false},
		{"Good 04", &gsConn{}, args{"012345", true}, false},
		{"Good 05", &gsConn{}, args{"!!", true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.g.Conn = db

			/*Setup results and expects*/
			rows := sqlmock.NewRows([]string{"COUNT(SCHEMA_NAME)"}).AddRow("0")
			q1 := fmt.Sprintf("SELECT COUNT(SCHEMA_NAME) FROM \"PUBLIC\".\"SCHEMAS\" WHERE SCHEMA_NAME = '%s'", tt.args.schema)
			mock.ExpectQuery(q1).WillReturnRows(rows)

			mock.ExpectBegin()

			//setup loop
			for _, statement := range GetSchemaStatements(tt.args.schema) {
				mock.ExpectExec(statement).WillReturnResult(sqlmock.NewResult(0, 0))
			}
			mock.ExpectCommit()

			if err := tt.g.CreateSchema(tt.args.schema, tt.args.drop); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CreateSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_CreateSchema_Schema_Present(t *testing.T) {
	type args struct {
		schema string
		drop   bool
	}
	tests := []struct {
		name    string
		g       *gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", &gsConn{}, args{"GS", true}, false},
		{"Good 02", &gsConn{}, args{"gs", true}, false},
		{"Good 03", &gsConn{}, args{"MySchema", true}, false},
		{"Good 04", &gsConn{}, args{"012345", true}, false},
		{"Good 05", &gsConn{}, args{"!!", true}, false},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.g.Conn = db

			/*Setup results and expects*/
			rows := sqlmock.NewRows([]string{"COUNT(SCHEMA_NAME)"}).AddRow("1")
			q1 := fmt.Sprintf("SELECT COUNT(SCHEMA_NAME) FROM \"PUBLIC\".\"SCHEMAS\" WHERE SCHEMA_NAME = '%s'", tt.args.schema)
			mock.ExpectQuery(q1).WillReturnRows(rows)

			//Found schema expect drop
			q2 := fmt.Sprintf("DROP SCHEMA \"%s\" CASCADE", tt.args.schema)
			mock.ExpectExec(q2).WillReturnResult(sqlmock.NewResult(0, 0))

			mock.ExpectBegin()

			//setup loop
			for _, statement := range GetSchemaStatements(tt.args.schema) {
				mock.ExpectExec(statement).WillReturnResult(sqlmock.NewResult(0, 0))
			}
			mock.ExpectCommit()

			if err := tt.g.CreateSchema(tt.args.schema, tt.args.drop); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CreateSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_CreateSchema_Schema_Drop_False(t *testing.T) {
	type args struct {
		schema string
		drop   bool
	}
	tests := []struct {
		name    string
		g       *gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", &gsConn{}, args{"GS", false}, true},
		{"Good 02", &gsConn{}, args{"gs", false}, true},
		{"Good 03", &gsConn{}, args{"MySchema", false}, true},
		{"Good 04", &gsConn{}, args{"012345", false}, true},
		{"Good 05", &gsConn{}, args{"!!", false}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.g.Conn = db

			/*Setup results and expects*/
			rows := sqlmock.NewRows([]string{"COUNT(SCHEMA_NAME)"}).AddRow("1")
			q1 := fmt.Sprintf("SELECT COUNT(SCHEMA_NAME) FROM \"PUBLIC\".\"SCHEMAS\" WHERE SCHEMA_NAME = '%s'", tt.args.schema)
			mock.ExpectQuery(q1).WillReturnRows(rows)

			if err := tt.g.CreateSchema(tt.args.schema, tt.args.drop); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CreateSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_CreateSchema_Schema_Present_Drop_Error(t *testing.T) {
	type args struct {
		schema string
		drop   bool
	}
	tests := []struct {
		name    string
		g       *gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", &gsConn{}, args{"GS", true}, true},
		{"Good 02", &gsConn{}, args{"gs", true}, true},
		{"Good 03", &gsConn{}, args{"MySchema", true}, true},
		{"Good 04", &gsConn{}, args{"012345", true}, true},
		{"Good 05", &gsConn{}, args{"!!", true}, true},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.g.Conn = db

			/*Setup results and expects*/
			rows := sqlmock.NewRows([]string{"COUNT(SCHEMA_NAME)"}).AddRow("1")
			q1 := fmt.Sprintf("SELECT COUNT(SCHEMA_NAME) FROM \"PUBLIC\".\"SCHEMAS\" WHERE SCHEMA_NAME = '%s'", tt.args.schema)
			mock.ExpectQuery(q1).WillReturnRows(rows)

			//Found schema expect drop
			q2 := fmt.Sprintf("DROP SCHEMA \"%s\" CASCADE", tt.args.schema)
			mock.ExpectExec(q2).WillReturnError(fmt.Errorf("drop error"))

			if err := tt.g.CreateSchema(tt.args.schema, tt.args.drop); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CreateSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_CreateSchema_Schema_Present_DB_Error(t *testing.T) {
	type args struct {
		schema string
		drop   bool
	}
	tests := []struct {
		name    string
		g       *gsConn
		args    args
		wantErr bool
	}{
		{"Good 01", &gsConn{}, args{"GS", true}, true},
		{"Good 02", &gsConn{}, args{"gs", true}, true},
		{"Good 03", &gsConn{}, args{"MySchema", true}, true},
		{"Good 04", &gsConn{}, args{"012345", true}, true},
		{"Good 05", &gsConn{}, args{"!!", true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.g.Conn = db

			/*Setup results and expects*/
			rows := sqlmock.NewRows([]string{"COUNT(SCHEMA_NAME)"}).AddRow("0")
			q1 := fmt.Sprintf("SELECT COUNT(SCHEMA_NAME) FROM \"PUBLIC\".\"SCHEMAS\" WHERE SCHEMA_NAME = '%s'", tt.args.schema)
			mock.ExpectQuery(q1).WillReturnRows(rows)

			/*Fail to start the transaction*/
			mock.ExpectBegin().WillReturnError(fmt.Errorf("transaction error"))

			if err := tt.g.CreateSchema(tt.args.schema, tt.args.drop); (err != nil) != tt.wantErr {
				t.Errorf("gsConn.CreateSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gsConn_WorkerInsertCustomers(t *testing.T) {
	type args struct {
		wid    int
		count  int
		schema string
		wg     *sync.WaitGroup
	}
	tests := []struct {
		name string
		g    gsConn
		args args
	}{
		{"Good 01", gsConn{}, args{0, 100, "Gs", &sync.WaitGroup{}}},
		{"Good 02", gsConn{}, args{0, 1000, "GtrShop", &sync.WaitGroup{}}},
		{"Good 03", gsConn{}, args{0, 500, "TEST", &sync.WaitGroup{}}},
		{"Good 04", gsConn{}, args{0, 1000, "tEST", &sync.WaitGroup{}}},
		{"Good 05", gsConn{}, args{0, 666, "123", &sync.WaitGroup{}}},
	}
	for _, tt := range tests {

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		mock.ExpectBegin()
		for i := 0; i < tt.args.count; i++ {
			mock.ExpectExec("INSERT INTO .+").WillReturnResult(sqlmock.NewResult(1, 1))
		}
		mock.ExpectCommit()

		tt.args.wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			tt.g.WorkerInsertCustomers(tt.args.wid, tt.args.count, tt.args.schema, tt.args.wg)
		})
	}
}

func Test_gsConn_GetCustomerIDs(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		want    []int
		wantErr bool
	}{
		{"Good 01", gsConn{}, args{"GS"}, []int{1, 2, 3}, false},
		{"DB Error 01", gsConn{}, args{"GS"}, []int{}, true},
		{"DB Error 02", gsConn{}, args{"GtrShop"}, []int{}, true},
	}
	for _, tt := range tests {

		/*Setup the DB*/
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		tt.g.Conn = db

		q1 := fmt.Sprintf("SELECT COUNT(ID) FROM \"%s\".\"CUSTOMERS\"", tt.args.schema)
		q2 := fmt.Sprintf("SELECT ID FROM \"%s\".\"CUSTOMERS\"", tt.args.schema)

		/*Set up the scenarios*/
		if tt.name == "Good 01" {
			/*The good one*/

			q1rows := sqlmock.NewRows([]string{"COUNT(ID)"}).AddRow(3)
			q2rows := sqlmock.NewRows([]string{"ID"}).AddRow(1).AddRow(2).AddRow(3)
			mock.ExpectQuery(q1).WillReturnRows(q1rows)
			mock.ExpectQuery(q2).WillReturnRows(q2rows)
		} else if tt.name == "DB Error 01" {
			mock.ExpectQuery(q1).WillReturnError(fmt.Errorf("DB Error"))

		} else if tt.name == "DB Error 02" {
			q1rows := sqlmock.NewRows([]string{"COUNT(ID)"}).AddRow(50)
			mock.ExpectQuery(q1).WillReturnRows(q1rows)
			mock.ExpectQuery(q2).WillReturnError(fmt.Errorf("DB Error"))
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.GetCustomerIDs(tt.args.schema)
			if (err != nil) != tt.wantErr {
				t.Errorf("gsConn.GetCustomerIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gsConn.GetCustomerIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gsConn_GetProductIDs(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name    string
		g       gsConn
		args    args
		wantErr bool
	}{
		{"Good", gsConn{}, args{"GS"}, false},
		{"DB Error 01", gsConn{}, args{"GtrShop"}, true},
		{"DB Error 02", gsConn{}, args{"GtrShop"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			/*Setup the DB*/
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.g.Conn = db

			/*Expected queries*/
			q1 := fmt.Sprintf("SELECT COUNT(ID) FROM \"%s\".\"PRODUCT\"", tt.args.schema)
			q2 := fmt.Sprintf("SELECT ID, RAND_WEIGHT FROM \"%s\".\"PRODUCT\"", tt.args.schema)

			/*Configure db scenarios*/
			if tt.name == "Good" {
				q1rows := sqlmock.NewRows([]string{"COUNT(ID)"}).AddRow(10)
				q2rows := sqlmock.NewRows([]string{"ID", "RAND_WEIGHT"})
				mock.ExpectQuery(q1).WillReturnRows(q1rows)
				mock.ExpectQuery(q2).WillReturnRows(q2rows)
			} else if tt.name == "DB Error 01" {
				mock.ExpectQuery(q1).WillReturnError(fmt.Errorf("DB Error"))
			} else if tt.name == "DB Error 02" {
				q1rows := sqlmock.NewRows([]string{"COUNT(ID)"}).AddRow(10)
				mock.ExpectQuery(q1).WillReturnRows(q1rows)
				mock.ExpectQuery(q2).WillReturnError(fmt.Errorf("DB Error"))
			}

			_, err = tt.g.GetProductIDs(tt.args.schema)
			if (err != nil) != tt.wantErr {
				t.Errorf("gsConn.GetProductIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
