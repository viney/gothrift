package rpc;

import java.util.HashMap;
import java.util.Map;
import java.util.List;

import org.apache.thrift.TException;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.transport.TFramedTransport;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;

/**
 * Thrift测试客户端
 */
public class Main {

	public static void main(String[] args) {
		
		try {
			TTransport transport = new TFramedTransport(new TSocket("192.168.1.241", 8000));
			
			TBinaryProtocol protocol = new TBinaryProtocol(transport);
			
			RpcService.Client client = new RpcService.Client(protocol);
			transport.open();
			
			Map<String, String> param = new HashMap<String, String>();
			param.put("name", "viney");
			param.put("email", "viney.chow@gmail.com");
			
	        List<String> res = client.findAll(1, "123456", param);

            System.out.println(res);
			
			transport.close();
		} catch (TException x) {
			x.printStackTrace();
		}
	}
}
