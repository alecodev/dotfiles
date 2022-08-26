package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const ShellToUse = "bash"

func Shell(command string) error {
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func installPackages() bool {
	// Update Packages
	fmt.Println("[Update Packages]")
	err := Shell("sudo pacman -Syu --noconfirm")
	if err != nil {
		log.Printf("Error Update Packages: %v\n", err)
		return false
	}

	// Install require package
	packageRequire := []string{
		"gcc", "make", "git", "base-devel",
		"numlockx",
		"nmap", "wget", "curl",
		"xclip",
		"p7zip", "unzip",
		"zsh", "zsh-autosuggestions", "zsh-syntax-highlighting",
		"tmux",
		"htop",
		"flameshot",
		"libsecret", "gnome-keyring",
		"xorg-server", "xorg-xev",
		"bspwm", "sxhkd",
		"alacritty",
		"rofi", "polybar", "picom",
		"lightdm", "lightdm-gtk-greeter",
		"bat", "lsd", "fzf", "jq",
		"feh",
		"neofetch",
		"bluez", "bluez-utils", "pipewire",
		"docker", "docker-compose",
		"openssh",
		"firefox", "firejail",
		"noto-fonts-emoji",
		"dbeaver",
	}

	fmt.Println("[Install Packages]")

	for _, p := range packageRequire {
		fmt.Println("- [" + p + "]")
		err := Shell("sudo pacman -S --needed --noconfirm " + p)
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
		"visual-studio-code-bin",
	}

	for _, p := range packageYay {
		fmt.Println("- [yay: " + p + "]")
		err := Shell("yay -S --needed --noconfirm " + p)
		if err != nil {
			log.Printf("Error Install Package Yay [%v]: %v\n", p, err)
			return false
		}
	}

	return true
}

func setupENV() bool {
	// Setup env
	fmt.Println("[File Settings]")
	file := map[string]string{
		"Creating directories":     "mkdir -p ~/Downloads/{Chrome,Firefox}",
		"Set up Chrome dark theme": "echo '--force-dark-mode' >> ~/.config/chrome-flags.conf",
		"Set system dark theme":    "gsettings set org.gnome.desktop.interface gtk-theme 'Adwaita-dark'",
	}

	for k, c := range file {
		fmt.Println("- [" + k + "]")
		err := Shell(c)
		if err != nil {
			log.Printf("Error File Settings [%v]: %v\n", k, err)
			return false
		}
	}

	return true
}

func setupFonts() bool {
	// Fonts
	fmt.Println("[Fonts]")

	_version := "v2.1.0"
	_url := "https://github.com/ryanoasis/nerd-fonts/releases/download/" + _version + "/"
	_path := "/usr/local/share/fonts/nerd-fonts/"
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

		err := Shell("sudo " + ShellToUse + " -c '" + command + "'")
		if err != nil {
			log.Printf("Error Fonts [%v]: %v\n", filename, err)
			return false
		}
	}

	/*
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
			command += "wget --quiet " + _url + filename + ".ttf --output-document=" + filename + ".ttf && "
		}

		err := Shell("sudo " + ShellToUse + " -c '" + strings.TrimRight(command, " && ") + "'")
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
			"wget --quiet " + _url + filename + ".ttf --output-document=" + filename + ".ttf"

		err = Shell("sudo " + ShellToUse + " -c '" + command + "'")
		if err != nil {
			log.Printf("Error Fonts [%v]: %v\n", "Material Design Icons Font", err)
			return false
		}
	*/

	fmt.Println("- [Fonts Reload]")

	var err = Shell("sudo " + ShellToUse + " -c 'fc-cache -vf'")
	if err != nil {
		log.Printf("Error Fonts Reload: %v\n", err)
		return false
	}

	return true
}

func main() {
	fmt.Println(strings.Repeat("_", 20))

	statusPackages := installPackages()
	if statusPackages {
		fmt.Println(strings.Repeat("#", 10) + " Install Packages OK " + strings.Repeat("#", 10))
	} else {
		fmt.Println(strings.Repeat("#", 10) + " Install Packages Error " + strings.Repeat("#", 10))
		os.Exit(1)
	}

	statusSetup := setupENV()
	if statusSetup {
		fmt.Println(strings.Repeat("#", 10) + " Setup env OK " + strings.Repeat("#", 10))
	} else {
		fmt.Println(strings.Repeat("#", 10) + " Setup env Error " + strings.Repeat("#", 10))
		os.Exit(1)
	}

	statusFonts := setupFonts()
	if statusFonts {
		fmt.Println(strings.Repeat("#", 10) + " Install Fonts OK " + strings.Repeat("#", 10))
	} else {
		fmt.Println(strings.Repeat("#", 10) + " Install Fonts Error " + strings.Repeat("#", 10))
		os.Exit(1)
	}
}
