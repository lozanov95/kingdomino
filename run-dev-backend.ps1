try {
    Set-Location ./backend
    go run main.go -port 8080
}
finally {
    Set-Location ../
}
