package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const ShellToUse = "bash"

var virtualMachine bool = false

func Shell(command string) error {
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func installPackages() bool {
	var err error = nil

	// Update Packages
	fmt.Println("[Update Packages]")
	err = Shell("sudo pacman -Syu --noconfirm")
	if err != nil {
		log.Printf("Error Update Packages: %v\n", err)
		return false
	}

	// Install require package
	packageRequire := []string{
		"gcc", "make", "git", "base-devel",
		"numlockx",
		"nmap", "wget", "curl", "inetutils", "dnsutils",
		"xclip",
		"p7zip", "unzip",
		"zsh", "tmux",
		"dunst", "ranger", "htop", "locate",
		"flameshot",
		"libsecret", "gnome-keyring", "gnome-themes-extra",
		"xorg-server", "xorg-xev",
		"bspwm", "sxhkd",
		"alacritty", "kitty",
		"rofi", "polybar", "picom",
		"lightdm", "lightdm-gtk-greeter",
		"bat", "lsd", "fzf", "jq",
		"feh",
		"neofetch",
		"pipewire", "pipewire-pulse",
		"docker", "docker-compose", "kubectl", "minikube",
		"udisks2", "udiskie",
		"openssh",
		"firefox", "firejail",
		"noto-fonts-emoji",
		"dbeaver",
		"keepassxc", "redshift",
	}

	if !virtualMachine {
		packageRequire = append(packageRequire, "bluez", "bluez-utils")
	}

	fmt.Println("[Install Packages]")

	for _, p := range packageRequire {
		fmt.Println("- [" + p + "]")
		err = Shell("sudo pacman -S --needed --noconfirm " + p)
		if err != nil {
			log.Printf("Error Install Package [%v]: %v\n", p, err)
			return false
		}
	}

	// Install Yay
	fmt.Println("- [Yay]")
	err = Shell("cd /tmp && git clone https://aur.archlinux.org/yay.git && cd yay && makepkg -si && cd .. && rm -rf /tmp/yay && yay -Yc")
	if err != nil {
		log.Printf("Error Install Package [%v]: %v\n", "Yay", err)
		return false
	}

	// Install package with Yay
	packageYay := []string{
		"google-chrome",
		"kubecolor",
		"openrgb",
		"visual-studio-code-bin",
	}

	for _, p := range packageYay {
		fmt.Println("- [yay: " + p + "]")
		err = Shell("yay -S --needed --noconfirm " + p)
		if err != nil {
			log.Printf("Error Install Package Yay [%v]: %v\n", p, err)
			return false
		}
	}

	return true
}

func setupFonts() bool {
	var err error = nil
	var _url string = ""
	var _path string = ""

	// Fonts
	fmt.Println("[Fonts]")

	_version := "v2.1.0"
	_url = "https://github.com/ryanoasis/nerd-fonts/releases/download/" + _version + "/"
	_path = "/usr/local/share/fonts/nerd-fonts/"
	nerdFonts := []string{
		"Hack",
		"Meslo",
		"JetBrainsMono",
	}

	for _, filename := range nerdFonts {
		fmt.Println("- [" + filename + " Nerd Font]")
		command := "mkdir -p " + _path + filename + " && " +
			"cd " + _path + filename + " && " +
			"wget --quiet " + _url + filename + ".zip && " +
			"unzip " + filename + ".zip && " +
			"rm " + filename + ".zip"

		err = Shell("sudo " + ShellToUse + " -c '" + command + "'")
		if err != nil {
			log.Printf("Error Fonts [%v]: %v\n", filename, err)
			return false
		}
	}

	_url = "https://github.com/adi1090x/rofi/raw/master/fonts/"
	_path = "/usr/local/share/fonts/Rofi/"
	rofiFonts := []string{
		"GrapeNuts-Regular",
		"Icomoon-Feather",
		"Iosevka-Nerd-Font-Complete",
		"JetBrains-Mono-Nerd-Font-Complete",
	}

	command := "mkdir -p " + _path + " && " +
		"cd " + _path + " && "

	for _, filename := range rofiFonts {
		fmt.Println("- [Rofi " + filename + " Font]")
		command += "wget --quiet " + _url + filename + ".ttf --output-document " + filename + ".ttf && "
	}

	err = Shell("sudo " + ShellToUse + " -c '" + strings.TrimRight(command, " && ") + "'")
	if err != nil {
		log.Printf("Error Fonts [%v]: %v\n", "Rofi", err)
		return false
	}

	_url = "https://github.com/Templarian/MaterialDesign-Font/raw/master/"
	_path = "/usr/local/share/fonts/MaterialDesign/"
	filename := "MaterialDesignIconsDesktop"

	fmt.Println("- [Material Design Icons Font]")
	command = "mkdir -p " + _path + " && " +
		"cd " + _path + " && " +
		"wget --quiet " + _url + filename + ".ttf --output-document " + filename + ".ttf"

	err = Shell("sudo " + ShellToUse + " -c '" + command + "'")
	if err != nil {
		log.Printf("Error Fonts [%v]: %v\n", "Material Design Icons Font", err)
		return false
	}

	fmt.Println("- [Fonts Reload]")

	err = Shell("sudo " + ShellToUse + " -c 'fc-cache -vf'")
	if err != nil {
		log.Printf("Error Fonts Reload: %v\n", err)
		return false
	}

	return true
}

func setupENV() bool {
	var err error = nil

	// Setup env
	fmt.Println("[File Settings]")
	file := map[string]string{
		"1) Creating directories Downloads":   "mkdir -p ~/Downloads/{Chrome,Firefox}",
		"1.1) Creating directories config":    "mkdir -p ~/{.config,.ssh,Images}",
		"2) Set up Chrome dark theme":         "echo '--force-dark-mode' >> ~/.config/chrome-flags.conf",
		"3) Set system dark theme":            "sudo gsettings set org.gnome.desktop.interface gtk-theme 'Adwaita-dark'",
		"4) Set system dark color":            "sudo gsettings set org.gnome.desktop.interface color-scheme 'prefer-dark'",
		"5) Set layout keyboard ":             "sudo localectl set-x11-keymap es",
		"6) Set lightdm greeter-session":      "sudo sed -i 's/#greeter-session=example-gtk-gnome/greeter-session=lightdm-gtk-greeter/' /etc/lightdm/lightdm.conf",
		"7) Set lightdm display-setup-script": "sudo sed -i 's/#display-setup-script=/display-setup-script=\\/usr\\/bin\\/setxkbmap es/' /etc/lightdm/lightdm.conf",
		"8) Set lightdm greeter-setup-script": "sudo sed -i 's/#greeter-setup-script=/greeter-setup-script=\\/usr\\/bin\\/numlockx on/' /etc/lightdm/lightdm.conf",
		"9) Enable lightdm service":           "sudo systemctl enable lightdm",
	}

	if !virtualMachine {
		file["Enable Bluetooth service"] = "sudo systemctl enable bluetooth"
		file["Auto Enable Bluetooth"] = "sudo sed -i 's/#AutoEnable=true/AutoEnable=true/' /etc/bluetooth/main.conf"
	}

	for k, c := range file {
		fmt.Println("- [" + k + "]")
		err = Shell(c)
		if err != nil {
			log.Printf("Error File Settings [%v]: %v\n", k, err)
			return false
		}
	}

	fmt.Println("- [Create xprofile]")

	if virtualMachine {
		err = Shell("printf \"VBoxClient-all &\n\" >> ~/.xprofile")
		if err != nil {
			log.Printf("Error Create xprofile: %v\n", err)
			return false
		}
	}

	err = Shell("printf \"dbus-update-activation-environment --systemd DISPLAY &\nsxhkd &\nexec bspwm\" >> ~/.xprofile")
	if err != nil {
		log.Printf("Error Create xprofile: %v\n", err)
		return false
	}

	if !virtualMachine {
		fmt.Println("- [Auto connection bluetooth]")

		err = Shell("sudo " + ShellToUse + " -c 'printf \"### Automatically switch to newly-connected devices\nload-module module-switch-on-connect\n\" >> /etc/pulse/default.pa'")
		if err != nil {
			log.Printf("Error Auto connection bluetooth: %v\n", err)
			return false
		}
	}

	fmt.Println("[Git Settings]")
	gitGlobals := []string{
		"alias.co checkout",
		"alias.br branch",
		"alias.ci commit",
		"alias.st status",
		"init.defaultBranch main",
		"core.autocrlf input",
	}

	for _, c := range gitGlobals {
		fmt.Println("- [Git: " + c + "]")
		err = Shell("git config --global " + c)
		if err != nil {
			log.Printf("Error Git Settings: %v\n", err)
			return false
		}
	}

	fmt.Println("[Docker Settings]")
	dockerSettings := map[string]string{
		"1) Enable service":                   "sudo systemctl enable docker",
		"2) Add group www-data":               "sudo groupadd -r -g 82 www-data",
		"3) Add user www-data":                "sudo useradd -M -r -u 82 -g 82 -c \"User HTTP files\" -s /usr/bin/nologin www-data",
		"4) Add permission to group www-data": "sudo usermod -aG docker,www-data $(whoami)",
		"5) Restart service":                  "sudo systemctl restart docker",
	}

	for k, c := range dockerSettings {
		fmt.Println("- [" + k + "]")
		err = Shell(c)
		if err != nil {
			log.Printf("Error Docker Settings [%v]: %v\n", k, err)
			return false
		}
	}

	fmt.Println("[Copying config files]")

	path, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting the path to change: %v\n", err)
		return false
	}
	path = regexp.MustCompile(`\/utils$`).ReplaceAllString(path, "")

	fmt.Println("- [Changing directory]")
	err = Shell("cd \"" + path + "\"")
	if err != nil {
		log.Printf("Error changing directory: %v\n", err)
		return false
	}

	fmt.Println("- [Add permissions to files]")
	pathToChangePermissions := []string{
		".config/bspwm/bspwmrc",
		".config/bspwm/scripts/{bspwm_count,bspwm_layout,bspwm_resize,bspwm_smart_move}",
	}

	for _, c := range pathToChangePermissions {
		fmt.Println("- [chmod: " + c + "]")
		err = Shell("sudo chmod +x $PWD/" + c)
		if err != nil {
			log.Printf("Error changing permissions: %v\n", err)
			return false
		}
	}

	fmt.Println("- [Create symbolic links of the files]")
	symlinkPath := []string{
		".config/{alacritty,bspwm,gtk-3.0,kitty,picom,polybar,sxhkd,systemd} ~/.config/",
		"{.aliases,.bashrc,.p10k.zsh,.xprofile,.zshrc,.tmux.conf} ~/",
	}

	for _, c := range symlinkPath {
		fmt.Println("- [ln: " + c + "]")
		err = Shell("ln -sf $PWD/" + c)
		if err != nil {
			log.Printf("Error creating symbolic link: %v\n", err)
			return false
		}
	}

	fmt.Println("- [Change default shell per user]")
	usersShell := []string{
		"$(whoami)",
		"root",
	}

	for _, c := range usersShell {
		fmt.Println("- [usermod: " + c + "]")
		err = Shell("sudo usermod --shell /usr/bin/zsh " + c)
		if err != nil {
			log.Printf("Error changing user shell: %v\n", err)
			return false
		}
	}

	/*
	   # Install Oh My Zsh
	   sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

	   # Add plugins and themes
	   git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
	   git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
	   git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/themes/powerlevel10k

	   # Install Oh My Zsh for root user
	   sudo sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

	   # Rest after installation
	   exit

	   # Create symbolic links of the files to the root user
	   sudo ln -sf ~/.oh-my-zsh/custom/plugins/{zsh-autosuggestions,zsh-syntax-highlighting} /root/.oh-my-zsh/custom/plugins/
	   sudo ln -sf ~/.oh-my-zsh/custom/themes/powerlevel10k /root/.oh-my-zsh/custom/themes/
	   sudo ln -sf ~/{.aliases,.bashrc,.p10k.zsh,.zshrc} /root/
	*/

	fmt.Println("- [Set Wallpaper]")
	err = Shell("wget --quiet https://wallpaperaccess.com/full/2098223.png --output-document ~/Images/wallpaper.png")
	if err != nil {
		log.Printf("Error downloading the wallpaper: %v\n", err)
		return false
	}

	fmt.Println("- [Change power button behavior]")
	err = Shell("sudo sed -i 's/#HandlePowerKey=poweroff/HandlePowerKey=ignore/' /etc/systemd/logind.conf")
	if err != nil {
		log.Printf("Error changing poweroff behavior: %v\n", err)
		return false
	}

	err = Shell("sudo systemctl restart systemd-logind")
	if err != nil {
		log.Printf("Error restart systemd-logind: %v\n", err)
		return false
	}

	return true
}

