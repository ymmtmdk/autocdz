function autocdz
  function on_fish_postexec --on-event fish_postexec
    if test -d $argv
      cd $argv
    else if z_jump $argv
      echo cd $PWD
    end
  end
end
