package main


type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int				`json:"page"`
	Fruits []string `json:"fruits"`
}

func main(){
	
	p:= fmt.Println
	bolB, _:= json.Marshal(true)
	p(string(bolB))

}
