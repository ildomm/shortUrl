aws_user="user"
aws_address="127.0.0.1"
db_address="127.0.0.1"
db_user="user"
db_pass="pass"

echo "UPLOAD"
scp -i key.pem /linx_challenge.zip $aws_user@$aws_address:~/


echo "INSTALL GO"
ssh -i key.pem $aws_user@$aws_address
cd ~/
sudo snap install --classic go
mkdir ~/go
echo "export GOPATH=\$HOME/go" >> ~/.bash_profile
source ~/.bash_profile


echo "UNPACK"
rm -rf linx_challenge || true
mkdir -p linx_challenge
unzip linx_challenge*.zip -d linx_challenge
rm linx_challenge*.zip
mkdir -p ~/go/src/github.com/ildomm/linx_challenge
cp -R linx_challenge/ ~/go/src/github.com/ildomm/linx_challenge


echo "BUILD"
export GOROOT=/usr/local/go
export GOPATH=~/go/
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
cd ~/go/src/github.com/ildomm/linx_challenge
go -o linx build cmd/poc-server/main.go


echo "MIGRATE DB"
wget https://github.com/golang-migrate/migrate/releases/download/v4.4.0/migrate.linux-amd64.tar.gz
tar -xzf migrate.linux-amd64.tar.gz
migrate.linux-amd64 -source file://db/migrations -database mysql://$db_user:$db_pass@tcp($db_address:3306)/poc goto 0202


echo "START"
sudo start-stop-daemon -b -m --start --name "linx" --pidfile "linx.pid" --exec "linx" --chdir "/home/ubuntu/go/src/github.com/ildomm/linx_challenge"