local thread_id = 1

function setup(thread)
    thread:set("id", thread_id)
    thread_id = thread_id + 1
end

function load_url_paths_from_file(file)
    local lines = {}
  
    local f = io.open(file,"r")
    if f ~= nil then
      io.close(f)
    else
      return lines
    end
  
    for line in io.lines(file) do
      if line ~= '' then
        lines[#lines + 1] = line
      end
    end
  
    return lines
  end
  
  -- Path is valid inside wrk container only
  -- (change it to run wrk on local host)
  paths = load_url_paths_from_file("/ego/paths.txt")
 
  function init(args)
    math.randomseed(id)
  end

  request = function()
    local index = math.random(#paths)
    url_path = paths[index]
    return wrk.format(nil, url_path)
  end
