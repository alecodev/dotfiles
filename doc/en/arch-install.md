# Arch Linux installation

***For more information on the installation process visit [Arch's guide](https://wiki.archlinux.org/title/Installation_guide)***

>**If you are installing on a virtual machine**
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

Select GPT and create the following partitions (the size depends on the use you want to give it), **remember to write before exiting**
```
512M      EFI System         (This will be the size of the system boot partition)
16G       Linux Swap         (This will be the size of the SWAP memory, it is recommended to double the size of the RAM memory)
40G       Linux filesystem   (This will be the size assigned to /)
63.5G     Linux filesystem   (This will be the size assigned to /home)
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
pacstrap /mnt base base-devel linux linux-firmware neovim dhcpcd
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

### Set Localization
Edit the file `/etc/locale.gen` with [text editor][1] and uncomment `en_US.UTF-8 UTF-8`

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

Add the domains to the file `/etc/hosts` with [text editor][1], replace the name of the computer with the one you established in the previous step
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

Activate the sudo group by executing the following command and uncomment `%wheel ALL=(ALL) ALL`
```bash
EDITOR=nvim
visudo
```

>**In case of running in a virtual machine like VirtualBox run the following command**
>```bash
>pacman -S virtualbox-guest-utils
>systemctl enable vboxservice
>usermod -aG vboxsf alejo
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

[1]:../../README.md#text-editor