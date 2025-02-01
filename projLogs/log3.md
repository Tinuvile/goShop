# 2025/2/1

## 配置管理

### 配置目录

```text
auth
└── conf    
    ├── dev             // 开发配置
    │   └── conf.yaml
    ├── online          // 生产配置
    │   └── conf.yaml
    ├── test            // 测试配置
    │   └── conf.yaml
    └── conf.go          // 解析所有配置文件
```

### 修改配置文件以占位符替代

```yaml
# demo/auth/conf/test/conf.yaml:
mysql:
#  dsn: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
  dsn: "%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
#  占位符（%s）顺序：用户名、密码、主机地址、数据库名
```

```go
// demo/auth/biz/dal/mysql/init.go
func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, 
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_ROOT_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
}
```

```powershell
PS F:\goShop\goShop> $env:MYSQL_USER = "root"
```

### .env文件编写

```dotenv
MYSQL_USER=root
MYSQL_PASSWORD=root
MYSQL_HOST=localhost
MYSQL_DATABASE=test
```

### 安装godotenv

```powershell
PS F:\goShop\goShop> go get github.com/joho/godotenv
go: added github.com/joho/godotenv v1.5.1
```

### 修改相应文件

```go
// demo/auth/main.go/func main
err := godotenv.Load()
if err != nil {
panic("Error loading .env file")
}
dal.Init()
```

```yaml
# docker-compose.yaml
version: "3"
services:
  consul:
    image: "consul:1.15.4"
    ports:
      - "8500:8500"
  mysql:
    image: "mysql:latest"
    ports:
      - "3306:3306"
    environment:
      - MYSQL_ROOT_PASSWORD=root
```

实际项目开发时最好固定版本号并设置一个较为安全的密码

### 启动容器

