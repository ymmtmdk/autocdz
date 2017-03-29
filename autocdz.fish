function autocdz
  function on_fish_postexec --on-event fish_postexec
    if test -z $argv
      ll
    end
  end
end
