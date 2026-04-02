# 🛡️ scanner

Scanner simples em Go para testes básicos de segurança em URLs (SQL Injection, XSS e headers HTTP de proteção).

## ✨ Funcionalidades

- Teste de SQL Injection via payloads:
  - `'
  - `' OR '1'='1`
  - `'--`
- Detecção de erros comuns de banco:
  - `SQL syntax`, `mysql_fetch`, `ORA-01756`, `syntax error`
- Teste de XSS via payload:
  - `<script>alert(1)</script>`
- Verificação de headers de segurança:
  - `X-Frame-Options`
  - `Content-Security-Policy`
  - `X-XSS-Protection`

---

## 🧱 Pré-requisitos

- Go instalado (recomendo `Go 1.21+` ou `Go 1.22`)
- Acesso à Internet

---

## 🛠️ Como instalar / rodar

1. Abra terminal em `c:\coding\scanner`

2. Instale dependências:

```powershell
go get github.com/go-resty/resty/v2
go get github.com/inancgumus/screen
```

3. Build e executar:

```powershell
go build -o scanner.exe
.\scanner.exe
```

ou direto:

```powershell
go run main.go
```

---

## ▶️ Uso

Ao rodar, ele pede:

- `Digite a url`
- Exemplo: `http://site.com/page?id=`

Basta informar URL com o parâmetro que você quer testar (`id=` ou similar).

---

## 📌 Exemplo

Input:

- `http://test.local/page?id=`

Resultados possivelmente mostrados no terminal:

- `[!] Possivel Sql Injection`
- `[!] Possivel XSS: ...]`
- `[!] Falta X-Frame-Options`
- `[!] Falta CSP`
- `[!] Falta X-XSS-Protection`

---

## 💡 Dicas

- Use em ambiente de teste / autorização apenas (pentest legal).
- Adeque o `baseURL` conforme endpoint correto (`?id=`, `search=`, etc).
- Se quiser adicionar mais payloads, mude `sqliPayloads` e `sqliErrors`.

---

## 📝 Arquivos

- `main.go` (código principal)
- `go.mod` (módulos)
- `README.md` (este arquivo)
