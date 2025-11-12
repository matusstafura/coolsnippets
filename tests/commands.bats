#!/usr/bin/env bats

APP="./coolsnippets"

@test "strip-tags-newline converts br to newline" {
    result=$(echo "hello<br>world" | $APP -u strip-tags-newline)
    [ "$result" = "$(echo -e 'hello\nworld')" ]
}

@test "strip-tags removes tags completely" {
    result=$(echo "hello<br>world" | $APP -u strip-tags)
    [ "$result" = "helloworld" ]
}

@test "unescape-html decodes entities" {
    result=$(echo "&lt;div&gt;" | $APP -u unescape-html)
    [ "$result" = "<div>" ]
}

@test "backlink wraps nth occurrence" {
    result=$(echo "example text with example" | $APP -u backlink example 2 "http://test.com")
    expected='example text with <a href="http://test.com">example</a>'
    [ "$result" = "$expected" ]
}

@test "backlink first occurrence" {
    result=$(echo "example text example" | $APP -u backlink example 1 "http://test.com")
    expected='<a href="http://test.com">example</a> text example'
    [ "$result" = "$expected" ]
}

@test "works with file input" {
    echo "example text" > /tmp/test_input.txt
    result=$($APP -u backlink example 1 "http://test.com" < /tmp/test_input.txt)
    expected='<a href="http://test.com">example</a> text'
    [ "$result" = "$expected" ]
    rm /tmp/test_input.txt
}

@test "works with -s flag" {
    result=$($APP -u backlink -s "example text" example 1 "http://test.com")
    expected='<a href="http://test.com">example</a> text'
    [ "$result" = "$expected" ]
}
