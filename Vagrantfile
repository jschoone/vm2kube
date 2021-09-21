ENV['VAGRANT_NO_PARALLEL'] = 'yes'
ENV['VAGRANT_DEFAULT_PROVIDER'] = 'libvirt'

Vagrant.configure("2") do |config|
  config.env.enable

  if ENV['APPSRVNODES']
    APPSRVNODES=ENV['APPSRVNODES'].to_i
  else
    APPSRVNODES=1
  end

  if ENV['LBNODES']
    LBNODES=ENV['LBNODES'].to_i
  else
    LBNODES=0
  end

  if ENV['DBNODES']
    DBNODES=ENV['DBNODES'].to_i
  else
    DBNODES=1
  end

  if ENV['OS']
    os=ENV['OS']
  else
    os="ubuntu2004"
  end

  node_name="node"
  #os = "ubuntu2004"
  #os = "rocky8"
  os_list = ["ubuntu2004", "rocky8", "alma8"]

  N = APPSRVNODES+LBNODES+DBNODES
  (1..N).each do |node_id|

    memory = "512"
    if node_id <= N-LBNODES-DBNODES
      node_name="appsrv#{node_id-1}"
      memory = "2048"
    elsif node_id > N-LBNODES-DBNODES && node_id <= N-DBNODES
      node_name="lb#{node_id-1-APPSRVNODES}"
    else
      node_name="db#{node_id-1-APPSRVNODES-LBNODES}"
    end
    config.vm.define "#{node_name}" do |node|

      node.vm.box = "generic/#{os}"
      node.vm.provider "libvirt" do |libvirt|
          libvirt.memory = "#{memory}"
          libvirt.cpus = 2
      end
      node.vm.provider "virtualbox" do |vb|
          vb.memory = "#{memory}"
          vb.cpus = 2
      end
      if node_id == N
        node.vm.provision :ansible do |ansible|
          ansible.groups = {
            "appsrv" => ["appsrv[0:#{APPSRVNODES-1}]"],
            "loadbalancer" => ["lb[0:#{LBNODES-1}]"],
            "database" => ["db[0:#{DBNODES-1}]"],
            "dns" => ["db[0:#{DBNODES-1}]"],
            "registry" => ["db[0:#{DBNODES-1}]"],
            "localhost" => ["localhost"],
          }
          ansible.limit = "all"
          ansible.playbook = "plays/site.yaml"
        end
      end
    end
  end
end
