function autocdz
  if test (count $argv) -gt 1
    echo "fish: Unknown command '$argv[1]'" >&2
  else if test -d $argv
    cd $argv
  else if __z $argv
    echo cd $PWD
  end
end
