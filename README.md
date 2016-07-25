# HERCULES [![License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://raw.githubusercontent.com/EgeBalci/HERCULES/master/LICENSE)  [![Donate](https://img.shields.io/badge/Donate-Patreon-green.svg)](http://patreon.com/user?u=3556027)

HERCULES is a customizable payload generator that can bypass all antivirus software.


		VERSION 3.0.0
		
	
#INSTALLATTION

SUPPORTED PLATFORMS:

- Kali Linux
- Ubuntu
- Arch Linux
- BlackArch Linux
- Manjero
- Parrot OS

		sudo chmod 777 Setup && sudo ./Setup


#USAGE

		sudo ./HERCULES


#SPECIAL FUNCTIONS


		Persistence : Persistence function adds the running binary to windows start-up registry (CurrentVersion/Run) for continious access.
		
		Migration : This function triggers a loop that tries to migrate to a remote process until it is successfully migrated. 

#WHAT IS UPX ?

		UPX (Ultimate Packer for Executables) is a free and open source executable packer supporting a number of file formats from different operating systems. UPX simply takes the binary file and compresses it, packed binary unpack(decompress) itself at runtime to memory.
		
#WHAT IS "AV EVASION SCORE"

		AV Evasion Score is a scale(1/10) for determining the effectiveness of the paylaods anti virus bypassing capabilities, 1 represends low possibility to pass AV softwares.
		
		Using special functions and packing the payloads with upx decreases the AV Evasion Score.
