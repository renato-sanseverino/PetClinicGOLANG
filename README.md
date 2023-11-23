# PetClinicGOLANG
Pet Clinic API written in GOLANG

![screenshot](assets/banner.png)

## Stack:
- GIN web framework
- Prisma ORM for GO
- Postgres


## Features:
- Agendamento de consultas veterinárias
- Histórico de vacinas e tratamento para os pets
- Pesquisa de satisfação e contatos através de SMS e Whatsapp

## Steps to run the project
- Crie as tabelas no banco com o comando:
    > go run github.com/steebchen/prisma-client-go db push
- go run main.go   (alternativamente:  docker compose up)
- follow the link http://localhost:3000/api/pets
