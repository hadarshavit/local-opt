// +build debug

package assert

import "log"

func Assert(should bool, msg ...interface{})  {
	if !should {
		log.Fatal(should, msg)
		panic("Assert Failure")
	}
}