```powershell
PS F:\goShop\goShop> docker-compose up
[+] Running 10/10
 ✔ mysql Pulled                                                                                                                                                                                   153.0s 
   ✔ d43055c38217 Download complete                                                                                                                                                                 6.3s 
   ✔ facc8f3107c1 Download complete                                                                                                                                                                 8.9s 
   ✔ c294da35c13e Download complete                                                                                                                                                                 2.3s 
   ✔ 139aca660b47 Download complete                                                                                                                                                                49.8s 
   ✔ 21577e00f2ba Download complete                                                                                                                                                                 3.3s 
   ✔ 4643f1cf56c2 Download complete                                                                                                                                                                 0.9s 
   ✔ de4342aa4ad8 Download complete                                                                                                                                                                 1.0s 
   ✔ b10e1082570e Download complete                                                                                                                                                                 1.0s 
   ✔ 26313a3e0799 Download complete                                                                                                                                                               140.3s 
[+] Running 2/2
 ✔ Container goshop-mysql-1   Created                                                                                                                                                               0.7s 
 ✔ Container goshop-consul-1  Created                                                                                                                                                               0.0s 
Attaching to consul-1, mysql-1
mysql-1   | 2025-02-01 10:53:48+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 9.2.0-1.el9 started.
consul-1  | ==> Starting Consul agent...
consul-1  |               Version: '1.15.4'
consul-1  |            Build Date: '2023-06-23 23:14:17 +0000 UTC'                                                                                                                                       
consul-1  |               Node ID: '2abdb849-47af-7029-198d-002ee4031f31'                                                                                                                                
consul-1  |             Node name: 'd34a957abaaa'                                                                                                                                                        
consul-1  |            Datacenter: 'dc1' (Segment: '<all>')                                                                                                                                              
consul-1  |                Server: true (Bootstrap: false)
consul-1  |           Client Addr: [0.0.0.0] (HTTP: 8500, HTTPS: -1, gRPC: 8502, gRPC-TLS: 8503, DNS: 8600)                                                                                              
consul-1  |          Cluster Addr: 127.0.0.1 (LAN: 8301, WAN: 8302)                                                                                                                                      
consul-1  |     Gossip Encryption: false                                                                                                                                                                 
consul-1  |      Auto-Encrypt-TLS: false
consul-1  |      Reporting Enabled: false                                                                                                                                                                
mysql-1   | 2025-02-01 10:53:48+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'                                                                                                           
consul-1  |             HTTPS TLS: Verify Incoming: false, Verify Outgoing: false, Min Version: TLSv1_2                                                                                                  
consul-1  |              gRPC TLS: Verify Incoming: false, Min Version: TLSv1_2                                                                                                                          
consul-1  |      Internal RPC TLS: Verify Incoming: false, Verify Outgoing: false (Verify Hostname: false), Min Version: TLSv1_2                                                                         
consul-1  | 
mysql-1   | 2025-02-01 10:53:48+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 9.2.0-1.el9 started.                                                                                       
consul-1  | ==> Log data will now stream in as it occurs:                                                                                                                                                
consul-1  |                                                                                                                                                                                              
consul-1  | 2025-02-01T10:53:48.440Z [DEBUG] agent.grpc.balancer: switching server: target=consul://dc1.2abdb849-47af-7029-198d-002ee4031f31/server.dc1 from=<none> to=<none>                            
consul-1  | 2025-02-01T10:53:48.461Z [INFO]  agent.server.raft: initial configuration: index=1 servers="[{Suffrage:Voter ID:2abdb849-47af-7029-198d-002ee4031f31 Address:127.0.0.1:8300}]"               
consul-1  | 2025-02-01T10:53:48.461Z [INFO]  agent.server.raft: entering follower state: follower="Node at 127.0.0.1:8300 [Follower]" leader-address= leader-id=
consul-1  | 2025-02-01T10:53:48.462Z [INFO]  agent.server.serf.wan: serf: EventMemberJoin: d34a957abaaa.dc1 127.0.0.1                                                                                    
consul-1  | 2025-02-01T10:53:48.463Z [INFO]  agent.server.serf.lan: serf: EventMemberJoin: d34a957abaaa 127.0.0.1                                                                                        
consul-1  | 2025-02-01T10:53:48.463Z [INFO]  agent.router: Initializing LAN area manager                                                                                                                 
consul-1  | 2025-02-01T10:53:48.463Z [DEBUG] agent.grpc.balancer: switching server: target=consul://dc1.2abdb849-47af-7029-198d-002ee4031f31/server.dc1 from=<none> to=dc1-127.0.0.1:8300                
consul-1  | 2025-02-01T10:53:48.463Z [INFO]  agent.server.autopilot: reconciliation now disabled
consul-1  | 2025-02-01T10:53:48.464Z [INFO]  agent.server: Handled event for server in area: event=member-join server=d34a957abaaa.dc1 area=wan                                                          
consul-1  | 2025-02-01T10:53:48.464Z [INFO]  agent.server: Adding LAN server: server="d34a957abaaa (Addr: tcp/127.0.0.1:8300) (DC: dc1)"                                                                 
consul-1  | 2025-02-01T10:53:48.466Z [DEBUG] agent.server.autopilot: autopilot is now running                                                                                                            
consul-1  | 2025-02-01T10:53:48.466Z [DEBUG] agent.server.autopilot: state update routine is now running                                                                                                 
consul-1  | 2025-02-01T10:53:48.466Z [DEBUG] agent.hcp_manager: HCP manager starting                                                                                                                     
consul-1  | 2025-02-01T10:53:48.466Z [INFO]  agent.server.cert-manager: initialized server certificate management
consul-1  | 2025-02-01T10:53:48.468Z [INFO]  agent: Started DNS server: address=0.0.0.0:8600 network=tcp                                                                                                 
consul-1  | 2025-02-01T10:53:48.468Z [INFO]  agent: Started DNS server: address=0.0.0.0:8600 network=udp                                                                                                 
consul-1  | 2025-02-01T10:53:48.469Z [INFO]  agent: Starting server: address=[::]:8500 network=tcp protocol=http                                                                                         
consul-1  | 2025-02-01T10:53:48.469Z [INFO]  agent: Started gRPC listeners: port_name=grpc_tls address=[::]:8503 network=tcp                                                                             
consul-1  | 2025-02-01T10:53:48.469Z [INFO]  agent: Started gRPC listeners: port_name=grpc address=[::]:8502 network=tcp                                                                                 
consul-1  | 2025-02-01T10:53:48.469Z [INFO]  agent: started state syncer                                                                                                                                 
consul-1  | 2025-02-01T10:53:48.469Z [INFO]  agent: Consul agent running!                                                                                                                                
consul-1  | 2025-02-01T10:53:48.515Z [WARN]  agent.server.raft: heartbeat timeout reached, starting election: last-leader-addr= last-leader-id=
consul-1  | 2025-02-01T10:53:48.515Z [INFO]  agent.server.raft: entering candidate state: node="Node at 127.0.0.1:8300 [Candidate]" term=2                                                               
consul-1  | 2025-02-01T10:53:48.515Z [DEBUG] agent.server.raft: voting for self: term=2 id=2abdb849-47af-7029-198d-002ee4031f31                                                                          
consul-1  | 2025-02-01T10:53:48.515Z [DEBUG] agent.server.raft: calculated votes needed: needed=1 term=2                                                                                                 
consul-1  | 2025-02-01T10:53:48.515Z [DEBUG] agent.server.raft: vote granted: from=2abdb849-47af-7029-198d-002ee4031f31 term=2 tally=1                                                                   
consul-1  | 2025-02-01T10:53:48.515Z [INFO]  agent.server.raft: election won: term=2 tally=1                                                                                                             
consul-1  | 2025-02-01T10:53:48.515Z [INFO]  agent.server.raft: entering leader state: leader="Node at 127.0.0.1:8300 [Leader]"                                                                          
consul-1  | 2025-02-01T10:53:48.515Z [DEBUG] agent.hcp_manager: HCP triggering status update
consul-1  | 2025-02-01T10:53:48.516Z [INFO]  agent.server: cluster leadership acquired                                                                                                                   
consul-1  | 2025-02-01T10:53:48.517Z [INFO]  agent.server: New leader elected: payload=d34a957abaaa                                                                                                      
consul-1  | 2025-02-01T10:53:48.518Z [INFO]  agent.server.autopilot: reconciliation now enabled                                                                                                          
consul-1  | 2025-02-01T10:53:48.519Z [INFO]  agent.leader: started routine: routine="federation state anti-entropy"                                                                                      
consul-1  | 2025-02-01T10:53:48.519Z [INFO]  agent.leader: started routine: routine="federation state pruning"
consul-1  | 2025-02-01T10:53:48.519Z [INFO]  agent.leader: started routine: routine="streaming peering resources"                                                                                        
consul-1  | 2025-02-01T10:53:48.519Z [INFO]  agent.leader: started routine: routine="metrics for streaming peering resources"                                                                            
consul-1  | 2025-02-01T10:53:48.519Z [INFO]  agent.leader: started routine: routine="peering deferred deletion"                                                                                          
consul-1  | 2025-02-01T10:53:48.519Z [DEBUG] connect.ca.consul: consul CA provider configured: id=fb:50:9b:45:1a:65:15:c1:68:57:73:5f:da:cd:b8:0d:0f:e2:26:eb:68:66:43:11:85:9d:67:a9:7a:56:9c:b9 is_primary=true
consul-1  | 2025-02-01T10:53:48.522Z [INFO]  connect.ca: updated root certificates from primary datacenter
consul-1  | 2025-02-01T10:53:48.522Z [INFO]  connect.ca: initialized primary datacenter CA with provider: provider=consul                                                                                
consul-1  | 2025-02-01T10:53:48.522Z [INFO]  agent.leader: started routine: routine="intermediate cert renew watch"                                                                                      
consul-1  | 2025-02-01T10:53:48.522Z [INFO]  agent.leader: started routine: routine="CA root pruning"                                                                                                    
consul-1  | 2025-02-01T10:53:48.522Z [INFO]  agent.leader: started routine: routine="CA root expiration metric"                                                                                          
consul-1  | 2025-02-01T10:53:48.522Z [INFO]  agent.leader: started routine: routine="CA signing expiration metric"                                                                                       
consul-1  | 2025-02-01T10:53:48.522Z [INFO]  agent.leader: started routine: routine="virtual IP version check"
consul-1  | 2025-02-01T10:53:48.522Z [INFO]  agent.leader: started routine: routine="config entry controllers"                                                                                           
consul-1  | 2025-02-01T10:53:48.522Z [DEBUG] agent.server: successfully established leadership: duration=5.519941ms                                                                                      
consul-1  | 2025-02-01T10:53:48.522Z [INFO]  agent.server: member joined, marking health alive: member=d34a957abaaa partition=default                                                                    
consul-1  | 2025-02-01T10:53:48.522Z [INFO]  agent.leader: stopping routine: routine="virtual IP version check"                                                                                          
consul-1  | 2025-02-01T10:53:48.522Z [INFO]  agent.leader: stopped routine: routine="virtual IP version check"                                                                                           
consul-1  | 2025-02-01T10:53:48.524Z [DEBUG] agent.server.xds_capacity_controller: updating drain rate limit: rate_limit=1                                                                               
consul-1  | 2025-02-01T10:53:48.592Z [DEBUG] agent.server.cert-manager: got cache update event: correlationID=leaf error=<nil>
consul-1  | 2025-02-01T10:53:48.592Z [DEBUG] agent.server.cert-manager: leaf certificate watch fired - updating auto TLS certificate: uri=spiffe://a493ddea-7dc5-674e-22db-5e20cee84ccd.consul/agent/server/dc/dc1                                                                                                                                                                                                
consul-1  | 2025-02-01T10:53:48.677Z [INFO]  agent.server: federation state anti-entropy synced
consul-1  | 2025-02-01T10:53:48.778Z [DEBUG] agent: Skipping remote check since it is managed automatically: check=serfHealth
consul-1  | 2025-02-01T10:53:48.793Z [INFO]  agent: Synced node info
mysql-1   | 2025-02-01 10:53:48+00:00 [Note] [Entrypoint]: Initializing database files
mysql-1   | 2025-02-01T10:53:48.813921Z 0 [System] [MY-015017] [Server] MySQL Server Initialization - start.                                                                                             
mysql-1   | 2025-02-01T10:53:48.816260Z 0 [System] [MY-013169] [Server] /usr/sbin/mysqld (mysqld 9.2.0) initializing of server in progress as process 80
mysql-1   | 2025-02-01T10:53:48.827100Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.                                                                                               
mysql-1   | 2025-02-01T10:53:49.140422Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
consul-1  | 2025-02-01T10:53:49.271Z [DEBUG] agent: Skipping remote check since it is managed automatically: check=serfHealth
consul-1  | 2025-02-01T10:53:49.271Z [DEBUG] agent: Node info in sync
consul-1  | 2025-02-01T10:53:49.271Z [DEBUG] agent: Node info in sync                                                                                                                                    
consul-1  | 2025-02-01T10:53:49.467Z [DEBUG] agent.server.cert-manager: CA config watch fired - updating auto TLS server name: name=server.dc1.peering.a493ddea-7dc5-674e-22db-5e20cee84ccd.consul       
mysql-1   | 2025-02-01T10:53:50.588577Z 6 [Warning] [MY-010453] [Server] root@localhost is created with an empty password ! Please consider switching off the --initialize-insecure option.
mysql-1   | 2025-02-01T10:53:53.109042Z 0 [System] [MY-015018] [Server] MySQL Server Initialization - end.
mysql-1   | 2025-02-01 10:53:53+00:00 [Note] [Entrypoint]: Database files initialized
mysql-1   | 2025-02-01 10:53:53+00:00 [Note] [Entrypoint]: Starting temporary server
mysql-1   | 2025-02-01T10:53:53.240139Z 0 [System] [MY-015015] [Server] MySQL Server - start.                                                                                                            
mysql-1   | 2025-02-01T10:53:53.483098Z 0 [System] [MY-010116] [Server] /usr/sbin/mysqld (mysqld 9.2.0) starting as process 125
mysql-1   | 2025-02-01T10:53:53.503546Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.                                                                                               
mysql-1   | 2025-02-01T10:53:53.873997Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.                                                                                                 
mysql-1   | 2025-02-01T10:53:54.297972Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
mysql-1   | 2025-02-01T10:53:54.298047Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
mysql-1   | 2025-02-01T10:53:54.302767Z 0 [Warning] [MY-011810] [Server] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.                                                                                                                                                                                    
mysql-1   | 2025-02-01T10:53:54.340985Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Socket: /var/run/mysqld/mysqlx.sock
mysql-1   | 2025-02-01T10:53:54.341286Z 0 [System] [MY-010931] [Server] /usr/sbin/mysqld: ready for connections. Version: '9.2.0'  socket: '/var/run/mysqld/mysqld.sock'  port: 0  MySQL Community Server - GPL.                                                                                                                                                                                                  
mysql-1   | 2025-02-01 10:53:54+00:00 [Note] [Entrypoint]: Temporary server started.
mysql-1   | '/var/lib/mysql/mysql.sock' -> '/var/run/mysqld/mysqld.sock'
mysql-1   | Warning: Unable to load '/usr/share/zoneinfo/iso3166.tab' as time zone. Skipping it.
mysql-1   | Warning: Unable to load '/usr/share/zoneinfo/leap-seconds.list' as time zone. Skipping it.
mysql-1   | Warning: Unable to load '/usr/share/zoneinfo/leapseconds' as time zone. Skipping it.                                                                                                         
mysql-1   | Warning: Unable to load '/usr/share/zoneinfo/tzdata.zi' as time zone. Skipping it.                                                                                                           
mysql-1   | Warning: Unable to load '/usr/share/zoneinfo/zone.tab' as time zone. Skipping it.
mysql-1   | Warning: Unable to load '/usr/share/zoneinfo/zone1970.tab' as time zone. Skipping it.                                                                                                        
mysql-1   |                                                                                                                                                                                              
mysql-1   | 2025-02-01 10:53:56+00:00 [Note] [Entrypoint]: Stopping temporary server
mysql-1   | 2025-02-01T10:53:56.244750Z 11 [System] [MY-013172] [Server] Received SHUTDOWN from user root. Shutting down mysqld (Version: 9.2.0).                                                        
mysql-1   | 2025-02-01T10:53:57.787125Z 0 [System] [MY-010910] [Server] /usr/sbin/mysqld: Shutdown complete (mysqld 9.2.0)  MySQL Community Server - GPL.
mysql-1   | 2025-02-01T10:53:57.787177Z 0 [System] [MY-015016] [Server] MySQL Server - end.
mysql-1   | 2025-02-01 10:53:58+00:00 [Note] [Entrypoint]: Temporary server stopped                                                                                                                      
mysql-1   | 
mysql-1   | 2025-02-01 10:53:58+00:00 [Note] [Entrypoint]: MySQL init process done. Ready for start up.                                                                                                  
mysql-1   |                                                                                                                                                                                              
mysql-1   | 2025-02-01T10:53:58.306226Z 0 [System] [MY-015015] [Server] MySQL Server - start.
mysql-1   | 2025-02-01T10:53:58.571130Z 0 [System] [MY-010116] [Server] /usr/sbin/mysqld (mysqld 9.2.0) starting as process 1
mysql-1   | 2025-02-01T10:53:58.598491Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.                                                                                               
mysql-1   | 2025-02-01T10:53:59.056600Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
mysql-1   | 2025-02-01T10:53:59.433635Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
mysql-1   | 2025-02-01T10:53:59.433718Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
mysql-1   | 2025-02-01T10:53:59.437545Z 0 [Warning] [MY-011810] [Server] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.                                                                                                                                                                                    
mysql-1   | 2025-02-01T10:53:59.472528Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /var/run/mysqld/mysqlx.sock
mysql-1   | 2025-02-01T10:53:59.472937Z 0 [System] [MY-010931] [Server] /usr/sbin/mysqld: ready for connections. Version: '9.2.0'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server - GPL.
ver - GPL.
consul-1  | 2025-02-01T10:55:00.240Z [DEBUG] agent: Skipping remote check since it is managed automatically: check=serfHealth
consul-1  | 2025-02-01T10:55:00.241Z [DEBUG] agent: Node info in sync
consul-1  | 2025-02-01T10:55:11.929Z [DEBUG] agent.http: Request finished: method=GET url=/v1/catalog/datacenters from=172.18.0.1:51656 latency="315.546µs"
consul-1  | 2025-02-01T10:55:12.121Z [DEBUG] agent.http: Request finished: method=GET url=/v1/internal/ui/services?dc=dc1 from=172.18.0.1:51656 latency="377.937µs"
consul-1  | 2025-02-01T10:55:40.456Z [DEBUG] agent.router.manager: Rebalanced servers, new active server: number_of_servers=1 active_server="d34a957abaaa.dc1 (Addr: tcp/127.0.0.1:8300) (DC: dc1)"
consul-1  | 2025-02-01T10:55:40.456Z [DEBUG] agent.router.manager: Rebalanced servers, new active server: number_of_servers=1 active_server="d34a957abaaa (Addr: tcp/127.0.0.1:8300) (DC: dc1)"
consul-1  | 2025-02-01T10:56:05.665Z [DEBUG] agent: Skipping remote check since it is managed automatically: check=serfHealth
consul-1  | 2025-02-01T10:56:05.665Z [DEBUG] agent: Node info in sync
consul-1  | 2025-02-01T10:57:03.362Z [DEBUG] agent: Skipping remote check since it is managed automatically: check=serfHealth
consul-1  | 2025-02-01T10:57:03.362Z [DEBUG] agent: Node info in sync
consul-1  | 2025-02-01T10:57:33.336Z [DEBUG] agent.router.manager: Rebalanced servers, new active server: number_of_servers=1 active_server="d34a957abaaa.dc1 (Addr: tcp/127.0.0.1:8300) (DC: dc1)"
consul-1  | 2025-02-01T10:58:20.166Z [DEBUG] agent.router.manager: Rebalanced servers, new active server: number_of_servers=1 active_server="d34a957abaaa (Addr: tcp/127.0.0.1:8300) (DC: dc1)"
consul-1  | 2025-02-01T10:58:53.002Z [DEBUG] agent: Skipping remote check since it is managed automatically: check=serfHealth
consul-1  | 2025-02-01T10:58:53.002Z [DEBUG] agent: Node info in sync
consul-1  | 2025-02-01T10:59:29.102Z [DEBUG] agent.router.manager: Rebalanced servers, new active server: number_of_servers=1 active_server="d34a957abaaa.dc1 (Addr: tcp/127.0.0.1:8300) (DC: dc1)"
```

