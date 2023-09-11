FROM ubuntu:latest

RUN  apt install update 
RUN apt-get install build-essential curl file git
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
RUN eval "$(homebrew/bin/brew shellenv)"
RUN brew update --force --quiet
RUN chmod -R go-w "$(brew --prefix)/share/zsh"
RUN test -d ~/.linuxbrew && eval "$(~/.linuxbrew/bin/brew shellenv)"
RUN test -d /home/linuxbrew/.linuxbrew && eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"
RUN test -r ~/.bash_profile && echo "eval \"\$($(brew --prefix)/bin/brew shellenv)\"" >> ~/.bash_profile
RUN echo "eval \"\$($(brew --prefix)/bin/brew shellenv)\"" >> ~/.profile
RUN brew install nuclei
RUN brew install httpx
RUN brew install amass

