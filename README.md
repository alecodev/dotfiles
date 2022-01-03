# Dotfiles in Arch Linux

##	Arch Linux installation
For more information on the installation process visit [Arch's guide](https://wiki.archlinux.org/title/Installation_guide)

>If you are installing on a virtual machine
>- Enable EFI
>- Disable 3D acceleration


Sets the console keyboard layout (in my case 'es'). You can validate the available keyboard layouts with the following command: `ls /usr/share/kbd/keymaps/**/*.map.gz`
```bash
loadkeys es
```

Verify the boot mode EFI
```bash
ls /sys/firmware/efi/efivars
```

Check connection
```bash
ip link
ping 8.8.8.8
```

Update the system clock
```bash
timedatectl set-ntp true
```

Identifies the device block usually is /dev/sda
```bash
lsblk
```

Set the partitions
```bash
cfdisk /dev/sda
```

Select GPT and create the following partitions (the size depends on the use you want to give it), remember to write before exiting
```
512M		EFI System			(This will be the size of the system boot partition)
16G			Linux Swap			(This will be the size of the SWAP memory, it is recommended to double the size of the RAM memory)
40G			Linux filesystem	(This will be the size assigned to /)
63.5G		Linux filesystem	(This will be the size assigned to /home)
write
```

Check the partitions
```bash
lsblk
```

Set the format of the partitions
```bash
mkfs.fat -F32 /dev/sda1
mkswap /dev/sda2
mkfs.ext4 /dev/sda3
mkfs.ext4 /dev/sda4
```

Mount the partitions
```bash
swapon /dev/sda2
mount /dev/sda3 /mnt
mkdir /mnt/{efi,home}
mount /dev/sda1 /mnt/efi
mount /dev/sda4 /mnt/home
```

Check the partitions
```bash
lsblk
```

Install the basic packages
```bash
pacstrap /mnt base linux linux-firmware nano vim dhcpcd
```

Generate the Fstab file
```bash
genfstab -U /mnt >> /mnt/etc/fstab
```

Change root into the new system
```bash
arch-chroot /mnt
```

Set the time zone (in my case 'America/Bogota'), you can see the available time zones with the following command: `timedatectl list-timezones`
```bash
ln -sf /usr/share/zoneinfo/America/Bogota /etc/localtime
hwclock --systohc
```

Set Localization, edit the file `/etc/locale.gen` and uncomment `en_US.UTF-8 UTF-8` with editor text (vim, nano, ...)

Generate the regional settings and set the default keyboard layout by running
```bash
locale-gen
echo "LANG=en_US.UTF-8" >> /etc/locale.conf
echo "KEYMAP=es" >> /etc/vconsole.conf
```

Create the hostname file (in my case the hostname will be 'Arch')
```bash
echo "Arch" >> /etc/hostname
```

Add the domains to the file `/etc/hosts` with the editor text (vim, nano, ...), replace the name of the computer with the one you established in the previous step
```bash
127.0.0.1   localhost
::1         localhost
127.0.1.1   Arch.localdomain  Arch
```

Activate the DHCPCD service
```bash
systemctl enable dhcpcd.service
```

Set the password for the root user
```bash
passwd
```

Create a new user (in my case 'alejo')
```bash
useradd -m alejo
passwd alejo
usermod -aG wheel,audio,video,optical,storage alejo
```

Activate the sudo group by executing the following command `EDITOR=vim visudo` and uncomment `%whell ALL=(ALL) ALL`

>In case of running in a virtual machine like VirtualBox run the following command
>```bash
>pacman -S virtualbox-guest-utils
>systemctl enable vboxservice
>```

Set the bootloader
```bash
pacman -S grub sudo efibootmgr os-prober intel-ucode
grub-install --target=x86_64-efi --efi-directory=/efi --bootloader-id=GRUB
echo "GRUB_DISABLE_OS_PROBER=false" >> /etc/default/grub
grub-mkconfig -o /boot/grub/grub.cfg
```

Finish the process with the following command
```bash
exit
umount -R /mnt
reboot
```

---
##	Install Window Manager

Login with the root user and run the following commands
```bash
pacman -Syu
pacman -S gcc make git xorg-server bspwm sxhkd alacritty rofi lightdm lightdm-gtk-greeter numlockx zsh
```

Edit the file `/etc/lightdm/lightdm.conf` with the editor text (vim, nano, ...) and set the following lines
```text
greeter-session=lightdm-gtk-greeter
greeter-setup-script=/usr/bin/numlockx on
display-setup-script=/usr/bin/setxkbmap es
```

Activate the lightdm service
```bash
systemctl enable lightdm
```

Create o edit file ```~/.xprofile``` with editor text (vim, nano, ...) and set the following lines
```text
VBoxClient-all &
sxhkd &
exec bspwm
```

Create the directories and configuration files of bspwm
```bash
mkdir -p ~/.config/{bspwm,sxhkd}
cd /usr/share/doc/bspwm/examples
cp bspwmrc ~/.config/bspwm
cp sxhkdrc ~/.config/sxhkd
cd ~
chmod +x ~/.config/bspwm/bspwmrc
```

Edit file ```~/.config/sxhkd/sxhkdrc``` with editor text (vim, nano, ...) and modify the following lines
```diff
# terminal emulator
super + Return
-	urxvt
+	alacritty

# program launcher
super + @space
-	dmenu_run
+	rofi -show run
```

Create the bspwm configuration directories in the other user and copy the files (in my case 'alejo')
```bash
su alejo
mkdir -p ~/.config/{bspwm,sxhkd}
sudo cp /root/.config/bspwmrc ~/.config/bspwm && sudo chown alejo:alejo ~/.config/bspwm/bspwmrc
sudo cp /root/.config/sxhkdrc ~/.config/sxhkd && sudo chown alejo:alejo ~/.config/sxhkd/sxhkdrc
sudo cp /root/.xprofile ~/.xprofile && sudo chown alejo:alejo ~/.xprofile
```

Reboot and log in with the other user
```bash
reboot
```

Ready now you can log in with the other user and use bspwm by pressing `Super + Enter`
