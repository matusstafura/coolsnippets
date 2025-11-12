# Brainstorm

general use:

```bash 
coolsnip -u remove-html -s "hello<br>world" # helloworld

echo "hello<br/>world" | ./coolsnippets -u remove-html # helloworld

./bin/coolsnippets-darwin-amd64 -u strip-tags-newline < file.txt

cat file.txt | ./bin/coolsnippets-darwin-amd64 -u strip-tags-newline | sort | uniq | ./bin/coolsnippets-darwin-amd64
 -u unescape-html

%!coolsnippets -u strip-tags-newline | awk '{$1=$1};1' | sed -r '/^\s*$/d' | sort -u
```

- [x] learn everything i dont understand in go code https://chatgpt.com/c/68f9e8dc-39f8-832d-8903-f9a93b8d53c9
- [x] flags
- [x] first snippet
- [x] test with coverage
- [x] nvim integration
- [x] bash integration
- [x] dev bash makes sense if neovim integration is done
- [x] deployment && builds
- [x] unescape html 
- [x] extract urls
- [x] backlinks
- [x] clean html from styles
- [x] extract href links
- [x] test benchmarks for every utility
- [x] integration tests
- [ ] url slug by line
- [?] multiple utility support like -u strip-tags-newline -u unescape-html OR -u strip-tags-newline,unescape-html
- [?] trim striptags fix
- [?] integration with VS Code, sublime text, etc
- [ ] readme with examples for every utility
- [!] extract skipped tags in stripattributes utility

### MVP
- [x] github actions for tests
- [ ] rename coolsnippets to ...?
- [ ] check readme for every utility
- [ ] check every utility
- [ ] check tests for every utility
- [ ] check benchmarks for every utility
- [ ] check deployment scripts
- [ ] check nvim integration
- [ ] check bash integration
- [ ] check go code for every utility
- [ ] check build scripts
- [ ] check examples in readme
- [ ] create GIF examples in readme
- [ ] check flags parsing
- [ ] check packaging for different OS
- [ ] check license
- [ ] check contributing guidelines

