
Componentes:
Sistema gerado a partir de padrao Swagger, contendo banco de dados Mysql.

ps: nao esta publicado no github, apenas mantive a referencia de diretorio
utilizada em outros projetos no meu ambiente de desenvolvimento.

Etapas para instalacao:
1) Criar banco de dados Mysql nomeado "poc"
2) Criar usuario com permisao para criacao de tabelas
3) Configurar db user/pass em config/config.json
4) Configurar db user/pass em deploy.sh
5) Configurar AWS user/ip em deploy.sh
6) Mover deploy.sh para diretorio anterior a linx_challenge
7) Executar deploy.sh
8) Acessar http://127.0.0.1/v1/docs , ou ip referente a instancia AWS
