package main
import {
	"Store"
}
func main() {
	store := NewStore(db)
	api := NewAPIServer(":3000",nil)	
}
