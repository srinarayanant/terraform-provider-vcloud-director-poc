terraform init
terraform plan -out o1
./killp.sh
terraform apply o1

