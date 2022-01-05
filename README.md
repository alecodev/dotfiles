# Dotfiles in Arch Linux

***Languages***
- [🇪🇸 - Español](./README.es.md)
- **🇪🇳 - English**

## Arch Linux installation
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
512M      EFI System         (This will be the size of the system boot partition)
16G       Linux Swap         (This will be the size of the SWAP memory, it is recommended to double the size of the RAM memory)
40G       Linux filesystem   (This will be the size assigned to /)
63.5G     Linux filesystem   (This will be the size assigned to /home)
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
pacstrap /mnt base base-devel linux linux-firmware nano vim dhcpcd
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
## Install Window Manager

Login with the root user and run the following commands
```bash
pacman -Syu
pacman -S gcc make git xorg-server bspwm sxhkd alacritty rofi lightdm lightdm-gtk-greeter numlockx zsh neovim
```

Configure zsh and history
```bash
touch ~/{.zshrc,.zsh_history}
sudo su
touch ~/{.zshrc,.zsh_history}
```

Create or edit the file `~/.zshrc` with editor text (vim, nano, ...) and add the following lines
```text
# Lines configured by zsh-newuser-install
HISTFILE=~/.zsh_history
HISTSIZE=1000
SAVEHIST=1000
bindkey -e

# End of lines configured by zsh-newuser-install
# The following lines were added by compinstall
zstyle :compinstall filename '/home/alejo/.zshrc'

autoload -Uz compinit
compinit
# End of lines added by compinstall

# Alias list files
alias ls='ls -lh --color=auto'

# Alias list all files
alias la='ls -Alh'

# Additional path for VScode
export PATH=$PATH:/opt/VSCode-linux-x64/bin
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

Change user
```bash
su alejo
```

Create or edit the file `~/.xprofile` with editor text (vim, nano, ...) and set the following lines
```text
VBoxClient-all &
sxhkd &
exec bspwm
```

Create or edit the file `~/.bashrc` with editor text (vim, nano, ...) and add the following lines
```text
# Alias list all files
alias la='ls -Alh'
```

Create the directories and configuration files of bspwm, sxhkd, polybar and alacritty
```bash
mkdir -p ~/.config/{bspwm,sxhkd,polybar,alacritty}
cd /usr/share/doc/bspwm/examples
cp bspwmrc ~/.config/bspwm
cp sxhkdrc ~/.config/sxhkd
cd ~
chmod +x ~/.config/bspwm/bspwmrc
mkdir ~/.config/bspwm/scripts
touch ~/.config/bspwm/scripts/bspwm_resize
chmod +x ~/.config/bspwm/scripts/bspwm_resize
```

Edit file `~/.config/bspwm/scripts/bspwm_resize` with editor text (vim, nano, ...) and add the following lines
```bash
#!/usr/bin/env bash

if bspc query -N -n focused.floating > /dev/null; then
	step=20
else
	step=100
fi

case "$1" in
	west) dir=right; falldir=left; x="-$step"; y=0;;
	east) dir=right; falldir=left; x="$step"; y=0;;
	north) dir=top; falldir=bottom; x=0; y="-$step";;
	south) dir=top; falldir=bottom; x=0; y="$step";;
esac

bspc node -z "$dir" "$x" "$y" || bspc node -z "$falldir" "$x" "$y"
```

Edit file `~/.config/sxhkd/sxhkdrc` with editor text (vim, nano, ...) and modify the following lines
```diff
# terminal emulator
super + Return
-	urxvt
+	alacritty

# program launcher
-super + @space
-	dmenu_run
+super + d
+	rofi -show run

...

# focus the node in the given direction
-super + {_,shift + }{h,j,k,l}
+super + {_,shift + }{Left,Down,Up,Right}

...

# preselect the direction
-super + ctrl + {h,j,k,l}
+super + ctrl + alt + {Left,Down,Up,Right}

...

# cancel the preselection for the focused desktop
-super + ctrl + shift + space
+super + ctrl + alt + space

...

# expand a window by moving one of its side outward
-super + alt + {h,j,k,l}
-	bspc node -z {left -20 0,bottom 0 20,top 0 -20,right 20 0}
+#super + alt + {h,j,k,l}
+#	bspc node -z {left -20 0,bottom 0 20,top 0 -20,right 20 0}

# contract a window by moving one of its side inward
-super + alt + shift + {h,j,k,l}
-	bspc node -z {right -20 0,top 0 20,bottom 0 -20,left 20 0}
+#super + alt + shift + {h,j,k,l}
+#	bspc node -z {right -20 0,top 0 20,bottom 0 -20,left 20 0}

# move a floating window
-super + {Left,Down,Up,Right}
+super + ctrl + {Left,Down,Up,Right}

+# Custom Move - Resize
+alt + super + {Left,Down,Up,Right}
+	/home/alejo/.config/bspwm/scripts/bspwm_resize {west,south,north,east}
```

Create the bspwm configuration directories in the other user and copy the files (in my case 'alejo')
```bash
sudo su
ln -s /home/alejo/.config/{bspwm,sxhkd,polybar,alacritty} ~/.config/
rm ~/.xprofile
ln -s /home/alejo/.xprofile ~/
rm ~/.bashrc
ln -s /home/alejo/.bashrc ~/
```

Reboot and log in with the other user
```bash
reboot
```

Ready now you can log in with the other user and use bspwm by pressing `Super + Enter`

---
## Setting up the work environment

### Install Firefox
```bash
sudo su
pacman -Syu
pacman -S wget curl libstdc++5 dbus-glib unzip firejail
cd /
chown alejo:alejo opt/
cd !$
su alejo
```

Get the download URL of the Firefox tar.* file `curl "https://download.mozilla.org/?product=firefox-latest-ssl&os=linux64&lang=en-US"` and change it in the following command
```bash
wget "https://download-installer.cdn.mozilla.net/pub/firefox/releases/95.0.2/linux-x86_64/en-US/firefox-95.0.2.tar.bz2"
tar -xf firefox-*
rm firefox-*
mkdir -p ~/Downloads/Firefox
```

Change the download directory in Firefox settings to `~/Downloads/Firefox`

Edit file `~/.config/sxhkd/sxhkdrc` with editor text (vim, nano, ...) and add the following lines
```text
# Open Firefox
super + shift + f
        firejail /opt/firefox/firefox
