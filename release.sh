set -e
sh build.sh
function_name=sns-notifier
d=`date +"%m-%d-%y-%H-%M-%S"`
release_dir=releases/$d
mkdir -p $release_dir
zip function.zip main.js notifier
mv function.zip $release_dir
zipfile=$release_dir/function.zip
aws lambda update-function-code --function-name $function_name --zip-file fileb://$zipfile
