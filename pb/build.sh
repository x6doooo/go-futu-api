#!bin/sh
for file in ./*.proto
do
    if test -f $file
    then
        new_dir=${file%.*}
        mkdir $new_dir
        # echo ${new_dir}

        protoc -I=./ --go_out=./$new_dir/ $file
    fi
    if test -d $file
    then
        echo $file "is dir"
    fi
done