# Instalación de Arch Linux

***Para obtener más información sobre el proceso de instalación, visite la [guía de Arch](https://wiki.archlinux.org/title/Installation_guide)***

>**Si está instalando en una máquina virtual**
>- Habilite EFI
>- Deshabilite la aceleración 3D

Establezca la distribución del teclado de la consola (en mi caso, 'es'). Puede validar las distribuciones de teclado disponibles con el siguiente comando : `ls /usr/share/kbd/keymaps/**/*.map.gz`
```bash
loadkeys es
```

Verifique el modo de arranque EFI este activo
```bash
ls /sys/firmware/efi/efivars
```

Verifica la conexión a internet
```bash
ip link
ping 8.8.8.8
```

Actualiza el reloj del sistema.
```bash
timedatectl set-ntp true
```

Identifica el bloque del dispositivo que normalmente es /dev/sda
```bash
lsblk
```

Establecer las particiones
```bash
cfdisk /dev/sda
```

Selecciona GPT y crea las siguientes particiones (el tamaño depende del uso que le quieras dar), **recuerda escribir las particiones antes de salir**
```
512M      EFI System         (Este será el tamaño de la partición de arranque del sistema)
16G       Linux Swap         (Este será el tamaño de la memoria SWAP, se recomienda duplicar el tamaño de la memoria RAM)
40G       Linux filesystem   (Este será el tamaño asignado a /)
63.5G     Linux filesystem   (Este será el tamaño asignado a /home)
```

Revisa las particiones
```bash
lsblk
```

Establezca el formato de las particiones
```bash
mkfs.fat -F32 /dev/sda1
mkswap /dev/sda2
mkfs.ext4 /dev/sda3
mkfs.ext4 /dev/sda4
```

Monte las particiones
```bash
swapon /dev/sda2
mount /dev/sda3 /mnt
mkdir /mnt/{efi,home}
mount /dev/sda1 /mnt/efi
mount /dev/sda4 /mnt/home
```

Revisa las particiones
```bash
lsblk
```

Instale los paquetes básicos
```bash
pacstrap /mnt base base-devel linux linux-firmware neovim dhcpcd
```

Genere el archivo Fstab
```bash
genfstab -U /mnt >> /mnt/etc/fstab
```

Cambie la raíz del nuevo sistema
```bash
arch-chroot /mnt
```

Configure la zona horaria (en mi caso 'America/Bogota'), puede ver las zonas horarias disponibles con el siguiente comando: `timedatectl list-timezones`
```bash
ln -sf /usr/share/zoneinfo/America/Bogota /etc/localtime
hwclock --systohc
```

### Establezca la localización
Edite el archivo `/etc/locale.gen` con el [editor de texto][1] y descomente `en_US.UTF-8 UTF-8`

Genere la configuración regional y configure la distribución del teclado predeterminado ejecutando
```bash
locale-gen
echo "LANG=en_US.UTF-8" >> /etc/locale.conf
echo "KEYMAP=es" >> /etc/vconsole.conf
```

Cree el archivo de nombre de host (en mi caso, el nombre de host será 'Arch')
```bash
echo "Arch" >> /etc/hostname
```

Agrega los dominios al archivo `/etc/hosts` con el [editor de texto][1], reemplaza el nombre de la computadora con el que estableciste en el paso anterior
```bash
127.0.0.1   localhost
::1         localhost
127.0.1.1   Arch.localdomain  Arch
```

Activa el servicio DHCPCD
```bash
systemctl enable dhcpcd.service
```

Establezca la contraseña para el usuario root
```bash
passwd
```

Cree un nuevo usuario (en mi caso 'alejo')
```bash
useradd -m alejo
passwd alejo
usermod -aG wheel,audio,video,optical,storage alejo
```

Active el grupo sudo ejecutando el siguiente comando y descomente `%wheel ALL=(ALL) ALL`
```bash
EDITOR=nvim
visudo
```

>**En caso de ejecutarse en una máquina virtual como VirtualBox ejecute el siguiente comando**
>```bash
>pacman -S virtualbox-guest-utils
>systemctl enable vboxservice
>usermod -aG vboxsf alejo
>```

Configurar el gestor de arranque
```bash
pacman -S grub sudo efibootmgr os-prober intel-ucode
grub-install --target=x86_64-efi --efi-directory=/efi --bootloader-id=GRUB
echo "GRUB_DISABLE_OS_PROBER=false" >> /etc/default/grub
grub-mkconfig -o /boot/grub/grub.cfg
```

Termina el proceso con el siguiente comando
```bash
exit
umount -R /mnt
reboot
```

[1]:../../README.es.md#editor-de-texto