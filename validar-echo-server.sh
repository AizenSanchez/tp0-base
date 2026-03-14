docker image pull alpine:latest > /dev/null
response=$(docker run --network tp0_testing_net alpine:latest sh -c "echo hola server | nc server 12345")
if [ "$response" = "hola server" ]; then
    printf "action: test_echo_server | result: success"
else
    printf "action: test_echo_server | result: fail"
fi