### 创建test数据库

```powershell
PS F:\goShop\goShop> docker-compose exec mysql bash
bash-5.1# mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 15
Server version: 9.2.0 MySQL Community Server - GPL

Copyright (c) 2000, 2025, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
mysql> SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
4 rows in set (0.03 sec)

mysql> CREATE DATABASE test;
Query OK, 1 row affected (0.03 sec)
```

### 启动项目

```powershell
PS F:\goShop\goShop\demo\auth> go run .
&{Env:test Kitex:{Service:auth Address::8888 LogLevel:info LogFileName:log/kitex.log LogMaxSize:10 LogMaxBackups:50 LogMaxAge:3} MySQL:{DSN:%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local} Redis:{Address:127.0.0.1:6379 Username: Password: DB:0} Registry:{RegistryAddress:[127.0.0.1:8500] Username: Password:}}
```

### 简单测试

```go
// demo/auth/biz/dal/mysql/init.go
func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_ROOT_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	// 测试使用
	type Version struct {
		Version string
	}

	var v Version

	err = DB.Raw("select version() as version").Scan(&v).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(v)
}
```

运行

```powershell
PS F:\goShop\goShop\demo\auth> go run .
&{Env:test Kitex:{Service:auth Address::8888 LogLevel:info LogFileName:log/kitex.log LogMaxSize:10 LogMaxBackups:50 LogMaxAge:3} MySQL:{DSN:%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local} Redis:{Address:127.0.0.1:6379 Username: Password: DB:0} Registry:{RegistryAddress:[127.0.0.1:8500] Username: Password:}}
{9.2.0}
```

## 配置中心

常用的有[Etcd](https://github.com/kitex-contrib/config-etcd)，
[Consul](https://github.com/kitex-contrib/config-consul)，
[Zookeeper](https://github.com/kitex-contrib/config-zookeeper)，
[Nacos](https://github.com/kitex-contrib/config-nacos)，
[Apollo](https://github.com/kitex-contrib/config-apollo)等

