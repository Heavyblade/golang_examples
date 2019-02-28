package main

import "os/exec"

func main() {
    app := "ls"

    arg0 := "-l"
    cmd := exec.Command(app, arg0)
    stdout, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return
    }

    print(string(stdout))
}

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html