

provider "vcloud-director" {
  user                 = "user1"
  password             = "Admin!23"
  org                  = "O1"
  ip                   = "10.112.83.27"
  vdc                  = "V03"
  max_retry_timeout      = "30"
  allow_unverified_ssl = "true"
}



resource "vcloud-director_catalog" "cata1" {
        name    = "ct24"
        description = "desc"
        shared  = "true"

}


resource "vcloud-director_catalog" "cata2" {
        name    = "ct11"
        description = "desc"
        shared  = "true"

}

