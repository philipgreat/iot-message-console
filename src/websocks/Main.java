package websocks;

import java.io.IOException;

import javax.websocket.DeploymentException;
import javax.xml.ws.Service;

public class Main{
	
	
	public static void main(String[] args) throws DeploymentException, IOException, InterruptedException {
		Server server = new Server("localhost", 8080, "", null, WebSocketTest.class);
        server.start();
        System.out.print("---- Server Started -----");
        //new CountDownLatch(1).await();
    }
	
	
}