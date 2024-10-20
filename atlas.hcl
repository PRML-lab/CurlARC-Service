data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/infra",
    "--dialect", "postgres", // | postgres | sqlite | sqlserver
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "local" {
  url = "postgres://app:password@localhost:5432/app?sslmode=disable"
  migration {
    dir = "atlas://curlarc"
  }
}

env "prod" {
  url = "postgres://postgres:EbmU6Q0LbRbe0LV@localhost:5432?sslmode=disable"
  migration {
    dir = "atlas://curlarc"
  }
}