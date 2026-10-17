# ponto

CLI simples para registrar a jornada de trabalho do dia: início, pausas (com motivo) e encerramento. Calcula o tempo trabalhado descontando as pausas.

## Instalação

Os comandos abaixo sempre baixam a release **mais recente** do [GitHub Releases](https://github.com/rejmann/ponto/releases/latest), sem número de versão para atualizar. O `ponto` vai para `~/.local/bin`, um diretório do seu usuário: não precisa de `sudo`. Para atualizar, rode o mesmo comando de novo.

### Linux (x86_64)

```sh
mkdir -p ~/.local/bin
curl -fsSL https://github.com/rejmann/ponto/releases/latest/download/ponto-linux-amd64.tar.gz \
  | tar -xz -C ~/.local/bin ponto
export PATH="$HOME/.local/bin:$PATH"   # adicione esta linha ao ~/.bashrc / ~/.zshrc também
```

### Linux (ARM64)

```sh
mkdir -p ~/.local/bin
curl -fsSL https://github.com/rejmann/ponto/releases/latest/download/ponto-linux-arm64.tar.gz \
  | tar -xz -C ~/.local/bin ponto
export PATH="$HOME/.local/bin:$PATH"   # adicione esta linha ao ~/.bashrc / ~/.zshrc também
```

> Na maioria das distribuições `~/.local/bin` já está no `PATH`; nesse caso o `export` é desnecessário.

### Instalando uma versão específica

Troque `latest/download` por `download/<tag>` em qualquer um dos comandos acima. Por exemplo:

```sh
VERSION=v1.0.1
curl -fsSL "https://github.com/rejmann/ponto/releases/download/${VERSION}/ponto-linux-amd64.tar.gz" \
  | tar -xz -C ~/.local/bin ponto
```

Todas as tags disponíveis estão na [página de releases](https://github.com/rejmann/ponto/releases).

### Verificar

```sh
ponto --help
```

### Com Go

Com Go 1.23+:

```sh
go install github.com/rejmann/ponto/cmd/ponto@latest
```

Para compilar a partir do código:

```sh
git clone https://github.com/rejmann/ponto.git
cd ponto
go build -o ponto ./cmd/ponto
```

## Uso

Um dia típico:

```sh
ponto iniciar          # bom dia! jornada iniciada as 09:00.
ponto pausar almoco    # inicia uma pausa com motivo "almoco"
ponto voltar           # de volta! pausa (almoco) durou 1h00m00s.
ponto status           # mostra o andamento do dia
ponto encerrar         # encerra a jornada
```

### Comandos

| Comando | Descrição |
|---|---|
| `ponto iniciar` | Inicia a jornada de hoje. Falha se a jornada de hoje já foi iniciada ou encerrada. |
| `ponto pausar [motivo]` | Inicia uma pausa. O motivo pode ter várias palavras (`ponto pausar cafe com o time`); sem motivo, usa `pausa`. Só é permitida uma pausa ativa por vez. |
| `ponto voltar` | Encerra a pausa em andamento e mostra quanto ela durou. |
| `ponto status` | Mostra data, horário de início, tempo trabalhado e a lista de pausas. |
| `ponto encerrar` | Encerra a jornada. Se houver uma pausa ativa, ela é encerrada no mesmo horário. |
| `ponto historico <AAAA-MM-DD>` | Mostra o registro de um dia específico (ex: `ponto historico 2026-10-15`). |

Exemplo de saída do resumo do dia (`status`):

```
data: 2026-10-16
inicio: 09:00
tempo trabalhado: 6h12m40s
pausas:
  - almoco - 12:00-13:00 1h00m00s

  - cafe - 15:30-15:47 0h17m03s (em andamento)
```

### Regras

- Existe no máximo **uma jornada por dia**, identificada pela data local (`AAAA-MM-DD`). Depois de encerrada, não é possível reabrir nem iniciar outra no mesmo dia.
- `pausar`, `voltar` e `encerrar` só funcionam com a jornada **de hoje** em andamento. Uma jornada que atravessa a meia-noite não pode ser encerrada no dia seguinte.
- **Tempo trabalhado** = (fim da jornada − início) − soma das pausas. Enquanto a jornada ou uma pausa estiver aberta, o horário atual é usado como fim.

## Armazenamento

Cada dia é salvo como um arquivo JSON em `~/.ponto/`, nomeado pela data:

```
~/.ponto/2026-10-16.json
```

```json
{
  "data": "2026-10-16",
  "pausas": [
    {
      "motivo": "almoco",
      "inicio": "2026-10-16T12:00:00-03:00",
      "fim": "2026-10-16T13:00:00-03:00"
    }
  ],
  "inicio": "2026-10-16T09:00:00-03:00",
  "fim": "2026-10-16T18:00:00-03:00"
}
```

O diretório é criado com permissão `0700` e os arquivos com `0600`. Os arquivos podem ser editados à mão para corrigir um horário; se o JSON ficar inválido, o comando avisa que o arquivo está corrompido.

## Estrutura do código

Tudo fica no pacote `main` em [cmd/ponto/](cmd/ponto/), usando [cobra](https://github.com/spf13/cobra):

| Arquivo | Conteúdo |
|---|---|
| `main.go` | Comando raiz e registro dos subcomandos |
| `ponto.go` | Tipos `Dia`, `Pausa`, `Periodo`, constantes de formato e cálculo do tempo trabalhado |
| `iniciar.go` | Comando `iniciar` e leitura/gravação dos arquivos (`carregarDia`, `salvarDia`) |
| `pausar.go`, `voltar.go`, `status.go`, `historico.go` | Um arquivo por comando |
| `encerrar.go` | Comando `encerrar` e os helpers de exibição (`diaAberto`, `imprimirDia`, `fmtDur`) |

## Release

O workflow [release.yml](.github/workflows/release.yml) roda quando um PR que altera arquivos Go (`**.go`, `go.mod`, `go.sum`) é mergeado na `main`. Ele gera a tag de versão, compila o binário para Linux e publica uma GitHub Release com notas geradas automaticamente, usando os workflows reutilizáveis de [rejmann/github-actions](https://github.com/rejmann/github-actions).

Também é possível disparar um release manualmente, pela aba **Actions → Release → Run workflow** ou com:

```sh
gh workflow run release.yml --ref main
```

Se o commit da `main` já tiver uma tag `v*`, a mesma versão é reaproveitada; caso contrário, é criada a próxima versão de patch.
