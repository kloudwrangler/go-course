function make_files(){
    FILE=$(basename "$PWD" | tr '[:upper:]' '[:lower:]' | gsed -r 's/^[0-9]+-//');
    for type in go md; do touch "$FILE.$type"; done;
    for type in slides exercise example; do touch "$FILE-$type.md"; done;
}

function omni_to_dirs(){
    dirs=$(cat - | sed "s/^- //" | sed 's/^\* //' | tr '[:lower:]' '[:upper:]' | sed 's/ /-/g')
    counter=1
    for dir ($=dirs); do
        num="$(printf "%02d" $counter)"
        echo $num-$dir
        counter=$((counter+1))

    done
}