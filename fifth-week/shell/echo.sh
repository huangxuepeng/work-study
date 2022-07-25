#!d:\git\Git\bin\bash
# printf "%d %s\n" 1 "abc"
# printf '%d %s\n' 1 "abc"
# printf %d 9
# printf %s abc def
# printf "%s\n" abc def
# printf "%s\n" a b c d e f g


# num1=100
# num2=100
# if test $[num1] -eq $[num2]
# then
#     echo '两个数相等'
# else
#     echo '不相等'
# fi
a=10
b=10
# if [ $a == $b ]
# then
#     echo "a 等于 b"
# elif [ $a -gt $b ]
# then
#     echo "a > b"
# elif [ $a -ge $b ]
# then
#     echo "a < b"
# else 
#     echo "没有符合条件的"
# fi
# if (( $a == $b ))
# then 
#     echo "a = b"
# elif (( $a > $b ))
# then
#    echo 'a > b'
# elif (( $a < $b ))
# then
#    echo 'a < b'
# fi
# if (( $a == $b )); then echo "knk"; elif (( $a > $b )); then echo 'ss'; else echo "kkk"; fi;

# num1=$[2*3]
# num2=$[1+5]
# if test $[num1] -eq $[num2]
# then
#    echo '='
# else
#    echo '!='
# fi

# for loop in 1 2 3 4 5 6 7; do echo $loop; done;
# for str in this is a string; do echo $str; done;
# a=1
# while [ $a -lt 5 ]; do
#     echo $a
#     let "a = a + 1"
# done
# echo -n '输入:'
# "${result}"
# read a
# echo $a
# si="runoob"
# case "$si" in
#     "runoob") echo "test"
#     ;;
#     "skll") echo "jjnj"
#     ;;
# esac

# funWithReturn(){
#     echo "运算"
#     echo "第一数字: "
#     read a 
#     echo "第二个数字: "
#     read b 
#     return $(($a + $b))
# }
# funWithReturn
# echo "$?"

#猜数字游戏
num=$[RANDOM%100+1]
echo "$num"

# -eq(等于) -ne(不等于) -gt(大于) -ge(大于等于) -lt(小于) -le(小于等于)
# while : 
# do 
#     read -p "菜: " cai
#     if [ "$cai" -eq "$num" ]
#     then
#        echo '对'
#        exit
#     elif [ "$cai" -gt "$num" ]
#     then
#         echo "大了"
#     else 
#         echo "小了"
#     fi     
# done
echo "$USER"

if [ $USER == "root" ]
    then
        echo "是的"
else
    echo "不是"
fi
