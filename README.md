# Politicos API

Aqui você vai encontrar os dados, ferramentas e recursos para realizar
pesquisas, desenvolver aplicativos, visualizações de dados e muito mais.

## Instalação

### Instalando as dependências

```
make setup
```

### Rodando a API

```bash
make run
```

### Pegando os CSV dos candidatos

```bash
make build-collector
```

```bash
./collector fetch-all-csv
```

### Rodando o coletor

```bash
make build-collector
```

```bash
./collector process-all
```

### Rotas da API (local)

http://localhost:8888/swagger/index.html

## Como Contribuir

Participe da [lista de email][lista] e também via IRC no canal
[#olhoneles][freenode] na Freenode.

Faça uma cópia do repositório e envie seus pull-requests.


## Licença

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, either version 3 of the License, or (at your option) any
later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.  See the GNU Affero General Public License for more
details.

You should have received a copy of the GNU Affero General Public License along
with this program.  If not, see <http://www.gnu.org/licenses/>.

[lista]: http://listas.olhoneles.org/cgi-bin/mailman/listinfo/montanha-dev
[freenode]: irc://irc.freenode.net:6667/olhoneles
