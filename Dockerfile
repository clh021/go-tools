# docker build -t leehom/docker.home:nging .
# TODO:
# sudo 权限
# 主机名


# FROM alpine:latest as builder
# MAINTAINER leehom Chen <clh021@gmail.com>
# LABEL maintainer="leehom Chen <clh021@gmail.com>"
# # Install dependencies
# RUN apk update \
#     && apk add --no-cache \
#     git \
#     vim \
#     bash \


FROM alpine:latest
MAINTAINER leehom Chen <clh021@gmail.com>
LABEL maintainer="leehom Chen <clh021@gmail.com>"

USER root
RUN apk --update add --no-cache \
    sudo \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && rm -rf \
    /usr/lib/go \
    /var/cache/* \
    /var/log/* \
    /var/tmp/* \
    && mkdir /var/cache/apk

    # tzdata \
    # tree \
    # git \
    # vim \
    # bash \
    # zsh \
    # curl \
    # go \
    # htop \
    # mosh-server \
    # openssh \
    # tmux \
    # fontconfig \
    # mkfontdir \
    # mkfontscale \
    # nodejs \
    # npm \
    # htop \
    # rsync \
    # fzf \
    # ranger \
    # bind-tools \
    # zsh-autosuggestions \
    # zsh-syntax-highlighting \
# docker 支持
# RUN apk --no-cache add \
#     docker-bash-completion \
#     docker-cli \
#     docker-compose-zsh-completion \
#     docker-compose-bash-completion \
#     docker-compose

# 工具链
# js 扩展所需
#RUN apk add --no-cache --virtual .build-deps g++ python3-dev libffi-dev openssl-dev && \
#    pip3 install --upgrade pip setuptools && \
# RUN npm i -g eslint-cli js-beautify \
#     bash-language-server \
#     javascript-typescript-langserver \
#     vue-language-server \
#     vscode-css-languageserver-bin
#RUN apk add --no-cache --update gcc py3-pip libc-dev && \
#    pip3 install wheel
# RUN pip3 install neovim python-language-server
# RUN echo "http://nl.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories \
# RUN echo "http://nl.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
#     && apk --no-cache add \
#            bat \
#            neovim \
#            fzf-neovim \
#            fzf-zsh-completion \
#            fzf-tmux \
#            fzf-bash-completion

# ENV GOROOT="/usr/lib/go"
# ENV GOBIN="$GOROOT/bin"
# ENV GOPATH="$HOME/CACHE/.go-global"
# ENV PATH="$PATH:$GOBIN:$GOPATH/bin"

# RUN addgroup -S --gid 233 docker \
#     && addgroup -S --gid 1000 li \
#     && adduser -S --uid 1000 li -s "/bin/zsh" -G li \
#     && sed -i "s/# %wheel/%wheel/g" /etc/sudoers \
#     && addgroup li docker \
#     && addgroup li wheel \
#     && rm -f /etc/ssh/ssh_*_key \
#     && ssh-keygen -A \
#     && sed -i "s/#*UsePrivilegeSeparation.*/UsePrivilegeSeparation no/g" /etc/ssh/sshd_config \
#     && chmod 600 /etc/ssh/ssh_*_key
# USER li
# RUN sh -c "$(wget https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)" \
#     && curl -sLf https://spacevim.org/cn/install.sh | bash \
#     && echo "set shell=/bin/zsh" >> ~/.vimrc \
#     && echo "source /usr/share/zsh/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh" >> ~/.zshrc \
#     && echo "source /usr/share/zsh/plugins/zsh-autosuggestions/zsh-autosuggestions.zsh" >> ~/.zshrc \
#     && echo "export TERM=screen-256color" >> ~/.zshrc \
#     && mkdir ~/.ssh \
#     && echo 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCoMF04gOgH0Qc5oBZKC6r1lzrSkSG1RTNnRkRJC20ie0pkRkKpUYEosEwL9MZtzZ81pDhm9yVC72PizSpR844l3ZCjKcSy6m0qg/zoxdYPgVVOVpby+eECFWAZFQzh82hNGBtFK30I7ZAXjX1yNVrkszhLAx5lyUuHoK8vcL4btxJLyXX61rzQfy7u9cYHWc5t9iMTpb+1Fam6yKzFtFna2GFmZHEgDQZo4fgtE3Dw8YsavS2wQP76lnW18IxUmSBD6mb9zz4gIWGIhweqLiGUSsRRXta24R3khZwHIFmfNSRLUKJkIvlRoIcv9/87viBONdfTGYMxd52v3uXN+ob1 user@houyi' > ~/.ssh/authorized_keys \
#     && chown -Rf $(whoami):$(whoami) ~
# COPY init.toml ~/.SpaceVim.d/init.toml
# RUN nvim --headless +'call dein#install()' +qall
#  COPY --from=builder /usr/local/bin/ /usr/local/bin
#  COPY --from=builder /usr/local/share/vim/ /usr/local/share/vim/



#      ssh mosh
# EXPOSE 22 60001-60010/udp

# 添加有关编译的依赖 YouCompleteMe python make perl build-deps libx11
# https://github.com/JAremko/alpine-vim/blob/master/Dockerfile

# RUN sudo sh -c "$(wget https://raw.githubusercontent.com/admpub/nging/master/nging-installer.sh -O -)"
RUN wget https://img.nging.coscms.com/nging/v3.5.4/nging_linux_amd64.tar.gz && \
    ls -lah && \
    pwd
# ENV TERM=xterm-256color
USER root
# eval "$(ssh-agent -s)"
# ssh-add $UHOME/.ssh/id_rsa
# ENTRYPOINT ["/usr/sbin/sshd", "-De"]
# CMD ["tmux"]
CMD ["/usr/sbin/sshd", "-De"]