func main() {
	fmt.Println(strings.Repeat("_", 20))

	var _userInput string

	fmt.Println("Are you working in a virtual machine? [y/N]: ")
	fmt.Scanln(&_userInput)
	_userInput = strings.Trim(strings.ToLower(_userInput), " ")

	if _userInput == "y" || _userInput == "yes" {
		virtualMachine = true
	}

	statusPackages := installPackages()
	if statusPackages {
		fmt.Println(strings.Repeat("#", 10) + " Install Packages OK " + strings.Repeat("#", 10))
	} else {
		fmt.Println(strings.Repeat("#", 10) + " Install Packages Error " + strings.Repeat("#", 10))
		os.Exit(1)
	}

	statusFonts := setupFonts()
	if statusFonts {
		fmt.Println(strings.Repeat("#", 10) + " Install Fonts OK " + strings.Repeat("#", 10))
	} else {
		fmt.Println(strings.Repeat("#", 10) + " Install Fonts Error " + strings.Repeat("#", 10))
		os.Exit(1)
	}

	statusSetup := setupENV()
	if statusSetup {
		fmt.Println(strings.Repeat("#", 10) + " Setup env OK " + strings.Repeat("#", 10))
	} else {
		fmt.Println(strings.Repeat("#", 10) + " Setup env Error " + strings.Repeat("#", 10))
		os.Exit(1)
	}
}
