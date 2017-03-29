function __fish_default_command_not_found_handler
  if test -d $argv
    cd $argv
  # else if __z $argv
  else if z_jump $argv
    echo cd $PWD
  end
end
