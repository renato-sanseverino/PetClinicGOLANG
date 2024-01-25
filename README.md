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
- Set DATABASE_URL in the .env file
- Run the script to generate prisma client and create the database:
    > go run github.com/steebchen/prisma-client-go db push
- docker compose up
- Follow the link http://localhost:3000/api/pets

## Deploy the project to a Kubernetes cluster
- Run
    > kubectl apply -f deploy/petclinic.yaml
