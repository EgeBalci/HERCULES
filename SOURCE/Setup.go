package main

import "os/exec"
import "strings"
import "runtime"
import "github.com/fatih/color"
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

  color.Green("\n+ -- --=[        HERCULES  FRAMEWORK        ]")
  BoldGreen.Println("+ -- --=[            Ege Balcı              ]")





  Priv := CheckSUDO()

  if Priv == false {
  	BoldRed.Println("[!] ERROR : Setup needs root privileges")
  	os.Exit(1)
  }






  BoldWhite.Println("\n\n[*] STARTING HERCULES SETUP \n")


  BoldYellow.Println("[*] Detecting OS...")
  if runtime.GOOS == "linux" {


    OsVersion, _ := exec.Command("sh", "-c", "uname -a").Output()
    BoldYellow.Println("[*] OS Detected : " + string(OsVersion))

    if strings.Contains(string(OsVersion), "Ubuntu") || strings.Contains(string(OsVersion), "kali") {
    	BoldYellow.Println("[*] Installing golang...")
    	Go := exec.Command("sh", "-c", "sudo apt-get install golang")
    	Go.Stdout = os.Stdout
      Go.Stderr = os.Stderr
    	Go.Stdin = os.Stdin
    	Go.Run()
    	BoldYellow.Println("[*] Installing upx...")
    	UPX := exec.Command("sh", "-c", "sudo apt-get install upx")
    	UPX.Stdout = os.Stdout
      UPX.Stderr = os.Stderr
    	UPX.Stdin = os.Stdin
    	UPX.Run()
    	BoldYellow.Println("[*] Installing openssl...")
    	OSSL := exec.Command("sh", "-c", "sudo apt-get install openssl")
    	OSSL.Stdout = os.Stdout
      OSSL.Stderr = os.Stderr
    	OSSL.Stdin = os.Stdin
    	OSSL.Run()
    	BoldYellow.Println("[*] Installing git...")
    	Git := exec.Command("sh", "-c", "sudo apt-get install git")
    	Git.Stdout = os.Stdout
      Git.Stderr = os.Stderr
    	Git.Stdin = os.Stdin
    	Git.Run()

    	BoldYellow.Println("[*] Cloning EGESPLOIT Library...")
    	exec.Command("sh", "-c", "sudo git clone https://github.com/EgeBalci/EGESPLOIT.git").Run()
    	exec.Command("sh", "-c", "sudo cp EGESPLOIT /usr/lib/go-1.6/src/").Run()
    	exec.Command("sh", "-c", "sudo mv EGESPLOIT /usr/lib/go/src/").Run()

    	BoldYellow.Println("[*] Moving binaries...")
    	exec.Command("sh", "-c", "sudo chmod 777 *").Run()
    	exec.Command("sh", "-c", "sudo mv HERCULES /bin/").Run()
    	exec.Command("sh", "-c", "sudo mv HERCULES_x64 /bin/").Run()
    	exec.Command("sh", "-c", "sudo mv Update /bin/").Run()
    }else if strings.Contains(string(OsVersion), "ARCH") || strings.Contains(string(OsVersion), "MANJARO") {
    	//pacman -S package_name1
    	BoldYellow.Println("[*] Installing golang...")
    	Go := exec.Command("sh", "-c", "pacman -S go")
    	Go.Stdout = os.Stdout
      Go.Stderr = os.Stderr
    	Go.Stdin = os.Stdin
    	Go.Run()
    	BoldYellow.Println("[*] Installing upx...")
    	UPX := exec.Command("sh", "-c", "pacman -S upx")
    	UPX.Stdout = os.Stdout
      UPX.Stderr = os.Stderr
    	UPX.Stdin = os.Stdin
    	UPX.Run()
    	BoldYellow.Println("[*] Installing openssl...")
    	OSSL := exec.Command("sh", "-c", "pacman -S openssl")
    	OSSL.Stdout = os.Stdout
      OSSL.Stderr = os.Stderr
    	OSSL.Stdin = os.Stdin
    	OSSL.Run()
    	BoldYellow.Println("[*] Installing git...")
    	Git := exec.Command("sh", "-c", "pacman -S git")
    	Git.Stdout = os.Stdout
      Git.Stderr = os.Stderr
    	Git.Stdin = os.Stdin
    	Git.Run()

    	BoldYellow.Println("[*] Cloning EGESPLOIT Library...")
    	exec.Command("sh", "-c", "sudo git clone https://github.com/EgeBalci/EGESPLOIT.git").Run()
    	exec.Command("sh", "-c", "sudo cp EGESPLOIT /usr/lib/go-1.6/src/").Run()
    	exec.Command("sh", "-c", "sudo mv EGESPLOIT /usr/lib/go/src/").Run()

    	BoldYellow.Println("[*] Moving binaries...")
    	exec.Command("sh", "-c", "sudo chmod 777 *").Run()
    	exec.Command("sh", "-c", "sudo mv HERCULES /bin/").Run()
    	exec.Command("sh", "-c", "sudo mv HERCULES_x64 /bin/").Run()
    	exec.Command("sh", "-c", "sudo mv Update /bin/").Run()
    }else if strings.Contains(string(OsVersion), "fc22") || strings.Contains(string(OsVersion), "fc23") || strings.Contains(string(OsVersion), "fc24") {
    	//The default package manager for Fedora 22+ is dnf, not yum... Syntax is:
    	//sudo dnf -y install package_name
    	BoldYellow.Println("[*] Installing golang...")
    	Go := exec.Command("sh", "-c", "sudo dnf -y install golang")
    	Go.Stdout = os.Stdout
      Go.Stderr = os.Stderr
    	Go.Stdin = os.Stdin
    	Go.Run()
    	BoldYellow.Println("[*] Installing upx...")
    	UPX := exec.Command("sh", "-c", "sudo dnf -y install upx")
    	UPX.Stdout = os.Stdout
      UPX.Stderr = os.Stderr
    	UPX.Stdin = os.Stdin
    	UPX.Run()
    	BoldYellow.Println("[*] Installing openssl...")
    	OSSL := exec.Command("sh", "-c", "sudo dnf -y install openssl")
    	OSSL.Stdout = os.Stdout
      OSSL.Stderr = os.Stderr
    	OSSL.Stdin = os.Stdin
    	OSSL.Run()
    	BoldYellow.Println("[*] Installing git...")
    	Git := exec.Command("sh", "-c", "sudo dnf -y install git")
    	Git.Stdout = os.Stdout
      Git.Stderr = os.Stderr
    	Git.Stdin = os.Stdin
    	Git.Run()

    	BoldYellow.Println("[*] Cloning EGESPLOIT Library...")
    	exec.Command("sh", "-c", "sudo git clone https://github.com/EgeBalci/EGESPLOIT.git").Run()
    	exec.Command("sh", "-c", "sudo mv EGESPLOIT /usr/lib/golang/src/").Run()

    	BoldYellow.Println("[*] Moving binaries...")
    	exec.Command("sh", "-c", "sudo chmod 777 *").Run()
    	exec.Command("sh", "-c", "sudo mv HERCULES /bin/").Run()
    	exec.Command("sh", "-c", "sudo mv HERCULES_x64 /bin/").Run()
    	exec.Command("sh", "-c", "sudo mv Update /bin/").Run()
    }else{
    	BoldRed.Println("[!] ERROR : HERCULES+ only supports linux distributions")
    }


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
  OutESP2, _ := exec.Command("sh", "-c", "cd /usr/lib/go-1.6/src/ && ls").Output()
  OutESP3, _ := exec.Command("sh", "-c", "cd /usr/lib/golang/src/ && ls").Output()
  if (!strings.Contains(string(OutESP), "EGESPLOIT")) && (!strings.Contains(string(OutESP2), "EGESPLOIT")) && (!strings.Contains(string(OutESP3), "EGESPLOIT")) {
    return false, "[!] ERROR : EGESPLOIT library is not installed"
  }

  OutUPX, _ := exec.Command("sh", "-c", "upx").Output()
  if (!strings.Contains(string(OutUPX), "Copyright")) {
    return false, "[!] ERROR : upx is not installed"
  }

  OutGO, _ := exec.Command("sh", "-c", "go version").Output()
  if (!strings.Contains(string(OutGO), "version")) {
    return false, "[!] ERROR : golang is not installed"
  }

  OutBin, _ := exec.Command("sh", "-c", "cd /bin/ && ls").Output()
  if (!strings.Contains(string(OutBin), "HERCULES")) || (!strings.Contains(string(OutBin), "HERCULES_x64")) || (!strings.Contains(string(OutBin), "Update")) {
    return false, "[!] ERROR : Unable to move HERCULES binaries "
  }

  return true, ""

}

func CheckSUDO() (bool){
	User, _ := exec.Command("sh", "-c", "whoami").Output()
	if strings.Contains(string(User), "root") {
		return true
	}else {
		return false
	}
	
}
