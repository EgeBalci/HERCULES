package main

import "os/exec"
import "strings"
import "runtime"
import "color"
import "os"





func main() {



  Green := color.New(color.FgGreen)
  BoldGreen := Green.Add(color.Bold)
  Yellow := color.New(color.FgYellow)
  BoldYellow := Yellow.Add(color.Bold)
  Red := color.New(color.FgRed)
  BoldRed := Red.Add(color.Bold)
  White := color.New(color.FgWhite)
  BoldWhite := White.Add(color.Bold)

  color.Red(" ██░ ██ ▓█████  ██▀███   ▄████▄   █    ██  ██▓    ▓█████   ██████ ")
  color.Red("▓██░ ██▒▓█   ▀ ▓██ ▒ ██▒▒██▀ ▀█   ██  ▓██▒▓██▒    ▓█   ▀ ▒██    ▒ ")
  color.Red("▒██▀▀██░▒███   ▓██ ░▄█ ▒▒▓█    ▄ ▓██  ▒██░▒██░    ▒███   ░ ▓██▄   ")
  color.Red("░▓█ ░██ ▒▓█  ▄ ▒██▀▀█▄  ▒▓▓▄ ▄██▒▓▓█  ░██░▒██░    ▒▓█  ▄   ▒   ██▒")
  color.Red("░▓█▒░██▓░▒████▒░██▓ ▒██▒▒ ▓███▀ ░▒▒█████▓ ░██████▒░▒████▒▒██████▒▒")
  color.Red(" ▒ ░░▒░▒░░ ▒░ ░░ ▒▓ ░▒▓░░ ░▒ ▒  ░░▒▓▒ ▒ ▒ ░ ▒░▓  ░░░ ▒░ ░▒ ▒▓▒ ▒ ░")
  color.Red(" ▒ ░▒░ ░ ░ ░  ░  ░▒ ░ ▒░  ░  ▒   ░░▒░ ░ ░ ░ ░ ▒  ░ ░ ░  ░░ ░▒  ░ ░")
  color.Red(" ░  ░░ ░   ░     ░░   ░ ░         ░░░ ░ ░   ░ ░      ░   ░  ░  ░  ")
  color.Red(" ░  ░  ░   ░  ░   ░     ░ ░         ░         ░  ░   ░  ░      ░  ")
  color.Red("                        ░                                         ")

  color.Green("\n+ -- --=[      HERCULES FRAMEWORK           ]")
  BoldGreen.Println("+ -- --=[            Ege Balcı              ]")








  BoldWhite.Println("\n\n[*] STARTING HERCULES SETUP \n")


  BoldYellow.Println("[*] Detecting OS...")

  if runtime.GOOS == "linux" {
    OsVersion, _ := exec.Command("sh", "-c", "uname -a").Output()
    BoldYellow.Println("[*] OS Detected : " + string(OsVersion))
    BoldYellow.Println("[*] Installing golang...")
    Go := exec.Command("sh", "-c", "apt-get install golang")
    Go.Stdout = os.Stdout
    Go.Stdin = os.Stdin
    Go.Run()
    BoldYellow.Println("[*] Installing upx...")
    UPX := exec.Command("sh", "-c", "apt-get install upx")
    UPX.Stdout = os.Stdout
    UPX.Stdin = os.Stdin
    UPX.Run()
    BoldYellow.Println("[*] Installing openssl...")
    OSSL := exec.Command("sh", "-c", "apt-get install openssl")
    OSSL.Stdout = os.Stdout
    OSSL.Stdin = os.Stdin
    OSSL.Run()
    BoldYellow.Println("[*] Installing git...")
    Git := exec.Command("sh", "-c", "apt-get install git")
    Git.Stdout = os.Stdout
    Git.Stdin = os.Stdin
    Git.Run()

    BoldYellow.Println("[*] Cloning EGESPLOIT Library...")
    exec.Command("sh", "-c", "git clone https://github.com/EgeBalci/EGESPLOIT.git").Run()
    exec.Command("sh", "-c", "cp EGESPLOIT /usr/lib/go-1.6/src/").Run()
    exec.Command("sh", "-c", "mv EGESPLOIT /usr/lib/go/src/").Run()

    BoldYellow.Println("[*] Cloning color Library...")
    exec.Command("sh", "-c", "git clone https://github.com/fatih/color.git").Run()
    exec.Command("sh", "-c", "cp color /usr/lib/go-1.6/src/").Run()
    exec.Command("sh", "-c", "mv color /usr/lib/go/src/").Run()// /usr/lib/go-1.6/src/


    Stat, Err := CheckValid()

    if Stat == false {
      BoldYellow.Println("\n")
      BoldRed.Println(Err)
    }else{
      BoldGreen.Println("\n\n[+] Setup completed successfully")
    }


  }else if runtime.GOOS != "linux" {
    BoldRed.Println("[!] ERROR : HERCULES+ only supports linux distributions")
  }

}


func CheckValid()  (bool, string){
  OutESP, _ := exec.Command("sh", "-c", "cd /usr/lib/go/src/ && ls").Output()
  if (!strings.Contains(string(OutESP), "EGESPLOIT")) {
    return false, "[!] ERROR : EGESPLOIT library is not installed"
  }

  OutESP2, _ := exec.Command("sh", "-c", "cd /usr/lib/go-1.6/src/ && ls").Output()
  if (!strings.Contains(string(OutESP2), "EGESPLOIT")) {
    return false, "[!] ERROR : EGESPLOIT library is not installed"
  }

  OutCL, _ := exec.Command("sh", "-c", "cd /usr/lib/go/src/ && ls").Output()
  if (!strings.Contains(string(OutCL), "color")) {
    return false, "[!] ERROR : color library is not installed"
  }

  OutCL2, _ := exec.Command("sh", "-c", "cd /usr/lib/go-1.6/src/ && ls").Output()
  if (!strings.Contains(string(OutCL2), "color")) {
    return false, "[!] ERROR : color library is not installed"
  }

  OutUPX, _ := exec.Command("sh", "-c", "upx").Output()
  if (!strings.Contains(string(OutUPX), "Copyright")) {
    return false, "[!] ERROR : upx is not installed"
  }

  OutGO, _ := exec.Command("sh", "-c", "go version").Output()
  if (!strings.Contains(string(OutGO), "version")) {
    return false, "[!] ERROR : golang is not installed"
  }

  return true, ""

}
