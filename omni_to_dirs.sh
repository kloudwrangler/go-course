function omni_to_dirs(){
    dirs=$(cat - | sed "s/^- //" | sed 's/^\* //' | tr '[:lower:]' '[:upper:]' | sed 's/ /-/g')
    counter=1
    for dir ($=dirs); do
        num="$(printf "%02d" $counter)"
        echo $num-$dir
        counter=$((counter+1))

    done
}