# cv-admin

## env

### OS

```
$ cat /etc/os-release 
PRETTY_NAME="Ubuntu 22.04.1 LTS"
```

### go

```
$ go version
go version go1.19 linux/amd64
```

## setup

### project

```
$ go mod init github.com/sky0621/cv-admin
go: creating new go.mod: module github.com/sky0621/cv-admin
```

### open-api-codegen

https://github.com/getkin/kin-openapi
https://github.com/deepmap/oapi-codegen

```
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

## use

### ent

```
cd src/interfaces/ent
go run -mod=mod entgo.io/ent/cmd/ent init User UserActivity UserQualification UserCareerGroup UserCareer CareerTask CareerSkillGroup CareerSkill UserNote UserNoteItem
```