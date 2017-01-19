# HERCULES [![License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://raw.githubusercontent.com/EgeBalci/HERCULES/master/LICENSE)   [![Support](https://img.shields.io/badge/Support-Mail-red.svg)](https://github.com/EgeBalci/HERCULES/wiki) [![Golang](https://img.shields.io/badge/Go-1.6-blue.svg)](https://golang.org)

HERCULES is a customizable payload generator that can bypass antivirus software.


		VERSION 3.0.5
		
	
![](http://i.imgur.com/SMU8WU4.png)
	
	
#INSTALLATTION

SUPPORTED PLATFORMS:

<table>
    <tr>
        <th>Operative system</th>
        <th> Version </th>
    </tr>
    <tr>
        <td>Ubuntu</td>
        <td> 16.04  / 15.10 </td>
    </tr>
    <tr>
        <td>Kali linux</td>
        <td> Rolling / Sana</td>
    </tr>
    <tr>
        <td>Manjaro</td>
        <td>* </td>
    </tr>
    <tr>
        <td>Arch Linux</td>
        <td>* </td>
    </tr>
    <tr>
        <td>Black Arch</td>
        <td>* </td>
    </tr>
    <tr>
        <td>Parrot OS</td>
        <td>3.1 </td>
    </tr>
</table>


		go run Setup.go


#USAGE

		HERCULES


#SPECIAL FUNCTIONS


		Persistence : Persistence function adds the running binary to windows start-up registry (CurrentVersion/Run) for continious access.
		
		Migration : This function triggers a loop that tries to migrate to a remote process until it is successfully migrated. 

#WHAT IS UPX ?

		UPX (Ultimate Packer for Executables) is a free and open source executable packer supporting a number of file formats from different operating systems. UPX simply takes the binary file and compresses it, packed binary unpack(decompress) itself at runtime to memory.
		
#WHAT IS "AV EVASION SCORE" ?

		AV Evasion Score is a scale(1/10) for determining the effectiveness of the payloads anti virus bypassing capabilities, 1 represents low possibility to pass AV softwares.
		
		Using special functions and packing the payloads with upx decreases the AV Evasion Score.

![](http://i.imgur.com/8L1wmjo.png)

![](https://blockchain.info/qr?data=14n1bJmCRNpLKAJVJ55jStH5VueFk6G1n4&size=200)

Bitcoin: 14n1bJmCRNpLKAJVJ55jStH5VueFk6G1n4&size
		
		
#COMING SOON...

- Binary infector
- Bypass AV functon
- AES payload encryption
- OSX support

		
