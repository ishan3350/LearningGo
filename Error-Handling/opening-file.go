package main

//opening file and reading data and also checking for error

	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	f,err := os.Open("test.txt")

	if err == nil{
		defer f.Close()
		data,err := ioutil.ReadAll(f)

		if err == nil{
			fmt.Println(string(data))
		}else{
			fmt.Println(err)
			return
		}
	}else{
		fmt.Println(err)
		return
	}
}