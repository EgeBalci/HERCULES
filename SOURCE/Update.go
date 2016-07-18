package main

import "net/http"
import "os/exec"
import "strings"
import "os"
import "io/ioutil"
import "color"




func main() {


	Repo := []string{"https://github.com/EgeBalci/HERCULES/blob/master/SOURCE/HERCULES.go", "https://github.com/EgeBalci/HERCULES/raw/master/HERCULES_x64", "https://github.com/EgeBalci/HERCULES/raw/master/HERCULES_x86", "https://github.com/EgeBalci/HERCULES/raw/master/README.md"}


	exec.Command("sh", "-c", "zenity --info --text=\"HERCULES Update Started... \"").Run()
	exec.Command("sh", "-c", "rm  -r SOURCE").Run()
	exec.Command("sh", "-c", "rm  -r STATISTICS").Run()
	exec.Command("sh", "-c", "rm HERCULES_x64").Run()
	exec.Command("sh", "-c", "rm HERCULES_x86").Run()
	color.Blue("[*] Updating HERCULES...\n\n")
	for i := 0; i < len(Repo); i++ {
		response, _ := http.Get(Repo[i])
		defer response.Body.Close();
    body, _ := ioutil.ReadAll(response.Body);

    Name := strings.Split(Repo[i], "/")
    color.Green("#	"+string(Name[(len(Name)-1)])+"		[OK]")
    File, _ := os.Create(string(Name[(len(Name)-1)]))

    File.WriteString(string(body))
	}
	exec.Command("sh", "-c", "zenity --info --text=\"HERCULES Update completed !\"").Run()


}
