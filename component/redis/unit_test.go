package main

import (
	"testing"
	"time"
)

func TestLock(t *testing.T){
	t.Run("test1", func(t *testing.T) {
		Lock()
		Lock()
		defer UnLock()
		time.Sleep(time.Second * 3)
	})

}
