# srl_agent_helloIntf

This is a starter example for using Nokia's SRLinux NDK; this example is a very simple application where you build an agent that is just continously reporting the total number of interfaces that are in an admin up state on the running system

It is simply using a GRPC subscription to recieve notifications when there is any interface state change; once a notification is recieved it is then used to update a total count & push via telemetry back to the NDK manager the desired state that just reports how many interfaces are admin-up 

### To install on an SRL system:

* copy the binary helloIntf to /etc/opt/srlinux/appmgr/user_agents/
* copy the yml file under yml directory to /etc/opt/srlinux/appmgr/
* copy the yang file under yang directory to /etc/opt/srlinux/appmgr/user_agents/yang
* run via CLI "tools system app-management application app_mgr reload"

### To use the agent:

The agent will be installed under helloIntf path in the configuration

```
--{ + candidate shared default }--[ helloIntf ]--
A:leaf1# info
    action enable
    debug enable
```

```
--{ candidate shared default }--[ helloIntf ]--
--{ + candidate shared default }--[ helloIntf ]--
A:leaf1# info from state
    action enable
    debug enable
    admin-up-count 3
    interfaces {
        interface ethernet-1/1 {
            state admin-up
        }
        interface ethernet-1/2 {
            state admin-up
        }
        interface mgmt0 {
            state admin-up
        }
    }

```    
The agent debug logs would be in /var/log/srlinux/stdout/

Since the agents are also exposed via GNMI then you can use your typical gnmi operations to get & set the agent state; for example for a get operation 

```
== getRequest:
path: <
  elem: <
    name: "helloIntf"
  >
>
encoding: JSON_IETF

== getResponse:
notification: <
  timestamp: 1619685289811545174
  update: <
    path: <
      elem: <
        name: "helloIntf:helloIntf"
      >
    >
    val: <
      json_ietf_val: "{\"action\": \"enable\", \"debug\": \"enable\", \"admin-up-count\": \"3\", \"interfaces\": {\"interface\": [{\"name\": \"ethernet-1/1\", \"state\": \"admin-up\"}, {\"name\": \"ethernet-1/2\", \"state\": \"admin-up\"}, {\"name\": \"mgmt0\", \"state\": \"admin-up\"}]}}"
    >
  >
>

```