```

Press `Super + Alt + r` and  `Super + esc`, open Firefox with `Super + Shift + f`

### Install Fonts
```bash
sudo su
```

- #### ***Hack Nerd Font***
	---
	Download the font ***Hack Nerd Font*** from the page `https://www.nerdfonts.com/font-downloads` and install with the following commands
	```bash
	mkdir -p /usr/local/share/fonts/nerd-fonts/Hack
	cd !$
	mv /home/alejo/Downloads/Firefox/Hack.zip .
	unzip Hack.zip
	rm Hack.zip
	```

- #### ***Meslo Nerd Font***
	---
	Download the font ***Meslo Nerd Font*** from the page `https://www.nerdfonts.com/font-downloads` and install with the following commands
	```bash
	mkdir -p /usr/local/share/fonts/nerd-fonts/Meslo
	cd !$
	mv /home/alejo/Downloads/Firefox/Meslo.zip .
	unzip Meslo.zip
	rm Meslo.zip
	```

Reload fonts
```bash
fc-cache -vf
```

Copy the configuration file from alacritty
```bash
su alejo
cp /usr/share/doc/alacritty/example/alacritty.yml ~/.config/alacritty/
```

Edit file `~/.config/alacritty/alacritty.yml` with editor text (vim, nano, ...) and modify the following lines
```diff
# Font configuration
-#font:
+font:
  # Normal (roman) font face
-  #normal:
+  normal:
    # Font family
    #
    # Default:
    #   - (macOS) Menlo
    #   - (Linux/BSD) monospace
    #   - (Windows) Consolas
-    #family: monospace
+    family: "Hack Nerd Font Mono"

    # The `style` can be specified to pick a specific face.
-    #style: Regular
+    style: Regular

  # Bold font face
-  #bold:
+  bold:
    # Font family
    #
    # If the bold family is not specified, it will fall back to the
    # value specified for the normal font.
-    #family: monospace
+    family: "Hack Nerd Font Mono"

    # The `style` can be specified to pick a specific face.
-    #style: Bold
+    style: Bold

  # Italic font face
-  #italic:
+  italic:
    # Font family
    #
    # If the italic family is not specified, it will fall back to the
    # value specified for the normal font.
-    #family: monospace
+    family: "Hack Nerd Font Mono"

    # The `style` can be specified to pick a specific face.
-    #style: Italic
+    style: Italic

  # Bold italic font face
-  #bold_italic:
+  bold_italic:
    # Font family
    #
    # If the bold italic family is not specified, it will fall back to the
    # value specified for the normal font.
-    #family: monospace
+    family: "Hack Nerd Font Mono"

    # The `style` can be specified to pick a specific face.
-    #style: Bold Italic
+    style: Bold Italic

  # Point size
-  #size: 11.0
+  size: 12
```

### Install Visual Studio Code
Download file tar.gz from the page `https://code.visualstudio.com/Download` and install with the following commands
```bash
sudo pacman -S electron
cd /opt
mv ~/Downloads/Firefox/code-*.tar.gz .
tar -xf code-*.tar.gz
rm code-*.tar.gz
cd ~
```

Edit file `~/.bashrc` with editor text (vim, nano, ...) and add the following lines
```text
# Additional path for VScode
export PATH=$PATH:/opt/VSCode-linux-x64/bin
```

### Set Wallpaper
```bash
sudo pacman -S feh
mkdir -p ~/Desktop/alejo/Images
```

Download wallpaper in Images and edit `~/.config/bspwm/bspwmrc` and add next line
```text
feh --bg-fill /home/alejo/Desktop/alejo/Images/wallpaper.jpg
```

### Install Neofetch
```bash
sudo pacman -S neofetch
```

### Install Polybar
```bash
sudo pacman -S cmake pkg-config libuv cairo libxcb python3 xcb-proto xcb-util-image xcb-util-wm python-sphinx python-packaging xcb-util-cursor xcb-util-xrm alsa-lib libpulse i3-wm jsoncpp libmpdclient libnl curl
cd ~/Downloads/Firefox/

git clone --recursive https://github.com/polybar/polybar
cd polybar

mkdir build
cd build
cmake ..
make -j$(nproc)
sudo make install
```

### Install Picom
```bash
sudo pacman -S meson libx11 libxext libconfig libdbus libev pixman uthash xcb-util-image xcb-util-renderutil libgl pcre asciidoc mesa ninja dbus xorg-xprop xorg-xwininfo
cd ~/Downloads/Firefox/

git clone https://github.com/yshui/picom.git
cd picom

git submodule update --init --recursive
meson --buildtype=release . build
ninja -C build
sudo ninja -C build install
```

### Install Powerlevel10k
```bash
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ~/powerlevel10k
echo 'source ~/powerlevel10k/powerlevel10k.zsh-theme' >>~/.zshrc
zsh

sudo su
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ~/powerlevel10k
echo 'source ~/powerlevel10k/powerlevel10k.zsh-theme' >>~/.zshrc
zsh

ln -sf /home/alejo/.zshrc ~/

usermod --shell /usr/bin/zsh alejo
usermod --shell /usr/bin/zsh root
```