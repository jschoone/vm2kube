Vagrant.configure("2") do |config|
  if Vagrant.has_plugin?("vagrant-hostmanager")
    config.hostmanager.enabled = true
    config.hostmanager.manage_host = true
    config.hostmanager.manage_guest = false
    config.hostmanager.ignore_private_ip = true
    config.hostmanager.include_offline = true
    config.vm.provision :hostmanager
  end


  ENV['VAGRANT_DEFAULT_PROVIDER'] = 'libvirt'

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

  node_name="node"
  memory = "1024"
  N = APPSRVNODES+LBNODES+DBNODES
  (1..N).each do |node_id|
    if node_id <= N-LBNODES-DBNODES
      node_name="appsrv-#{node_id-1}"
    elsif node_id > N-LBNODES-DBNODES && node_id <= N-DBNODES
      node_name="lb-#{node_id-1-APPSRVNODES}"
    else
      node_name="db-#{node_id-1-APPSRVNODES-LBNODES}"
    end
    config.vm.define "#{node_name}" do |node|
      node.vm.box = "generic/ubuntu2004"
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
            "appsrv" => ["appsrv-[0:#{APPSRVNODES-1}]"],
            "loadbalancer" => ["lb-[0:#{LBNODES-1}]"],
            "database" => ["db-[0:#{DBNODES-1}]"],
          }
          ansible.limit = "all"
          ansible.playbook = "plays/app.yaml"
        end
      end
    end
  end
end
