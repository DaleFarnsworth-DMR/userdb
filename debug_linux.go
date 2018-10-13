// Copyright 2017-2018 Dale Farnsworth. All rights reserved.

// Dale Farnsworth
// 1007 W Mendoza Ave
// Mesa, AZ  85210
// USA
//
// dale@farnsworth.org

// This file is part of UserDB.
//
// UserDB is free software: you can redistribute it and/or modify
// it under the terms of version 3 of the GNU General Public License
// as published by the Free Software Foundation.
//
// UserDB is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with UserDB.  If not, see <http://www.gnu.org/licenses/>.

package userdb

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
)

func dprint(v ...interface{}) {
	skip := 1

	_, filename, line, ok := runtime.Caller(skip)

	str := ""
	if ok {
		dir := filepath.Base(filepath.Dir(filename))
		filename = filepath.Base(filename)
		filename = filepath.Join(dir, filename)
		str = fmt.Sprintf("%s:%d", filename, line)
	}

	for i := range v {
		v[i] = fmt.Sprintf("%#v", v[i])

	}

	v = append([]interface{}{str}, v...)

	fmt.Fprintln(os.Stderr, v...)
}

func printStack() {
	fmt.Fprintln(os.Stderr, "start stack trace")
	debug.PrintStack()
	fmt.Fprintln(os.Stderr)
}

func logFatalf(s string, v ...interface{}) {
	log.Fatalf(s, v...)
}

func logFatal(v ...interface{}) {
	log.Fatal(v...)
}

func logPrintf(s string, v ...interface{}) {
	log.Printf(s, v...)
}

func logPrint(v ...interface{}) {
	log.Print(v...)
}
