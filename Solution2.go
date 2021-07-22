package main
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
func averageRating(wg *sync.WaitGroup,totalRating *int){
	*totalRating+= rand.Intn(10)+1
	r:= rand.Intn(10)
	time.Sleep(time.Duration(r)*time.Microsecond)
	wg.Done()

}
func main() {
	var wg sync.WaitGroup
	totalRating:=0
	for i:=1;i<=200;i++ {
		wg.Add(1)
		go averageRating(&wg,&totalRating)
	}
	averageRating:=totalRating/200
	fmt.Println(averageRating)
	wg.Wait()
}