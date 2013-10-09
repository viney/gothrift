namespace go rpc
namespace java rpc

service RpcService {
    list<string> findAll(1:i64 userid, 2:string password 3:map<string, string> param)
}
