#!/usr/bin/env fish

function announce
    set_color purple --bold
    echo " +++ $argv +++ "
    set_color normal
end

function announce_and_execute
    announce $argv

    sleep 1

    eval $argv
end

set commands \
    'http GET :1323/products' \
    'http GET :1323/products/12' \
    'http PUT :1323/products/12 name="Hair dryer" price:=39.99 categoryId:=2' \
    'http GET :1323/products/12' \
    'http DELETE :1323/products/12' \
    'http GET :1323/products/12' \

for command in $commands
    announce_and_execute $command
end

announce 'http POST :1323/carts'
set response (eval 'http POST :1323/carts')
echo $response

set cartID (echo $response | jq .ID)

announce "Using cart ID: $cartID"

set commands \
    "http GET :1323/carts/$cartID/products" \
    "http POST :1323/carts/$cartID/products/18" \
    "http POST :1323/carts/$cartID/products/21" \
    "http POST :1323/carts/$cartID/products/23" \
    "http GET :1323/carts/$cartID/products" \
    "http DELETE :1323/carts/$cartID/products/18" \
    "http GET :1323/carts/$cartID/products" \
    "http DELETE :1323/carts/$cartID/products" \
    "http GET :1323/carts/$cartID/products" \
    "http DELETE :1323/carts/$cartID" \
    "http GET :1323/carts/$cartID" \

for command in $commands
    announce_and_execute $command
end
