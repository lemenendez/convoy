package convoy_test

import (
	"github.com/lemenendez/convoy"
	"testing"
)

const srvName string = "dbserver1"
const usrName string = "root"
const usrPass string = "password"

var options = []struct {
	opt      convoy.Options
	expected string
}{
	{
		opt:      convoy.NewOps(convoy.Options{Host: srvName, User: usrName, Pass: usrPass}),
		expected: "root:password@tcp(dbserver1:3306)/?parseTime=false&multiStatements=false",
	},
	{
		opt:      convoy.NewOps(convoy.Options{Host: srvName, User: usrName, Pass: usrPass, MultiStatements: true}),
		expected: "root:password@tcp(dbserver1:3306)/?parseTime=false&multiStatements=true",
	},
	{
		opt:      convoy.NewOps(convoy.Options{Host: srvName, User: usrName, Pass: usrPass, ParseTime: true, MultiStatements: true}),
		expected: "root:password@tcp(dbserver1:3306)/?parseTime=true&multiStatements=true",
	},
}

func TestOptons(t *testing.T) {
	for _, s := range options {
		dsn := convoy.NewDSN(s.opt)
		if dsn != s.expected {
			t.Error("dsn is not eq to expected:", s.expected, ":[", dsn, "]")
		}
	}

}
