package neoKg

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"io"
)

type ServiceGroup struct {
	// Code generated by github.com/flipped-aurora/gin-vue-admin/server Begin; DO NOT EDIT.

	NeoPersonService
	// Code generated by github.com/flipped-aurora/gin-vue-admin/server End; DO NOT EDIT.
}

func unsafeClose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		global.GVA_LOG.Fatal(fmt.Errorf("could not close resource: %w", err).Error())
	}
}
