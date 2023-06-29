function my_fun(){
    dirs="A\nB\nC" 
    echo $dirs
    # for dir (${(f)dirs}); do
    #     echo $dir
    #     echo "hello"
    # done
    for num (${(@s/\n/)dirs}); do 
        print -r New item: $num
    done
}
my_fun