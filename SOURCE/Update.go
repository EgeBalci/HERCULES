package main

import "net/http"
import "os/exec"
import "strings"
import "os"
import "io/ioutil"
import "color"




func main() {


	Repo := [3]string{"https://github.com/EgeBalci/HERCULES/raw/master/HERCULES", "https://github.com/EgeBalci/HERCULES/raw/master/HERCULES_x64", "https://github.com/EgeBalci/HERCULES/raw/master/Update"}


	exec.Command("sh", "-c", "zenity --info --text=\"HERCULES Update Started... \"").Run()
	exec.Command("sh", "-c", "sudo rm HERCULES").Run()
	exec.Command("sh", "-c", "sudo rm HERCULES_x86").Run()
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
	Result, Status := CheckValid()
	if Result == false {
		color.Red(Status)
		exec.Command("sh", "-c", "zenity --warning --text=\"HERCULES Update Failed !\"").Run()
	}else{
		exec.Command("sh", "-c", "zenity --info --text=\"HERCULES Update completed !\"").Run()
	}
	


}


func CheckValid()  (bool, string){
  OutESP, _ := exec.Command("sh", "-c", "ls").Output()
  if (!strings.Contains(string(OutESP), "HERCULES")) || (!strings.Contains(string(OutESP), "HERCULES_x64")) || (!strings.Contains(string(OutESP), "Update")) {
    return false, "[!] ERROR : Update failed"
  }
  return true, ""
}
