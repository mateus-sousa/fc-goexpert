Para fazer download de modulos privados voce deve:
Adicionar a url do repo privado a sua variavel GOPRIVATE

export GOPRIVATE=<repositorioprivado>

Se você ja estiver autenticado nesse repositorio, o download ja funcionará, senão.

VocÊ pode autenticar via http:

criar arquivo na raiz do computador:
vim ~/.netrc
Para fazer download de modulos privados voce deve:
Adicionar a url do repo privado a sua variavel GOPRIVATE

export GOPRIVATE=<repositorioprivado>

Se você ja estiver autenticado nesse repositorio, o download ja funcionará, senão.

VocÊ pode autenticar via http:

criar arquivo na raiz do computador:
vim ~/.netrc

dentro do arquivo adicionar o conteudo

machine <repositorio raiz do modulo privado> (bitbucket tem uma pegadinha, deve se adicionar api como prefixo, ficando: api.bitbucket.org)
login <login>
password <token de acesso ao repositorio>

para fazer autenticaçao via SSH, deve tambem adicionar ao seu g
dentro do arquivo adicionar o conteudo

machine <repositorio raiz do modulo privado> (bitbucket tem uma pegadinha, deve se adicionar api como prefixo, ficando: api.bitbucket.org)
login <login>
password <token de acesso ao repositorio>

para fazer autenticaçao via SSH, deve tambem adicionar ao seu gitconfig

vim ~/.gitconfig

adicionar conteudo: (exemplo se fosse github)

[url "ssh://git@github.com/"]
    insteadOf = https://github.com/