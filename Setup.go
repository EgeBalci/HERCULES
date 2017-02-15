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

  BoldWhite.Println("\n\n[*] STARTING HERCULES SETUP \n")


  BoldYellow.Println("[*] Detecting OS...")

  if runtime.GOOS == "linux" {


    OsVersion, _ := exec.Command("sh", "-c", "uname -a").Output()
    BoldYellow.Println("[*] OS Detected : " + string(OsVersion))
    BoldYellow.Println("[*] Setting HERCULES path...")


    Path, _ := exec.Command("sh", "-c", "pwd").Output()
	BoldYellow.Println("[*] HERCULES_PATH="+string(Path))
	_Path := strings.Trim(string(Path), "\n")
    var HERCULES_PATH string = string("echo 'export HERCULES_PATH="+_Path+"' >> ~/.bashrc")
    exec.Command("sh", "-c", HERCULES_PATH).Run()
    exec.Command("sh", "-c", string("export HERCULES_PATH="+string(Path))).Run()
    if strings.Contains(string(OsVersion), "Ubuntu") || strings.Contains(string(OsVersion), "kali") {
    	BoldYellow.Println("[*] Installing golang...")
    	 if Priv == false {
  			BoldRed.Println("[!] ERROR : Setup needs root privileges")
  		}
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
    	BoldYellow.Println("[*] Installing git...")
    	Git := exec.Command("sh", "-c", "sudo apt-get install git")
    	Git.Stdout = os.Stdout
     	Git.Stderr = os.Stderr
    	Git.Stdin = os.Stdin
    	Git.Run()

    	BoldYellow.Println("[*] Cloning EGESPLOIT Library...")
      	exec.Command("sh", "-c", "cd src && git clone https://github.com/EgeBalci/EGESPLOIT.git").Run()
      	exec.Command("sh", "-c", "export GOPATH=$HERCULES_PATH").Run()
      	BoldYellow.Println("[*] Cloning color Library...")
      	exec.Command("sh", "-c", "go get github.com/fatih/color").Run()
    	
      	exec.Command("sh", "-c", "cd SOURCE && go build HERCULES.go").Run()

    	BoldYellow.Println("[*] Createing shoutcut...")
    	exec.Command("sh", "-c", "sudo cp HERCULES /bin/").Run()
    	exec.Command("sh", "-c", "sudo chmod 777 /bin/HERCULES").Run()

    }else if strings.Contains(string(OsVersion), "ARCH") || strings.Contains(string(OsVersion), "MANJARO") {
    	//pacman -S package_name1
    	BoldYellow.Println("[*] Installing golang...")
    	BoldYellow.Println("[*] Installing golang...")
    	 if Priv == false {
  			BoldRed.Println("[!] ERROR : Setup needs root privileges")
  		}
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
    	BoldYellow.Println("[*] Installing git...")
    	Git := exec.Command("sh", "-c", "pacman -S git")
    	Git.Stdout = os.Stdout
      Git.Stderr = os.Stderr
    	Git.Stdin = os.Stdin
    	Git.Run()

    	BoldYellow.Println("[*] Cloning EGESPLOIT Library...")
    	exec.Command("sh", "-c", "cd SOURCE && git clone https://github.com/EgeBalci/EGESPLOIT.git").Run()
    	exec.Command("sh", "-c", "export GOPATH=$HERCULES_PATH").Run()
    	BoldYellow.Println("[*] Cloning color Library...")
    	exec.Command("sh", "-c", "go get github.com/fatih/color").Run()

		exec.Command("sh", "-c", "cd SOURCE && go build HERCULES.go").Run()

    	BoldYellow.Println("[*] Createing shoutcut...")
    	exec.Command("sh", "-c", "sudo cp HERCULES /bin/").Run()
    	exec.Command("sh", "-c", "sudo chmod 777 /bin/HERCULES").Run()

    }else{
    	BoldRed.Println("[!] ERROR : HERCULES does not support this OS")
    }


    Stat, Err := CheckValid()

    if Stat == false {
      BoldYellow.Println("\n")
      BoldRed.Println(Err)
    }else{
      BoldGreen.Println("\n\n[+] Setup completed successfully")
      exec.Command("sh", "-c", "gnome-terminal").Run()
      exec.Command("sh", "-c", "exit").Run()
    }


  	}else if runtime.GOOS != "linux" {
    	BoldRed.Println("[!] ERROR : HERCULES only supports linux distributions")
  	}

}


func CheckValid()  (bool, string){

  OutUPX, _ := exec.Command("sh", "-c", "upx").Output()
  if (!strings.Contains(string(OutUPX), "Copyright")) {
    return false, "[!] ERROR : upx is not installed"
  }

  OutGO, _ := exec.Command("sh", "-c", "go version").Output()
  if (!strings.Contains(string(OutGO), "version")) {
    return false, "[!] ERROR : golang is not installed"
  }

  OutBin, _ := exec.Command("sh", "-c", "cd /bin/ && ls").Output()
  if (!strings.Contains(string(OutBin), "HERCULES"))  {
    return false, "[!] ERROR : Unable to create shoutcut "
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
