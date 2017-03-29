function autocdz
  if test -d $argv
    cd $argv
  else if z_jump $argv
    echo cd $PWD
  end
end
