rm master.zip
rm -rf bats-core-master
wget https://github.com/bats-core/bats-core/archive/master.zip
unzip master.zip
docker build .