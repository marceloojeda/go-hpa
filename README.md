# Desafio CI

Implementação do desafio Kubernetes com HPA proposto no curso Full Cycle, módulo Devops.

Trata-se de um simples HTTP server que irá realizar o somatório da raiz quadrada de 0.0001, com 10000000 interações.

A aplicação deverá rodar na porta 8000.

O objetivo é rodar essa aplicação através de um loop infinito para testar o autoscaling de pods.

# Image Docker do Square Root in Go
URL: https://hub.docker.com/r/marceloojeda/go-hpa

![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)