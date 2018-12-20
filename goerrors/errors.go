package goerrors

import (
	"database/sql"
	"io"

	"github.com/coolmsg/go-coolmsg"
)

const (
	ERRCODE_SQL_ERR_NO_ROWS = 0xe3040eaf8487b183
	ERRCODE_IO_EOF          = 0x98a2d040496efa20
)

func RegisterGoErrors(reg *coolmsg.Registry) {
	reg.RegisterError(ERRCODE_SQL_ERR_NO_ROWS, func(*coolmsg.Error) error { return sql.ErrNoRows })
	reg.RegisterError(ERRCODE_IO_EOF, func(*coolmsg.Error) error { return io.EOF })
}

func init() {
	RegisterGoErrors(coolmsg.DefaultRegistry)
}
