
Several easy steps:

 - parse config file
 - prepare/format/divide disk
   - nvme
   - ssd/hdd: GPT, BIOS
 - configure/check network
 - download and install stage3
 - install base system + chroot
 - eselects
 - configure:
   - use flags
   - make.conf
   - time zone
   - locale
 - rebuild @system and @world
 - configure and build kernel
 - fstab
 - install system tools
 - create users and groups
 - install and configure GRUB2

REBOOT

- install software
- update the system