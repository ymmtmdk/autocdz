if test -z "$Z_DATA"
  set -U Z_DATA "$HOME/.z"
end

if test ! -f "$Z_DATA"
  touch "$Z_DATA"
end

function __z_add -d "Add PATH to .z file"
  set -l path (dirname (status -f))
  set -l tmpfile (mktemp $Z_DATA.XXXXXX)

  if test -f $tmpfile
    zadd "$argv" (date +%s) $Z_DATA ^ /dev/null > $tmpfile
    mv -f $tmpfile $Z_DATA
  end
end

function __z_move -d "Jump to a recent directory"
  set -g path (dirname (status -f))
  set -l target

  set target (zmove "$argv" (date +%s) "$Z_DATA")

  if test "$status" -gt 0
    return 1
  end

  if test -z "$target"
    printf "'%s' did not match any z results" "$argv"
    return 1
  end

  if test -d $target
    pushd "$target"
    echo cd $target
  end
end

function __z_on_variable_pwd --on-variable PWD
  __z_add $PWD
end

function fish_command_not_found
  autocdz $argv
end

function autocdz
  if test (count $argv) -gt 1
    echo "fish: Unknown command '$argv[1]'" >&2
  else if test -d $argv
    cd $argv
  else
    __z_move $argv
  end
end

