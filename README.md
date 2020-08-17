# Bitcoin PUB Converter
 <!-- A Bitcoin PUB converter in Golang by #https://github.com/Godtide  inspiration from #https://github.com/codaelux/pubkeyConverter   -->
 
>Requirement

Go env is set apart from GOPATH

> Deployment

 **go get -u "github.com/Godtide/tester/PubConverter.go"**

# Full Example

 package convert

 import ( 

     "fmt"
     "github.com/Godtide/tester/pubConverter.go"
        )
        
func main() {

	pub := "6CUGRUonZSQ4TWtTMmzXdrXDtypWKiKrhko4egpiMZbpiaQL2jkwSB1icqYh2cfDfVxdx4df189oLKnC5fSwqPfgyP3hooxujYzAu3fDVmz"

	xpub := pubConverter.Xpub(pub)
    fmt.Println(xpub)
    
    ypub = pubConverter.Ypub(pub)
    fmt.Println(ypub)

    zpub = pubConverter.Zpub(pub)
    fmt.Println(zpub)

}
 
 

