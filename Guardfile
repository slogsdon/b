group :tests do
  guard :shell do
    watch(%r{.*\.go}) { `echo "" && go test ./...` }
  end
end