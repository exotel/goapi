package helpers

import "log"

func protect(g func()) {
	defer func() {
		log.Println("done") // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}
