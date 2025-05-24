//Mostly for codee reuse
//package increase compile time 
//better to givw a package name like this github.com/username/package name
package main //main papckage
//go mod init name -> in terminal so make it a create a go mod file
//create file  based on package a folder with files a pckage
// like ./auth/credentials.go give name package auth
/*
have function for username and password and logic
and need to scope it -> this function can use in this folder only but if you want to use outside so make it capital start naming
//even naming rule for struct fields also
*/
import "github/......name of auth"
func main(){
  auth.LoginCredential(name,password)
}


//Third packages also there
//package go as for npm like go packages
//install ->go get module_name
color.Red("Hello")

//go.sum is same like log file for packages versions and all
//go mod tidy//manage packages and dependency same like npm i
