// From http://stackoverflow.com/questions/6182369/exec-a-shell-command-in-go

package main

import "os/exec"

func main() {
    app := "echo"
    //app := "buah"

    arg0 := "-e"
    arg1 := "Hello world"
    arg2 := "\n\tfrom"
    arg3 := "golang"

    cmd := exec.Command(app, arg0, arg1, arg2, arg3)
    out, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return
    }

    print(string(out))
}
