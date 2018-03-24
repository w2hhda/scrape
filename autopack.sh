
dir=$(pwd)

git pull

bee pack

mkdir ~/scrape

cd ~/scrape

rm -rf *

cd $dir

mv scrape.tar.gz ~/scrape

cd ~/scrape

tar zxvf scrape.tar.gz

nohup ./scrape &
