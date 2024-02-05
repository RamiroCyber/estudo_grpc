# Estudo gRPC

## Visão Geral

Este projeto demonstra a implementação de um serviço gRPC usando a linguagem Go. Ele oferece um exemplo prático de como o gRPC facilita a construção de comunicação eficiente baseada em RPC entre cliente e servidor. O serviço, chamado `Invoice`, permite a criação de faturas que são retornadas nos formatos PDF e DOCX.

## O que é gRPC?

O gRPC é um framework moderno e de código aberto para Chamada de Procedimento Remoto (RPC) que permite a comunicação transparente e eficiente entre aplicações cliente e servidor. Desenvolvido pelo Google, ele usa o HTTP/2 para transporte, Protocol Buffers como a linguagem de descrição de interface e oferece recursos como autenticação, balanceamento de carga e mais.

### Principais Características do gRPC:

- **Baseado em HTTP/2**: Utiliza as funcionalidades modernas do HTTP/2, como fluxos, gerenciamento eficiente de conexões e mais.
- **Linguagem de Descrição de Interface (IDL)**: Usa Protocol Buffers, mecanismo maduro e de código aberto do Google para serialização de dados estruturados.
- **Suporte Multi-Idiomas**: Oferece suporte a múltiplos idiomas, facilitando a criação de serviços que funcionam de forma eficiente e segura entre diferentes linguagens.
- **Suporte a Streaming**: Fornece capacidades de streaming (lado do cliente, lado do servidor ou bidirecional) para troca de dados em tempo real.

## Estrutura do Projeto

- `invoice.proto`: Define o serviço gRPC e as mensagens usando Protocol Buffers.
- `server/main.go`: Implementa a lógica do lado do servidor do serviço de Faturas.

## Serviço de Faturas

O serviço de Faturas é definido no arquivo `invoice.proto`. Ele consiste em um único método RPC:

- `Create(CreateRequest) returns (CreateResponse)`: Aceita uma solicitação com detalhes da fatura (valor, remetente, destinatário) e retorna a fatura gerada nos formatos PDF e DOCX, além de um status de sucesso.

### Mensagens:

- `CreateRequest`: Contém o valor da fatura, moeda, detalhes do remetente e do destinatário.
- `CreateResponse`: Inclui a fatura gerada nos formatos PDF e DOCX e um indicador de sucesso.

## Executando o Serviço

Para executar o serviço, certifique-se de ter o Go e o compilador de Protocol Buffers instalados. Gere o código gRPC e, em seguida, inicie o servidor executando `cmd/server/main.go`. Se existir uma implementação de cliente, ela pode se comunicar com o servidor através dos métodos gRPC definidos.
